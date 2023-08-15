package main

import (
	"log"
	"net/http"
	"proxy-golang/internal/routes"
)

func main() {

	// Setup the Router from Gorilla Mux
	router := routes.SetupRouter()
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
