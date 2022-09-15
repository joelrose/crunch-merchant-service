package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/joelrose/crunch-merchant-service/utils"
)

func (db *DB) GetStore(id uuid.UUID) (models.Store, error) {
	store := models.Store{}
	err := db.Sqlx.Get(&store, "SELECT * FROM stores WHERE id = $1", id)

	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (db *DB) GetOpenStores(id uuid.UUID) (models.Store, error) {
	store := models.Store{}
	err := db.Sqlx.Get(&store, "SELECT * FROM stores WHERE id = $1 AND is_open = true", id)

	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (db *DB) GetAvailableStore(id uuid.UUID) (models.Store, error) {
	storeQuery := `
	SELECT *
	FROM stores s
	WHERE s.id = $1 
	  AND is_open = true
	  AND EXISTS AS (
			SELECT o.id
			FROM store_opening_hours o
			WHERE s.id = o.store_id
			  AND o.day_of_week = $2
			  AND o.start_timestamp <= $3
			  AND o.end_timestamp >= $3
		);`

	weekday, timestamp := utils.GetDayAndTimestamp()

	store := models.Store{}
	err := db.Sqlx.Get(&store, storeQuery, id, weekday, timestamp)

	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (db *DB) GetAvailableStores() ([]dtos.GetStoresOverviewResponse, error) {
	storesQuery := `
	SELECT id, name, description, address, average_pickup_time, average_review, review_count, google_maps_link, phone_number, image_url
	FROM stores s
	WHERE is_open = true
	  AND EXISTS (
			SELECT o.id
			FROM store_opening_hours o
			WHERE s.id = o.store_id
			  AND o.day_of_week = $1
			  AND o.start_timestamp <= $2
			  AND o.end_timestamp >= $2
		);`

	weekday, timestamp := utils.GetDayAndTimestamp()

	stores := []dtos.GetStoresOverviewResponse{}
	err := db.Sqlx.Select(&stores, storesQuery, weekday, timestamp)

	if err != nil {
		return []dtos.GetStoresOverviewResponse{}, err
	}

	return stores, nil
}

func (db *DB) SetIsOpen(isOpen bool, id uuid.UUID) error {
	_, err := db.Sqlx.Exec(
		"UPDATE stores SET is_open = $1 WHERE id = $2",
		isOpen, id,
	)

	if err != nil {
		return err
	}

	return nil
}
