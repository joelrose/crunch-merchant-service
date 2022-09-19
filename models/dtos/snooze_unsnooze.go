package dtos

import "time"

type SnoozeUnsnzoozeRequestDto struct {
	AccountId     string                    `json:"accountId"`
	LocationId    string                    `json:"locationId"`
	ChannelLinkId string                    `json:"channelLinkId"`
	Operations    []SnoozeUnsnoozeOperation `json:"operations"`
}

type SnoozeUnsnoozeOperation struct {
	Action string             `json:"action"`
	Data   SnoozeUnsnoozeData `json:"data"`
}

type SnoozeUnsnoozeData struct {
	Items []SnoozeUnsnoozeItem `json:"items"`
}

type SnoozeUnsnoozeItem struct {
	Id          string    `json:"_id"`
	Plu         string    `json:"plu"`
	SnoozeStart time.Time `json:"snoozeStart"`
	SnoozeEnd   time.Time `json:"snoozeEnd"`
}

type SnoozeUnsnoozeResponseDto struct {
	Results []SnoozeUnsnoozeResponseResult `json:"results"`
}

type SnoozeUnsnoozeResponseResult struct {
	Action string                     `json:"action"`
	Data   SnoozeUnsnoozeResponseData `json:"data"`
}

type SnoozeUnsnoozeResponseData struct {
	LocationId      string   `json:"locationId"`
	AllSnoozedItems []string `json:"allSnoozedItems"`
}
