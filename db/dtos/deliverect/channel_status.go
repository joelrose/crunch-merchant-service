package deliverect

type ChannelStatus int64

const (
	Register ChannelStatus = iota
	Active
	Inactive
)

type ChannelStatusRequest struct {
	ChannelLinkId     string `db:"channel_link_id" json:"channel_link_id"` // UNUSED
	ChannelLocationId int    `db:"channel_location_id" json:"channel_location_id"`
	LocationId        string `db:"location_id" json:"location_id"`
	Status            string `db:"status" json:"status"`
}
