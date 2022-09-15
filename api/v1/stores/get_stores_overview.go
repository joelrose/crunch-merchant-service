package stores

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetStoresOverview godoc
// @Summary      Get all open stores
// @Tags         stores
// @Accept       json
// @Produce      json
// @Security 	 FirebaseToken
// @Success      200  {object}  []dtos.GetStoresOverviewResponse
// @Failure      500  {object}  error
// @Router       /stores [get]
func GetStoresOverview(c echo.Context) error {
	db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

	stores, err := db.GetAvailableStores()

	if err != nil {
		log.Errorf("failed to get stores: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, stores)
}
