package storage

import (
	"github.com/krivenkov/order/internal/storage/es"
	"github.com/krivenkov/order/internal/storage/pg"
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	pg.FXModule,
	es.FXModule,
)
