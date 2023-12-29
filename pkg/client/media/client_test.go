package media

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"

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

var ret = &npool.Media{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	Name:      uuid.NewString(),
	Ext:       ".jpg",
	MediaURL:  uuid.NewString(),
	CreatedBy: uuid.NewString(),
}

var req = &npool.MediaReq{
	EntID:     &ret.EntID,
	AppID:     &ret.AppID,
	Name:      &ret.Name,
	Ext:       &ret.Ext,
	MediaURL:  &ret.MediaURL,
	CreatedBy: &ret.CreatedBy,
}

func createMedia(t *testing.T) {
	info, err := CreateMedia(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func getMedia(t *testing.T) {
	info, err := GetMedia(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getMedias(t *testing.T) {
	infos, total, err := GetMedias(context.Background(), &npool.Conds{
		ID:       &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		MediaURL: &basetypes.StringVal{Op: cruder.EQ, Value: ret.MediaURL},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getMediaOnly(t *testing.T) {
	info, err := GetMediaOnly(context.Background(), &npool.Conds{
		ID:       &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		MediaURL: &basetypes.StringVal{Op: cruder.EQ, Value: ret.MediaURL},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteMedia(t *testing.T) {
	info, err := DeleteMedia(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}

	info, err = GetMedia(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMedia(t *testing.T) {
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

	t.Run("createMedia", createMedia)
	t.Run("getMedia", getMedia)
	t.Run("getMedias", getMedias)
	t.Run("getMediaOnly", getMediaOnly)
	t.Run("deleteMedia", deleteMedia)
}
