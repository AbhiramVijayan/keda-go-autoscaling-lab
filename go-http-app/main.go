package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(" ok ✅\n"))
}

func workHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate some work
	time.Sleep(500 * time.Millisecond)

	hostname, _ := os.Hostname()
	response := fmt.Sprintf("Work done by pod : %s\n", hostname)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/work", workHandler)

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
