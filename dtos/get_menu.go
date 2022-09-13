package dtos

type GetMenuRequest struct {
	StoreId int `param:"id"`
}

type GetMenuCategory struct {
	Id              int    `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	ImageUrl        string `db:"image_url"`
	SortOrder       int    `db:"sort_order"`
	ProductChildren []int
}

type GetMenuProduct struct {
	Id              int    `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	Price           int    `db:"price"`
	Max             int    `db:"max"`
	Min             int    `db:"min"`
	Multiply        int    `db:"multiply"`
	Plu             string `db:"plu"`
	Snoozed         bool   `db:"snoozed"`
	Tax             int    `db:"tax"`
	ProductType     int    `db:"product_type"`
	ImageUrl        string `db:"image_url"`
	SortOrder       int    `db:"sort_order"`
	Visible         bool   `db:"visible"`
	ProductChildren []int
}

type GetMenuOpeningHour struct {
	DayOfWeek      int `db:"day_of_week"`
	StartTimestamp int `db:"start_timestamp"`
	EndTimestamp   int `db:"end_timestamp"`
}

type GetMenuResponse struct {
	Categories   []GetMenuCategory
	Products     []GetMenuProduct
	OpeningHours []GetMenuOpeningHour
}
