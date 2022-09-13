package menus

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetMenu(c echo.Context) error {
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

	categories, err := db.GetCategories(r.StoreId)
	if err != nil {
		log.Errorf("failed to get categories: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for ind, category := range categories {
		childrenProductIds, err := db.GetCategoryChildren(category.Id)

		if err != nil {
			log.Errorf("failed to get category children: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		categories[ind].ProductChildren = childrenProductIds
	}

	products, err := db.GetProducts(r.StoreId)
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for ind, product := range products {
		childrenProductIds, err := db.GetProductChildren(product.Id)

		if err != nil {
			log.Errorf("failed to get product children: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		products[ind].ProductChildren = childrenProductIds
	}

	openingHours, err := db.GetOpeningHours(r.StoreId)
	if err != nil {
		log.Errorf("failed to get opening hours: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, dtos.GetMenuResponse{
		Categories:   categories,
		Products:     products,
		OpeningHours: openingHours,
	})
}
