package order

import "go.uber.org/fx"

var FXModule = fx.Options(
	fx.Provide(
		fx.Annotate(NewCommander, fx.ResultTags(`name:"order_pg_cmd"`)),
		fx.Annotate(NewQuerier, fx.ResultTags(`name:"order_pg_qr"`)),
	),
)
