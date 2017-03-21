package routing

import (
	"net/http"

	"github.com/asaintgenis/catapi/app"
)

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
		app.Index,
	},
	Route{
		"CatsIndex",
		"GET",
		"/cats",
		app.CatsIndex,
	},
	Route{
		"CatShow",
		"GET",
		"/cats/{catId}",
		app.CatShow,
	},
}
