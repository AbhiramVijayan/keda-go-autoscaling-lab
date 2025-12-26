package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type OrderRequest struct {
	OrderID  string `json:"order_id"`
	UserID   string `json:"user_id"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

type OrderResponse struct {
	Status  string `json:"status"`
	OrderID string `json:"order_id"`
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		logger.Error("Invalid method for CreateOrderHandler", "request_uri", r.RequestURI, "method", r.Method)
		// logger.Error("Invalid method for CreateOrderHandler", "method", r.Method)
		return
	}
	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		logger.Error("invalid JSON payload", "request_uri", r.RequestURI, "error", err)
		return
	}

	if req.OrderID == "" || req.UserID == "" || req.Item == "" || req.Quantity <= 0 {
		http.Error(w, "invalid request fields", http.StatusBadRequest)
		logger.Error("invalid request fields", "request_uri", r.RequestURI, "request", req)
		return
	}
	_ = time.Now().UTC()
	resp := OrderResponse{
		Status:  "Order Created Successfully",
		OrderID: req.OrderID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
	logger.Info("Order created", "request_uri", r.RequestURI, "order_id", req.OrderID, "user_id", req.UserID)

}
