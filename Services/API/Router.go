package apiService

import (
	"github.com/gorilla/mux"
)

func NewRouter(model dao) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes(model) {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
