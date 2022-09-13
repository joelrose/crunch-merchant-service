package db

import (
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (db *DB) CreateOrder(order models.CreateOrder) (int, error) {
	var lastInsertId int
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO orders (status, estimated_pickup_time, price, stripe_order_id, is_paid, store_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		order.Status,
		order.EstimatedPickupTime,
		order.Price,
		order.StripeOrderId,
		order.IsPaid,
		order.StoreId,
		order.UserId,
	)

	return lastInsertId, err
}

func (db *DB) CreateOrderItem(orderItem models.CreateOrderItem) (int, error) {
	var lastInsertId int
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO order_items (plu, name, price, quantity, order_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		orderItem.Plu,
		orderItem.Name,
		orderItem.Price,
		orderItem.Quantity,
		orderItem.OrderId,
	)

	return lastInsertId, err
}
