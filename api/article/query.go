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

func (s *Server) GetArticle(ctx context.Context, in *npool.GetArticleRequest) (*npool.GetArticleResponse, error) {
	handler, err := article1.NewHandler(
		ctx,
		article1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetArticle",
			"In", in,
			"Error", err,
		)
		return &npool.GetArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetArticle(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetArticle",
			"In", in,
			"Error", err,
		)
		return &npool.GetArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetArticleResponse{
		Info: info,
	}, nil
}

func (s *Server) GetArticles(ctx context.Context, in *npool.GetArticlesRequest) (*npool.GetArticlesResponse, error) {
	handler, err := article1.NewHandler(
		ctx,
		article1.WithConds(in.GetConds()),
		article1.WithOffset(in.GetOffset()),
		article1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetArticles",
			"In", in,
			"Error", err,
		)
		return &npool.GetArticlesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetArticles(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetArticles",
			"In", in,
			"Error", err,
		)
		return &npool.GetArticlesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetArticlesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
