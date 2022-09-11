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
