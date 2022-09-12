package db

import (
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (db *DB) CreateStoreOpeningHour(openingHour models.StoreOpeningHour) error {
	_, err := db.Sqlx.NamedExec(
		"INSERT INTO store_opening_hours (store_id, day_of_week, start_timestamp, end_timestamp) VALUES (:store_id, :day_of_week, :start_timestamp, :end_timestamp)",
		openingHour,
	)

	return err
}

func (db *DB) DeleteOpeningHours(storeId int) error {
	_, err := db.Sqlx.Exec("DELETE FROM store_opening_hours WHERE store_id = $1", storeId)

	return err
}
