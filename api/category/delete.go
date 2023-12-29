package category

import (
	"context"

	category1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
)

func (s *Server) DeleteCategory(ctx context.Context, in *npool.DeleteCategoryRequest) (*npool.DeleteCategoryResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCategory",
			"In", in,
		)
		return &npool.DeleteCategoryResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := category1.NewHandler(
		ctx,
		category1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCategory",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteCategory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCategory",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCategoryResponse{
		Info: info,
	}, nil
}
