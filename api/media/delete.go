package media

import (
	"context"

	media1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/media"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"
)

func (s *Server) DeleteMedia(ctx context.Context, in *npool.DeleteMediaRequest) (*npool.DeleteMediaResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteMedia",
			"In", in,
		)
		return &npool.DeleteMediaResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := media1.NewHandler(
		ctx,
		media1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMedia",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteMedia(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMedia",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteMediaResponse{
		Info: info,
	}, nil
}
