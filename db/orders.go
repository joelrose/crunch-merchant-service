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

func (database *DB) MarkOrderAsPaid(orderId int) error {
	_, err := database.Sqlx.Exec("UPDATE orders SET is_paid=true WHERE id = $1", orderId)
	if err != nil {
		return err
	}

	return nil
}
