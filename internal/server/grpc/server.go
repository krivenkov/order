package grpc

import (
	"context"
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/krivenkov/pkg/mlog"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Params struct {
	fx.In

	Logger *zap.Logger
	Cfg    Config
}

func NewServer(lc fx.Lifecycle, p Params) (*grpc.Server, error) {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcRecovery.UnaryServerInterceptor(),
				func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
					return handler(mlog.CtxWithLogger(ctx, p.Logger), req)
				},
			),
		),
	)

	socket, err := net.Listen("tcp", p.Cfg.Addr())
	if err != nil {
		return nil, fmt.Errorf("create tcp socket: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.Logger.Info("staring GRPC server", zap.String("addr", p.Cfg.Addr()))

			go func() {
				if errServe := server.Serve(socket); errServe != nil {
					p.Logger.Fatal("serve GRPC server failed", zap.Error(errServe))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer p.Logger.Info("GRPC server stopped")

			server.GracefulStop()
			return socket.Close()
		},
	})

	return server, nil
}
