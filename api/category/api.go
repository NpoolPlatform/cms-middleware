package category

import (
	"github.com/NpoolPlatform/message/npool/cms/mw/v1/category"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	category.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	category.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
