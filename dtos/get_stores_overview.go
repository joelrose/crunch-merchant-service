package dtos

type GetStoresOverviewResponse struct {
	Id                int     `db:"id"`
	Name              string  `db:"name"`
	Description       string  `db:"description"`
	Address           string  `db:"address"`
	AveragePickupTime int     `db:"average_pickup_time"`
	AverageReview     float64 `db:"average_review"`
	ReviewCount       int     `db:"review_count"`
	GoogleMapsLink    string  `db:"google_maps_link"`
	PhoneNumber       string  `db:"phone_number"`
	ImageUrl          string  `db:"image_url"`
} //@name GetStoresOverviewResponse
