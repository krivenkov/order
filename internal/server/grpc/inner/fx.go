package inner

import (
	"github.com/krivenkov/order/pkg/api"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var FXModule = fx.Options(
	fx.Provide(NewServer),

	fx.Invoke(func(server *grpc.Server, service api.OrderServiceServer) {
		api.RegisterOrderServiceServer(server, service)
	}),
)
