package article

import (
	"github.com/NpoolPlatform/message/npool/cms/mw/v1/article"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	article.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	article.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
