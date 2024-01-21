package media

import (
	"context"

	mediacrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/media"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"

	"github.com/google/uuid"
)

func (h *Handler) CreateMedia(ctx context.Context) (*npool.Media, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := mediacrud.CreateSet(
			cli.Media.Create(),
			&mediacrud.Req{
				EntID:     h.EntID,
				AppID:     h.AppID,
				Name:      h.Name,
				Ext:       h.Ext,
				MediaURL:  h.MediaURL,
				CreatedBy: h.CreatedBy,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetMedia(ctx)
}
