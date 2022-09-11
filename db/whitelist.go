package db

import (
	"database/sql"

	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/labstack/gommon/log"
)

func (db *DB) IsWhitelisted(identifier string) bool {
	model := models.Whitelist{}
	err := db.Sqlx.Get(&model, "SELECT * FROM whitelist WHERE identifier = $1", identifier)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		log.Errorf("error getting user by id: %v", err)
		return false
	}

	return true
}
