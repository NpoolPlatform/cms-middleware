package article

import (
	"context"
	"fmt"

	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	categorycrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/category"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	categorymw "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	oldArticleInfo *npool.Article
}

func (h *createHandler) checkArticleExist(ctx context.Context) error {
	latest := true
	h.Conds = &articlecrud.Conds{
		Title:      &cruder.Cond{Op: cruder.EQ, Val: *h.Title},
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		Latest:     &cruder.Cond{Op: cruder.EQ, Val: latest},
		ArticleKey: &cruder.Cond{Op: cruder.NEQ, Val: *h.ArticleKey},
	}
	exist, err := h.ExistArticleConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("arleady exists")
	}
	return nil
}

func (h *createHandler) checkCategoryExist(ctx context.Context) error {
	handler, err := categorymw.NewHandler(ctx)
	if err != nil {
		return err
	}
	handler.Conds = &categorycrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.CategoryID},
		AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
	}
	exist, err := handler.ExistCategoryConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid categoryid")
	}
	return nil
}

func (h *createHandler) getLatestedArticle(infos []*npool.Article) *npool.Article {
	version := uint32(1)
	latestInfo := &npool.Article{}
	for _, info := range infos {
		if info.Version >= version {
			version = info.Version
			latestInfo = info
		}
	}
	return latestInfo
}

func (h *createHandler) checkVersion(ctx context.Context, tx *ent.Tx) error {
	if h.Version == nil {
		return fmt.Errorf("invalid version")
	}
	h.Conds = &articlecrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		ArticleKey: &cruder.Cond{Op: cruder.NEQ, Val: *h.ArticleKey},
	}
	infos, _, err := h.GetArticles(ctx)
	if err != nil {
		return err
	}

	if len(infos) == 0 {
		if *h.Version == 1 {
			return nil
		}
		return fmt.Errorf("invalid version")
	}

	latestInfo := h.getLatestedArticle(infos)
	h.oldArticleInfo = latestInfo
	newVersion := latestInfo.Version + 1
	if *h.Version != newVersion {
		return fmt.Errorf("invalid version")
	}

	if err := h.updateArticleLatest(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createArticle(ctx context.Context, tx *ent.Tx) error {
	if _, err := articlecrud.CreateSet(
		tx.Article.Create(),
		&articlecrud.Req{
			EntID:      h.EntID,
			AppID:      h.AppID,
			CategoryID: h.CategoryID,
			AuthorID:   h.AuthorID,
			ArticleKey: h.ArticleKey,
			Title:      h.Title,
			Subtitle:   h.Subtitle,
			Digest:     h.Digest,
			Status:     h.Status,
			Version:    h.Version,
			Latest:     h.Latest,
			Host:       h.Host,
			ISO:        h.ISO,
			ContentURL: h.ContentURL,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateArticleLatest(ctx context.Context, tx *ent.Tx) error {
	latest := false
	if _, err := articlecrud.UpdateSet(
		tx.Article.UpdateOneID(h.oldArticleInfo.ID),
		&articlecrud.Req{
			Latest: &latest,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateArticle(ctx context.Context) (*npool.Article, error) {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppGood, *h.AppID, *h.Title)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler := &createHandler{
		Handler: h,
	}

	if err := handler.checkCategoryExist(ctx); err != nil {
		return nil, err
	}

	if err := handler.checkArticleExist(ctx); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkVersion(ctx, tx); err != nil {
			return err
		}
		if err := handler.createArticle(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetArticle(ctx)
}
