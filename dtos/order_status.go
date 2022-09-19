package dtos

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db/models"
)

type UpdateOrderStatusRequest struct {
	OrderId        string             `json:"orderId"`
	Status         models.OrderStatus `json:"status"`
	ReceiptId      string             `json:"receiptId"`
	Reason         string             `json:"reason"`
	ChannelOrderId uuid.UUID          `json:"channelOrderId"`
	Location       string             `json:"location"`
	ChannelLink    string             `json:"channelLink"`
}
