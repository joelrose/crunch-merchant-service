package stripe

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/services/deliverect"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/webhook"
)

func WebhookHandler(c echo.Context) error {
	request := c.Request()
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		log.Errorf("failed to read request body: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	config := c.Get(middleware.CONFIG_CONTEXT_KEY).(*config.Config)
	stripeSignature := request.Header["Stripe-Signature"]
	if stripeSignature == nil {
		log.Errorf("No stripe signature found")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	event, err := webhook.ConstructEvent(payload, stripeSignature[0], config.Stripe.WebhookSecret)
	if err != nil {
		log.Errorf("Error verifying webhook signature: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if event.Type == "charge.succeeded" {
		var charge stripe.Charge
		err = json.Unmarshal(event.Data.Raw, &charge)
		if err != nil {
			log.Errorf("Error parsing webhook JSON: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

		order, err := db.GetOrderByStripeOrderId(charge.PaymentIntent.ID)
		if err != nil {
			log.Errorf("Error getting order from database: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		channel, err := db.GetChannelByStoreId(order.StoreId)
		if err != nil {
			log.Errorf("error getting channel from database: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = db.MarkOrderAsPaid(order.Id)
		if err != nil {
			log.Errorf("Error marking order as paid: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		orderItems, err := db.GetOrderItems(order.Id)
		if err != nil {
			log.Errorf("Error getting order items from database: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		users, err := db.GetUserByUserId(order.UserId)
		if err != nil {
			log.Errorf("failed to get user: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		store, err := db.GetStoreById(order.StoreId)
		if err != nil {
			log.Errorf("failed to get store: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		orderItemsDto := utils.ConvertOrderItemsToDto(orderItems)
		amount := utils.CalculateOrderPrice(orderItemsDto)
		orderId := order.Id.String()

		pickupTime := utils.GetPickupTime(store.AveragePickupTime)
		pickupTimeString := pickupTime.Format(utils.DeliverectTimeFormat)

		createOrderRequest := deliverect.CreateOrderRequest{
			ChannelOrderId:        orderId,
			ChannelOrderDisplayId: users.Firstname + "#" + orderId[len(orderId)-3:],
			Items:                 orderItemsDto,
			OrderType:             deliverect.PICKUP,
			OrderIsAlreadyPaid:    true,
			DecimalDigits:         2,
			Payment: deliverect.PaymentModel{
				Amount: amount,
				Type:   deliverect.CREDIT_CARD_ONLINE,
			},
			Customer: deliverect.CustomerModel{
				Name: users.Firstname,
			},
			EstimatedPickupTime: string(pickupTimeString),
			PickupTime:          string(pickupTimeString),
		}

		log.Debugf("Sending order to Deliverect: %v\n", pickupTimeString)

		deliverectService := c.Get(middleware.DELIVERECT_SERVICE_CONTEXT_KEY).(deliverect.DeliverectInterface)
		err = deliverectService.CreateOrder(createOrderRequest, channel.DeliverectLinkId)
		if err != nil {
			log.Errorf("Error creating order in deliverect: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}

	log.Errorf("Unhandled event type: %v\n", event.Type)
	return echo.NewHTTPError(http.StatusBadRequest)
}
