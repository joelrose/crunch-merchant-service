package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-redis/redis/v9"
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

func HandleStripe(c echo.Context) error {
	request := c.Request()
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		log.Errorf("failed to read request body: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	config := c.Get(middleware.CONFIG_CONTEXT_KEY).(*config.Config)
	stripeSignature := request.Header["Stripe-Signature"]

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

		db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

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

		orderItemsDto := utils.ConvertOrderItemsToDto(orderItems)
		amount := utils.CalculateOrderPrice(orderItemsDto)
		createOrderRequest := deliverect.CreateOrderRequest{
			ChannelOrderId:        fmt.Sprint(order.Id),
			ChannelOrderDisplayId: users.Firstname + "#" + fmt.Sprint(order.Id),
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
		}

		redisClient := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)
		deliverectService := deliverect.NewDeliverectService(*config, redisClient, channel.DeliverectLinkId, "crunch")

		err = deliverectService.CreateOrder(createOrderRequest)
		if err != nil {
			log.Errorf("Error creating order in deliverect: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	} else {
		log.Errorf("Unhandled event type: %v\n", event.Type)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
}
