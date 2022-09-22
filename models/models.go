package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           int       `db:"id" json:"id"`
	FirebaseId   string    `db:"firebase_id" json:"firebaseId"`
	LanguageCode string    `db:"language_code" json:"languageCode"`
	Firstname    string    `db:"firstname" json:"firstname"`
	Lastname     string    `db:"lastname" json:"lastname"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
} //@name User

type Merchant struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

type OrderStatus int

const (
	Parsed OrderStatus = iota + 1
	Received
	New               = 10
	Accepted          = 20
	Duplicate         = 30
	Printed           = 40
	Preparing         = 50
	Prepared          = 60
	ReadyForPickup    = 70
	InDelivery        = 80
	Finalized         = 90
	AutoFinalized     = 95
	Cancel            = 100
	Canceled          = 110
	Failed            = 120
	PosReceivedFailed = 121
	ParseFailed       = 124
)

type Order struct {
	Id                  int         `db:"id"`
	Status              OrderStatus `db:"status"`
	EstimatedPickupTime time.Time   `db:"estimated_pickup_time"`
	Price               int         `db:"price"`
	StripeOrderId       string      `db:"stripe_order_id"`
	IsPaid              bool        `db:"is_paid"`
	CreatedAt           time.Time   `db:"created_at"`
	StoreId             uuid.UUID   `db:"store_id"`
	UserId              int         `db:"user_id"`
}

type OrderItem struct {
	Id       uuid.UUID     `db:"id"`
	Plu      string        `db:"plu"`
	Name     string        `db:"name"`
	Price    int           `db:"price"`
	Quantity int           `db:"quantity"`
	OrderId  int           `db:"order_id"`
	ParentId uuid.NullUUID `db:"parent_id"`
}

type Store struct {
	Id                  uuid.UUID      `db:"id"`
	Name                string         `db:"name"`
	Description         string         `db:"description"`
	Address             string         `db:"address"`
	AveragePickupTime   int            `db:"average_pickup_time"`
	AverageReview       float32        `db:"average_review"`
	ReviewCount         int            `db:"review_count"`
	GoogleMapsLink      string         `db:"google_maps_link"`
	PhoneNumber         string         `db:"phone_number"`
	StripeAccountId     sql.NullString `db:"stripe_account_id"`
	StripeAccountStatus sql.NullString `db:"stripe_account_status"`
	Fee                 float32        `db:"fee"`
	IsOpen              bool           `db:"is_open"`
	ImageUrl            string         `db:"image_url"`
	MerchantUserId      string         `db:"merchant_user_id"`
}

type StoreOpeningHour struct {
	Id             uuid.UUID    `db:"id"`
	DayOfWeek      time.Weekday `db:"day_of_week"`
	StartTimestamp int          `db:"start_timestamp"`
	EndTimestamp   int          `db:"end_timestamp"`
	StoreId        uuid.UUID    `db:"store_id"`
}

type DeliverectChannel struct {
	StoreId          uuid.UUID `db:"store_id"`
	DeliverectLinkId string    `db:"deliverect_link_id"`
	LocationId       string    `db:"location_id"`
	Status           int       `db:"status"`
}

type MenuCategory struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	ImageUrl    string    `db:"image_url"`
	SortOrder   int       `db:"sort_order"`
	StoreId     uuid.UUID `db:"store_id"`
}

type MenuProduct struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	Max         int       `db:"max"`
	Min         int       `db:"min"`
	Multiply    int       `db:"multiply"`
	MultiMax    int       `db:"multi_max"`
	Plu         string    `db:"plu"`
	Snoozed     bool      `db:"snoozed"`
	Tax         int       `db:"tax"`
	ProductType int       `db:"product_type"`
	ImageUrl    string    `db:"image_url"`
	SortOrder   int       `db:"sort_order"`
	Visible     bool      `db:"visible"`
	StoreId     uuid.UUID `db:"store_id"`
}

type ProductProductRelation struct {
	ParentProductId uuid.UUID `db:"parent_product_id"`
	ChildProductId  uuid.UUID `db:"child_product_id"`
}

type CategoryProductRelation struct {
	MenuCategoryId uuid.UUID `db:"menu_category_id"`
	MenuProductId  uuid.UUID `db:"menu_product_id"`
}

type Whitelist struct {
	Identifier string    `db:"identifier" json:"identifier"`
	CreatedAt  time.Time `db:"created_at"`
}
