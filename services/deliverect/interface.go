package deliverect

type DeliverectInterface interface {
	CreateOrder(order CreateOrderRequest, channelLinkId string) error
}
