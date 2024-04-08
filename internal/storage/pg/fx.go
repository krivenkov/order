package pg

import (
	"github.com/krivenkov/order/internal/storage/pg/order"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	order.FXModule,
)
