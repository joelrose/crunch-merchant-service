package webhook

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/joelrose/crunch-merchant-service/services/deliverect"
	"github.com/joelrose/crunch-merchant-service/test_helper"
	"github.com/joelrose/crunch-merchant-service/test_helper/mock_deliverect"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/webhook"
)

const (
	stripeWebhookSecretCorrect = "whsec_123"
	stripeWebhookSecretWrong   = "whsec_456"
)

var (
	mockOrderId     = uuid.MustParse("a7286987-72b8-459a-9a06-223ef7418be4")
	mockStoreId     = uuid.MustParse("783ed769-dca9-4031-9ff6-03ce6f126a4c")
	mockOrderItemId = uuid.MustParse("dd419c16-57b1-4c83-af39-6bc63e54e0fd")
	mockUserId      = uuid.MustParse("5550a287-87b7-473f-87d1-04b6a351454f")
	mockRequestBody = stripe.Event{
		Type:       "charge.succeeded",
		APIVersion: "2022-08-01",
		Data:       &stripe.EventData{},
	}
	mockConfig = config.Config{
		Stripe: config.Stripe{
			WebhookSecret: stripeWebhookSecretCorrect,
		},
	}
	mockCharge = stripe.Charge{
		ID: "ch_123",
		PaymentIntent: &stripe.PaymentIntent{
			ID: "pi_123",
		},
	}
	mockOrder = models.Order{
		Id:      mockOrderId,
		StoreId: mockStoreId,
		UserId:  mockUser.Id,
	}
	mockOrderItems = []models.OrderItem{
		{
			Id:       mockOrderItemId,
			OrderId:  mockOrder.Id,
			Plu:      "123",
			Quantity: 1,
			Price:    2000,
			Name:     "Test Item",
		},
	}
	mockOrderItemsDto = []dtos.OrderItem{
		{
			Id:       mockOrderItems[0].Id,
			Quantity: mockOrderItems[0].Quantity,
			Price:    mockOrderItems[0].Price,
			Name:     mockOrderItems[0].Name,
			Plu:      mockOrderItems[0].Plu,
			SubItems: []dtos.OrderItem{},
		},
	}
	mockUser = models.User{
		Id:        mockUserId,
		Firstname: "John",
	}
	mockChannel = models.DeliverectChannel{
		DeliverectLinkId: "delivered_link_id",
	}
	mockDeliverectObject = deliverect.CreateOrderRequest{
		ChannelOrderId:        mockOrder.Id.String(),
		ChannelOrderDisplayId: "John#be4",
		Items:                 mockOrderItemsDto,
		OrderType:             deliverect.PICKUP,
		OrderIsAlreadyPaid:    true,
		DecimalDigits:         2,
		Payment: deliverect.PaymentModel{
			Amount: 2000,
			Type:   deliverect.CREDIT_CARD_ONLINE,
		},
		Customer: deliverect.CustomerModel{
			Name: mockUser.Firstname,
		},
	}
)

func TestStripeWebhookWrongWebhookSecret(t *testing.T) {
	requestJSON, _ := json.Marshal(mockRequestBody)
	_, c, _ := test_helper.NewRequest(t, http.MethodPost, string(requestJSON))

	c.Request().Header["Stripe-Signature"] = []string{stripeWebhookSecretWrong}
	c.Set(middleware.CONFIG_CONTEXT_KEY, &mockConfig)

	err := HandleStripe(c)
	if assert.NotNil(t, err) {
		assert.Equal(t, echo.NewHTTPError(http.StatusBadRequest), err)
	}
}

func TestStripeWebhookNoHeaderSet(t *testing.T) {
	requestJSON, _ := json.Marshal(mockRequestBody)
	_, c, _ := test_helper.NewRequest(t, http.MethodPost, string(requestJSON))

	c.Set(middleware.CONFIG_CONTEXT_KEY, &mockConfig)

	err := HandleStripe(c)
	if assert.NotNil(t, err) {
		assert.Equal(t, echo.NewHTTPError(http.StatusBadRequest), err)
	}
}

func TestStripeWebhookSucceeds(t *testing.T) {
	chargeJSON, _ := json.Marshal(mockCharge)
	mockRequestBody.Data.Raw = chargeJSON

	requestJSON, _ := json.Marshal(mockRequestBody)
	rec, c, mockDB := test_helper.NewRequest(t, http.MethodPost, string(requestJSON))

	signedPayload := webhook.GenerateTestSignedPayload(&webhook.UnsignedPayload{Payload: requestJSON, Secret: stripeWebhookSecretCorrect})
	c.Request().Header["Stripe-Signature"] = []string{signedPayload.Header}
	c.Set(middleware.CONFIG_CONTEXT_KEY, &mockConfig)

	mockDB.EXPECT().GetOrderByStripeOrderId(mockCharge.PaymentIntent.ID).Return(mockOrder, nil)

	mockDB.EXPECT().GetChannelByStoreId(mockOrder.StoreId).Return(mockChannel, nil)

	mockDB.EXPECT().MarkOrderAsPaid(mockOrder.Id).Return(nil)

	mockDB.EXPECT().GetOrderItems(mockOrder.Id).Return(mockOrderItems, nil)

	mockDB.EXPECT().GetUserByUserId(mockOrder.UserId).Return(mockUser, nil)

	ctrl := gomock.NewController(t)
	mockDeliverect := mock_deliverect.NewMockDeliverectInterface(ctrl)

	c.Set(middleware.DELIVERECT_SERVICE_CONTEXT_KEY, mockDeliverect)

	mockDeliverect.EXPECT().CreateOrder(mockDeliverectObject, mockChannel.DeliverectLinkId).Return(nil)

	err := HandleStripe(c)
	if assert.Nil(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
