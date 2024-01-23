package media

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/cms-middleware/pkg/const"
	mediacrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/media"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	CreatedBy *uuid.UUID
	Name      *string
	Ext       *string
	MediaURL  *string
	Conds     *mediacrud.Conds
	Offset    int32
	Limit     int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithCreatedBy(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CreatedBy = &_id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}

func WithExt(ext *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ext == nil {
			if must {
				return fmt.Errorf("invalid ext")
			}
			return nil
		}

		h.Ext = ext
		return nil
	}
}

func WithMediaURL(url *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if url == nil {
			if must {
				return fmt.Errorf("invalid mediaurl")
			}
			return nil
		}
		if *url == "" {
			return fmt.Errorf("invalid mediaurl")
		}
		h.MediaURL = url
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &mediacrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{Op: conds.GetName().GetOp(), Val: conds.GetName().GetValue()}
		}
		if conds.MediaURL != nil {
			h.Conds.MediaURL = &cruder.Cond{Op: conds.GetMediaURL().GetOp(), Val: conds.GetMediaURL().GetValue()}
		}
		if len(conds.GetEntIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}
		if len(conds.GetIDs().GetValue()) > 0 {
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: conds.GetIDs().GetValue()}
		}

		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
