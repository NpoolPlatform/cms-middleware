package category

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"

	categorycrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/category"
	entcategory "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/category"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
)

type queryHandler struct {
	*Handler
	stm   *ent.CategorySelect
	infos []*npool.Category
	total uint32
}

func (h *queryHandler) selectCategory(stm *ent.CategoryQuery) {
	h.stm = stm.Select(
		entcategory.FieldID,
		entcategory.FieldEntID,
		entcategory.FieldAppID,
		entcategory.FieldParentID,
		entcategory.FieldName,
		entcategory.FieldSlug,
		entcategory.FieldEnabled,
		entcategory.FieldCreatedAt,
		entcategory.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCategory(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Category.Query().Where(entcategory.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcategory.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcategory.EntID(*h.EntID))
	}
	h.selectCategory(stm)
	return nil
}

func (h *queryHandler) queryCategories(ctx context.Context, cli *ent.Client) error {
	stm, err := categorycrud.SetQueryConds(cli.Category.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectCategory(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetCategory(ctx context.Context) (*npool.Category, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCategory(cli); err != nil {
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

func (h *Handler) GetCategories(ctx context.Context) ([]*npool.Category, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCategories(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Asc(entcategory.FieldCreatedAt)).
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
