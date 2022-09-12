package deliverect

type BusyModeRequest struct {
	AccountId     string `json:"accountId"`
	LocationId    string `json:"locationId"`
	ChannelLinkId string `json:"channelLinkId"`
	Status        string `json:"status"`
}

type BusyModeResponse struct {
	Status string `json:"status"`
}
