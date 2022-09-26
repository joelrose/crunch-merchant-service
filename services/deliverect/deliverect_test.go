package deliverect

import (
	"testing"

	"github.com/go-redis/redis/v9"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/services/http_client"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	mockConfig := config.Config{
		Deliverect: config.Deliverect{
			BaseUrl:      "https://deliverect.getcrunch.tech",
			ChannelName:  "crunch",
			ClientId:     "client_id",
			ClientSecret: "secret",
		},
	}
	mockRedisClient := redis.Client{}
	mockHttpClient := http_client.HttpClient{}

	service := NewDeliverectService(mockConfig, &mockRedisClient, &mockHttpClient)

	assert.Equal(t, service.(*DeliverectService).Config.BaseUrl, mockConfig.Deliverect.BaseUrl)
	assert.Equal(t, service.(*DeliverectService).Config.ChannelName, mockConfig.Deliverect.ChannelName)
	assert.Equal(t, service.(*DeliverectService).Config.ClientId, mockConfig.Deliverect.ClientId)
	assert.Equal(t, service.(*DeliverectService).Config.ClientSecret, mockConfig.Deliverect.ClientSecret)

	assert.Equal(t, service.(*DeliverectService).RedisClient, &mockRedisClient)
	assert.Equal(t, service.(*DeliverectService).HttpClient, &mockHttpClient)
}
