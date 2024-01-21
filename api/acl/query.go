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

func (s *Server) GetACL(ctx context.Context, in *npool.GetACLRequest) (*npool.GetACLResponse, error) {
	handler, err := acl1.NewHandler(
		ctx,
		acl1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetACL",
			"In", in,
			"Error", err,
		)
		return &npool.GetACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetACL(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetACL",
			"In", in,
			"Error", err,
		)
		return &npool.GetACLResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetACLResponse{
		Info: info,
	}, nil
}

func (s *Server) GetACLs(ctx context.Context, in *npool.GetACLsRequest) (*npool.GetACLsResponse, error) {
	handler, err := acl1.NewHandler(
		ctx,
		acl1.WithConds(in.GetConds()),
		acl1.WithOffset(in.GetOffset()),
		acl1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetACLs",
			"In", in,
			"Error", err,
		)
		return &npool.GetACLsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetACLs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetACLs",
			"In", in,
			"Error", err,
		)
		return &npool.GetACLsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetACLsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
