package service

import (
	"github.com/krivenkov/order/internal/model/order"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	fx.Provide(
		order.New,
	),
)
