package controller

import (
	"net/http"

	"github.com/chonchen/TestRestApi/utils"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = utils.Logger(handler, route.Name)
		router.Methods(route.Method).Name(route.Name).Path(route.Pattern).Handler(route.HandlerFunc)
	}

	return router
}
