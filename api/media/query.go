//nolint:dupl
package media

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	media1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/media"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/media"
)

func (s *Server) GetMedia(ctx context.Context, in *npool.GetMediaRequest) (*npool.GetMediaResponse, error) {
	handler, err := media1.NewHandler(
		ctx,
		media1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMedia",
			"In", in,
			"Error", err,
		)
		return &npool.GetMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetMedia(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMedia",
			"In", in,
			"Error", err,
		)
		return &npool.GetMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMediaResponse{
		Info: info,
	}, nil
}

func (s *Server) GetMedias(ctx context.Context, in *npool.GetMediasRequest) (*npool.GetMediasResponse, error) {
	handler, err := media1.NewHandler(
		ctx,
		media1.WithConds(in.GetConds()),
		media1.WithOffset(in.GetOffset()),
		media1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMedias",
			"In", in,
			"Error", err,
		)
		return &npool.GetMediasResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetMedias(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMedias",
			"In", in,
			"Error", err,
		)
		return &npool.GetMediasResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMediasResponse{
		Infos: infos,
		Total: total,
	}, nil
}
