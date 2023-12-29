package category

import (
	"context"

	category1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
)

func (s *Server) UpdateCategory(ctx context.Context, in *npool.UpdateCategoryRequest) (*npool.UpdateCategoryResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCategory",
			"In", in,
		)
		return &npool.UpdateCategoryResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := category1.NewHandler(
		ctx,
		category1.WithID(req.ID, true),
		category1.WithParentID(req.ParentID, false),
		category1.WithName(req.Name, false),
		category1.WithEnabled(req.Enabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCategory",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateCategory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCategory",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateCategoryResponse{
		Info: info,
	}, nil
}
