//nolint:dupl
package article

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	article1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/article"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
)

func (s *Server) ExistArticle(ctx context.Context, in *npool.ExistArticleRequest) (*npool.ExistArticleResponse, error) {
	handler, err := article1.NewHandler(
		ctx,
		article1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistArticle",
			"In", in,
			"Error", err,
		)
		return &npool.ExistArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistArticle(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistArticle",
			"In", in,
			"Error", err,
		)
		return &npool.ExistArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistArticleResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistArticleConds(ctx context.Context, in *npool.ExistArticleCondsRequest) (*npool.ExistArticleCondsResponse, error) {
	handler, err := article1.NewHandler(
		ctx,
		article1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistArticleConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistArticleCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistArticleConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistArticleConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistArticleCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistArticleCondsResponse{
		Info: exist,
	}, nil
}
