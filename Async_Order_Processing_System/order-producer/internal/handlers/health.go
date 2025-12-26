package handlers

import (
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logger.Error("Invalid method for HealthCheckHandler", "method", r.Method)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	logger.Info("Health check accessed", "request_uri", r.RequestURI, "method", r.Method)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy"}`))
}
