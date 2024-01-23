package acl

import (
	"fmt"

	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entacl "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	RoleID     *uuid.UUID
	ArticleKey *uuid.UUID
	DeletedAt  *uint32
}

func CreateSet(c *ent.ACLCreate, req *Req) *ent.ACLCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.RoleID != nil {
		c.SetRoleID(*req.RoleID)
	}
	if req.ArticleKey != nil {
		c.SetArticleKey(*req.ArticleKey)
	}

	return c
}

func UpdateSet(u *ent.ACLUpdateOne, req *Req) *ent.ACLUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	RoleID     *cruder.Cond
	ArticleKey *cruder.Cond
	IDs        *cruder.Cond
	EntIDs     *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.ACLQuery, conds *Conds) (*ent.ACLQuery, error) {
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
				entacl.ID(id),
				entacl.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entacl.IDNEQ(id),
				entacl.DeletedAt(0),
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
				entacl.EntID(id),
				entacl.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entacl.EntIDNEQ(id),
				entacl.DeletedAt(0),
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
				entacl.EntIDIn(ids...),
				entacl.DeletedAt(0),
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
			q.Where(entacl.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.RoleID != nil {
		roleid, ok := conds.RoleID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid roleid")
		}
		switch conds.RoleID.Op {
		case cruder.EQ:
			q.Where(
				entacl.RoleID(roleid),
				entacl.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid roleid field")
		}
	}
	if conds.ArticleKey != nil {
		articlekey, ok := conds.ArticleKey.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid articlekey")
		}
		switch conds.ArticleKey.Op {
		case cruder.EQ:
			q.Where(
				entacl.ArticleKey(articlekey),
				entacl.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid articlekey field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entacl.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid id filed")
		}
	}

	return q, nil
}
