package models

import "time"

type CreateOrder struct {
	Status              int       `db:"status"`
	EstimatedPickupTime time.Time `db:"estimated_pickup_time"`
	Price               int       `db:"price"`
	StripeOrderId       string    `db:"stripe_order_id"`
	IsPaid              bool      `db:"is_paid"`
	StoreId             int       `db:"store_id"`
	UserId              int       `db:"user_id"`
}
