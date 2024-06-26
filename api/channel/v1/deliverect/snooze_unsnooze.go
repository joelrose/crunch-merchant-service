package deliverect

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func convertToString(uuidArray []uuid.UUID) []string {
	var str []string
	for _, value := range uuidArray {
		str = append(str, value.String())
	}
	return str
}

func SnoozeUnsnooze(c echo.Context) error {
	var request dtos.SnoozeUnsnzoozeRequestDto
	err := c.Bind(&request)
	if err != nil {
		log.Errorf("failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	channel, err := db.GetChannelByDeliverectLinkId(request.ChannelLinkId)
	if err != nil {
		log.Errorf("failed to get channel: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	response := dtos.SnoozeUnsnoozeResponseDto{
		Results: []dtos.SnoozeUnsnoozeResponseResult{},
	}

	for _, operation := range request.Operations {
		shouldSnooze := operation.Action == "snooze"

		for _, item := range operation.Data.Items {
			productIds, err := db.GetProductsByPlu(item.Plu, channel.StoreId)
			if err != nil {
				log.Errorf("failed to get products by plu: %v", err)
				return echo.NewHTTPError(http.StatusBadRequest)
			}

			if len(productIds) != 0 {
				err := db.UpdateProductsSnooze(productIds, shouldSnooze)
				if err != nil {
					log.Errorf("failed to update products snooze: %v", err)
					return echo.NewHTTPError(http.StatusBadRequest)
				}
			}

			result := dtos.SnoozeUnsnoozeResponseResult{
				Action: operation.Action,
				Data: dtos.SnoozeUnsnoozeResponseData{
					LocationId:      request.LocationId,
					AllSnoozedItems: convertToString(productIds),
				},
			}

			response.Results = append(response.Results, result)
		}

	}

	return c.JSON(http.StatusOK, response)
}
