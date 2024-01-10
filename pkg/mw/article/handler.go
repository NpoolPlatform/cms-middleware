package article

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/cms-middleware/pkg/const"
	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	CategoryID  *uuid.UUID
	AuthorID    *uuid.UUID
	ArticleKey  *uuid.UUID
	Title       *string
	Subtitle    *string
	Digest      *string
	Status      *basetypes.ArticleStatus
	Version     *uint32
	Latest      *bool
	ContentURL  *string
	Host        *string
	ISO         *string
	PublishedAt *uint32
	ACLEnabled  *bool
	Conds       *articlecrud.Conds
	Offset      int32
	Limit       int32
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

func WithCategoryID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid categoryid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CategoryID = &_id
		return nil
	}
}

func WithAuthorID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid authorid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AuthorID = &_id
		return nil
	}
}

func WithArticleKey(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid articlekey")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ArticleKey = &_id
		return nil
	}
}

func WithTitle(title *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if title == nil {
			if must {
				return fmt.Errorf("invalid title")
			}
			return nil
		}
		if *title == "" {
			return fmt.Errorf("invalid title")
		}
		h.Title = title
		return nil
	}
}

func WithSubtitle(subtitle *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if subtitle == nil {
			if must {
				return fmt.Errorf("invalid subtitle")
			}
			return nil
		}
		if *subtitle == "" {
			return fmt.Errorf("invalid subtitle")
		}
		h.Subtitle = subtitle
		return nil
	}
}

func WithDigest(digest *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if digest == nil {
			if must {
				return fmt.Errorf("invalid digest")
			}
			return nil
		}
		if *digest == "" {
			return fmt.Errorf("invalid digest")
		}
		h.Digest = digest
		return nil
	}
}

func WithStatus(e *basetypes.ArticleStatus, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid status")
			}
			return nil
		}
		switch *e {
		case basetypes.ArticleStatus_Draft:
		case basetypes.ArticleStatus_Published:
		default:
			return fmt.Errorf("invalid cancelmode")
		}
		h.Status = e
		return nil
	}
}

func WithVersion(version *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if version == nil {
			if must {
				return fmt.Errorf("invalid version")
			}
			return nil
		}

		h.Version = version
		return nil
	}
}

func WithLatest(latest *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if latest == nil {
			if must {
				return fmt.Errorf("invalid latest")
			}
			return nil
		}

		h.Latest = latest
		return nil
	}
}

func WithACLEnabled(aclenabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if aclenabled == nil {
			if must {
				return fmt.Errorf("invalid aclenabled")
			}
			return nil
		}

		h.ACLEnabled = aclenabled
		return nil
	}
}

func WithContentURL(url *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if url == nil {
			if must {
				return fmt.Errorf("invalid contenturl")
			}
			return nil
		}
		if *url == "" {
			return fmt.Errorf("invalid contenturl")
		}
		h.ContentURL = url
		return nil
	}
}

func WithHost(host *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if host == nil {
			if must {
				return fmt.Errorf("invalid host")
			}
			return nil
		}
		if *host == "" {
			return fmt.Errorf("invalid host")
		}
		h.Host = host
		return nil
	}
}

func WithISO(iso *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if iso == nil {
			if must {
				return fmt.Errorf("invalid iso")
			}
			return nil
		}
		if *iso == "" {
			return fmt.Errorf("invalid iso")
		}
		h.ISO = iso
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &articlecrud.Conds{}
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
		if conds.CategoryID != nil {
			id, err := uuid.Parse(conds.GetCategoryID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CategoryID = &cruder.Cond{Op: conds.GetCategoryID().GetOp(), Val: id}
		}
		if conds.AuthorID != nil {
			id, err := uuid.Parse(conds.GetAuthorID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AuthorID = &cruder.Cond{Op: conds.GetAuthorID().GetOp(), Val: id}
		}
		if conds.ArticleKey != nil {
			id, err := uuid.Parse(conds.GetArticleKey().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ArticleKey = &cruder.Cond{Op: conds.GetArticleKey().GetOp(), Val: id}
		}
		if conds.Title != nil {
			h.Conds.Title = &cruder.Cond{Op: conds.GetTitle().GetOp(), Val: conds.GetTitle().GetValue()}
		}
		if conds.Subtitle != nil {
			h.Conds.Subtitle = &cruder.Cond{Op: conds.GetSubtitle().GetOp(), Val: conds.GetSubtitle().GetValue()}
		}
		if conds.Digest != nil {
			h.Conds.Digest = &cruder.Cond{Op: conds.GetDigest().GetOp(), Val: conds.GetDigest().GetValue()}
		}
		if conds.Status != nil {
			h.Conds.Status = &cruder.Cond{Op: conds.GetStatus().GetOp(), Val: basetypes.ArticleStatus(conds.GetStatus().GetValue())}
		}
		if conds.Version != nil {
			h.Conds.Version = &cruder.Cond{Op: conds.GetVersion().GetOp(), Val: conds.GetVersion().GetValue()}
		}
		if conds.Host != nil {
			h.Conds.Host = &cruder.Cond{Op: conds.GetHost().GetOp(), Val: conds.GetHost().GetValue()}
		}
		if conds.ISO != nil {
			h.Conds.ISO = &cruder.Cond{Op: conds.GetISO().GetOp(), Val: conds.GetISO().GetValue()}
		}
		if conds.Latest != nil {
			h.Conds.Latest = &cruder.Cond{Op: conds.GetLatest().GetOp(), Val: conds.GetLatest().GetValue()}
		}
		if conds.ContentURL != nil {
			h.Conds.ContentURL = &cruder.Cond{Op: conds.GetContentURL().GetOp(), Val: conds.GetContentURL().GetValue()}
		}
		if conds.ACLEnabled != nil {
			h.Conds.ACLEnabled = &cruder.Cond{Op: conds.GetACLEnabled().GetOp(), Val: conds.GetACLEnabled().GetValue()}
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
