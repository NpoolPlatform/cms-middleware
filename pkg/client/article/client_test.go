package article

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"

	category1 "github.com/NpoolPlatform/cms-middleware/pkg/client/category"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	cmstypes "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	categorymwpb "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/cms-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.Article{
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

var req = &npool.ArticleReq{
	EntID:      &ret.EntID,
	AppID:      &ret.AppID,
	CategoryID: &ret.CategoryID,
	AuthorID:   &ret.AuthorID,
	ArticleKey: &ret.ArticleKey,
	Title:      &ret.Title,
	Subtitle:   &ret.Subtitle,
	Digest:     &ret.Digest,
	ContentURL: &ret.ContentURL,
	Host:       &ret.Host,
	ISO:        &ret.ISO,
	Version:    &ret.Version,
}

var slug = uuid.NewString()

func setup(t *testing.T) func(*testing.T) {
	info1, err := category1.CreateCategory(context.Background(), &categorymwpb.CategoryReq{
		EntID: &ret.CategoryID,
		AppID: &ret.AppID,
		Name:  &ret.CategoryName,
		Slug:  &slug,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info1)

	return func(*testing.T) {
		_, _ = category1.DeleteCategory(context.Background(), info1.ID)
	}
}

func createArticle(t *testing.T) {
	info, err := CreateArticle(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.PublishedAt = info.PublishedAt
		assert.Equal(t, ret, info)
	}
}

func updateArticle(t *testing.T) {
	ret.Title = uuid.NewString()
	info, err := UpdateArticle(context.Background(), &npool.ArticleReq{
		ID:    &ret.ID,
		AppID: &ret.AppID,
		Title: &ret.Title,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getArticle(t *testing.T) {
	info, err := GetArticle(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getArticles(t *testing.T) {
	infos, total, err := GetArticles(context.Background(), &npool.Conds{
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
		Host:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Host},
		ContentURL: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ContentURL},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getArticleOnly(t *testing.T) {
	info, err := GetArticleOnly(context.Background(), &npool.Conds{
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
		Host:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Host},
		ContentURL: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ContentURL},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteArticle(t *testing.T) {
	info, err := DeleteArticle(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}

	info, err = GetArticle(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestArticle(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setup(t)
	defer teardown(t)

	t.Run("createArticle", createArticle)
	t.Run("updateArticle", updateArticle)
	t.Run("getArticle", getArticle)
	t.Run("getArticles", getArticles)
	t.Run("getArticleOnly", getArticleOnly)
	t.Run("deleteArticle", deleteArticle)
}
