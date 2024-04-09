package grpc

import (
	"github.com/krivenkov/order/internal/server/grpc/inner"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var FXModule = fx.Options(
	fx.Provide(NewServer),

	inner.FXModule,

	fx.Invoke(func(server *grpc.Server) {
		grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	}),
)
