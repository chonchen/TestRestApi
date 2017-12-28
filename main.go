package main

import (
	"log"
	"net/http"

	"github.com/chonchen/TestRestApi/controller"
)

func main() {

	router := controller.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
