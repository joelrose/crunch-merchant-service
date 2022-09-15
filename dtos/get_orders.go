package dtos

import (
	"time"

	"github.com/google/uuid"
)

type GetOrdersResponse struct {
	Id                  int         `json:"id"`
	Status              int         `json:"status"`
	Price               int         `json:"price"`
	IsPaid              bool        `json:"is_paid"`
	EstimatedPickupTime time.Time   `json:"estimated_pickup_time"`
	CreatedAt           time.Time   `json:"created_at"`
	OrderItems          []OrderItem `json:"order_items"`
	StoreId             uuid.UUID   `json:"merchant_id"`
} //@name GetOrdersResponse
