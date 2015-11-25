package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route is a struct representing all the information mux needs to route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is the interface for our router
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Checkin",
		"POST",
		"/checkin",
		Checkin,
	},
	Route{
		"Current",
		"GET",
		"/current",
		Current,
	},
}

//NewRouter controls routing for the server
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
