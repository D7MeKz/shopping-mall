package main

import (
	"log"
	"net/http"
)

func main() {

	// Go handler
	http.HandleFunc(
		"/setup/user",
		SetupUserHandler,
	)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
