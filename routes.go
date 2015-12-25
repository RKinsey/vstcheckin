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
		IndexHandler,
	},
	Route{
		"Checkin",
		"POST",
		"/checkin",
		CheckinHandler,
	},
	Route{
		"Current",
		"GET",
		"/current",
		CurrentHandler,
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
	router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/",http.FileServer(http.Dir("./static/css/"))))
	return router
}
