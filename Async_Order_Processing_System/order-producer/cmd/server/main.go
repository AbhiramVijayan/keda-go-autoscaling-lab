package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Producer API is running "))
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("🚀 Order Producer API listening on port 8080")

	log.Fatal(server.ListenAndServe())

}
