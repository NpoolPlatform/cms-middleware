package category

import (
	"context"

	category1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/category"
)

func (s *Server) CreateCategory(ctx context.Context, in *npool.CreateCategoryRequest) (*npool.CreateCategoryResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCategory",
			"In", in,
		)
		return &npool.CreateCategoryResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := category1.NewHandler(
		ctx,
		category1.WithEntID(req.EntID, false),
		category1.WithAppID(req.AppID, true),
		category1.WithParentID(req.ParentID, false),
		category1.WithName(req.Name, true),
		category1.WithEnabled(req.Enabled, false),
		category1.WithSlug(req.Slug, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCategory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateCategory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCategory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCategoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCategoryResponse{
		Info: info,
	}, nil
}
