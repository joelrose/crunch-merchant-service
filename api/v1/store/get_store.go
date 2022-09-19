package store

import (
	"context"
	"encoding/json"
	"fmt"
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
// @Router       /store/{id} [get]
func GetStore(c echo.Context) error {
	ctx := context.Background()
	db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

	r := dtos.GetStoreRequest{}
	err := c.Bind(&r)
	if err != nil {
		log.Debugf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	store, err := db.GetOpenStore(r.StoreId)
	if err != nil {
		log.Debugf("failed to get store: %v", err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	rdb := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)

	value, err := rdb.Get(ctx, fmt.Sprint(r.StoreId)).Result()
	if err == nil {
		var menu menus.MenuRedisModel
		err := json.Unmarshal([]byte(value), &menu)
		if err != nil {
			log.Errorf("failed to marshal menu model: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		log.Debug("serving cached menu")
		return c.JSON(http.StatusOK, menus.ConvertToGetStoreResponse(store, &menu))
	} else {
		log.Debugf("rebuilding menu: %v", err)
	}

	menu, err := menus.Build(db, r.StoreId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	json, err := json.Marshal(menu)
	if err != nil {
		log.Errorf("failed to marshal menu: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = rdb.Set(ctx, fmt.Sprint(r.StoreId), json, 0).Err()
	if err != nil {
		log.Errorf("failed to save menu to redis: %v", err)
	}

	return c.JSON(http.StatusOK, menus.ConvertToGetStoreResponse(store, menu))
}
