package whitelist

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
)

type WhitelistRequest struct {
	Id string `json:"identifier"`
}

func IsWhitelisted(c echo.Context) error {
	db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

	request := WhitelistRequest{}

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	isWhitelisted := db.IsWhitelisted(request.Id)

	return c.JSON(http.StatusOK, isWhitelisted)
}
