package acl

import (
	"context"
	"fmt"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entacl "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
)

func (h *Handler) ExistACL(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			ACL.
			Query().
			Where(
				entacl.EntID(*h.EntID),
				entacl.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistACLConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := aclcrud.SetQueryConds(cli.ACL.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
