// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/article"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/category"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/categorylang"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/media"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 5)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   acl.Table,
			Columns: acl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: acl.FieldID,
			},
		},
		Type: "ACL",
		Fields: map[string]*sqlgraph.FieldSpec{
			acl.FieldCreatedAt:  {Type: field.TypeUint32, Column: acl.FieldCreatedAt},
			acl.FieldUpdatedAt:  {Type: field.TypeUint32, Column: acl.FieldUpdatedAt},
			acl.FieldDeletedAt:  {Type: field.TypeUint32, Column: acl.FieldDeletedAt},
			acl.FieldEntID:      {Type: field.TypeUUID, Column: acl.FieldEntID},
			acl.FieldAppID:      {Type: field.TypeUUID, Column: acl.FieldAppID},
			acl.FieldRoleID:     {Type: field.TypeUUID, Column: acl.FieldRoleID},
			acl.FieldArticleKey: {Type: field.TypeUUID, Column: acl.FieldArticleKey},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   article.Table,
			Columns: article.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: article.FieldID,
			},
		},
		Type: "Article",
		Fields: map[string]*sqlgraph.FieldSpec{
			article.FieldCreatedAt:   {Type: field.TypeUint32, Column: article.FieldCreatedAt},
			article.FieldUpdatedAt:   {Type: field.TypeUint32, Column: article.FieldUpdatedAt},
			article.FieldDeletedAt:   {Type: field.TypeUint32, Column: article.FieldDeletedAt},
			article.FieldEntID:       {Type: field.TypeUUID, Column: article.FieldEntID},
			article.FieldAppID:       {Type: field.TypeUUID, Column: article.FieldAppID},
			article.FieldCategoryID:  {Type: field.TypeUUID, Column: article.FieldCategoryID},
			article.FieldAuthorID:    {Type: field.TypeUUID, Column: article.FieldAuthorID},
			article.FieldArticleKey:  {Type: field.TypeUUID, Column: article.FieldArticleKey},
			article.FieldTitle:       {Type: field.TypeString, Column: article.FieldTitle},
			article.FieldSubtitle:    {Type: field.TypeString, Column: article.FieldSubtitle},
			article.FieldDigest:      {Type: field.TypeString, Column: article.FieldDigest},
			article.FieldStatus:      {Type: field.TypeString, Column: article.FieldStatus},
			article.FieldHost:        {Type: field.TypeString, Column: article.FieldHost},
			article.FieldVersion:     {Type: field.TypeUint32, Column: article.FieldVersion},
			article.FieldContentURL:  {Type: field.TypeString, Column: article.FieldContentURL},
			article.FieldLatest:      {Type: field.TypeBool, Column: article.FieldLatest},
			article.FieldPublishedAt: {Type: field.TypeUint32, Column: article.FieldPublishedAt},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: category.FieldID,
			},
		},
		Type: "Category",
		Fields: map[string]*sqlgraph.FieldSpec{
			category.FieldCreatedAt: {Type: field.TypeUint32, Column: category.FieldCreatedAt},
			category.FieldUpdatedAt: {Type: field.TypeUint32, Column: category.FieldUpdatedAt},
			category.FieldDeletedAt: {Type: field.TypeUint32, Column: category.FieldDeletedAt},
			category.FieldEntID:     {Type: field.TypeUUID, Column: category.FieldEntID},
			category.FieldAppID:     {Type: field.TypeUUID, Column: category.FieldAppID},
			category.FieldParentID:  {Type: field.TypeUUID, Column: category.FieldParentID},
			category.FieldName:      {Type: field.TypeString, Column: category.FieldName},
			category.FieldSlug:      {Type: field.TypeString, Column: category.FieldSlug},
			category.FieldEnabled:   {Type: field.TypeBool, Column: category.FieldEnabled},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   categorylang.Table,
			Columns: categorylang.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: categorylang.FieldID,
			},
		},
		Type: "CategoryLang",
		Fields: map[string]*sqlgraph.FieldSpec{
			categorylang.FieldCreatedAt:  {Type: field.TypeUint32, Column: categorylang.FieldCreatedAt},
			categorylang.FieldUpdatedAt:  {Type: field.TypeUint32, Column: categorylang.FieldUpdatedAt},
			categorylang.FieldDeletedAt:  {Type: field.TypeUint32, Column: categorylang.FieldDeletedAt},
			categorylang.FieldEntID:      {Type: field.TypeUUID, Column: categorylang.FieldEntID},
			categorylang.FieldAppID:      {Type: field.TypeUUID, Column: categorylang.FieldAppID},
			categorylang.FieldLangID:     {Type: field.TypeUUID, Column: categorylang.FieldLangID},
			categorylang.FieldCategoryID: {Type: field.TypeUUID, Column: categorylang.FieldCategoryID},
			categorylang.FieldDisplay:    {Type: field.TypeString, Column: categorylang.FieldDisplay},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   media.Table,
			Columns: media.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: media.FieldID,
			},
		},
		Type: "Media",
		Fields: map[string]*sqlgraph.FieldSpec{
			media.FieldCreatedAt: {Type: field.TypeUint32, Column: media.FieldCreatedAt},
			media.FieldUpdatedAt: {Type: field.TypeUint32, Column: media.FieldUpdatedAt},
			media.FieldDeletedAt: {Type: field.TypeUint32, Column: media.FieldDeletedAt},
			media.FieldEntID:     {Type: field.TypeUUID, Column: media.FieldEntID},
			media.FieldAppID:     {Type: field.TypeUUID, Column: media.FieldAppID},
			media.FieldName:      {Type: field.TypeString, Column: media.FieldName},
			media.FieldExt:       {Type: field.TypeString, Column: media.FieldExt},
			media.FieldMediaURL:  {Type: field.TypeString, Column: media.FieldMediaURL},
			media.FieldCreatedBy: {Type: field.TypeUUID, Column: media.FieldCreatedBy},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *ACLQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ACLQuery builder.
func (aq *ACLQuery) Filter() *ACLFilter {
	return &ACLFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *ACLMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ACLMutation builder.
func (m *ACLMutation) Filter() *ACLFilter {
	return &ACLFilter{config: m.config, predicateAdder: m}
}

// ACLFilter provides a generic filtering capability at runtime for ACLQuery.
type ACLFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ACLFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *ACLFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(acl.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ACLFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(acl.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ACLFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(acl.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ACLFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(acl.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *ACLFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(acl.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ACLFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(acl.FieldAppID))
}

// WhereRoleID applies the entql [16]byte predicate on the role_id field.
func (f *ACLFilter) WhereRoleID(p entql.ValueP) {
	f.Where(p.Field(acl.FieldRoleID))
}

// WhereArticleKey applies the entql [16]byte predicate on the article_key field.
func (f *ACLFilter) WhereArticleKey(p entql.ValueP) {
	f.Where(p.Field(acl.FieldArticleKey))
}

// addPredicate implements the predicateAdder interface.
func (aq *ArticleQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ArticleQuery builder.
func (aq *ArticleQuery) Filter() *ArticleFilter {
	return &ArticleFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *ArticleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ArticleMutation builder.
func (m *ArticleMutation) Filter() *ArticleFilter {
	return &ArticleFilter{config: m.config, predicateAdder: m}
}

// ArticleFilter provides a generic filtering capability at runtime for ArticleQuery.
type ArticleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ArticleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *ArticleFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(article.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ArticleFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(article.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ArticleFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(article.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ArticleFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(article.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *ArticleFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(article.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ArticleFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(article.FieldAppID))
}

// WhereCategoryID applies the entql [16]byte predicate on the category_id field.
func (f *ArticleFilter) WhereCategoryID(p entql.ValueP) {
	f.Where(p.Field(article.FieldCategoryID))
}

// WhereAuthorID applies the entql [16]byte predicate on the author_id field.
func (f *ArticleFilter) WhereAuthorID(p entql.ValueP) {
	f.Where(p.Field(article.FieldAuthorID))
}

// WhereArticleKey applies the entql [16]byte predicate on the article_key field.
func (f *ArticleFilter) WhereArticleKey(p entql.ValueP) {
	f.Where(p.Field(article.FieldArticleKey))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *ArticleFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(article.FieldTitle))
}

// WhereSubtitle applies the entql string predicate on the subtitle field.
func (f *ArticleFilter) WhereSubtitle(p entql.StringP) {
	f.Where(p.Field(article.FieldSubtitle))
}

// WhereDigest applies the entql string predicate on the digest field.
func (f *ArticleFilter) WhereDigest(p entql.StringP) {
	f.Where(p.Field(article.FieldDigest))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *ArticleFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(article.FieldStatus))
}

// WhereHost applies the entql string predicate on the host field.
func (f *ArticleFilter) WhereHost(p entql.StringP) {
	f.Where(p.Field(article.FieldHost))
}

// WhereVersion applies the entql uint32 predicate on the version field.
func (f *ArticleFilter) WhereVersion(p entql.Uint32P) {
	f.Where(p.Field(article.FieldVersion))
}

// WhereContentURL applies the entql string predicate on the content_url field.
func (f *ArticleFilter) WhereContentURL(p entql.StringP) {
	f.Where(p.Field(article.FieldContentURL))
}

// WhereLatest applies the entql bool predicate on the latest field.
func (f *ArticleFilter) WhereLatest(p entql.BoolP) {
	f.Where(p.Field(article.FieldLatest))
}

// WherePublishedAt applies the entql uint32 predicate on the published_at field.
func (f *ArticleFilter) WherePublishedAt(p entql.Uint32P) {
	f.Where(p.Field(article.FieldPublishedAt))
}

// addPredicate implements the predicateAdder interface.
func (cq *CategoryQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CategoryQuery builder.
func (cq *CategoryQuery) Filter() *CategoryFilter {
	return &CategoryFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CategoryMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CategoryMutation builder.
func (m *CategoryMutation) Filter() *CategoryFilter {
	return &CategoryFilter{config: m.config, predicateAdder: m}
}

// CategoryFilter provides a generic filtering capability at runtime for CategoryQuery.
type CategoryFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CategoryFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *CategoryFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(category.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CategoryFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(category.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CategoryFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(category.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CategoryFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(category.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *CategoryFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(category.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CategoryFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(category.FieldAppID))
}

// WhereParentID applies the entql [16]byte predicate on the parent_id field.
func (f *CategoryFilter) WhereParentID(p entql.ValueP) {
	f.Where(p.Field(category.FieldParentID))
}

// WhereName applies the entql string predicate on the name field.
func (f *CategoryFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(category.FieldName))
}

// WhereSlug applies the entql string predicate on the slug field.
func (f *CategoryFilter) WhereSlug(p entql.StringP) {
	f.Where(p.Field(category.FieldSlug))
}

// WhereEnabled applies the entql bool predicate on the enabled field.
func (f *CategoryFilter) WhereEnabled(p entql.BoolP) {
	f.Where(p.Field(category.FieldEnabled))
}

// addPredicate implements the predicateAdder interface.
func (clq *CategoryLangQuery) addPredicate(pred func(s *sql.Selector)) {
	clq.predicates = append(clq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CategoryLangQuery builder.
func (clq *CategoryLangQuery) Filter() *CategoryLangFilter {
	return &CategoryLangFilter{config: clq.config, predicateAdder: clq}
}

// addPredicate implements the predicateAdder interface.
func (m *CategoryLangMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CategoryLangMutation builder.
func (m *CategoryLangMutation) Filter() *CategoryLangFilter {
	return &CategoryLangFilter{config: m.config, predicateAdder: m}
}

// CategoryLangFilter provides a generic filtering capability at runtime for CategoryLangQuery.
type CategoryLangFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CategoryLangFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *CategoryLangFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(categorylang.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CategoryLangFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(categorylang.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CategoryLangFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(categorylang.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CategoryLangFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(categorylang.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *CategoryLangFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(categorylang.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *CategoryLangFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(categorylang.FieldAppID))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *CategoryLangFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(categorylang.FieldLangID))
}

// WhereCategoryID applies the entql [16]byte predicate on the category_id field.
func (f *CategoryLangFilter) WhereCategoryID(p entql.ValueP) {
	f.Where(p.Field(categorylang.FieldCategoryID))
}

// WhereDisplay applies the entql string predicate on the display field.
func (f *CategoryLangFilter) WhereDisplay(p entql.StringP) {
	f.Where(p.Field(categorylang.FieldDisplay))
}

// addPredicate implements the predicateAdder interface.
func (mq *MediaQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MediaQuery builder.
func (mq *MediaQuery) Filter() *MediaFilter {
	return &MediaFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MediaMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MediaMutation builder.
func (m *MediaMutation) Filter() *MediaFilter {
	return &MediaFilter{config: m.config, predicateAdder: m}
}

// MediaFilter provides a generic filtering capability at runtime for MediaQuery.
type MediaFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MediaFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *MediaFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(media.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *MediaFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(media.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *MediaFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(media.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *MediaFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(media.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *MediaFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(media.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *MediaFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(media.FieldAppID))
}

// WhereName applies the entql string predicate on the name field.
func (f *MediaFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(media.FieldName))
}

// WhereExt applies the entql string predicate on the ext field.
func (f *MediaFilter) WhereExt(p entql.StringP) {
	f.Where(p.Field(media.FieldExt))
}

// WhereMediaURL applies the entql string predicate on the media_url field.
func (f *MediaFilter) WhereMediaURL(p entql.StringP) {
	f.Where(p.Field(media.FieldMediaURL))
}

// WhereCreatedBy applies the entql [16]byte predicate on the created_by field.
func (f *MediaFilter) WhereCreatedBy(p entql.ValueP) {
	f.Where(p.Field(media.FieldCreatedBy))
}
