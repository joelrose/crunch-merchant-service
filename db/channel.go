package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

func (db *DB) GetChannelByStoreId(storeId uuid.UUID) (models.DeliverectChannel, error) {
	channel := models.DeliverectChannel{}
	err := db.Sqlx.Get(&channel, "SELECT * FROM deliverect_channels WHERE store_id = $1", storeId)

	if err != nil {
		return models.DeliverectChannel{}, err
	}

	return channel, nil
}

func (db *DB) GetChannelByDeliverectLinkId(deliverectLinkId string) (models.DeliverectChannel, error) {
	channel := models.DeliverectChannel{}
	err := db.Sqlx.Get(&channel, "SELECT * FROM deliverect_channels WHERE deliverect_link_id = $1", deliverectLinkId)

	if err != nil {
		return models.DeliverectChannel{}, err
	}

	return channel, nil
}

func (db *DB) CreateChannel(storeId uuid.UUID, locationId string, deliverectChannelLinkId string, status dtos.ChannelStatus) error {
	_, err := db.Sqlx.Exec(
		"INSERT INTO deliverect_channels (store_id, location_id, deliverect_link_id, status) VALUES ($1, $2, $3, $4)",
		storeId, locationId, deliverectChannelLinkId, status,
	)

	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateChannelStatus(status dtos.ChannelStatus, storeId uuid.UUID) error {
	_, err := db.Sqlx.Exec(
		"UPDATE deliverect_channels SET status = $1 WHERE store_id = $2",
		status, storeId,
	)

	if err != nil {
		return err
	}

	return nil
}
