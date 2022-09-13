package users

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateUser(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	token := c.Get("token").(*auth.Token)

	_, err := db.GetUserByFirebaseId(token.UID)

	if err == nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	var userRequest dtos.CreateUserRequest
	if c.Bind(&userRequest) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user model is invalid")
	}

	err = db.CreateUser(token.UID, userRequest)
	if err != nil {
		log.Errorf("failed to create user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}
