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
		return
	}
	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	if req.OrderID == "" || req.UserID == "" || req.Item == "" || req.Quantity <= 0 {
		http.Error(w, "invalid request fields", http.StatusBadRequest)
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
}
