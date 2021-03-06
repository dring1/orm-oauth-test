package middleware

import (
	"net/http"

	"github.com/dring1/jwt-oauth/config"
	"github.com/dring1/jwt-oauth/services"
)

type Middleware func(http.Handler) http.Handler

func Handlers(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, mdlware := range middlewares {
		handler = mdlware(handler)
	}
	return handler
}

func DefaultMiddleWare(config *config.Cfg, svcs *services.Services) []Middleware {
	// order from last to first - LIFO
	globalMiddlewares := []Middleware{
		//JsonResponseHandler,
		NewApacheLoggingHandler(svcs.Logger),
		AddUUID,
		ContextCreate,
		RecoverHandler,
	}
	return globalMiddlewares
}
