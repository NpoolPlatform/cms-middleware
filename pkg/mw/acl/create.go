package acl

import (
	"context"
	"fmt"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	articlemw "github.com/NpoolPlatform/cms-middleware/pkg/mw/article"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkArticleExist(ctx context.Context) error {
	handler, err := articlemw.NewHandler(ctx)
	if err != nil {
		return err
	}
	handler.Conds = &articlecrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		ArticleKey: &cruder.Cond{Op: cruder.EQ, Val: *h.ArticleKey},
	}
	exist, err := handler.ExistArticleConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid articlekey")
	}
	return nil
}

func (h *createHandler) checkACLExist(ctx context.Context) error {
	h.Conds = &aclcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		RoleID:     &cruder.Cond{Op: cruder.EQ, Val: *h.RoleID},
		ArticleKey: &cruder.Cond{Op: cruder.EQ, Val: *h.ArticleKey},
	}
	exist, err := h.ExistACLConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("acl exist")
	}
	return nil
}

func (h *createHandler) createACL(ctx context.Context) error {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := aclcrud.CreateSet(
			cli.ACL.Create(),
			&aclcrud.Req{
				EntID:      h.EntID,
				AppID:      h.AppID,
				RoleID:     h.RoleID,
				ArticleKey: h.ArticleKey,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (h *Handler) CreateACL(ctx context.Context) (*npool.ACL, error) {
	lockKey := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateACL, *h.AppID, *h.RoleID, *h.ArticleKey)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	handler := &createHandler{
		Handler: h,
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	if err := handler.checkArticleExist(ctx); err != nil {
		return nil, err
	}
	if err := handler.checkACLExist(ctx); err != nil {
		return nil, err
	}

	if err := handler.createACL(ctx); err != nil {
		return nil, err
	}

	return h.GetACL(ctx)
}
