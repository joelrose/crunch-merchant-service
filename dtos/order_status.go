package dtos

const (
	ORDER_STATUS_PARSED              = 1
	ORDER_STATUS_RECEIVED            = 2
	ORDER_STATUS_NEW                 = 10
	ORDER_STATUS_ACCEPTED            = 20
	ORDER_STATUS_DUPLICATE           = 30
	ORDER_STATUS_PRINTED             = 40
	ORDER_STATUS_PREPARING           = 50
	ORDER_STATUS_PREPARED            = 60
	ORDER_STATUS_READY_FOR_PICKUP    = 70
	ORDER_STATUS_INDELIVERY          = 80
	ORDER_STATUS_FINALIZED           = 90
	ORDER_STATUS_AUTO_FINALIZED      = 95
	ORDER_STATUS_CANCEL              = 100
	ORDER_STATUS_CANCELED            = 110
	ORDER_STATUS_FAILED              = 120
	ORDER_STATUS_POS_RECEIVED_FAILED = 121
	ORDER_STATUS_PARSE_FAILED        = 124
)

type UpdateOrderStatusRequest struct {
	OrderId        string `json:"orderId"`
	Status         int    `json:"status"`
	ReceiptId      string `json:"receiptId"`
	Reason         string `json:"reason"`
	ChannelOrderId int    `json:"channelOrderId"`
	Location       string `json:"location"`
	ChannelLink    string `json:"channelLink"`
}
