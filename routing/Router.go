package routing

import (
	"net/http"

	"github.com/asaintgenis/catapi/utils"
	"github.com/gorilla/mux"
)

//NewRouter create a new router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = utils.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Handle("/", http.FileServer(http.Dir("./web/")))

	return router
}
