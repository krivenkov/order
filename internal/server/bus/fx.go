package bus

import (
	"github.com/krivenkov/order/internal/model/user"
	"github.com/krivenkov/order/internal/server/bus/user_handler"
	busBuilder "github.com/krivenkov/pkg/bus/builder"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	fx.Provide(
		newRouter,
		user_handler.New,
	),

	fx.Provide(
		fx.Annotate(busBuilder.NewFXPublisher[user.User](user.UpdateUserTopic), fx.ResultTags(`name:"user_bus_update"`)),
	),

	fx.Invoke(registerRoutes),
)
