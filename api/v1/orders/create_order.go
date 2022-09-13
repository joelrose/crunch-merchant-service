package orders

import (
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/paymentintent"
)

func CreateOrder(c echo.Context) error {
	db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

	token := c.Get("token").(*auth.Token)

	user, err := db.GetUserByFirebaseId(token.UID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var orderRequest dtos.CreateOrderRequest
	err = c.Bind(&orderRequest)
	if err != nil {
		log.Debugf("failed to bind order request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "order model is invalid")
	}

	// make this check for
	store, err := db.GetAvailableStore(orderRequest.StoreId)
	if err != nil {
		log.Errorf("failed to retrieve store: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	/*if !store.StripeAccountId.Valid {
		log.Errorf("store has no stripe account id: %v", orderRequest.StoreId)
		return echo.NewHTTPError(http.StatusBadRequest)
	}*/

	// check if order items exists (recursively)
	if len(orderRequest.OrderItems) == 0 {
		log.Errorf("order items is empty")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	price := utils.CalculateOrderPrice(orderRequest.OrderItems)
	_ = float64(price) * store.Fee

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(price)),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		// ApplicationFeeAmount: stripe.Int64(int64(fee)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}

	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		log.Errorf("failed to create payment intent: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	order := models.CreateOrder{
		Status:              1,
		EstimatedPickupTime: time.Now(),
		Price:               price,
		StripeOrderId:       paymentIntent.ID,
		IsPaid:              false,
		StoreId:             orderRequest.StoreId,
		UserId:              user.Id,
	}

	orderDatabaseId, err := db.CreateOrder(order)
	if err != nil {
		log.Errorf("failed to create order: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for _, orderItem := range orderRequest.OrderItems {
		// TODO: parent_id
		orderItem := models.CreateOrderItem{
			Plu:      orderItem.Plu,
			Name:     orderItem.Name,
			Price:    orderItem.Price,
			Quantity: orderItem.Quantity,
			OrderId:  orderDatabaseId,
		}

		_, err = db.CreateOrderItem(orderItem)
		if err != nil {
			log.Errorf("failed to create order item: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	response := dtos.CreateOrderResponse{
		StripeClientSecret: paymentIntent.ClientSecret,
	}

	return c.JSON(http.StatusCreated, response)
}
