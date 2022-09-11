package users

import (
	"log"
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := c.Get("dbContextKey").(*db.DB)

	users := []models.User{}

	err := db.Sqlx.Select(&users, "SELECT * FROM users")

	if err != nil {
		log.Fatal("failed to select users")
	}

	return c.JSON(http.StatusOK, users)
}
