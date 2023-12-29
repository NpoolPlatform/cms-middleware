package acl

import (
	"context"
	"fmt"
	"time"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
)

func (h *Handler) DeleteACL(ctx context.Context) (*npool.ACL, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetACL(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := aclcrud.UpdateSet(
			cli.ACL.UpdateOneID(*h.ID),
			&aclcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
