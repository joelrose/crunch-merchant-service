package db

import (
	"github.com/joelrose/crunch-merchant-service/db/dtos"
	"github.com/joelrose/crunch-merchant-service/db/models"
)

func (database *DB) GetAllUsers() ([]models.User, error) {
	users := []models.User{}

	err := database.Sqlx.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (database *DB) GetUserByFirebaseId(firebaseId string) (models.User, error) {
	user := models.User{}
	err := database.Sqlx.Get(&user, "SELECT * FROM users WHERE firebase_id = $1", firebaseId)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (database *DB) CreateUser(firebaseId string, user dtos.CreateUserRequest) error {
	_, err := database.Sqlx.Exec(
		"INSERT INTO users (firebase_id, language_code, firstname, lastname) VALUES ($1, $2, $3, $4)",
		firebaseId, user.LanguageCode, user.Firstname, user.Lastname,
	)

	return err
}
