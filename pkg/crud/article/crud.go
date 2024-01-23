package article

import (
	"fmt"

	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entarticle "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/article"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	"github.com/google/uuid"
)

type Req struct {
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
	Host        *string
	ISO         *string
	ContentURL  *string
	ACLEnabled  *bool
	PublishedAt *uint32
	DeletedAt   *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.ArticleCreate, req *Req) *ent.ArticleCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CategoryID != nil {
		c.SetCategoryID(*req.CategoryID)
	}
	if req.AuthorID != nil {
		c.SetAuthorID(*req.AuthorID)
	}
	if req.ArticleKey != nil {
		c.SetArticleKey(*req.ArticleKey)
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Subtitle != nil {
		c.SetSubtitle(*req.Subtitle)
	}
	if req.Digest != nil {
		c.SetDigest(*req.Digest)
	}
	if req.Status != nil {
		c.SetStatus(req.Status.String())
	}
	if req.Version != nil {
		c.SetVersion(*req.Version)
	}
	if req.Latest != nil {
		c.SetLatest(*req.Latest)
	}
	if req.ContentURL != nil {
		c.SetContentURL(*req.ContentURL)
	}
	if req.Host != nil {
		c.SetHost(*req.Host)
	}
	if req.PublishedAt != nil {
		c.SetPublishedAt(*req.PublishedAt)
	}
	if req.ISO != nil {
		c.SetIso(*req.ISO)
	}
	if req.ACLEnabled != nil {
		c.SetACLEnabled(*req.ACLEnabled)
	}

	return c
}

func UpdateSet(u *ent.ArticleUpdateOne, req *Req) *ent.ArticleUpdateOne {
	if req.CategoryID != nil {
		u.SetCategoryID(*req.CategoryID)
	}
	if req.AuthorID != nil {
		u.SetAuthorID(*req.AuthorID)
	}
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Subtitle != nil {
		u.SetSubtitle(*req.Subtitle)
	}
	if req.Digest != nil {
		u.SetDigest(*req.Digest)
	}
	if req.Status != nil {
		u.SetStatus(req.Status.String())
	}
	if req.Latest != nil {
		u.SetLatest(*req.Latest)
	}
	if req.PublishedAt != nil {
		u.SetPublishedAt(*req.PublishedAt)
	}
	if req.ACLEnabled != nil {
		u.SetACLEnabled(*req.ACLEnabled)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	CategoryID *cruder.Cond
	AuthorID   *cruder.Cond
	ArticleKey *cruder.Cond
	Title      *cruder.Cond
	Subtitle   *cruder.Cond
	Digest     *cruder.Cond
	Status     *cruder.Cond
	Version    *cruder.Cond
	Host       *cruder.Cond
	Latest     *cruder.Cond
	ContentURL *cruder.Cond
	ISO        *cruder.Cond
	IDs        *cruder.Cond
	EntIDs     *cruder.Cond
	ACLEnabled *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.ArticleQuery, conds *Conds) (*ent.ArticleQuery, error) {
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
				entarticle.ID(id),
				entarticle.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entarticle.IDNEQ(id),
				entarticle.DeletedAt(0),
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
				entarticle.EntID(id),
				entarticle.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entarticle.EntIDNEQ(id),
				entarticle.DeletedAt(0),
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
				entarticle.EntIDIn(ids...),
				entarticle.DeletedAt(0),
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
			q.Where(entarticle.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.CategoryID != nil {
		id, ok := conds.CategoryID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid categoryid")
		}
		switch conds.CategoryID.Op {
		case cruder.EQ:
			q.Where(
				entarticle.CategoryID(id),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid categoryid field")
		}
	}
	if conds.AuthorID != nil {
		id, ok := conds.AuthorID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid authorid")
		}
		switch conds.AuthorID.Op {
		case cruder.EQ:
			q.Where(
				entarticle.AuthorID(id),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid authorid field")
		}
	}
	if conds.ArticleKey != nil {
		id, ok := conds.ArticleKey.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid articlekey")
		}
		switch conds.ArticleKey.Op {
		case cruder.EQ:
			q.Where(
				entarticle.ArticleKey(id),
				entarticle.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entarticle.ArticleKeyNEQ(id),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid articlekey field")
		}
	}
	if conds.Title != nil {
		title, ok := conds.Title.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid title")
		}
		switch conds.Title.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Title(title),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid title field")
		}
	}
	if conds.Subtitle != nil {
		subtitle, ok := conds.Subtitle.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid subtitle")
		}
		switch conds.Subtitle.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Subtitle(subtitle),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid subtitle field")
		}
	}
	if conds.Digest != nil {
		digest, ok := conds.Digest.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid digest")
		}
		switch conds.Digest.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Digest(digest),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid digest field")
		}
	}

	if conds.Status != nil {
		status, ok := conds.Status.Val.(basetypes.ArticleStatus)
		if !ok {
			return nil, fmt.Errorf("invalid status")
		}
		switch conds.Status.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Status(status.String()),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid status field")
		}
	}

	if conds.Version != nil {
		version, ok := conds.Version.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid version")
		}
		switch conds.Version.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Version(version),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid version field")
		}
	}
	if conds.Latest != nil {
		latest, ok := conds.Latest.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid latest")
		}
		switch conds.Latest.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Latest(latest),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid latest field")
		}
	}
	if conds.ContentURL != nil {
		contenturl, ok := conds.ContentURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid contenturl")
		}
		switch conds.ContentURL.Op {
		case cruder.EQ:
			q.Where(
				entarticle.ContentURL(contenturl),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid contenturl field")
		}
	}
	if conds.Host != nil {
		host, ok := conds.Host.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid host")
		}
		switch conds.Host.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Host(host),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid host field")
		}
	}
	if conds.ISO != nil {
		iso, ok := conds.ISO.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid iso")
		}
		switch conds.ISO.Op {
		case cruder.EQ:
			q.Where(
				entarticle.Iso(iso),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid iso field")
		}
	}
	if conds.ACLEnabled != nil {
		aclenabled, ok := conds.ACLEnabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid aclenabled")
		}
		switch conds.ACLEnabled.Op {
		case cruder.EQ:
			q.Where(
				entarticle.ACLEnabled(aclenabled),
				entarticle.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid aclenabled field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entarticle.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid id filed")
		}
	}
	return q, nil
}
