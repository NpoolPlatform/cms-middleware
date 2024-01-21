package media

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"

	servicename "github.com/NpoolPlatform/cms-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateMedia(ctx context.Context, in *npool.MediaReq) (*npool.Media, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateMedia(ctx, &npool.CreateMediaRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Media), nil
}

func GetMedia(ctx context.Context, id string) (*npool.Media, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetMedia(ctx, &npool.GetMediaRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Media), nil
}

func GetMedias(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Media, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetMedias(ctx, &npool.GetMediasRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Media), total, nil
}

func GetMediaOnly(ctx context.Context, conds *npool.Conds) (*npool.Media, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetMedias(ctx, &npool.GetMediasRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Media)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Media)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Media)[0], nil
}

func DeleteMedia(ctx context.Context, id uint32) (*npool.Media, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteMedia(ctx, &npool.DeleteMediaRequest{
			Info: &npool.MediaReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Media), nil
}

func ExistMediaConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistMediaConds(ctx, &npool.ExistMediaCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistMedia(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistMedia(ctx, &npool.ExistMediaRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}
