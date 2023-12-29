package category

import (
	"context"
	"fmt"

	categorycrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/category"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkCategoryExist(ctx context.Context) error {
	h.Conds = &categorycrud.Conds{
		Name:  &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
		AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
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

func (h *createHandler) checkParentExist(ctx context.Context) error {
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

func (h *Handler) CreateCategory(ctx context.Context) (*npool.Category, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid name")
	}

	if h.ParentID != nil {
		key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateAppGood, *h.AppID, *h.ParentID, *h.Slug)
		if err := redis2.TryLock(key, 0); err != nil {
			return nil, err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()
	}

	handler := &createHandler{
		Handler: h,
	}

	if err := handler.checkCategoryExist(ctx); err != nil {
		return nil, err
	}

	if err := handler.checkParentExist(ctx); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := categorycrud.CreateSet(
			cli.Category.Create(),
			&categorycrud.Req{
				EntID:    h.EntID,
				AppID:    h.AppID,
				ParentID: h.ParentID,
				Name:     h.Name,
				Slug:     h.Slug,
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
