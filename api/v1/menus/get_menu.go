package menus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func buildMenu(db *db.DB, storeId uuid.UUID) (*dtos.GetMenuResponse, error) {
	categories, err := db.GetCategories(storeId)
	if err != nil {
		log.Errorf("failed to get categories: %v", err)
		return nil, err
	}

	for ind, category := range categories {
		childrenProductIds, err := db.GetCategoryChildren(category.Id)

		if err != nil {
			log.Errorf("failed to get category children: %v", err)
			return nil, err
		}

		categories[ind].ProductChildren = childrenProductIds
	}

	products, err := db.GetProducts(storeId)
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return nil, err
	}

	for ind, product := range products {
		childrenProductIds, err := db.GetProductChildren(product.Id)

		if err != nil {
			log.Errorf("failed to get product children: %v", err)
			return nil, err
		}

		products[ind].ProductChildren = childrenProductIds
	}

	openingHours, err := db.GetOpeningHours(storeId)
	if err != nil {
		log.Errorf("failed to get opening hours: %v", err)
		return nil, err
	}

	return &dtos.GetMenuResponse{
		Categories:   categories,
		Products:     products,
		OpeningHours: openingHours,
	}, nil
}

func GetMenu(c echo.Context) error {
	ctx := context.Background()
	db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

	r := dtos.GetMenuRequest{}
	err := c.Bind(&r)
	if err != nil {
		log.Debug("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err = db.GetAvailableStore(r.StoreId)
	if err != nil {
		log.Debugf("failed to get store: %v", err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	rdb := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)

	value, err := rdb.Get(ctx, fmt.Sprint(r.StoreId)).Result()
	if err == nil {
		log.Debug("serving cached menu")
		return c.Blob(http.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, []byte(value))
	} else {
		log.Debugf("rebuilding menu: %v", err)
	}

	menu, err := buildMenu(db, r.StoreId)
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

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, []byte(json))
}
