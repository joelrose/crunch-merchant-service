package models

import "time"

type User struct {
	Id           int       `db:"id"`
	FirebaseId   string    `db:"firebase_id"`
	LanguageCode string    `db:"language_code"`
	Firstname    string    `db:"firstname"`
	Lastname     string    `db:"lastname"`
	CreatedAt    time.Time `db:"created_at"`
}

type Merchant struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

type Order struct {
	Id                  int       `db:"id"`
	Status              int       `db:"status"`
	EstimatedPickupTime time.Time `db:"estimated_pickup_time"`
	Price               int       `db:"price"`
	StripeOrderId       string    `db:"stripe_order_id"`
	IsPaid              bool      `db:"is_paid"`
	CreatedAt           time.Time `db:"created_at"`
	MerchantId          int       `db:"merchant_id"`
	UserId              int       `db:"user_id"`
}

type OrderItem struct {
	Id        int       `db:"id"`
	Plu       string    `db:"plu"`
	Name      string    `db:"name"`
	Price     int       `db:"price"`
	Quantity  int       `db:"quantity"`
	OrderId   int       `db:"order_id"`
	ParentId  int       `db:"parent_id"`
	CreatedAt time.Time `db:"created_at"`
}

type Store struct {
	Id                  int       `db:"id"`
	Name                string    `db:"name"`
	Description         string    `db:"description"`
	Address             string    `db:"address"`
	AveragePickupTime   int       `db:"average_pickup_time"`
	AverageReview       float64   `db:"average_review"`
	ReviewCount         int       `db:"review_count"`
	GoogleMapsLink      string    `db:"google_maps_link"`
	PhoneNumber         string    `db:"phone_number"`
	StripeAccountId     string    `db:"stripe_account_id"`
	StripeAccountStatus int       `db:"stripe_account_status"`
	Fee                 float64   `db:"fee"`
	IsOpen              bool      `db:"is_open"`
	ImageUrl            string    `db:"image_url"`
	MerchantId          int       `db:"merchant_id"`
	CreatedAt           time.Time `db:"created_at"`
}

type StoreOpeningHours struct {
	Id             int       `db:"id"`
	DayOfWeek      int       `db:"day_of_week"`
	StartTimestamp int       `db:"start_timestamp"`
	EndTimestamp   int       `db:"end_timestamp"`
	StoreId        int       `db:"store_id"`
	CreatedAt      time.Time `db:"created_at"`
}

type DeliverectChannel struct {
	StoreId              int    `db:"store_id"`
	DeliverectLocationId string `db:"deliverect_location_id"`
	Status               int    `db:"status"`
}

type Menu struct {
	Id          int       `db:"id"`
	Description string    `db:"description"`
	Name        string    `db:"name"`
	ImageUrl    string    `db:"image_url"`
	CreatedAt   time.Time `db:"created_at"`
}

type MenuCategory struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	ImageUrl    string    `db:"image_url"`
	SortOrder   int       `db:"sort_order"`
	MenuId      int       `db:"menu_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type MenuProduct struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Plu         string    `db:"plu"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
	Snoozed     bool      `db:"snoozed"`
	Tax         int       `db:"tax"`
	ImageUrl    string    `db:"image_url"`
	Max         int       `db:"max"`
	Min         int       `db:"min"`
	Multiply    int       `db:"multiply"`
	SortOrder   int       `db:"sort_order"`
	Visible     bool      `db:"visible"`
	MenuId      int       `db:"menu_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type ProductProductRelation struct {
	ParentProductId int `db:"parent_product_id"`
	ChildProductId  int `db:"child_product_id"`
}

type CategoryProductRelation struct {
	MenuCategoryId int `db:"menu_category_id"`
	MenuProductId  int `db:"menu_product_id"`
}

type Whitelist struct {
	Identifier string    `db:"identifier" json:"identifier"`
	CreatedAt  time.Time `db:"created_at"`
}