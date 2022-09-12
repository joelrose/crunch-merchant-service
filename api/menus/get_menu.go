package menus

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetMenu(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	r := dtos.GetMenuRequest{}
	err := c.Bind(&r)
	if err != nil {
		log.Debug("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// check if menu exists

	categories, err := db.GetCategories(r.StoreId)
	if err != nil {
		log.Debugf("failed to get categories: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for _, category := range categories {
		childrenProductIds, err := db.GetCategoryChildren(category.Id)

		if err != nil {
			log.Errorf("failed to get category children: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		log.Debugf("childrenProductIds: %v", childrenProductIds)
	}

	products, err := db.GetProducts(r.StoreId)
	if err != nil {
		log.Debugf("failed to get products: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for _, product := range products {
		childrenProductIds, err := db.GetProductChildren(product.Id)

		if err != nil {
			log.Errorf("failed to get product children: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		log.Debugf("childrenProductIds: %v", childrenProductIds)
	}

	return c.JSON(http.StatusOK, dtos.GetMenuResponse{
		Categories: categories,
		Products:   products,
	})
}
