package main

import (
	"log"
	"net/http"

	"github.com/AbhiramVijayan/order-producer/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handlers.CreateOrderHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("🚀 Order Producer API listening on port 8080")

	log.Fatal(server.ListenAndServe())

}
