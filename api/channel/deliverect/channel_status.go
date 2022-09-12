package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/dtos"
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
		return dtos.Inactive
	}
}

func DeliverectChannelStatus(c echo.Context) error {
	// Bind request body
	channelStatusRequest := dtos.ChannelStatusRequest{}

	err := c.Bind(&channelStatusRequest)
	if err != nil {
		log.Errorf("failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db := c.Get("db").(*db.DB)

	// Check if [ChannelLocationId=StoreId] exists
	_, err = db.GetStore(channelStatusRequest.ChannelLocationId)

	if err != nil {
		log.Errorf("failed to get store: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Check if channel exists
	_, err = db.GetChannelByStoreId(channelStatusRequest.ChannelLocationId)

	channelStatus := convertToEnum(channelStatusRequest.Status)

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

	response := dtos.ChannelStatusReponse{
		StatusUpdateURL:   "",
		MenuUpdateURL:     "",
		SnoozeUnsnoozeURL: "",
		BusyModeURL:       "",
		UpdatePrepTimeURL: "",
	}

	return c.JSON(http.StatusOK, response)
}
