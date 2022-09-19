package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
)

func (db *DB) CreateOrderItemWithoutParent(orderItem models.CreateOrderItem) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
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

func (db *DB) CreateOrderItemWithParent(orderItem models.CreateOrderItem) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.Get(
		&lastInsertId,
		"INSERT INTO order_items (plu, name, price, quantity, order_id, parent_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		orderItem.Plu,
		orderItem.Name,
		orderItem.Price,
		orderItem.Quantity,
		orderItem.OrderId,
		(*orderItem.ParentId),
	)

	return lastInsertId, err
}

func (database *DB) GetOrderItems(orderId uuid.UUID) ([]models.OrderItem, error) {
	orderItems := []models.OrderItem{}
	err := database.Sqlx.Select(&orderItems, "SELECT * FROM order_items WHERE order_id = $1", orderId)
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
