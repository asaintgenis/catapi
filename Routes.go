package main

import "net/http"

//Route route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CatsIndex",
		"GET",
		"/cats",
		CatsIndex,
	},
	Route{
		"CatShow",
		"GET",
		"/cats/{catId}",
		CatShow,
	},
}
