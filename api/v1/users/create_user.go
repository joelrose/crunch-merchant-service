package users

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// CreateUser godoc
// @Summary      Create User associated with the given Firebase Token
// @Tags         users
// @Accept       json
// @Produce      json
// @Security 	 FirebaseToken
// @Param 		 request body dtos.CreateUserRequest true "body"
// @Success      201
// @Failure      400  {object}  error
// @Failure      401  {object}  error
// @Failure      409  {object}  error
// @Failure      500  {object}  error
// @Router       /users [post]
func CreateUser(c echo.Context) error {
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)
	token := c.Get(middleware.FIREBASE_CONTEXT_KEY).(*auth.Token)

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
