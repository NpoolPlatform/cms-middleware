package article

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	category1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cms-middleware/pkg/testinit"
	cmstypes "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	slug = uuid.NewString()
	ret  = npool.Article{
		EntID:        uuid.NewString(),
		AppID:        uuid.NewString(),
		CategoryID:   uuid.NewString(),
		CategoryName: uuid.NewString(),
		AuthorID:     uuid.NewString(),
		ArticleKey:   uuid.NewString(),
		Title:        uuid.NewString(),
		Subtitle:     uuid.NewString(),
		Digest:       uuid.NewString(),
		Host:         "api.site.top",
		ISO:          "en-US",
		Status:       cmstypes.ArticleStatus_Draft,
		StatusStr:    cmstypes.ArticleStatus_Draft.String(),
		Version:      1,
		Latest:       true,
		ContentURL:   uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := category1.NewHandler(
		context.Background(),
		category1.WithEntID(&ret.CategoryID, true),
		category1.WithAppID(&ret.AppID, true),
		category1.WithName(&ret.CategoryName, true),
		category1.WithSlug(&slug, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateCategory(context.Background())
	assert.Nil(t, err)
	h1.ID = &info1.ID

	return func(*testing.T) {
		_, _ = h1.DeleteCategory(context.Background())
	}
}

func createArticle(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCategoryID(&ret.CategoryID, true),
		WithAuthorID(&ret.AuthorID, true),
		WithArticleKey(&ret.ArticleKey, true),
		WithTitle(&ret.Title, true),
		WithSubtitle(&ret.Subtitle, true),
		WithDigest(&ret.Digest, true),
		WithContentURL(&ret.ContentURL, true),
		WithHost(&ret.Host, true),
		WithISO(&ret.ISO, true),
		WithVersion(&ret.Version, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateArticle(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.PublishedAt = info.PublishedAt
		assert.Equal(t, info, &ret)
	}
}

func updateArticle(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithCategoryID(&ret.CategoryID, true),
		WithTitle(&ret.Title, true),
		WithSubtitle(&ret.Subtitle, true),
		WithDigest(&ret.Digest, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateArticle(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func publishArticle(t *testing.T) {
	ret.Status = cmstypes.ArticleStatus_Published
	ret.StatusStr = cmstypes.ArticleStatus_Published.String()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithStatus(&ret.Status, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateArticle(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getArticle(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetArticle(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getArticles(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CategoryID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CategoryID},
		AuthorID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AuthorID},
		ArticleKey: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ArticleKey},
		Title:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.Title},
		Subtitle:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.Subtitle},
		Digest:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Digest},
		Status:     &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.Status)},
		Version:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Version},
		Latest:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Latest},
		ContentURL: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ContentURL},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetArticles(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteArticle(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteArticle(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetArticle(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestArticle(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createArticle", createArticle)
	t.Run("updateArticle", updateArticle)
	t.Run("publishArticle", publishArticle)
	t.Run("getArticle", getArticle)
	t.Run("getArticles", getArticles)
	t.Run("deleteArticle", deleteArticle)
}
