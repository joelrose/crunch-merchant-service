package dtos

import "time"

type GetOrdersResponse struct {
	Id                  int         `json:"id"`
	Status              int         `json:"status"`
	Price               int         `json:"price"`
	IsPaid              bool        `json:"is_paid"`
	EstimatedPickupTime time.Time   `json:"estimated_pickup_time"`
	CreatedAt           time.Time   `json:"created_at"`
	OrderItems          []OrderItem `json:"order_items"`
	StoreId             int         `json:"merchant_id"`
}
