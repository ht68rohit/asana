package main

import (
	route "github.com/heaptracetechnology/microservice-asana/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
