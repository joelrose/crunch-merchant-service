package users

import (
	"database/sql"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	
)

// GetUser godoc
// @Summary      Get User associated with the given Firebase Token
// @Tags         users
// @Accept       json
// @Produce      json
// @Security FirebaseToken
// @Success      200  {object}  models.User
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /users [get]
func GetUser(c echo.Context) error {
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)
	token := c.Get(middleware.FIREBASE_CONTEXT_KEY).(*auth.Token)

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
