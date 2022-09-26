package deliverect

import (
	"github.com/go-redis/redis/v9"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/services/http_client"
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
	HttpClient  http_client.CustomHttpClient
}

func NewDeliverectService(config config.Config, redisClient *redis.Client, httpClient http_client.CustomHttpClient) DeliverectInterface {
	var deliverectInterface DeliverectInterface = &DeliverectService{
		RedisClient: redisClient,
		Config: DeliverectServiceConfig{
			BaseUrl:      config.Deliverect.BaseUrl,
			ChannelName:  config.Deliverect.ChannelName,
			ClientId:     config.Deliverect.ClientId,
			ClientSecret: config.Deliverect.ClientSecret,
		},
		HttpClient: httpClient,
	}
	return deliverectInterface
}
