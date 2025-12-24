package main

import (
	"log"
	"net/http"

	"github.com/AbhiramVijayan/order-producer/internal/handlers"
)

func main() {
	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/orders", handlers.CreateOrderHandler)
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	// Create HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("🚀 Order Producer API listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
