package deliverect

import "github.com/joelrose/crunch-merchant-service/config"

type DeliverectService struct {
	BaseUrl       string
	ClientId      string
	ClientSecret  string
	ChannelLinkId string
	ChannelName   string
}

func NewDeliverectService(config config.Config, channelLinkId string, channelName string) *DeliverectService {
	return &DeliverectService{
		BaseUrl:       config.Deliverect.BaseUrl,
		ClientId:      config.Deliverect.ClientId,
		ClientSecret:  config.Deliverect.ClientSecret,
		ChannelLinkId: channelLinkId,
		ChannelName:   channelName,
	}
}
