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

func (database *DB) GetOrderByStripeOrderId(stripeOrderId string) (models.Order, error) {
	order := models.Order{}
	err := database.Sqlx.Get(&order, "SELECT * FROM orders WHERE stripe_order_id = $1", stripeOrderId)

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (database *DB) GetOrdersByUserId(userId int) ([]models.Order, error) {
	orders := []models.Order{}
	err := database.Sqlx.Select(&orders, "SELECT * FROM orders WHERE user_id = $1", userId)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (database *DB) GetOrderById(orderId int) (models.Order, error) {
	order := models.Order{}
	err := database.Sqlx.Get(&order, "SELECT * FROM orders WHERE id = $1", orderId)

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (db *DB) UpdateOrderStatus(orderId int, orderStatus int) error {
	_, err := db.Sqlx.Exec(
		"UPDATE orders SET status = $1 WHERE id = $2",
		orderId, orderStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (database *DB) MarkOrderAsPaid(orderId int) error {
	_, err := database.Sqlx.Exec("UPDATE orders SET is_paid=true WHERE id = $1", orderId)
	if err != nil {
		return err
	}

	return nil
}
