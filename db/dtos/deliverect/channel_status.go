package deliverect

type ChannelStatus int64

const (
	Register ChannelStatus = iota
	Active
	Inactive
)

type ChannelStatusRequest struct {
	ChannelLinkId     string `db:"channel_link_id" json:"channelLinkId"` // UNUSED
	ChannelLocationId int    `db:"channel_location_id" json:"channelLocationId"`
	LocationId        string `db:"location_id" json:"locationId"`
	Status            string `db:"status" json:"status"`
}

type ChannelStatusReponse struct {
	StatusUpdateURL   string `json:"statusUpdateURL"`
	MenuUpdateURL     string `json:"menuUpdateURL"`
	SnoozeUnsnoozeURL string `json:"snoozeUnsnoozeURL"`
	BusyModeURL       string `json:"busyModeURL"`
	UpdatePrepTimeURL string `json:"updatePrepTimeURL"`
}
