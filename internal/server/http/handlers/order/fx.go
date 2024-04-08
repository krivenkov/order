package order

import (
	"github.com/krivenkov/order/internal/server/http/handlers/order/count"
	"github.com/krivenkov/order/internal/server/http/handlers/order/create"
	"github.com/krivenkov/order/internal/server/http/handlers/order/item"
	"github.com/krivenkov/order/internal/server/http/handlers/order/list"
	"github.com/krivenkov/order/internal/server/http/handlers/order/remove"
	"github.com/krivenkov/order/internal/server/http/handlers/order/update"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	create.FXModule,
	update.FXModule,
	remove.FXModule,
	list.FXModule,
	item.FXModule,
	count.FXModule,
)
