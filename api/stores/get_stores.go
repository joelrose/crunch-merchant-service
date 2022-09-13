package stores

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetStores(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	stores, err := db.GetAvailableStores()

	if err != nil {
		log.Errorf("failed to get stores: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, stores)
}
