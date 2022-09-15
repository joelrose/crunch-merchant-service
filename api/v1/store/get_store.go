package store

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	MenuModel struct {
		Categories   []dtos.GetStoreCategory
		Products     []dtos.GetStoreProduct
		OpeningHours []dtos.GetStoreOpeningHour
	}
)

func buildMenu(db *db.DB, storeId uuid.UUID) (*MenuModel, error) {
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

	return &MenuModel{
		Categories:   categories,
		Products:     products,
		OpeningHours: openingHours,
	}, nil
}

func buildStore(store models.Store, menu *MenuModel) dtos.GetStoreResponse {
	return dtos.GetStoreResponse{
		Id:                store.Id,
		Name:              store.Name,
		Description:       store.Description,
		Address:           store.Address,
		AveragePickupTime: store.AveragePickupTime,
		AverageReview:     store.AverageReview,
		ReviewCount:       store.ReviewCount,
		GoogleMapsLink:    store.GoogleMapsLink,
		PhoneNumber:       store.PhoneNumber,
		ImageUrl:          store.ImageUrl,
		Categories:        menu.Categories,
		Products:          menu.Products,
		OpeningHours:      menu.OpeningHours,
	}
}

// GetStore godoc
// @Summary      Get menu for a given store
// @Tags         store
// @Accept       json
// @Produce      json
// @Security 	 FirebaseToken
// @Param storeId path string true "Id of the store"
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
		log.Debug("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	store, err := db.GetAvailableStore(r.StoreId)
	if err != nil {
		log.Debugf("failed to get store: %v", err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	rdb := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)

	value, err := rdb.Get(ctx, fmt.Sprint(r.StoreId)).Result()
	if err == nil {
		var menu MenuModel
		err := json.Unmarshal([]byte(value), &menu)
		if err != nil {
			log.Errorf("failed to marshal menu model: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		log.Debug("serving cached menu")
		return c.JSON(http.StatusOK, buildStore(store, &menu))
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

	return c.JSON(http.StatusOK, buildStore(store, menu))
}
