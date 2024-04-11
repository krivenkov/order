package di

import (
	"github.com/krivenkov/order/internal/server"
	"github.com/krivenkov/pkg/auth"
	busBuilder "github.com/krivenkov/pkg/bus/builder"
	"github.com/krivenkov/pkg/clients/database"
	"github.com/krivenkov/pkg/clients/es"
	"github.com/krivenkov/pkg/mlog"
	"go.uber.org/fx"
)

const appName = "order-api"

type Config struct {
	fx.Out

	Log  mlog.Config       `json:"log" yaml:"log" envPrefix:"LOG_"`
	Bus  busBuilder.Config `json:"bus" yaml:"bus" envPrefix:"BUS_"`
	DB   database.Config   `json:"db" yaml:"db" envPrefix:"DB_"`
	ES   es.Config         `json:"es" yaml:"es" envPrefix:"ES_"`
	Auth auth.Config       `json:"auth" yaml:"auth" envPrefix:"AUTH_"`

	Server server.Config `json:"server" yaml:"server" envPrefix:"SERVER_"`
}
