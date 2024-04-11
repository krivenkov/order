package di

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/krivenkov/pkg/auth"
	"github.com/krivenkov/pkg/clients/database"
	"github.com/krivenkov/pkg/clients/es"
	"github.com/krivenkov/pkg/global"
	"github.com/krivenkov/pkg/mcfg"
	"github.com/krivenkov/pkg/mlog"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var FXBaseModule = fx.Options(
	fx.Provide(
		global.NewInfo(appName),
		mlog.NewC,
		mcfg.NewConfig[Config],

		database.NewPgxPool,
		database.NewTXerFX,
		es.NewClient,
		auth.NewClient,

		func(logger *zap.Logger) context.Context {
			return mlog.CtxWithLogger(context.Background(), logger)
		},
	),

	fx.Supply(
		time.Now,
		uuid.New,
	),
)
