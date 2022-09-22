package dtos

import (
	"time"
)

type GetOrdersResponse struct {
	Id                  int         `db:"id" json:"id"`
	Status              int         `db:"status" json:"status"`
	Price               int         `db:"price" json:"price"`
	IsPaid              bool        `db:"is_paid" json:"isPaid"`
	EstimatedPickupTime time.Time   `db:"estimated_pickup_time" json:"estimatedPickupTime"`
	CreatedAt           time.Time   `db:"created_at" json:"createdAt"`
	StoreName           string      `db:"name" json:"storeName"`
	StoreDescription    string      `db:"description" json:"storeDescription"`
	StoreImageUrl       string      `db:"image_url" json:"storeImageUrl"`
	StoreAdress         string      `db:"address" json:"storeAddress"`
	StorePhoneNumber    string      `db:"phone_number" json:"storePhoneNumber"`
	StoreGoogleMapsLink string      `db:"google_maps_link" json:"googleMapsLink"`
	OrderItems          []OrderItem `db:"order_items" json:"orderItems"`
} //@name GetOrdersResponse
