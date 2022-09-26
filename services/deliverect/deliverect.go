package deliverect

import (
	"github.com/go-redis/redis/v9"
	"github.com/joelrose/crunch-merchant-service/config"
)

type DeliverectServiceConfig struct {
	BaseUrl      string
	ChannelName  string
	ClientId     string
	ClientSecret string
}

type DeliverectService struct {
	Config      DeliverectServiceConfig
	RedisClient *redis.Client
}

func NewDeliverectService(config config.Config, redisClient *redis.Client) DeliverectInterface {
	var deliverectInterface DeliverectInterface = &DeliverectService{
		RedisClient: redisClient,
		Config: DeliverectServiceConfig{
			BaseUrl:      config.Deliverect.BaseUrl,
			ChannelName:  config.Deliverect.ChannelName,
			ClientId:     config.Deliverect.ClientId,
			ClientSecret: config.Deliverect.ClientSecret,
		},
	}
	return deliverectInterface
}
