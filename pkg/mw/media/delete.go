package media

import (
	"context"
	"fmt"
	"time"

	mediacrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/media"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"
)

func (h *Handler) DeleteMedia(ctx context.Context) (*npool.Media, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetMedia(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := mediacrud.UpdateSet(
			cli.Media.UpdateOneID(*h.ID),
			&mediacrud.Req{
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
