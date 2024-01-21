package acl

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	article1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/article"
	category1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"
	cmstypes "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
	articlemw "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
	categorymw "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cms-middleware/pkg/testinit"
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
	articleRet = articlemw.Article{
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
		ACLEnabled:   true,
	}
	categoryRet = categorymw.Category{
		EntID:   articleRet.CategoryID,
		AppID:   articleRet.AppID,
		Name:    articleRet.CategoryName,
		Slug:    uuid.NewString(),
		Enabled: false,
		Index:   uint32(0),
	}
	ret = npool.ACL{
		EntID:      uuid.NewString(),
		AppID:      articleRet.AppID,
		RoleID:     uuid.NewString(),
		ArticleKey: articleRet.ArticleKey,
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := category1.NewHandler(
		context.Background(),
		category1.WithEntID(&categoryRet.EntID, true),
		category1.WithAppID(&categoryRet.AppID, true),
		category1.WithName(&categoryRet.Name, true),
		category1.WithSlug(&categoryRet.Slug, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateCategory(context.Background())
	assert.Nil(t, err)
	h1.ID = &info1.ID

	h2, err := article1.NewHandler(
		context.Background(),
		article1.WithEntID(&articleRet.EntID, true),
		article1.WithEntID(&articleRet.EntID, true),
		article1.WithAppID(&articleRet.AppID, true),
		article1.WithCategoryID(&articleRet.CategoryID, true),
		article1.WithAuthorID(&articleRet.AuthorID, true),
		article1.WithArticleKey(&articleRet.ArticleKey, true),
		article1.WithTitle(&articleRet.Title, true),
		article1.WithSubtitle(&articleRet.Subtitle, true),
		article1.WithDigest(&articleRet.Digest, true),
		article1.WithContentURL(&articleRet.ContentURL, true),
		article1.WithHost(&articleRet.Host, true),
		article1.WithISO(&articleRet.ISO, true),
		article1.WithVersion(&articleRet.Version, true),
		article1.WithACLEnabled(&articleRet.ACLEnabled, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateArticle(context.Background())
	assert.Nil(t, err)
	h2.ID = &info2.ID

	return func(*testing.T) {
		_, _ = h1.DeleteCategory(context.Background())
		_, _ = h2.DeleteArticle(context.Background())
	}
}

func createACL(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithRoleID(&ret.RoleID, true),
		WithArticleKey(&ret.ArticleKey, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateACL(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getACL(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetACL(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getACLs(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RoleID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.RoleID},
		ArticleKey: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ArticleKey},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetACLs(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteACL(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteACL(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetACL(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestACL(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createACL", createACL)
	t.Run("getACL", getACL)
	t.Run("getACLs", getACLs)
	t.Run("deleteACL", deleteACL)
}
