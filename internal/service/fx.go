package service

import (
	"github.com/krivenkov/order/internal/service/order"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	fx.Provide(
		order.New,
	),
)
