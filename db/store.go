package db

import (
	"github.com/joelrose/crunch-merchant-service/db/dtos"
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (db *DB) GetStore(id int) (models.Store, error) {
	store := models.Store{}
	err := db.Sqlx.Get(&store, "SELECT * FROM stores WHERE id = $1", id)

	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (db *DB) GetAvailableStores(dayOfWeek int, timestamp int) ([]dtos.GetStoresResponse, error) {
	storeQuery := `
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

	stores := []dtos.GetStoresResponse{}
	err := db.Sqlx.Select(&stores, storeQuery, dayOfWeek, timestamp)

	if err != nil {
		return []dtos.GetStoresResponse{}, err
	}

	return stores, nil
}

func (db *DB) SetIsOpen(isOpen bool, id int) error {
	_, err := db.Sqlx.Exec(
		"UPDATE stores SET is_open = $1 WHERE id = $2",
		isOpen, id,
	)

	if err != nil {
		return err
	}

	return nil
}
