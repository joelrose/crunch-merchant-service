package db

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

func (database *DB) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.Sqlx.Select(&users, "SELECT * FROM users")

	return users, err
}

func (database *DB) GetUserByFirebaseId(firebaseId string) (models.User, error) {
	var user models.User
	err := database.Sqlx.Get(&user, "SELECT * FROM users WHERE firebase_id = $1", firebaseId)

	return user, err
}

func (database *DB) GetUserByUserId(userId uuid.UUID) (models.User, error) {
	var user models.User
	err := database.Sqlx.Get(&user, "SELECT * FROM users WHERE id = $1", userId)

	return user, err
}

func (database *DB) CreateUser(firebaseId string, user dtos.CreateUserRequest) error {
	_, err := database.Sqlx.Exec(
		"INSERT INTO users (firebase_id, language_code, firstname, lastname) VALUES ($1, $2, $3, $4)",
		firebaseId, user.LanguageCode, user.Firstname, user.Lastname,
	)

	return err
}
