package article

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entarticle "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/article"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) ExistArticle(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Article.
			Query().
			Where(
				entarticle.EntID(*h.EntID),
				entarticle.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistArticleConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := articlecrud.SetQueryConds(cli.Article.Query(), h.Conds)
		if err != nil {
			return err
		}
		var condErr error
		exist, err = stm.Modify(func(s *sql.Selector) {
			if h.Conds != nil && h.Conds.Title != nil {
				title, ok := h.Conds.Title.Val.(string)
				if !ok {
					return
				}
				switch h.Conds.Title.Op {
				case cruder.EQ:
					s.Where(
						sql.EQ(sql.Lower(entarticle.FieldTitle), strings.ToLower(title)),
					)
				default:
					condErr = fmt.Errorf("invalid title field")
				}
			}
		}).Exist(_ctx)
		if err != nil {
			return err
		}
		if condErr != nil {
			return condErr
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
