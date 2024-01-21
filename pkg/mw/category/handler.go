package category

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/cms-middleware/pkg/const"
	categorycrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/category"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID       *uint32
	EntID    *uuid.UUID
	AppID    *uuid.UUID
	ParentID *uuid.UUID
	Name     *string
	Slug     *string
	Enabled  *bool
	Index    *uint32
	Conds    *categorycrud.Conds
	Offset   int32
	Limit    int32
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

func WithParentID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid parentid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		nilUUID := uuid.Nil.String()
		if *id == nilUUID {
			return nil
		}
		h.ParentID = &_id
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

func WithSlug(slug *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if slug == nil {
			if must {
				return fmt.Errorf("invalid slug")
			}
			return nil
		}
		if *slug == "" {
			return fmt.Errorf("invalid slug")
		}
		h.Slug = slug
		return nil
	}
}

func WithEnabled(enabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if enabled == nil {
			if must {
				return fmt.Errorf("invalid enabled")
			}
			return nil
		}

		h.Enabled = enabled
		return nil
	}
}

func WithIndex(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid index")
			}
			return nil
		}
		h.Index = u
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &categorycrud.Conds{}
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
		if conds.ParentID != nil {
			id, err := uuid.Parse(conds.GetParentID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ParentID = &cruder.Cond{Op: conds.GetParentID().GetOp(), Val: id}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{Op: conds.GetName().GetOp(), Val: conds.GetName().GetValue()}
		}
		if conds.Slug != nil {
			h.Conds.Slug = &cruder.Cond{Op: conds.GetSlug().GetOp(), Val: conds.GetSlug().GetValue()}
		}
		if conds.Enabled != nil {
			h.Conds.Enabled = &cruder.Cond{Op: conds.GetEnabled().GetOp(), Val: conds.GetEnabled().GetValue()}
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
