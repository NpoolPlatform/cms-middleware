package category

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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

var ret = &npool.Category{
	EntID:   uuid.NewString(),
	AppID:   uuid.NewString(),
	Name:    uuid.NewString(),
	Slug:    uuid.NewString(),
	Enabled: false,
	Index:   uint32(0),
}

var req = &npool.CategoryReq{
	EntID:   &ret.EntID,
	AppID:   &ret.AppID,
	Name:    &ret.Name,
	Slug:    &ret.Slug,
	Enabled: &ret.Enabled,
	Index:   &ret.Index,
}

func createCategory(t *testing.T) {
	info, err := CreateCategory(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.ParentID = info.ParentID
		assert.Equal(t, ret, info)
	}
}

func updateCategory(t *testing.T) {
	req.ID = &ret.ID
	info, err := UpdateCategory(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getCategory(t *testing.T) {
	info, err := GetCategory(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getCategories(t *testing.T) {
	infos, total, err := GetCategories(context.Background(), &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		Enabled: &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getCategoryOnly(t *testing.T) {
	info, err := GetCategoryOnly(context.Background(), &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		Enabled: &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteCategory(t *testing.T) {
	info, err := DeleteCategory(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}

	info, err = GetCategory(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCategory(t *testing.T) {
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

	t.Run("createCategory", createCategory)
	t.Run("updateCategory", updateCategory)
	t.Run("getCategory", getCategory)
	t.Run("getCategories", getCategories)
	t.Run("getCategoryOnly", getCategoryOnly)
	t.Run("deleteCategory", deleteCategory)
}
