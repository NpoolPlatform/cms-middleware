package acl

import (
	"context"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"

	"github.com/google/uuid"
)

func (h *Handler) CreateACL(ctx context.Context) (*npool.ACL, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

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
	if err != nil {
		return nil, err
	}

	return h.GetACL(ctx)
}
