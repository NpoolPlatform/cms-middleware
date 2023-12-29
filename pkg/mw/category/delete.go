package category

import (
	"context"
	"fmt"
	"time"

	categorycrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/category"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
)

func (h *Handler) DeleteCategory(ctx context.Context) (*npool.Category, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetCategory(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := categorycrud.UpdateSet(
			cli.Category.UpdateOneID(*h.ID),
			&categorycrud.Req{
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
