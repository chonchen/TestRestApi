package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, world")
}

func todo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Todo Index!")
}

func todoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Fprintln(writer, "Todo show:", id)
}
