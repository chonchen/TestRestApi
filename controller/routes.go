package controller

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/todos", todo)
	router.HandleFunc("/todos/{id}", todoShow)
	return router
}
