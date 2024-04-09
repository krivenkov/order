package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/krivenkov/order/internal/server/http/auth"
	"github.com/krivenkov/order/internal/server/http/handlers"
	"github.com/krivenkov/order/internal/server/http/middlewares"
	"github.com/krivenkov/order/internal/server/http/operations"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var FXModule = fx.Options(
	handlers.FXModule,
	middlewares.FXModule,

	fx.Provide(
		newApi,
		newServer,
		auth.NewJWT,
	),

	fx.Invoke(
		// middlewares
		invokeMiddlewares,

		// server
		invokeApi,
	),
)

func newApi(logger *zap.Logger, authJWT *auth.JWT) (*operations.OrderAPIAPI, error) {
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("load specs: %w", err)
	}

	api := operations.NewOrderAPIAPI(swaggerSpec)

	api.Logger = logger.Sugar().Infof
	api.ServeError = serveError
	api.UseSwaggerUI()
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.JWTAuth = authJWT.Handle

	return api, nil
}

func newServer(cfg Config, lc fx.Lifecycle, api *operations.OrderAPIAPI) (*Server, error) {
	server := NewServer(nil)

	server.Host = cfg.Host
	server.Port = cfg.Port
	server.EnabledListeners = []string{schemeHTTP}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if errLoc := server.Listen(); errLoc != nil {
				return fmt.Errorf("listen fail: %w", errLoc)
			}

			go func() {
				_ = server.Serve()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})

	return server, nil
}

func invokeApi(api *operations.OrderAPIAPI, server *Server) {
	server.api = api
}

func invokeMiddlewares(
	api *operations.OrderAPIAPI,
	server *Server, logger *middlewares.Logger) {

	log := logger.Provide(api.Serve(nil))

	mux := http.NewServeMux()

	mux.Handle("/", log)

	server.handler = mux
}
