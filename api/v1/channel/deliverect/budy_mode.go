package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func BusyMode(c echo.Context) error {
	var request dtos.BusyModeRequest
	err := c.Bind(&request)
	if err != nil {
		log.Errorf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	channel, err := db.GetChannelByDeliverectLinkId(request.ChannelLinkId)
	if err != nil {
		log.Errorf("failed to get channel: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	isOpen := request.Status == "ONLINE"
	err = db.SetIsOpen(channel.StoreId, isOpen)
	if err != nil {
		log.Errorf("failed to set is_open: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, dtos.BusyModeResponse{
		Status: request.Status,
	})
}
