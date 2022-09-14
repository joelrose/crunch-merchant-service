package deliverect

import "github.com/joelrose/crunch-merchant-service/dtos"

type CreateMachineMachineTokenRequest struct {
	CliendId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type CreateMachineMachineTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type CustomerModel struct {
	Name string `json:"name"`
}

type PaymentModel struct {
	Amount int `json:"amount"`
	Type   int `json:"type"`
}

type CreateOrderRequest struct {
	ChannelOrderId        string           `json:"channel_order_id"`
	ChannelOrderDisplayId string           `json:"channel_order_display_id"`
	Items                 []dtos.OrderItem `json:"items"`
	Payment               PaymentModel     `json:"payment"`
	Customer              CustomerModel    `json:"customer"`
	OrderType             int              `json:"order_type"` 
	OrderIsAlreadyPaid    bool             `json:"order_is_already_paid"`
	DecimalDigits         int              `json:"decimal_digits"`
}
