package category

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

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

func CreateCategory(ctx context.Context, in *npool.CategoryReq) (*npool.Category, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCategory(ctx, &npool.CreateCategoryRequest{
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
	return info.(*npool.Category), nil
}

func GetCategory(ctx context.Context, id string) (*npool.Category, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCategory(ctx, &npool.GetCategoryRequest{
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
	return info.(*npool.Category), nil
}

func GetCategories(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Category, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCategories(ctx, &npool.GetCategoriesRequest{
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
	return infos.([]*npool.Category), total, nil
}

func GetCategoryOnly(ctx context.Context, conds *npool.Conds) (*npool.Category, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCategories(ctx, &npool.GetCategoriesRequest{
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
	if len(infos.([]*npool.Category)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Category)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Category)[0], nil
}

func DeleteCategory(ctx context.Context, id uint32) (*npool.Category, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteCategory(ctx, &npool.DeleteCategoryRequest{
			Info: &npool.CategoryReq{
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
	return info.(*npool.Category), nil
}

func ExistCategoryConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistCategoryConds(ctx, &npool.ExistCategoryCondsRequest{
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

func ExistCategory(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistCategory(ctx, &npool.ExistCategoryRequest{
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

func UpdateCategory(ctx context.Context, in *npool.CategoryReq) (*npool.Category, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateCategory(ctx, &npool.UpdateCategoryRequest{
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
	return info.(*npool.Category), nil
}
