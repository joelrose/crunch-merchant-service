package db

import (
	"github.com/joelrose/crunch-merchant-service/db/dtos/deliverect"
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (db *DB) GetChannel(storeId int, deliverectLocationId string) (models.DeliverectChannel, error) {
	channel := models.DeliverectChannel{}
	err := db.Sqlx.Get(&channel, "SELECT * FROM deliverect_channels WHERE store_id = $1 AND deliverect_location_id = $2", storeId, deliverectLocationId)

	if err != nil {
		return models.DeliverectChannel{}, err
	}

	return channel, nil
}

func (db *DB) CreateChannel(storeId int, deliverectLocationId string, status deliverect.ChannelStatus) error {
	_, err := db.Sqlx.Exec("INSERT INTO deliverect_channels (store_id, deliverect_location_id, status) VALUES ($1, $2, $3)", storeId, deliverectLocationId, status)

	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateChannelStatus(status deliverect.ChannelStatus, storeId int, deliverectLocationId string) error {
	_, err := db.Sqlx.Exec(
		"UPDATE deliverect_channels SET status = $1 WHERE store_id = $2 AND deliverect_location_id = $3",
		status, storeId, deliverectLocationId,
	)

	if err != nil {
		return err
	}

	return nil
}
