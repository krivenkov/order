package handlers

import (
	"github.com/krivenkov/order/internal/server/http/handlers/order"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	order.FXModule,
)
