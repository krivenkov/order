package item

import "go.uber.org/fx"

var FXModule = fx.Options(
	fx.Provide(
		NewCommander,
		NewQuerier,
	),
)
