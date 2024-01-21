package media

import (
	"github.com/NpoolPlatform/message/npool/cms/mw/v1/media"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	media.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	media.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
