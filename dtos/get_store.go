package dtos

import "github.com/google/uuid"

type GetStoreRequest struct {
	StoreId uuid.UUID `param:"id"`
} //@name GetStoreRequest

type GetStoreCategory struct {
	Id          uuid.UUID         `db:"id" json:"id"`
	Name        string            `db:"name" json:"name"`
	Description string            `db:"description" json:"description"`
	ImageUrl    string            `db:"image_url" json:"imageUrl"`
	SortOrder   int               `db:"sort_order" json:"sortOrder"`
	Products    []GetStoreProduct `json:"products"`
} //@name GetStoreCategory

type GetStoreProduct struct {
	Id          uuid.UUID         `db:"id" json:"id"`
	Name        string            `db:"name" json:"name"`
	Description string            `db:"description" json:"description"`
	Price       int               `db:"price" json:"price"`
	Max         int               `db:"max" json:"max"`
	Min         int               `db:"min" json:"min"`
	Multiply    int               `db:"multiply" json:"multiply"`
	MultiMax    int               `db:"multi_max" json:"multiMax"`
	Plu         string            `db:"plu" json:"plu"`
	Snoozed     bool              `db:"snoozed" json:"snoozed"`
	Tax         int               `db:"tax" json:"tax"`
	ProductType int               `db:"product_type" json:"productType"`
	ImageUrl    string            `db:"image_url" json:"imageUrl"`
	SortOrder   int               `db:"sort_order" json:"sortOrder"`
	Visible     bool              `db:"visible" json:"visible"`
	Products    []GetStoreProduct `json:"products"`
} //@name GetStoreProduct

type GetStoreOpeningHour struct {
	DayOfWeek      int `db:"day_of_week" json:"dayOfWeek"`
	StartTimestamp int `db:"start_timestamp" json:"startTimestamp"`
	EndTimestamp   int `db:"end_timestamp" json:"endTimestamp"`
} //@name GetStoreOpeningHour

type GetStoreResponse struct {
	Id                uuid.UUID             `db:"id" json:"id"`
	Name              string                `db:"name" json:"name"`
	Description       string                `db:"description" json:"description"`
	Address           string                `db:"address" json:"address"`
	AveragePickupTime int                   `db:"average_pickup_time" json:"averagePickupTime"`
	AverageReview     float32               `db:"average_review" json:"averageReview"`
	ReviewCount       int                   `db:"review_count" json:"reviewCount"`
	GoogleMapsLink    string                `db:"google_maps_link" json:"googleMapsLink"`
	PhoneNumber       string                `db:"phone_number" json:"phoneNumber"`
	ImageUrl          string                `db:"image_url" json:"imageUrl"`
	Categories        []GetStoreCategory    `json:"categories"`
	IsAvailable       bool                  `json:"isAvailable"`
	OpeningHours      []GetStoreOpeningHour `json:"openingHours"`
} //@name GetStoreResponse
