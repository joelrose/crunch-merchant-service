package dtos

import "github.com/joelrose/crunch-merchant-service/db/models"

type UpdateOrderStatusRequest struct {
	OrderId        string             `json:"orderId"`
	Status         models.OrderStatus `json:"status"`
	ReceiptId      string             `json:"receiptId"`
	Reason         string             `json:"reason"`
	ChannelOrderId int                `json:"channelOrderId"`
	Location       string             `json:"location"`
	ChannelLink    string             `json:"channelLink"`
}
