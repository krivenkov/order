package middlewares

import "go.uber.org/fx"

var FXModule = fx.Options(
	fx.Provide(
		NewLogger,
	),
)
