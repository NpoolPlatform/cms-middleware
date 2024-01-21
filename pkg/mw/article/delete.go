package article

import (
	"context"
	"fmt"
	"time"

	aclcrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/acl"
	articlecrud "github.com/NpoolPlatform/cms-middleware/pkg/crud/article"
	"github.com/NpoolPlatform/cms-middleware/pkg/db"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent"
	entacl "github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
	"github.com/google/uuid"
)

func (h *Handler) DeleteArticle(ctx context.Context) (*npool.Article, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetArticle(ctx)
	if err != nil {
		return nil, err
	}
	articleKey := uuid.MustParse(info.ArticleKey)

	now := uint32(time.Now().Unix())
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := articlecrud.UpdateSet(
			tx.Article.UpdateOneID(*h.ID),
			&articlecrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		aclInfos, err := tx.
			ACL.
			Query().
			Where(
				entacl.ArticleKey(articleKey),
				entacl.DeletedAt(0),
			).
			ForUpdate().
			All(ctx)
		if err != nil {
			return err
		}
		if aclInfos == nil {
			return nil
		}

		for _, info := range aclInfos {
			if _, err := aclcrud.UpdateSet(
				tx.ACL.UpdateOneID(info.ID),
				&aclcrud.Req{
					DeletedAt: &now,
				},
			).Save(ctx); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
