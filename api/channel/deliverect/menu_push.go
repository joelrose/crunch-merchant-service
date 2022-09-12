package deliverect

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/dtos"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/maps"
)

func convertTimestamp(time string) int {
	splitTime := strings.Split(time, ":")

	hour, hErr := strconv.Atoi(splitTime[0])
	minute, mErr := strconv.Atoi(splitTime[1])

	if hErr != nil || mErr != nil {
		log.Errorf("failed to convert time to timestamp: %v", time)
		return 0
	}

	return utils.ConvertToTimestamp(hour, minute)
}

func DeliverectMenuPush(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	// Bind request body
	d := dtos.MenuPushRequest{}

	err := c.Bind(&d)
	if err != nil {
		log.Errorf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	menu := d[0]

	channel, err := db.GetChannelByDeliverectLinkId(menu.ChannelLinkID)
	if err != nil {
		log.Errorf("failed to get channel: %v", err)
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

	products := map[string]dtos.DeliverectMenuProduct{}

	maps.Copy(products, menu.Products)
	maps.Copy(products, menu.ModifierGroups)
	maps.Copy(products, menu.Modifiers)
	maps.Copy(products, menu.Bundles)

	// Save availabilities
	for _, openingHour := range menu.StoreOpeningHours {
		strings.Split(openingHour.StartTime, ":")

		openingHour := models.StoreOpeningHour{
			StoreId:        channel.StoreId,
			DayOfWeek:      openingHour.DayOfWeek,
			StartTimestamp: convertTimestamp(openingHour.StartTime),
			EndTimestamp:   convertTimestamp(openingHour.EndTime),
		}

		err = db.CreateStoreOpeningHour(openingHour)
		if err != nil {
			log.Errorf("failed to create opening hour: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	// Save products
	productIds := make(map[string]int)
	for _, product := range products {
		productModel := models.MenuProduct{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Max:         product.Max,
			Min:         product.Min,
			Multiply:    product.Multiply,
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

		for _, subProductId := range product.SubProducts {
			err := db.CreateProductRelation(productIds[subProductId], productIds[product.Id])

			if err != nil {
				log.Errorf("failed to save product relation: %v", err)
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

	return c.NoContent(http.StatusOK)
}
