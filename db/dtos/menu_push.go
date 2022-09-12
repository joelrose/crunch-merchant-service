package dtos

import (
	"time"
)

type DeliverectMenuProduct struct {
	Id          string    `db:"id" json:"_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Max         int       `db:"max" json:"max"`
	Min         int       `db:"min" json:"min"`
	Multiply    int       `db:"multiply" json:"multiply"`
	Plu         string    `db:"plu" json:"plu"`
	Snoozed     bool      `db:"snoozed" json:"snoozed"`
	Tax         int       `db:"tax" json:"takeawayTax"`
	ProductType int       `db:"product_type" json:"productType"`
	ImageURL    string    `db:"image_url" json:"imageUrl"`
	SortOrder   int       `db:"sort_order" json:"sortOrder"`
	SubProducts []string  `json:"subProducts"`
	Visible     bool      `db:"visible" json:"visible"`
	Price       int       `db:"price" json:"price"`
	StoreId     int       `db:"store_id" json:"storeId"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

type DeliverectCategory struct {
	ID                      string        `json:"_id"`
	Name                    string        `json:"name"`
	Description             string        `json:"description"`
	Account                 string        `json:"account"`
	PosLocationID           string        `json:"posLocationId"`
	PosCategoryType         string        `json:"posCategoryType"`
	PosCategoryID           string        `json:"posCategoryId"`
	ImageURL                string        `json:"imageUrl"`
	SubCategories           []interface{} `json:"subCategories"`
	Products                []string      `json:"products"`
	Availabilities          []interface{} `json:"availabilities"`
	Level                   int           `json:"level"`
	Menu                    string        `json:"menu"`
	SortedChannelProductIds []interface{} `json:"sortedChannelProductIds"`
	SubProducts             []string      `json:"subProducts"`
	SubProductSortOrder     []interface{} `json:"subProductSortOrder"`
}

type DeliverectStoreOpeningHour struct {
	DayOfWeek int    `json:"dayOfWeek"`
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`
}

type MenuPushRequest []struct {
	StoreOpeningHours []DeliverectStoreOpeningHour     `json:"availabilities"`
	Bundles           map[string]DeliverectMenuProduct `json:"bundles"`
	Categories        []DeliverectCategory             `json:"categories"`
	ChannelLinkID     string                           `json:"channelLinkId"`
	Currency          int                              `json:"currency"`
	Description       string                           `json:"description"`
	Menu              string                           `json:"menu"`
	MenuID            string                           `json:"menuId"`
	MenuImageURL      string                           `json:"menuImageURL"`
	MenuType          int                              `json:"menuType"`
	ModifierGroups    map[string]DeliverectMenuProduct `json:"modifierGroups"`
	Modifiers         map[string]DeliverectMenuProduct `json:"modifiers"`
	NestedModifiers   bool                             `json:"nestedModifiers"`
	Products          map[string]DeliverectMenuProduct `json:"products"`
	ProductTags       []int                            `json:"productTags"`
	SnoozedProducts   struct{}                         `json:"snoozedProducts"`
}
