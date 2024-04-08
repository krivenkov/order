package es

import (
	"github.com/krivenkov/order/internal/storage/es/order"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	order.FXModule,
)
