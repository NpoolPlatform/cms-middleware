package article

import (
	"context"

	article1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/article"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/article"
)

func (s *Server) CreateArticle(ctx context.Context, in *npool.CreateArticleRequest) (*npool.CreateArticleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateArticle",
			"In", in,
		)
		return &npool.CreateArticleResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := article1.NewHandler(
		ctx,
		article1.WithEntID(req.EntID, false),
		article1.WithAppID(req.AppID, true),
		article1.WithCategoryID(req.CategoryID, true),
		article1.WithAuthorID(req.AuthorID, true),
		article1.WithArticleKey(req.ArticleKey, true),
		article1.WithTitle(req.Title, true),
		article1.WithSubtitle(req.Subtitle, false),
		article1.WithDigest(req.Digest, false),
		article1.WithStatus(req.Status, false),
		article1.WithHost(req.Host, true),
		article1.WithISO(req.ISO, true),
		article1.WithContentURL(req.ContentURL, true),
		article1.WithVersion(req.Version, true),
		article1.WithACLEnabled(req.ACLEnabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateArticle",
			"In", in,
			"Error", err,
		)
		return &npool.CreateArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateArticle(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateArticle",
			"In", in,
			"Error", err,
		)
		return &npool.CreateArticleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateArticleResponse{
		Info: info,
	}, nil
}
