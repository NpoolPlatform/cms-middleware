package media

import (
	"fmt"

	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entmedia "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/media"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	Name      *string
	Ext       *string
	MediaURL  *string
	CreatedBy *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.MediaCreate, req *Req) *ent.MediaCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Ext != nil {
		c.SetExt(*req.Ext)
	}
	if req.MediaURL != nil {
		c.SetMediaURL(*req.MediaURL)
	}
	if req.CreatedBy != nil {
		c.SetCreatedBy(*req.CreatedBy)
	}

	return c
}

func UpdateSet(u *ent.MediaUpdateOne, req *Req) *ent.MediaUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	EntID    *cruder.Cond
	AppID    *cruder.Cond
	ParentID *cruder.Cond
	Name     *cruder.Cond
	MediaURL *cruder.Cond
	IDs      *cruder.Cond
	EntIDs   *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.MediaQuery, conds *Conds) (*ent.MediaQuery, error) {
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
				entmedia.ID(id),
				entmedia.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entmedia.IDNEQ(id),
				entmedia.DeletedAt(0),
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
				entmedia.EntID(id),
				entmedia.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entmedia.EntIDNEQ(id),
				entmedia.DeletedAt(0),
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
				entmedia.EntIDIn(ids...),
				entmedia.DeletedAt(0),
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
			q.Where(entmedia.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
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
				entmedia.Name(name),
				entmedia.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.MediaURL != nil {
		url, ok := conds.MediaURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid mediaurl")
		}
		switch conds.MediaURL.Op {
		case cruder.EQ:
			q.Where(
				entmedia.MediaURL(url),
				entmedia.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid mediaurl field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entmedia.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid id filed")
		}
	}

	return q, nil
}
