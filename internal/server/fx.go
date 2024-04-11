package server

import (
	"github.com/krivenkov/order/internal/server/bus"
	"github.com/krivenkov/order/internal/server/grpc"
	"github.com/krivenkov/order/internal/server/http"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	bus.FXModule,
	http.FXModule,
	grpc.FXModule,
)
