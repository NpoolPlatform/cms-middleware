package acl

import (
	"context"

	acl1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/acl"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
)

func (s *Server) DeleteACL(ctx context.Context, in *npool.DeleteACLRequest) (*npool.DeleteACLResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteACL",
			"In", in,
		)
		return &npool.DeleteACLResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := acl1.NewHandler(
		ctx,
		acl1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteACL",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteACL(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteACL",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteACLResponse{
		Info: info,
	}, nil
}
