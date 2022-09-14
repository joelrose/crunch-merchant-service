package dtos

import "github.com/google/uuid"

type CreateOrderRequest struct {
	StoreId    uuid.UUID   `json:"storeId"`
	OrderItems []OrderItem `json:"orderItems"`
}

type OrderItem struct {
	Id       int         `json:"id"`
	Plu      string      `json:"plu"`
	Name     string      `json:"name"`
	Price    int         `json:"price"`
	Quantity int         `json:"quantity"`
	SubItems []OrderItem `json:"subItems"`
}

type CreateOrderResponse struct {
	StripeClientSecret string `json:"stripeClientSecret"`
}
