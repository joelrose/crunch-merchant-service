package users

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateUser(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	var userRequest dtos.CreateUserRequest
	err := c.Bind(&userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user model is invalid")
	}

	token := c.Get("token").(*auth.Token)

	_, err = db.GetUserByFirebaseId(token.UID)

	if err == nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	err = db.CreateUser(token.UID, userRequest)
	if err != nil {
		log.Errorf("failed to create user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}
