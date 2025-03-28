package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// Set up ans app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print out a message
	log.Println("Starting server on port 8080...	")

	// start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
