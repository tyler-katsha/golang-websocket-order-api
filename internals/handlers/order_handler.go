package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tyler-katsha/golang-websocket-order-api/internals"
	"uuid"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	var order internals.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w,"Invalid JSON Body",http.StatusBadRequest)
		return
	}

	order.OrderID = fmt.Sprintf("ORD-%d", time.Now().UnixNano() % 100_000)
	order.Status = "Pending"
	order.PurchasedAt = time.Now().Format("15:04:05")

	for i := range order.Products {
		if order.Products[i].ProductId == uuid.Nil() {
			order.Products[i].ProductId = uuid.New()
		}
	}
	BroadcastOrders(order)

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)

}