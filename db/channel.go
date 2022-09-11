package db

import (
	"github.com/joelrose/crunch-merchant-service/db/dtos/deliverect"
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (db *DB) GetChannel(storeId int, deliverectLocationId string) (models.DeliverectChannel, error) {
	channel := models.DeliverectChannel{}
	err := db.Sqlx.Get(&channel, "SELECT * FROM deliverect_channel WHERE store_id = $1 AND deliverect_location_id = $2", storeId, deliverectLocationId)

	if err != nil {
		return models.DeliverectChannel{}, err
	}

	return channel, nil
}

func (db *DB) CreateChannel(storeId int, deliverectLocationId string, status deliverect.ChannelStatus) error {
	_, err := db.Sqlx.Exec("INSERT INTO deliverect_channel (store_id, deliverect_location_id, status) VALUES ($1, $2, $3)", storeId, deliverectLocationId, status)

	if err != nil {
		return err
	}
	return nil
}
