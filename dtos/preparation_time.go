package dtos 

type PreparationTimeRequestDto struct {
	ChannelOrderId string `json:"channelOrderId"`
	OrderId        string `json:"orderId"`
	Location       string `json:"location"`
	Status         int    `json:"status"`
	PickupTime     string `json:"pickupTime"`
}