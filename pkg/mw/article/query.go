package article

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entacl "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
	entarticle "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/article"
	entcategory "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/category"
	types "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.ArticleSelect
	stmCount  *ent.ArticleSelect
	infos     []*npool.Article
	total     uint32
}

func (h *queryHandler) selectArticle(stm *ent.ArticleQuery) *ent.ArticleSelect {
	return stm.Select(entarticle.FieldID)
}

func (h *queryHandler) queryArticle(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Article.Query().Where(entarticle.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entarticle.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entarticle.EntID(*h.EntID))
	}
	h.stmSelect = h.selectArticle(stm)
	return nil
}

func (h *queryHandler) queryArticles(cli *ent.Client) (*ent.ArticleSelect, error) {
	stm, err := articlecrud.SetQueryConds(cli.Article.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectArticle(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entarticle.Table)
	s.LeftJoin(t).
		On(
			s.C(entarticle.FieldID),
			t.C(entarticle.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entarticle.FieldEntID), "ent_id"),
			sql.As(t.C(entarticle.FieldAppID), "app_id"),
			sql.As(t.C(entarticle.FieldCategoryID), "category_id"),
			sql.As(t.C(entarticle.FieldAuthorID), "author_id"),
			sql.As(t.C(entarticle.FieldArticleKey), "article_key"),
			sql.As(t.C(entarticle.FieldTitle), "title"),
			sql.As(t.C(entarticle.FieldSubtitle), "subtitle"),
			sql.As(t.C(entarticle.FieldDigest), "digest"),
			sql.As(t.C(entarticle.FieldStatus), "status"),
			sql.As(t.C(entarticle.FieldVersion), "version"),
			sql.As(t.C(entarticle.FieldLatest), "latest"),
			sql.As(t.C(entarticle.FieldHost), "host"),
			sql.As(t.C(entarticle.FieldIso), "iso"),
			sql.As(t.C(entarticle.FieldContentURL), "content_url"),
			sql.As(t.C(entarticle.FieldPublishedAt), "published_at"),
			sql.As(t.C(entarticle.FieldACLEnabled), "acl_enabled"),
		)
}

func (h *queryHandler) queryJoinCategory(s *sql.Selector) {
	t := sql.Table(entcategory.Table)
	s.LeftJoin(t).
		On(
			s.C(entarticle.FieldCategoryID),
			t.C(entcategory.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entcategory.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entcategory.FieldName), "category_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCategory(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCategory(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Status = types.ArticleStatus(types.ArticleStatus_value[info.StatusStr])
	}
}
func (h *queryHandler) queryACLRoles(ctx context.Context) error {
	if len(h.infos) == 0 {
		return nil
	}

	type acl struct {
		ArticleKey uuid.UUID `json:"article_key"`
		RoleID     uuid.UUID `json:"role_id"`
	}

	acls := []*acl{}
	akeys := []uuid.UUID{}

	for _, info := range h.infos {
		akeys = append(akeys, uuid.MustParse(info.ArticleKey))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return cli.
			ACL.
			Query().
			Where(
				entacl.ArticleKeyIn(akeys...),
				entacl.DeletedAt(0),
			).
			Select(
				entacl.FieldRoleID,
				entacl.FieldArticleKey,
			).
			Scan(_ctx, &acls)
	})
	if err != nil {
		return err
	}

	for _, info := range h.infos {
		for _, acl := range acls {
			if info.ArticleKey == acl.ArticleKey.String() {
				info.ACLRoleIDs = append(info.ACLRoleIDs, acl.RoleID.String())
			}
		}
	}

	return nil
}

func (h *Handler) GetArticle(ctx context.Context) (*npool.Article, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryArticle(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	if err := handler.queryACLRoles(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetArticles(ctx context.Context) ([]*npool.Article, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryArticles(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryArticles(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	if err := handler.queryACLRoles(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
