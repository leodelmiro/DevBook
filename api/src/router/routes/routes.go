package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, routeLogin)
	routes = append(routes, routesPosts...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
