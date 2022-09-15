package dtos

import "github.com/google/uuid"

type GetStoreRequest struct {
	StoreId uuid.UUID `param:"id"`
} //@name GetStoreRequest

type GetStoreCategory struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ImageUrl    string `db:"image_url"`
	SortOrder   int    `db:"sort_order"`
	Products    []GetStoreProduct
} //@name GetStoreCategory

type GetStoreProduct struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	Max         int    `db:"max"`
	Min         int    `db:"min"`
	Multiply    int    `db:"multiply"`
	Plu         string `db:"plu"`
	Snoozed     bool   `db:"snoozed"`
	Tax         int    `db:"tax"`
	ProductType int    `db:"product_type"`
	ImageUrl    string `db:"image_url"`
	SortOrder   int    `db:"sort_order"`
	Visible     bool   `db:"visible"`
	Products    []GetStoreProduct
} //@name GetStoreProduct

type GetStoreOpeningHour struct {
	DayOfWeek      int `db:"day_of_week"`
	StartTimestamp int `db:"start_timestamp"`
	EndTimestamp   int `db:"end_timestamp"`
} //@name GetStoreOpeningHour

type GetStoreResponse struct {
	Id                uuid.UUID `db:"id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	Address           string    `db:"address"`
	AveragePickupTime int       `db:"average_pickup_time"`
	AverageReview     float64   `db:"average_review"`
	ReviewCount       int       `db:"review_count"`
	GoogleMapsLink    string    `db:"google_maps_link"`
	PhoneNumber       string    `db:"phone_number"`
	ImageUrl          string    `db:"image_url"`
	Categories        []GetStoreCategory
	IsAvailable       bool
	OpeningHours      []GetStoreOpeningHour
} //@name GetStoreResponse
