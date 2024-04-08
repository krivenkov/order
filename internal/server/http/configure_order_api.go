// This file is safe to edit. Once it exists it will not be overwritten

package http

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/krivenkov/order/internal/server/http/operations"
	"github.com/krivenkov/order/internal/server/http/operations/order"
)

//go:generate swagger generate server --target ../../../../order --name OrderAPI --spec ../../../api-spec/swagger.json --model-package internal/server/http/models --server-package internal/server/http --principal interface{} --exclude-main

func configureFlags(api *operations.OrderAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OrderAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.JWTAuth == nil {
		api.JWTAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.OrderCreateOrderHandler == nil {
		api.OrderCreateOrderHandler = order.CreateOrderHandlerFunc(func(params order.CreateOrderParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation order.CreateOrder has not yet been implemented")
		})
	}
	if api.OrderDeleteOrderHandler == nil {
		api.OrderDeleteOrderHandler = order.DeleteOrderHandlerFunc(func(params order.DeleteOrderParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation order.DeleteOrder has not yet been implemented")
		})
	}
	if api.OrderGetOrderHandler == nil {
		api.OrderGetOrderHandler = order.GetOrderHandlerFunc(func(params order.GetOrderParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation order.GetOrder has not yet been implemented")
		})
	}
	if api.OrderGetOrdersHandler == nil {
		api.OrderGetOrdersHandler = order.GetOrdersHandlerFunc(func(params order.GetOrdersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation order.GetOrders has not yet been implemented")
		})
	}
	if api.OrderGetOrdersCountHandler == nil {
		api.OrderGetOrdersCountHandler = order.GetOrdersCountHandlerFunc(func(params order.GetOrdersCountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation order.GetOrdersCount has not yet been implemented")
		})
	}
	if api.OrderUpdateOrderHandler == nil {
		api.OrderUpdateOrderHandler = order.UpdateOrderHandlerFunc(func(params order.UpdateOrderParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation order.UpdateOrder has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
