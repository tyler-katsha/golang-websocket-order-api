package internals

import "uuid"


type Product struct {
	ProductId uuid.UUID `json:"product_id"`
	Name string `json:"name"`
	Price float64 `json:"price"`

}

type Order struct {
	OrderID string `json:"order_id"`
	Products []Product `json:"product"`
	PurchasedAt string `json:"purchased_at"`
	Status string `json:"status"`
	Quantity int `json:"quantity"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
	TimeStamp string `json:"timestamp,omitempty"`
}