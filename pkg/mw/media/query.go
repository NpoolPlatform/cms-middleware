package media

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"

	mediacrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/media"
	entmedia "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/media"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"
)

type queryHandler struct {
	*Handler
	stm   *ent.MediaSelect
	infos []*npool.Media
	total uint32
}

func (h *queryHandler) selectMedia(stm *ent.MediaQuery) {
	h.stm = stm.Select(
		entmedia.FieldID,
		entmedia.FieldEntID,
		entmedia.FieldAppID,
		entmedia.FieldName,
		entmedia.FieldExt,
		entmedia.FieldMediaURL,
		entmedia.FieldCreatedBy,
		entmedia.FieldCreatedAt,
		entmedia.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryMedia(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Media.Query().Where(entmedia.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entmedia.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entmedia.EntID(*h.EntID))
	}
	h.selectMedia(stm)
	return nil
}

func (h *queryHandler) queryMedias(ctx context.Context, cli *ent.Client) error {
	stm, err := mediacrud.SetQueryConds(cli.Media.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectMedia(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetMedia(ctx context.Context) (*npool.Media, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryMedia(cli); err != nil {
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

func (h *Handler) GetMedias(ctx context.Context) ([]*npool.Media, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryMedias(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Asc(entmedia.FieldCreatedAt)).
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
