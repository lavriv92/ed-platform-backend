package routes

import (
	"fmt"
	"net/http"

	"app/auth"
	"app/courses"
	"app/errors"
	"app/logger"

	"github.com/gorilla/mux"
)

const (
	API_VERSION = "1"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(errors.RouteNotFoundHandler)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func _NewRoute(namespace string, name string, method string, endpoint string, handler http.HandlerFunc) Route {
	return Route{
		name,
		method,
		fmt.Sprintf("/%s%s", namespace, endpoint),
		handler,
	}
}

func _NewApiRoute(namespace string, name string, method string, endpoint string, handler http.HandlerFunc) Route {
	return Route{
		name,
		method,
		fmt.Sprintf("/%s/v%s%s", namespace, API_VERSION, endpoint),
		handler,
	}
}

var routes = Routes{
	_NewRoute(
		"auth",
		"Auth",
		"GET",
		"/token",
		auth.AuthTokenHandler,
	),
	_NewApiRoute(
		"api",
		"Courses",
		"GET",
		"/courses",
		courses.CoursesIndexHandler,
	),
}
