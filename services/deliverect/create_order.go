package deliverect

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

func (d DeliverectService) CreateOrder(order CreateOrderRequest) error {
	reqJson, err := json.Marshal(order)
	if err != nil {
		log.Errorf("Error marshalling request: %v\n", err)
		return err
	}

	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	orderUrl := d.BaseUrl + "/" + d.ChannelName + "/order/" + d.ChannelLinkId

	req, err := http.NewRequest("POST", orderUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		log.Errorf("error creating request: %v\n", err)
		return err
	}

	token, err := d.getMachineToMachineToken()
	if err != nil {
		log.Errorf("error getting machine to machine token: %v\n", err)
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+*token)

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("failed to send request: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		log.Errorf("failed to create order, status code: %v", resp.StatusCode)
		return err
	}

	return nil
}
