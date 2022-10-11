package db

import (
	"github.com/joelrose/crunch-merchant-service/models"
)

func (db *DB) IsWhitelisted(identifier string) bool {
	var model models.Whitelist
	err := db.Sqlx.Get(&model, "SELECT * FROM whitelist WHERE identifier = $1", identifier)

	return err == nil
}
