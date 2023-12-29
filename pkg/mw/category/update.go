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
)

func (h *Handler) UpdateCategory(ctx context.Context) (*npool.Category, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	if h.ParentID != nil {
		key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateAppGood, *h.AppID, *h.ParentID, *h.Slug)
		if err := redis2.TryLock(key, 0); err != nil {
			return nil, err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()

		h.Conds = &categorycrud.Conds{
			ID:    &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
			Name:  &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		}
		exist, err := h.ExistCategoryConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("arleady exists")
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := categorycrud.UpdateSet(
			cli.Category.UpdateOneID(*h.ID),
			&categorycrud.Req{
				Name:    h.Name,
				Enabled: h.Enabled,
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
