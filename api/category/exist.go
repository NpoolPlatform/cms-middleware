//nolint:dupl
package category

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	category1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
)

func (s *Server) ExistCategory(ctx context.Context, in *npool.ExistCategoryRequest) (*npool.ExistCategoryResponse, error) {
	handler, err := category1.NewHandler(
		ctx,
		category1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCategory",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCategory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCategory",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCategoryResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCategoryConds(ctx context.Context, in *npool.ExistCategoryCondsRequest) (*npool.ExistCategoryCondsResponse, error) {
	handler, err := category1.NewHandler(
		ctx,
		category1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCategoryConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCategoryCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCategoryConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCategoryConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCategoryCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCategoryCondsResponse{
		Info: exist,
	}, nil
}
