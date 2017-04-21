package routes

import (
	"app/auth"
	"app/errors"
	"app/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type AppRouter struct {
	router *mux.Router
}

func NewAppRouter() *AppRouter {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(errors.RouteNotFoundHandler)
	return &AppRouter{
		router: router,
	}
}

func (app_router *AppRouter) AddNamespace(namespaceName string, routes Routes, middlewares ...func(http.HandlerFunc) http.HandlerFunc) *AppRouter {
	namespace := app_router.router.PathPrefix(namespaceName).Subrouter()
	for _, route := range routes {
		var handler http.HandlerFunc
		handler = logger.Logger(route.HandlerFunc, route.Name)
		handler = ApplyMiddlewares(handler, middlewares)
		namespace.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return app_router
}

func ApplyMiddlewares(handler http.HandlerFunc, middlewares []func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func NewRouter() *mux.Router {
	r := NewAppRouter()
	r.
		AddNamespace("/api/v1", api_routes, auth.AuthenticateMiddleware).
		AddNamespace("/auth", auth_routes)
	return r.router
}
