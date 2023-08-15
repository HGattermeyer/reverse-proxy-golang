package main

import (
	"fmt"
	"log"
	"net/http"
	"proxy-golang/internal/data"
	"proxy-golang/internal/routes"
)

func main() {
	// Initialize DB
	fmt.Println("Initializing DB")
	db, err := data.InitializeDb()
	if err != nil {
		log.Fatal("Error initializing DB: ", err)
	}

	// Setup the Router from Gorilla Mux
	router := routes.SetupRouter(db)
	http.Handle("/", router)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
