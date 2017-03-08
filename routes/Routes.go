package routes

import (
	"fmt"
	"net/http"

	"app/cources"
	"app/logger"

	"github.com/gorilla/mux"
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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"Cources",
		"GET",
		"/cources",
		cources.CourcesIndexHandler,
	},
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ed platform rest api home")
}
