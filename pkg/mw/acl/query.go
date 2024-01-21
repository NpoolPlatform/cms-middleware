package acl

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	entacl "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
)

type queryHandler struct {
	*Handler
	stm   *ent.ACLSelect
	infos []*npool.ACL
	total uint32
}

func (h *queryHandler) selectACL(stm *ent.ACLQuery) {
	h.stm = stm.Select(
		entacl.FieldID,
		entacl.FieldEntID,
		entacl.FieldAppID,
		entacl.FieldRoleID,
		entacl.FieldArticleKey,
		entacl.FieldCreatedAt,
		entacl.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryACL(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.ACL.Query().Where(entacl.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entacl.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entacl.EntID(*h.EntID))
	}
	h.selectACL(stm)
	return nil
}

func (h *queryHandler) queryACLs(ctx context.Context, cli *ent.Client) error {
	stm, err := aclcrud.SetQueryConds(cli.ACL.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectACL(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetACL(ctx context.Context) (*npool.ACL, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryACL(cli); err != nil {
			return err
		}
		const limit = 2
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(limit).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetACLs(ctx context.Context) ([]*npool.ACL, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryACLs(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Asc(entacl.FieldCreatedAt)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
