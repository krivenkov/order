package api

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=order.api.pb.go -package=mock -destination=mock/order.api.pb.go

type ClientConfig struct {
	Addr string `json:"addr" yaml:"addr" env:"ADDR"`
}

func NewClient(c ClientConfig, lc fx.Lifecycle) (OrderServiceClient, error) {
	cli, err := dial(c, lc)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	return NewOrderServiceClient(cli), nil
}

func dial(c ClientConfig, lc fx.Lifecycle) (*grpc.ClientConn, error) {
	cli, err := grpc.Dial(c.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}

	if lc != nil {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return cli.Close()
			},
		})
	}

	return cli, nil
}
