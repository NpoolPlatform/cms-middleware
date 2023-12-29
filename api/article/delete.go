package article

import (
	"context"

	article1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/article"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
)

func (s *Server) DeleteArticle(ctx context.Context, in *npool.DeleteArticleRequest) (*npool.DeleteArticleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteArticle",
			"In", in,
		)
		return &npool.DeleteArticleResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := article1.NewHandler(
		ctx,
		article1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteArticle",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteArticle(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteArticle",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteArticleResponse{
		Info: info,
	}, nil
}
