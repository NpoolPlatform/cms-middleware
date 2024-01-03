package category

import (
	"context"
	"fmt"

	categorycrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/category"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) checkCategoryExist(ctx context.Context) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	info, err := h.GetCategory(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid category")
	}
	appID := uuid.MustParse(info.AppID)
	h.AppID = &appID
	return nil
}

func (h *updateHandler) checkParentExist(ctx context.Context) error {
	if h.ParentID == nil {
		return nil
	}
	h.Conds = &categorycrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.ParentID},
		AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
	}
	exist, err := h.ExistCategoryConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("parentid not exist")
	}
	return nil
}

func (h *updateHandler) checkNameRepeated(ctx context.Context) error {
	if h.Name == nil {
		return nil
	}
	parentID := uuid.Nil
	if h.ParentID == nil {
		info, err := h.GetCategory(ctx)
		if err != nil {
			return err
		}
		parentID = uuid.MustParse(info.ParentID)
	}
	h.Conds = &categorycrud.Conds{
		ID:       &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
		Name:     &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
		AppID:    &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		ParentID: &cruder.Cond{Op: cruder.EQ, Val: parentID},
	}
	exist, err := h.ExistCategoryConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("arleady exists")
	}
	return nil
}

func (h *Handler) UpdateCategory(ctx context.Context) (*npool.Category, error) {
	handler := &updateHandler{
		Handler: h,
	}

	if err := handler.checkCategoryExist(ctx); err != nil {
		return nil, err
	}

	if err := handler.checkParentExist(ctx); err != nil {
		return nil, err
	}

	if err := handler.checkNameRepeated(ctx); err != nil {
		return nil, err
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := categorycrud.UpdateSet(
			cli.Category.UpdateOneID(*h.ID),
			&categorycrud.Req{
				Name:     h.Name,
				ParentID: h.ParentID,
				Enabled:  h.Enabled,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCategory(ctx)
}
