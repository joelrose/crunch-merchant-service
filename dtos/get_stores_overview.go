package dtos

import "github.com/google/uuid"

type GetStoresOverviewResponse struct {
	Id                uuid.UUID `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	Description       string    `db:"description" json:"description"`
	Address           string    `db:"address" json:"address"`
	AveragePickupTime int       `db:"average_pickup_time" json:"averagePickupTime"`
	AverageReview     float32   `db:"average_review" json:"averageReview"`
	ReviewCount       int       `db:"review_count" json:"reviewCount"`
	GoogleMapsLink    string    `db:"google_maps_link" json:"googleMapsLink"`
	PhoneNumber       string    `db:"phone_number" json:"phoneNumber"`
	ImageUrl          string    `db:"image_url" json:"imageUrl"`
} //@name GetStoresOverviewResponse
