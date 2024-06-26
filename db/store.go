package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

func (db *DB) GetStoreById(storeId uuid.UUID) (models.Store, error) {
	var store models.Store
	err := db.Sqlx.Get(&store, "SELECT * FROM stores WHERE id = $1", storeId)

	return store, err
}

func (db *DB) GetStoreByMerchantUserId(merchantUserId string) (uuid.UUID, error) {
	var storeId uuid.UUID
	err := db.Sqlx.Get(&storeId, "SELECT id FROM stores WHERE merchant_user_id = $1", merchantUserId)

	return storeId, err
}

func (db *DB) GetOpenStore(storeId uuid.UUID) (models.Store, error) {
	var store models.Store
	err := db.Sqlx.Get(&store, "SELECT * FROM stores WHERE id = $1 AND is_open = true", storeId)

	return store, err
}

func (db *DB) GetAvailableStore(storeId uuid.UUID, weekday time.Weekday, timestamp int) (models.Store, error) {
	storeQuery := `
	SELECT *
	FROM stores s
	WHERE s.id = $1 
	  AND is_open = true
	  AND EXISTS (
			SELECT o.id
			FROM store_opening_hours o
			WHERE s.id = o.store_id
			  AND o.day_of_week = $2
			  AND o.start_timestamp <= $3
			  AND o.end_timestamp >= $3
		);`

	var store models.Store
	err := db.Sqlx.Get(&store, storeQuery, storeId, weekday, timestamp)

	return store, err
}

func (db *DB) GetOpenStores() ([]dtos.GetStoresOverviewResponse, error) {
	storesQuery := `
	SELECT id, name, description, address, average_pickup_time, average_review, review_count, google_maps_link, phone_number, image_url
	FROM stores
	WHERE is_open = true`

	stores := []dtos.GetStoresOverviewResponse{}
	err := db.Sqlx.Select(&stores, storesQuery)

	return stores, err
}

func (db *DB) SetIsOpen(storeId uuid.UUID, isOpen bool) error {
	_, err := db.Sqlx.Exec(
		"UPDATE stores SET is_open = $1 WHERE id = $2",
		isOpen, storeId,
	)

	return err
}

func (db *DB) SetStoreImageUrl(storeId uuid.UUID, imageUrl string) error {
	_, err := db.Sqlx.Exec(
		"UPDATE stores SET image_url = $1 WHERE id = $2",
		imageUrl, storeId,
	)

	return err
}
