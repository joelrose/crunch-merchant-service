package stores

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/utils"
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
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	stores, err := db.GetAvailableStores()
	if err != nil {
		log.Errorf("failed to get stores: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for ind, store := range stores {
		openingHours, err := db.GetOpeningHours(store.Id)
		if err != nil {
			log.Debugf("failed to get opening hours: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		stores[ind].IsAvailable = utils.IsStoreAvailable(openingHours)
	}

	return c.JSON(http.StatusOK, stores)
}
