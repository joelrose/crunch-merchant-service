package deliverect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (d DeliverectService) CreateOrder(order CreateOrderRequest, channelLinkId string) error {
	requestJson, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error marshalling request: %v", err)
	}

	token, err := d.getCachedMachineToMachineToken()
	if err != nil {
		return fmt.Errorf("error getting machine to machine token: %v", err)
	}

	orderUrl := d.Config.BaseUrl + "/" + d.Config.ChannelName + "/order/" + channelLinkId

	resp, err := d.HttpClient.SendPost(requestJson, orderUrl, map[string]string{
		"Authorization": "Bearer " + *token,
	})
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		payload, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read request body: %v", err)
		}

		return fmt.Errorf("failed to create order, status code: %v, error: %v", resp.StatusCode, string(payload))
	}

	return nil
}
