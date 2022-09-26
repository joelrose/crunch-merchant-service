package dashboard

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetOrders godoc
// @Summary      Get all orders from a store
// @Tags         dashboard
// @Accept       json
// @Produce      json
// @Security 	 Auth0Token
// @Success      200  {object}  []dtos.GetOrdersResponse
// @Success      400  {object} 	error
// @Failure      500  {object}  error
// @Router       /dashboard/orders [get]
func GetOrders(c echo.Context) error {
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	userId := c.Get(middleware.AUTH0_USER_ID_CONTEXT_KEY)
	userIdString := userId.(string)

	storeId, err := db.GetStoreByMerchantUserId(userIdString)
	if err != nil {
		log.Errorf("Failed to get store by merchant user id: %v", err)
		return echo.NewHTTPError(http.StatusForbidden)
	}

	orders, err := db.GetOrdersByStoreId(storeId)
	if err != nil {
		log.Errorf("failed to get orders: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for ind, order := range orders {
		orderItems, err := db.GetOrderItems(order.Id)

		if err != nil {
			log.Errorf("failed to get order items: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		orders[ind].OrderItems = utils.ConvertOrderItemsToDto(orderItems)
	}

	return c.JSON(http.StatusOK, orders)
}
