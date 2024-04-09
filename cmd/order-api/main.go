package main

import (
	"github.com/krivenkov/order/internal/di"
	"github.com/krivenkov/order/internal/server"
	"github.com/krivenkov/order/internal/service"
	"github.com/krivenkov/order/internal/storage"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.FXBaseModule,

		storage.FXModule,
		service.FXModule,
		server.FXModule,
	).Run()
}
