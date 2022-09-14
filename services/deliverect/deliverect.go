package deliverect

import (
	"github.com/go-redis/redis/v9"
	"github.com/joelrose/crunch-merchant-service/config"
)

type DeliverectServiceConfig struct {
	BaseUrl       string
	ClientId      string
	ClientSecret  string
	ChannelLinkId string
	ChannelName   string
}

type DeliverectService struct {
	Config      DeliverectServiceConfig
	RedisClient *redis.Client
}

func NewDeliverectService(config config.Config, redisClient *redis.Client, channelLinkId string, channelName string) *DeliverectService {
	return &DeliverectService{
		RedisClient: redisClient,
		Config: DeliverectServiceConfig{
			BaseUrl:       config.Deliverect.BaseUrl,
			ClientId:      config.Deliverect.ClientId,
			ClientSecret:  config.Deliverect.ClientSecret,
			ChannelLinkId: channelLinkId,
			ChannelName:   channelName,
		},
	}
}
