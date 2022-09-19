package dtos

import "github.com/google/uuid"

type ChannelStatus int64

const (
	Register ChannelStatus = iota
	Active
	InActive
)

type ChannelStatusRequest struct {
	ChannelLinkId     string    `db:"deliverect_link_id" json:"channelLinkId"`
	ChannelLocationId uuid.UUID `db:"deliverect_location_id" json:"channelLocationId"`
	LocationId        string    `db:"location_id" json:"locationId"`
	Status            string    `db:"status" json:"status"`
}

type ChannelStatusReponse struct {
	OrderStatusUpdateURL string `json:"statusUpdateURL"`
	MenuUpdateURL        string `json:"menuUpdateURL"`
	SnoozeUnsnoozeURL    string `json:"snoozeUnsnoozeURL"`
	BusyModeURL          string `json:"busyModeURL"`
	UpdatePrepTimeURL    string `json:"updatePrepTimeURL"`
}
