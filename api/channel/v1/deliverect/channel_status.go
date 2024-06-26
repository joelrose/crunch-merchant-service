package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func convertToEnum(status string) dtos.ChannelStatus {
	switch status {
	case "register":
		return dtos.Register
	case "active":
		return dtos.Active
	default:
		return dtos.InActive
	}
}

func ChannelStatus(c echo.Context) error {
	var channelStatusRequest dtos.ChannelStatusRequest
	err := c.Bind(&channelStatusRequest)
	if err != nil {
		log.Errorf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	// [ChannelLocationId=StoreId]
	_, err = db.GetStoreById(channelStatusRequest.ChannelLocationId)
	if err != nil {
		log.Errorf("failed to get store [%v]: %v", channelStatusRequest.ChannelLocationId, err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	channelStatus := convertToEnum(channelStatusRequest.Status)

	_, err = db.GetChannelByStoreId(channelStatusRequest.ChannelLocationId)
	if err != nil {
		log.Debugf("channel does not exist, creating channel %v", err)
		// Create new channel
		err := db.CreateChannel(
			channelStatusRequest.ChannelLocationId,
			channelStatusRequest.LocationId,
			channelStatusRequest.ChannelLinkId,
			channelStatus,
		)

		if err != nil {
			log.Errorf("error creating channel: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else {
		// Update existing channel
		err := db.UpdateChannelStatus(
			channelStatus,
			channelStatusRequest.ChannelLocationId,
		)

		if err != nil {
			log.Errorf("error updating channel status: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	requestHost := c.Request().Host
	buildUrl := func(path string) string {
		return "https://" + requestHost + "/v1/channel/deliverect/" + path
	}
	response := dtos.ChannelStatusReponse{
		MenuUpdateURL:        buildUrl("menu_push"),
		SnoozeUnsnoozeURL:    buildUrl("snooze_unsnooze"),
		BusyModeURL:          buildUrl("busy_mode"),
		UpdatePrepTimeURL:    buildUrl("prep_time"),
		OrderStatusUpdateURL: buildUrl("order_status"),
	}

	return c.JSON(http.StatusOK, response)
}
