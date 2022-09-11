package users

import (
	"database/sql"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetUser(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	token := c.Get("token").(*auth.Token)

	users, err := db.GetUserByFirebaseId(token.UID)

	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		log.Errorf("failed to get user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, users)
}
