package acl

import (
	"context"
	"fmt"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateACL(ctx context.Context) (*npool.ACL, error) {
	lockKey := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateAppLang, *h.AppID, *h.RoleID, *h.ArticleKey)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &aclcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		RoleID:     &cruder.Cond{Op: cruder.EQ, Val: *h.RoleID},
		ArticleKey: &cruder.Cond{Op: cruder.EQ, Val: *h.ArticleKey},
	}
	exist, err := h.ExistACLConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("acl exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
	if err != nil {
		return nil, err
	}

	return h.GetACL(ctx)
}
