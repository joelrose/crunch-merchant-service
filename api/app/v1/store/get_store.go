package store

import (
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/domain/menus"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetStore godoc
// @Summary      Get menu for a given store
// @Tags         store
// @Accept       json
// @Produce      json
// @Security 	 FirebaseToken
// @Param id path string true "Id of the store"
// @Success      200  {object}  dtos.GetStoreResponse
// @Success      400  {object} 	error
// @Success      404  {object} 	error
// @Failure      500  {object}  error
// @Router       /app/v1/store/{id} [get]
func GetStore(c echo.Context) error {
	var request dtos.GetStoreRequest
	err := c.Bind(&request)
	if err != nil {
		log.Debugf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	store, err := db.GetOpenStore(request.StoreId)
	if err != nil {
		log.Debugf("failed to get store: %v", err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	rdb := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)
	menuService := menus.NewMenuService(db, rdb, request.StoreId)

	menu, err := menuService.GetMenu()
	if err != nil {
		log.Debugf("failed to get menu: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, menus.ConvertToGetStoreResponse(store, menu))
}
