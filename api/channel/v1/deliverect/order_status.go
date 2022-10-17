package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func OrderStatus(c echo.Context) error {
	var request dtos.UpdateOrderStatusRequest
	err := c.Bind(&request)
	if err != nil {
		log.Errorf("failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)
	order, err := db.GetOrderById(request.ChannelOrderId)
	if err != nil {
		log.Errorf("failed to get order: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = db.UpdateOrderStatus(order.Id, request.Status)
	if err != nil {
		log.Errorf("failed to update order status: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
