package main

import (
	"log"
	"net/http"

	"github.com/asaintgenis/catapi/routing"
)

func main() {
	router := routing.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
