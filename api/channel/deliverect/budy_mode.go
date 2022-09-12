package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func DeliverectBusyMode(c echo.Context) error {
	db := c.Get("db").(*db.DB)

	// Bind request body
	busyModeRequest := dtos.BusyModeRequest{}

	err := c.Bind(&busyModeRequest)
	if err != nil {
		log.Errorf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	channel, err := db.GetChannelByDeliverectLinkId(busyModeRequest.ChannelLinkId)
	if err != nil {
		log.Errorf("failed to get channel: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	isOpen := busyModeRequest.Status == "ONLINE"
	err = db.SetIsOpen(isOpen, channel.StoreId)
	if err != nil {
		log.Errorf("failed to set is_open: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, dtos.BusyModeResponse{
		Status: busyModeRequest.Status,
	})
}
