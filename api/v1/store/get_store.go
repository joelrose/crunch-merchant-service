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
	"github.com/joelrose/crunch-merchant-service/utils"
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

func convertHelperMenuCategory(db *db.DB, productMap map[int]dtos.GetStoreProduct, productId int) dtos.GetStoreProduct {
	childrenIds, err := db.GetProductChildren(productId)
	if err != nil {
		log.Errorf("failed to get product children: %v", err)
		return dtos.GetStoreProduct{}
	}

	product := productMap[productId]
	result := []dtos.GetStoreProduct{}
	for _, childId := range childrenIds {
		result = append(result, convertHelperMenuCategory(db, productMap, childId))
	}

	if len(result) > 0 {
		product.Products = result
	}

	return product
}

func buildMenu(db *db.DB, storeId uuid.UUID) (*MenuModel, error) {
	categories, err := db.GetCategories(storeId)
	if err != nil {
		log.Errorf("failed to get categories: %v", err)
		return nil, err
	}

	products, err := db.GetProducts(storeId)
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return nil, err
	}

	productMap := make(map[int]dtos.GetStoreProduct)
	for _, product := range products {
		productMap[product.Id] = product
	}

	for ind, category := range categories {
		childrenProductIds, err := db.GetCategoryChildren(category.Id)

		if err != nil {
			log.Errorf("failed to get category children: %v", err)
			return nil, err
		}

		children := []dtos.GetStoreProduct{}
		for _, childrenProductId := range childrenProductIds {
			children = append(children, convertHelperMenuCategory(db, productMap, childrenProductId))
		}

		if len(children) > 0 {
			categories[ind].Products = children
		}
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

func buildStore(store models.Store, menu *MenuModel, isAvailable bool) dtos.GetStoreResponse {
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
		IsAvailable:       isAvailable,
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

	store, err := db.GetOpenStores(r.StoreId)
	if err != nil {
		log.Debugf("failed to get store: %v", err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	openingHours, err := db.GetOpeningHours(store.Id)
	if err != nil {
		log.Debugf("failed to get opening hours: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	isAvailable := utils.IsStoreAvailable(openingHours)

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
		return c.JSON(http.StatusOK, buildStore(store, &menu, isAvailable))
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

	return c.JSON(http.StatusOK, buildStore(store, menu, isAvailable))
}
