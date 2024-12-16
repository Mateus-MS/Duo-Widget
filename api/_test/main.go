package main

import (
	"log"
	"net/http"

	handler "github.com/Mateus-MS/Duo-Widget"
)

func main() {
	// Add the Vercel handler
	http.HandleFunc("/", handler.Handler)

	// Start the server locally on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
