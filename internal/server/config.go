package server

import (
	"github.com/krivenkov/order/internal/server/bus"
	"github.com/krivenkov/order/internal/server/grpc"
	"github.com/krivenkov/order/internal/server/http"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	Bus  bus.Config  `json:"bus" yaml:"bus" envPrefix:"BUS_"`
	HTTP http.Config `json:"http" yaml:"http" envPrefix:"HTTP_"`
	GRPC grpc.Config `json:"grpc" yaml:"grpc" envPrefix:"GRPC_"`
}
