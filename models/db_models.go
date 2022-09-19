package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrder struct {
	Status              int       `db:"status"`
	EstimatedPickupTime time.Time `db:"estimated_pickup_time"`
	Price               int       `db:"price"`
	StripeOrderId       string    `db:"stripe_order_id"`
	IsPaid              bool      `db:"is_paid"`
	StoreId             uuid.UUID `db:"store_id"`
	UserId              int       `db:"user_id"`
	Fee                 float32   `db:"fee"`
}

type CreateOrderItem struct {
	Plu      string     `db:"plu"`
	Name     string     `db:"name"`
	Price    int        `db:"price"`
	Quantity int        `db:"quantity"`
	OrderId  int        `db:"order_id"`
	ParentId *uuid.UUID `db:"parent_id"`
}
