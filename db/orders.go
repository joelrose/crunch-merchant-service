package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

func (db *DB) CreateOrder(order models.CreateOrder) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO orders (status, estimated_pickup_time, price, stripe_order_id, is_paid, store_id, user_id, fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		order.Status,
		order.EstimatedPickupTime,
		order.Price,
		order.StripeOrderId,
		order.IsPaid,
		order.StoreId,
		order.UserId,
		order.Fee,
	)

	return lastInsertId, err
}

func (database *DB) GetOrderByStripeOrderId(stripeOrderId string) (models.Order, error) {
	var order models.Order
	err := database.Sqlx.Get(&order, "SELECT * FROM orders WHERE stripe_order_id = $1", stripeOrderId)

	return order, err
}

func (database *DB) GetOrdersByUserId(userId uuid.UUID) ([]dtos.GetOrdersResponse, error) {
	query := `
		SELECT o.id, status, price, is_paid, estimated_pickup_time, created_at, name, description, image_url, address, phone_number, google_maps_link
		FROM orders o
		LEFT JOIN stores m on o.store_id = m.id
		WHERE user_id = $1 AND is_paid = true
		ORDER BY created_at DESC`

	orders := []dtos.GetOrdersResponse{}
	err := database.Sqlx.Select(&orders, query, userId)

	return orders, err
}

func (database *DB) GetOrdersByStoreId(storeId uuid.UUID) ([]dtos.GetOrdersResponse, error) {
	query := `
		SELECT o.id, status, price, is_paid, estimated_pickup_time, created_at, name, description, image_url, address, phone_number, google_maps_link
		FROM orders o
		LEFT JOIN stores m on o.store_id = m.id
		WHERE store_id = $1
		ORDER BY created_at DESC`

	var orders []dtos.GetOrdersResponse
	err := database.Sqlx.Select(&orders, query, storeId)

	return orders, err
}

func (database *DB) GetOrderById(orderId uuid.UUID) (models.Order, error) {
	var order models.Order
	err := database.Sqlx.Get(&order, "SELECT * FROM orders WHERE id = $1", orderId)

	return order, err
}

func (db *DB) UpdateOrderStatus(orderId uuid.UUID, orderStatus models.OrderStatus) error {
	_, err := db.Sqlx.Exec(
		"UPDATE orders SET status = $1 WHERE id = $2",
		int(orderStatus), orderId,
	)

	return err
}

func (database *DB) MarkOrderAsPaid(orderId uuid.UUID) error {
	_, err := database.Sqlx.Exec("UPDATE orders SET is_paid=true WHERE id = $1", orderId)

	return err
}
