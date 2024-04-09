package server

import (
	"github.com/krivenkov/order/internal/server/http"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	HTTP http.Config `json:"http" yaml:"http" envPrefix:"HTTP_"`
}
