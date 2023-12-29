// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/article"
	"github.com/google/uuid"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID uuid.UUID `json:"category_id,omitempty"`
	// AuthorID holds the value of the "author_id" field.
	AuthorID uuid.UUID `json:"author_id,omitempty"`
	// ArticleKey holds the value of the "article_key" field.
	ArticleKey uuid.UUID `json:"article_key,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Subtitle holds the value of the "subtitle" field.
	Subtitle string `json:"subtitle,omitempty"`
	// Digest holds the value of the "digest" field.
	Digest string `json:"digest,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Host holds the value of the "host" field.
	Host string `json:"host,omitempty"`
	// Version holds the value of the "version" field.
	Version uint32 `json:"version,omitempty"`
	// ContentURL holds the value of the "content_url" field.
	ContentURL string `json:"content_url,omitempty"`
	// Latest holds the value of the "latest" field.
	Latest bool `json:"latest,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt uint32 `json:"published_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldLatest:
			values[i] = new(sql.NullBool)
		case article.FieldID, article.FieldCreatedAt, article.FieldUpdatedAt, article.FieldDeletedAt, article.FieldVersion, article.FieldPublishedAt:
			values[i] = new(sql.NullInt64)
		case article.FieldTitle, article.FieldSubtitle, article.FieldDigest, article.FieldStatus, article.FieldHost, article.FieldContentURL:
			values[i] = new(sql.NullString)
		case article.FieldEntID, article.FieldAppID, article.FieldCategoryID, article.FieldAuthorID, article.FieldArticleKey:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint32(value.Int64)
		case article.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = uint32(value.Int64)
			}
		case article.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = uint32(value.Int64)
			}
		case article.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = uint32(value.Int64)
			}
		case article.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				a.EntID = *value
			}
		case article.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				a.AppID = *value
			}
		case article.FieldCategoryID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value != nil {
				a.CategoryID = *value
			}
		case article.FieldAuthorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value != nil {
				a.AuthorID = *value
			}
		case article.FieldArticleKey:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field article_key", values[i])
			} else if value != nil {
				a.ArticleKey = *value
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldSubtitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subtitle", values[i])
			} else if value.Valid {
				a.Subtitle = value.String
			}
		case article.FieldDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field digest", values[i])
			} else if value.Valid {
				a.Digest = value.String
			}
		case article.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = value.String
			}
		case article.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				a.Host = value.String
			}
		case article.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				a.Version = uint32(value.Int64)
			}
		case article.FieldContentURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_url", values[i])
			} else if value.Valid {
				a.ContentURL = value.String
			}
		case article.FieldLatest:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field latest", values[i])
			} else if value.Valid {
				a.Latest = value.Bool
			}
		case article.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				a.PublishedAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", a.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AppID))
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("author_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AuthorID))
	builder.WriteString(", ")
	builder.WriteString("article_key=")
	builder.WriteString(fmt.Sprintf("%v", a.ArticleKey))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("subtitle=")
	builder.WriteString(a.Subtitle)
	builder.WriteString(", ")
	builder.WriteString("digest=")
	builder.WriteString(a.Digest)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(a.Status)
	builder.WriteString(", ")
	builder.WriteString("host=")
	builder.WriteString(a.Host)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", a.Version))
	builder.WriteString(", ")
	builder.WriteString("content_url=")
	builder.WriteString(a.ContentURL)
	builder.WriteString(", ")
	builder.WriteString("latest=")
	builder.WriteString(fmt.Sprintf("%v", a.Latest))
	builder.WriteString(", ")
	builder.WriteString("published_at=")
	builder.WriteString(fmt.Sprintf("%v", a.PublishedAt))
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}