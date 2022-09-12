package stores

import (
	"net/http"
	"time"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetStores(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	time := time.Now()

	timestamp := utils.ConvertToTimestamp(time.Hour(), time.Minute())

	stores, err := db.GetAvailableStores(int(time.Weekday()), timestamp)

	if err != nil {
		log.Errorf("failed to get stores: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, stores)
}
