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

func (s *Server) ExistMedia(ctx context.Context, in *npool.ExistMediaRequest) (*npool.ExistMediaResponse, error) {
	handler, err := media1.NewHandler(
		ctx,
		media1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMedia",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistMedia(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMedia",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMediaResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistMediaResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistMediaConds(ctx context.Context, in *npool.ExistMediaCondsRequest) (*npool.ExistMediaCondsResponse, error) {
	handler, err := media1.NewHandler(
		ctx,
		media1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMediaConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMediaCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistMediaConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMediaConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMediaCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistMediaCondsResponse{
		Info: exist,
	}, nil
}
