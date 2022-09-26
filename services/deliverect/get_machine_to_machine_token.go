package deliverect

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/labstack/gommon/log"
)

const (
	DeliverectGrantType    = "client_credentials"
	DeliverectMachineToken = "deliverect_machine_token"
	MachineTokenPath       = "/oauth/token"
)

func (d DeliverectService) getCachedMachineToMachineToken() (*string, error) {
	context := context.Background()
	tokenJson, err := d.RedisClient.Get(context, DeliverectMachineToken).Result()
	if err != nil {
		log.Debug("requesting machine to machine token from deliverect")
		return d.getMachineToMachineToken()
	}

	token := CreateMachineMachineTokenResponse{}
	err = json.Unmarshal([]byte(tokenJson), &token)
	if err != nil {
		log.Errorf("failed to unmarshal token body: %v", err)
		return nil, err
	}

	return &token.AccessToken, nil
}

func (d DeliverectService) getMachineToMachineToken() (*string, error) {
	request := CreateMachineMachineTokenRequest{
		CliendId:     d.Config.ClientId,
		ClientSecret: d.Config.ClientSecret,
		Audience:     d.Config.BaseUrl,
		GrantType:    DeliverectGrantType,
	}

	requestJson, err := json.Marshal(request)
	if err != nil {
		log.Errorf("Error marshalling request: %v\n", err)
		return nil, err
	}

	resp, err := d.HttpClient.SendPost(requestJson, d.Config.BaseUrl+MachineTokenPath, nil)
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

	expiresIn := token.ExpiresAt - time.Now().Unix()
	d.RedisClient.Set(context.Background(), DeliverectMachineToken, string(body), time.Duration(expiresIn)*time.Second)

	return &token.AccessToken, nil
}
