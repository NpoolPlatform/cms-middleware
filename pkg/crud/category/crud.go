package category

import (
	"fmt"

	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entcategory "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/category"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	ParentID  *uuid.UUID
	Name      *string
	Slug      *string
	Enabled   *bool
	Index     *uint32
	DeletedAt *uint32
}

func CreateSet(c *ent.CategoryCreate, req *Req) *ent.CategoryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.ParentID != nil {
		c.SetParentID(*req.ParentID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Slug != nil {
		c.SetSlug(*req.Slug)
	}
	if req.Enabled != nil {
		c.SetEnabled(*req.Enabled)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}

	return c
}

func UpdateSet(u *ent.CategoryUpdateOne, req *Req) *ent.CategoryUpdateOne {
	if req.ParentID != nil {
		u.SetParentID(*req.ParentID)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Enabled != nil {
		u.SetEnabled(*req.Enabled)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	if req.Slug != nil {
		u.SetSlug(*req.Slug)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	EntID    *cruder.Cond
	AppID    *cruder.Cond
	ParentID *cruder.Cond
	Name     *cruder.Cond
	Enabled  *cruder.Cond
	Slug     *cruder.Cond
	IDs      *cruder.Cond
	EntIDs   *cruder.Cond
	Index    *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CategoryQuery, conds *Conds) (*ent.CategoryQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(
				entcategory.ID(id),
				entcategory.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entcategory.IDNEQ(id),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entcategory.EntID(id),
				entcategory.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entcategory.EntIDNEQ(id),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(
				entcategory.EntIDIn(ids...),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcategory.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.ParentID != nil {
		id, ok := conds.ParentID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid parentid")
		}
		switch conds.ParentID.Op {
		case cruder.EQ:
			q.Where(
				entcategory.ParentID(id),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid parentid field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(
				entcategory.Name(name),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.Slug != nil {
		slug, ok := conds.Slug.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid slug")
		}
		switch conds.Slug.Op {
		case cruder.EQ:
			q.Where(
				entcategory.Slug(slug),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid slug field")
		}
	}
	if conds.Enabled != nil {
		enabled, ok := conds.Enabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid enabled")
		}
		switch conds.Enabled.Op {
		case cruder.EQ:
			q.Where(
				entcategory.Enabled(enabled),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid enabled field")
		}
	}
	if conds.Index != nil {
		index, ok := conds.Index.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid index")
		}
		switch conds.Index.Op {
		case cruder.EQ:
			q.Where(
				entcategory.Index(index),
				entcategory.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid index field")
		}
	}

	return q, nil
}
