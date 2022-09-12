package db

import "github.com/joelrose/crunch-merchant-service/db/models"

func (db *DB) GetStore(id int) (models.Store, error) {
	store := models.Store{}
	err := db.Sqlx.Get(&store, "SELECT * FROM stores WHERE id = $1", id)

	if err != nil {
		return models.Store{}, err
	}

	return store, nil
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
