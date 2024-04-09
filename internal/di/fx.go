package di

import (
	"time"

	"github.com/google/uuid"
	"github.com/krivenkov/pkg/clients/database"
	"github.com/krivenkov/pkg/global"
	"github.com/krivenkov/pkg/mcfg"
	"github.com/krivenkov/pkg/mlog"
	"go.uber.org/fx"
)

var FXBaseModule = fx.Options(
	fx.Provide(
		global.NewInfo(appName),
		mlog.NewC,
		mcfg.NewConfig[Config],

		database.NewPgxPool,
		database.NewTXerFX,
	),

	fx.Supply(
		time.Now,
		uuid.New,
	),
)
