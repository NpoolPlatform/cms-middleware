package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	"github.com/google/uuid"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("category_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("author_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("article_key", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			String("subtitle").
			Optional().
			Default(""),
		field.
			Text("digest").
			Optional().
			Default(""),
		field.
			String("status").
			Optional().
			Default(types.ArticleStatus_Draft.String()),
		field.
			Text("host").
			Optional().
			Default(""),
		field.
			Uint32("version").
			Optional().
			Default(1),
		field.
			String("iso").
			Optional().
			Default(""),
		field.
			Text("content_url").
			Optional().
			Default(""),
		field.
			Bool("latest").
			Optional().
			Default(true),
		field.
			Uint32("published_at").
			Optional().
			Default(0),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
