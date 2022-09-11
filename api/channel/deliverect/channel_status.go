package deliverect

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/dtos/deliverect"
	"github.com/labstack/echo/v4"
)

func convertToEnum(status string) deliverect.ChannelStatus {
	switch status {
	case "register":
		return deliverect.Register
	case "active":
		return deliverect.Active
	default:
		return deliverect.Inactive
	}
}

func DeliverectChannelStatus(c echo.Context) error {
	db := c.Get("dbContextKey").(*db.DB)

	channelStatusRequest := deliverect.ChannelStatusRequest{}
	err := c.Bind(&channelStatusRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "channel model is invalid")
	}

	// store_id = channelLocationId

	// Check if merchant with channel_location_id exists
	_, err = db.GetStore(channelStatusRequest.ChannelLocationId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "cannot find chanenl location")
	}
	/*
		// Check if Channel exists
		channel, err := db.GetChannel(channelStatusRequest.ChannelLocationId, channelStatusRequest.LocationId)

		channelStatus := convertToEnum(channelStatusRequest.Status)

		if err != nil {
			// Create new channel
			// INSERTR

			err := db.CreateChannel()
		} else {
			// UPDATE
		}*/

	/*

			var merchant = await _merchantRepository.GetById(Guid.Parse(request.Model.ChannelLocationId));

		            if (merchant == null)
		            {
		                throw new NotFoundException();
		            }

		            var channelStatus = request.Model.Status switch
		            {
		                "register" => ChannelStatus.Register,
		                "active" => ChannelStatus.Active,
		                "inactive" => ChannelStatus.Inactive,
		                _ => throw new BadRequestException()
		            };

		            if (merchant.Channel == null)
		            {
		                if (channelStatus != ChannelStatus.Register)
		                {
		                    throw new BadRequestException();
		                }

		                merchant.Channel = new ChannelModel
		                {
		                    ChannelLinkId = request.Model.ChannelLinkId,
		                    ChannelLocationId = request.Model.ChannelLocationId,
		                    Status = channelStatus,
		                    LocationId = request.Model.LocationId,
		                };
		            }
		            else
		            {
		                merchant.Channel.ChannelLocationId = request.Model.ChannelLocationId;
		                merchant.Channel.Status = channelStatus;
		                merchant.Channel.LocationId = request.Model.LocationId;
		            }

		            await _merchantRepository.Save();

		            var baseUrl = _handlerRepository.GetBaseUrl();

		            string BuildUrl(string mode) =>
		                $"{baseUrl}/api/deliverect/{mode}";

		            return new DeliverectWebhookUrlDto
		            {
		                StatusUpdateURL = BuildUrl("orderstatus"),
		                MenuUpdateURL = BuildUrl("menupush"),
		                SnoozeUnsnoozeURL = BuildUrl("snoozeUnsnoozeProduct"),
		                BusyModeURL = BuildUrl("busyMode"),
		                UpdatePrepTimeURL = BuildUrl("preparationTimeUpdate"),
		            };*/
	return nil
}
