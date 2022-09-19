package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func PreparationTime(c echo.Context) error {
	request := dtos.PreparationTimeRequestDto{}
	err := c.Bind(&request)
	if err != nil {
		log.Errorf("failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
