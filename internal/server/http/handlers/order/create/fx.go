package create

import (
	"github.com/krivenkov/order/internal/server/http/operations"
	"github.com/krivenkov/order/internal/server/http/operations/order"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	fx.Provide(New),

	fx.Invoke(
		func(handler order.CreateOrderHandler, api *operations.OrderAPIAPI) {
			api.OrderCreateOrderHandler = handler
		},
	),
)
