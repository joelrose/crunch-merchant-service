package dtos

type CreateOrderRequest struct {
	StoreId    int         `json:"storeId"`
	OrderItems []OrderItem `json:"orderItems"`
}

type OrderItem struct {
	Plu      int         `json:"plu"`
	Name     string      `json:"name"`
	Price    int         `json:"price"`
	Quantity int         `json:"quantity"`
	Children []OrderItem `json:"children"`
}
