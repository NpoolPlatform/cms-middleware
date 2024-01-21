package category

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
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
	ret = npool.Category{
		EntID:   uuid.NewString(),
		AppID:   uuid.NewString(),
		Name:    uuid.NewString(),
		Slug:    uuid.NewString(),
		Enabled: false,
		Index:   uint32(0),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCategory(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithName(&ret.Name, true),
		WithEnabled(&ret.Enabled, false),
		WithSlug(&ret.Slug, true),
		WithIndex(&ret.Index, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCategory(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.ParentID = info.ParentID
		assert.Equal(t, info, &ret)
	}
}

func updateCategory(t *testing.T) {
	ret.Enabled = true
	ret.Name = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithName(&ret.Name, false),
		WithEnabled(&ret.Enabled, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCategory(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCategory(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCategory(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCategories(t *testing.T) {
	conds := &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		Enabled: &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCategories(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCategory(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCategory(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCategory(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCategory(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCategory", createCategory)
	t.Run("updateCategory", updateCategory)
	t.Run("getCategory", getCategory)
	t.Run("getCategories", getCategories)
	t.Run("deleteCategory", deleteCategory)
}
