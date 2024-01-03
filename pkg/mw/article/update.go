package article

import (
	"context"
	"fmt"
	"time"

	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) checkStatus() error {
	if h.Status == nil {
		return nil
	}
	switch *h.Status {
	case types.ArticleStatus_Draft:
	case types.ArticleStatus_Published:
		now := uint32(time.Now().Unix())
		h.PublishedAt = &now
	default:
		return fmt.Errorf("invalid status")
	}
	return nil
}

func (h *updateHandler) updateArticle(ctx context.Context, tx *ent.Tx) error {
	if _, err := articlecrud.UpdateSet(
		tx.Article.UpdateOneID(*h.ID),
		&articlecrud.Req{
			CategoryID: h.CategoryID,
			AuthorID:   h.AuthorID,
			Title:      h.Title,
			Subtitle:   h.Subtitle,
			Digest:     h.Digest,
			Status:     h.Status,
			Latest:     h.Latest,
			ContentURL: h.ContentURL,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateArticle(ctx context.Context) (*npool.Article, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetArticle(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	if !info.Latest {
		return info, nil
	}
	appID := uuid.MustParse(info.AppID)
	h.AppID = &appID

	if h.Title != nil {
		latest := true
		h.Conds = &articlecrud.Conds{
			ID:     &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
			Title:  &cruder.Cond{Op: cruder.EQ, Val: *h.Title},
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			Latest: &cruder.Cond{Op: cruder.EQ, Val: latest},
		}
		exist, err := h.ExistArticleConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("arleady exists")
		}
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkStatus(); err != nil {
			return err
		}
		if err := handler.updateArticle(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetArticle(ctx)
}
