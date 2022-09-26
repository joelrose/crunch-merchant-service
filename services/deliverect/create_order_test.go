package deliverect

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/golang/mock/gomock"
	"github.com/joelrose/crunch-merchant-service/services/http_client"
	redisService "github.com/joelrose/crunch-merchant-service/services/redis"
	"github.com/joelrose/crunch-merchant-service/test_helper/mock_http_client"
	"github.com/stretchr/testify/assert"
)

var (
	machineTokenResponse = CreateMachineMachineTokenResponse{
		AccessToken: "access_token",
		ExpiresAt:   time.Now().Add(time.Hour * 24).Unix(),
		TokenType:   "Bearer",
		Scope:       "scope",
	}
	deliverectConfig = DeliverectServiceConfig{
		BaseUrl:      "https://deliverect.getcrunch.tech",
		ChannelName:  "crunch",
		ClientId:     "client_id",
		ClientSecret: "client_secret",
	}
	service = DeliverectService{
		RedisClient: nil,
		Config:      deliverectConfig,
		HttpClient:  nil,
	}
	channelLinkId = "channel_link_id"
)

func createServices(t *testing.T) (*mock_http_client.MockCustomHttpClient, *miniredis.Miniredis) {
	miniRedis := miniredis.RunT(t)

	redisClient := redisService.NewClient("redis://" + miniRedis.Addr())
	httpClient := http_client.NewMockClient(t)

	service.HttpClient = httpClient
	service.RedisClient = redisClient

	return httpClient, miniRedis
}

func TestCreateOrderRequestToken(t *testing.T) {
	httpClient, miniRedis := createServices(t)

	machineTokenResponseJson, _ := json.Marshal(machineTokenResponse)
	machineTokenReader := io.NopCloser(bytes.NewReader([]byte(machineTokenResponseJson)))

	httpClient.EXPECT().SendPost(gomock.Any(), deliverectConfig.BaseUrl+MachineTokenPath, gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       machineTokenReader,
	}, nil).Times(1)

	httpClient.EXPECT().SendPost(
		gomock.Any(),
		deliverectConfig.BaseUrl+"/"+deliverectConfig.ChannelName+"/order/"+channelLinkId,
		map[string]string{
			"Authorization": "Bearer " + machineTokenResponse.AccessToken,
		}).Return(&http.Response{
		StatusCode: 201,
	}, nil).Times(1)

	assert.Nil(t, service.CreateOrder(CreateOrderRequest{}, channelLinkId))

	token, _ := miniRedis.Get(DeliverectMachineToken)
	assert.Equal(t, string(machineTokenResponseJson), token)
}

func TestCreateOrderCachedToken(t *testing.T) {
	httpClient, miniRedis := createServices(t)

	machineTokenResponseJson, _ := json.Marshal(machineTokenResponse)

	miniRedis.Set(DeliverectMachineToken, string(machineTokenResponseJson))

	httpClient.EXPECT().SendPost(
		gomock.Any(),
		deliverectConfig.BaseUrl+"/"+deliverectConfig.ChannelName+"/order/"+channelLinkId,
		map[string]string{
			"Authorization": "Bearer " + machineTokenResponse.AccessToken,
		}).Return(&http.Response{
		StatusCode: 201,
	}, nil).Times(1)

	assert.Nil(t, service.CreateOrder(CreateOrderRequest{}, channelLinkId))
}