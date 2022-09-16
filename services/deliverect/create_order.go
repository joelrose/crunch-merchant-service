package deliverect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

func (d DeliverectService) CreateOrder(order CreateOrderRequest) error {
	reqJson, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error marshalling request: %v", err)
	}

	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	orderUrl := d.Config.BaseUrl + "/" + d.Config.ChannelName + "/order/" + d.Config.ChannelLinkId

	log.Debugf("creating order with deliverect api: %v", orderUrl)

	req, err := http.NewRequest("POST", orderUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	token, err := d.getCachedMachineToMachineToken()
	if err != nil {
		return fmt.Errorf("error getting machine to machine token: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+*token)

	resp, err := client.Do(req)
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
