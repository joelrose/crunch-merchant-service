package whitelist

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
)

// Whitelist
// @Summary      Check if identifier is whitelisted
// @Description  get string by ID
// @Tags         whitelist
// @Accept       json
// @Produce      json
// @Param request body WhitelistRequest true "body"
// @Success      200  {object}  bool
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /whitelist [post]
func IsWhitelisted(c echo.Context) error {
	var request dtos.WhitelistRequest
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err = c.Validate(request); err != nil {
		return err
	}

	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)
	isWhitelisted := db.IsWhitelisted(request.Id)

	return c.JSON(http.StatusOK, isWhitelisted)
}
