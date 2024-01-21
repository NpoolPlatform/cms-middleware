package media

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"
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
	ret = npool.Media{
		EntID:     uuid.NewString(),
		AppID:     uuid.NewString(),
		Name:      uuid.NewString(),
		Ext:       ".jpg",
		MediaURL:  uuid.NewString(),
		CreatedBy: uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createMedia(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithName(&ret.Name, true),
		WithExt(&ret.Ext, false),
		WithMediaURL(&ret.MediaURL, true),
		WithCreatedBy(&ret.CreatedBy, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateMedia(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getMedia(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetMedia(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getMedias(t *testing.T) {
	conds := &npool.Conds{
		ID:       &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		MediaURL: &basetypes.StringVal{Op: cruder.EQ, Value: ret.MediaURL},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetMedias(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteMedia(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteMedia(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetMedia(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMedia(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createMedia", createMedia)
	t.Run("getMedia", getMedia)
	t.Run("getMedias", getMedias)
	t.Run("deleteMedia", deleteMedia)
}
