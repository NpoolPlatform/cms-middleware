//nolint:dupl
package acl

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	acl1 "github.com/NpoolPlatform/cms-middleware/pkg/mw/acl"
	npool "github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"
)

func (s *Server) ExistACL(ctx context.Context, in *npool.ExistACLRequest) (*npool.ExistACLResponse, error) {
	handler, err := acl1.NewHandler(
		ctx,
		acl1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistACL",
			"In", in,
			"Error", err,
		)
		return &npool.ExistACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistACL(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistACL",
			"In", in,
			"Error", err,
		)
		return &npool.ExistACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistACLResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistACLConds(ctx context.Context, in *npool.ExistACLCondsRequest) (*npool.ExistACLCondsResponse, error) {
	handler, err := acl1.NewHandler(
		ctx,
		acl1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistACLConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistACLCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistACLConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistACLConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistACLCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistACLCondsResponse{
		Info: exist,
	}, nil
}
