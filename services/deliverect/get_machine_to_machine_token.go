package deliverect

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

const (
	DeliverectGrantType = "client_credentials"
	MachineTokenPath    = "/oauth/token"
)

func (d DeliverectService) getMachineToMachineToken() (*string, error) {
	request := CreateMachineMachineTokenRequest{
		CliendId:     d.ClientId,
		ClientSecret: d.ClientSecret,
		Audience:     d.BaseUrl,
		GrantType:    DeliverectGrantType,
	}

	requestJson, err := json.Marshal(request)
	if err != nil {
		log.Errorf("Error marshalling request: %v\n", err)
		return nil, err
	}

	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	resp, err := client.Post(MachineTokenPath, "application/json", bytes.NewBuffer(requestJson))
	if err != nil {
		log.Errorf("failed to send request: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to read response body: %v", err)
		return nil, err
	}

	token := CreateMachineMachineTokenResponse{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Errorf("failed to unmarshal response body: %v", err)
		return nil, err
	}

	return &token.AccessToken, nil
}
