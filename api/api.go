package api

import (
	"context"

	"github.com/NpoolPlatform/cms-middleware/api/acl"
	"github.com/NpoolPlatform/cms-middleware/api/article"
	"github.com/NpoolPlatform/cms-middleware/api/category"
	"github.com/NpoolPlatform/cms-middleware/api/media"
	cms "github.com/NpoolPlatform/message/npool/cms/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	cms.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	cms.RegisterMiddlewareServer(server, &Server{})
	article.Register(server)
	category.Register(server)
	media.Register(server)
	acl.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := cms.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}
