package deliverect

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/maps"
)

func MenuPush(c echo.Context) error {
	var request dtos.MenuPushRequest
	err := c.Bind(&request)
	if err != nil {
		log.Debugf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if len(request) != 1 {
		log.Errorf("deliverect menu push request body must have exactly one element")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	menu := request[0]
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	channel, err := db.GetChannelByDeliverectLinkId(menu.ChannelLinkID)
	if err != nil {
		log.Debugf("failed to get channel: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Delete existing products
	pErr := db.DeleteProducts(channel.StoreId)
	hErr := db.DeleteOpeningHours(channel.StoreId)
	cErr := db.DeleteCategories(channel.StoreId)

	if pErr != nil || hErr != nil || cErr != nil {
		log.Errorf("failed to delete existing products: %v", channel.StoreId)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	productsLength := len(menu.Products)+len(menu.ModifierGroups)+len(menu.Modifiers)+len(menu.Bundles)
	products := make(map[string]dtos.DeliverectMenuProduct, productsLength)

	maps.Copy(products, menu.Products)
	maps.Copy(products, menu.ModifierGroups)
	maps.Copy(products, menu.Modifiers)
	maps.Copy(products, menu.Bundles)

	// Save availabilities
	for _, openingHour := range menu.StoreOpeningHours {
		strings.Split(openingHour.StartTime, ":")

		openingHour := models.StoreOpeningHour{
			StoreId:        channel.StoreId,
			DayOfWeek:      utils.ParseDeliverectDayOfWeek(openingHour.DayOfWeek),
			StartTimestamp: utils.ParseTimestamp(openingHour.StartTime),
			EndTimestamp:   utils.ParseTimestamp(openingHour.EndTime),
		}

		err = db.CreateStoreOpeningHour(openingHour)
		if err != nil {
			log.Errorf("failed to create opening hour: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	// Save products
	productIds := make(map[string]uuid.UUID, len(products))
	for _, product := range products {
		productModel := models.MenuProduct{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Max:         product.Max,
			Min:         product.Min,
			Multiply:    product.Multiply,
			MultiMax:    product.MultiMax,
			Plu:         product.Plu,
			Snoozed:     product.Snoozed,
			Tax:         product.Tax,
			ProductType: product.ProductType,
			ImageUrl:    product.ImageURL,
			SortOrder:   product.SortOrder,
			Visible:     product.Visible,
			StoreId:     channel.StoreId,
		}
		insertId, err := db.CreateProduct(productModel)

		if err != nil {
			log.Errorf("failed to save product: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		productIds[product.Id] = insertId
	}

	// Save production relations
	for _, product := range products {
		if product.SubProducts == nil {
			continue
		}

		for ind, subProductId := range product.SubProducts {
			childProductId := productIds[subProductId]
			err := db.CreateProductRelation(childProductId, productIds[product.Id])
			if err != nil {
				log.Errorf("failed to save product relation: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			err = db.UpdateProductSortOrder(childProductId, ind)
			if err != nil {
				log.Errorf("failed to update product sort order: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
	}

	// Save categories
	for ind, category := range menu.Categories {
		categoryId, err := db.CreateCategory(models.MenuCategory{
			Name:        category.Name,
			Description: category.Description,
			ImageUrl:    category.ImageURL,
			SortOrder:   ind,
			StoreId:     channel.StoreId,
		})

		if err != nil {
			log.Errorf("failed to save category: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		for _, productId := range category.Products {
			err := db.CreateProductCategoryRelation(categoryId, productIds[productId])

			if err != nil {
				log.Errorf("failed to save product category relation: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
	}

	// Save Store Details
	err = db.SetStoreImageUrl(channel.StoreId, menu.MenuImageURL)
	if err != nil {
		log.Errorf("failed to save store image url: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	rdb := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)
	cmd := rdb.Del(context.Background(), channel.StoreId.String())
	if cmd.Err() != nil {
		log.Errorf("failed to delete redis cache: %v", cmd)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
