package article

import (
	"context"

	article1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/article"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
)

func (s *Server) UpdateArticle(ctx context.Context, in *npool.UpdateArticleRequest) (*npool.UpdateArticleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateArticle",
			"In", in,
		)
		return &npool.UpdateArticleResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := article1.NewHandler(
		ctx,
		article1.WithID(req.ID, true),
		article1.WithCategoryID(req.CategoryID, false),
		article1.WithAuthorID(req.AuthorID, false),
		article1.WithTitle(req.Title, false),
		article1.WithSubtitle(req.Subtitle, false),
		article1.WithDigest(req.Digest, false),
		article1.WithStatus(req.Status, false),
		article1.WithContentURL(req.ContentURL, false),
		article1.WithLatest(req.Latest, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateArticle",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateArticle(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateArticle",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateArticleResponse{
		Info: info,
	}, nil
}
