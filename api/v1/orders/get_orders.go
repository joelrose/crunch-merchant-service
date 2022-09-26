package orders

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetOrders godoc
// @Summary      Get all orders from a user
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security 	 FirebaseToken
// @Success      200  {object}  []dtos.GetOrdersResponse
// @Success      400  {object} 	error
// @Failure      500  {object}  error
// @Router       /orders [get]
func GetOrders(c echo.Context) error {
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)
	token := c.Get(middleware.FIREBASE_CONTEXT_KEY).(*auth.Token)

	user, err := db.GetUserByFirebaseId(token.UID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	orders, err := db.GetOrdersByUserId(user.Id)
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
