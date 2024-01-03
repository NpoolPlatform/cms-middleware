package acl

import (
	"github.com/NpoolPlatform/message/npool/cms/mw/v1/acl"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	acl.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	acl.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
