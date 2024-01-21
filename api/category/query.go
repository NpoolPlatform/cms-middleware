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

func (s *Server) GetCategory(ctx context.Context, in *npool.GetCategoryRequest) (*npool.GetCategoryResponse, error) {
	handler, err := category1.NewHandler(
		ctx,
		category1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCategory",
			"In", in,
			"Error", err,
		)
		return &npool.GetCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetCategory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCategory",
			"In", in,
			"Error", err,
		)
		return &npool.GetCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCategoryResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCategories(ctx context.Context, in *npool.GetCategoriesRequest) (*npool.GetCategoriesResponse, error) {
	handler, err := category1.NewHandler(
		ctx,
		category1.WithConds(in.GetConds()),
		category1.WithOffset(in.GetOffset()),
		category1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCategories",
			"In", in,
			"Error", err,
		)
		return &npool.GetCategoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetCategories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCategories",
			"In", in,
			"Error", err,
		)
		return &npool.GetCategoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCategoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
