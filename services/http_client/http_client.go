package http_client

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/joelrose/crunch-merchant-service/test_helper/mock_http_client"
	"github.com/labstack/gommon/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type CustomHttpClient interface {
	SendPost(requestBody []byte, url string, headers map[string]string) (*http.Response, error)
}

type HttpClient struct{}

func (httpClient *HttpClient) SendPost(requestBody []byte, url string, headers map[string]string) (*http.Response, error) {
	client := http.Client{
		Timeout:   time.Duration(2) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	log.Debugf("sending http post request to: %v", url)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	return resp, nil
}

func NewClient() CustomHttpClient {
	return &HttpClient{}
}

func NewMockClient(t *testing.T) *mock_http_client.MockCustomHttpClient {
	ctrl := gomock.NewController(t)
	return mock_http_client.NewMockCustomHttpClient(ctrl)
}
