package media

import (
	"context"

	media1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/media"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"
)

func (s *Server) CreateMedia(ctx context.Context, in *npool.CreateMediaRequest) (*npool.CreateMediaResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateMedia",
			"In", in,
		)
		return &npool.CreateMediaResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := media1.NewHandler(
		ctx,
		media1.WithEntID(req.EntID, false),
		media1.WithAppID(req.AppID, true),
		media1.WithName(req.Name, true),
		media1.WithMediaURL(req.MediaURL, true),
		media1.WithExt(req.Ext, true),
		media1.WithCreatedBy(req.CreatedBy, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMedia",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateMedia(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMedia",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateMediaResponse{
		Info: info,
	}, nil
}
