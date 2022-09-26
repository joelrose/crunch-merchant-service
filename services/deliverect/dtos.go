package deliverect

import "github.com/joelrose/crunch-merchant-service/models/dtos"

type CreateMachineMachineTokenRequest struct {
	CliendId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type CreateMachineMachineTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type CustomerModel struct {
	Name string `json:"name"`
}

type PaymentModel struct {
	Amount int         `json:"amount"`
	Type   PaymentType `json:"type"`
}

type CreateOrderRequest struct {
	ChannelOrderId        string           `json:"channelOrderId"`
	ChannelOrderDisplayId string           `json:"channelOrderDisplayId"`
	Items                 []dtos.OrderItem `json:"items"`
	Payment               PaymentModel     `json:"payment"`
	Customer              CustomerModel    `json:"customer"`
	OrderType             OrderType        `json:"orderType"`
	OrderIsAlreadyPaid    bool             `json:"orderIsAlreadyPaid"`
	DecimalDigits         int              `json:"decimalDigits"`
}
