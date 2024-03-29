package acl

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	article1 "github.com/NpoolPlatform/cms-middleware/pkg/client/article"
	category1 "github.com/NpoolPlatform/cms-middleware/pkg/client/category"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
	articlemw "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
	categorymw "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	cmstypes "github.com/NpoolPlatform/message/npool/basetypes/cms/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

	ret = &npool.ACL{
		EntID:      uuid.NewString(),
		AppID:      articleRet.AppID,
		RoleID:     uuid.NewString(),
		ArticleKey: articleRet.ArticleKey,
	}

	req = &npool.ACLReq{
		EntID:      &ret.EntID,
		AppID:      &ret.AppID,
		RoleID:     &ret.RoleID,
		ArticleKey: &ret.ArticleKey,
	}
)

func setup(t *testing.T) func(*testing.T) {
	info1, err := category1.CreateCategory(context.Background(), &categorymw.CategoryReq{
		EntID: &categoryRet.EntID,
		AppID: &categoryRet.AppID,
		Name:  &categoryRet.Name,
		Slug:  &categoryRet.Slug,
		Index: &categoryRet.Index,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info1)

	info2, err := article1.CreateArticle(context.Background(), &articlemw.ArticleReq{
		EntID:      &articleRet.EntID,
		AppID:      &articleRet.AppID,
		CategoryID: &articleRet.CategoryID,
		AuthorID:   &articleRet.AuthorID,
		ArticleKey: &articleRet.ArticleKey,
		Title:      &articleRet.Title,
		Subtitle:   &articleRet.Subtitle,
		Digest:     &articleRet.Digest,
		ContentURL: &articleRet.ContentURL,
		Host:       &articleRet.Host,
		ISO:        &articleRet.ISO,
		Version:    &articleRet.Version,
		ACLEnabled: &articleRet.ACLEnabled,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info2)

	return func(*testing.T) {
		_, _ = category1.DeleteCategory(context.Background(), info1.ID)
		_, _ = article1.DeleteArticle(context.Background(), info2.ID)
	}
}

func createACL(t *testing.T) {
	info, err := CreateACL(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func getACL(t *testing.T) {
	info, err := GetACL(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getACLs(t *testing.T) {
	infos, total, err := GetACLs(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RoleID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.RoleID},
		ArticleKey: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ArticleKey},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getACLOnly(t *testing.T) {
	info, err := GetACLOnly(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RoleID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.RoleID},
		ArticleKey: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ArticleKey},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteACL(t *testing.T) {
	info, err := DeleteACL(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}

	info, err = GetACL(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestACL(t *testing.T) {
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

	t.Run("createACL", createACL)
	t.Run("getACL", getACL)
	t.Run("getACLs", getACLs)
	t.Run("getACLOnly", getACLOnly)
	t.Run("deleteACL", deleteACL)
}
