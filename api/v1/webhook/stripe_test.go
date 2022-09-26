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
		Id:      1,
		StoreId: uuid.New(),
		UserId:  1,
	}
	mockOrderItems = []models.OrderItem{
		{
			Id:       uuid.New(),
			OrderId:  mockOrder.Id,
			Quantity: 1,
			Price:    2000,
		},
	}
	mockUser = models.User{
		Id:        1,
		Firstname: "John",
	}
	mockChannel = models.DeliverectChannel{
		DeliverectLinkId: "delivered_link_id",
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

	mockDeliverect.EXPECT().CreateOrder(gomock.Any(), mockChannel.DeliverectLinkId).Return(nil)

	err := HandleStripe(c)
	if assert.Nil(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
