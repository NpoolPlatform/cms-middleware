package acl

import (
	"context"

	acl1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/acl"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
)

func (s *Server) CreateACL(ctx context.Context, in *npool.CreateACLRequest) (*npool.CreateACLResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateACL",
			"In", in,
		)
		return &npool.CreateACLResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := acl1.NewHandler(
		ctx,
		acl1.WithEntID(req.EntID, false),
		acl1.WithAppID(req.AppID, true),
		acl1.WithRoleID(req.RoleID, true),
		acl1.WithArticleKey(req.ArticleKey, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateACL",
			"In", in,
			"Error", err,
		)
		return &npool.CreateACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateACL(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateACL",
			"In", in,
			"Error", err,
		)
		return &npool.CreateACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateACLResponse{
		Info: info,
	}, nil
}
