package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AbhiramVijayan/order-producer/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/orders", handlers.CreateOrderHandler)
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	// Create HTTP server
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("🚀 Order Producer API listening on port " + port)
	log.Fatal(server.ListenAndServe())
}
