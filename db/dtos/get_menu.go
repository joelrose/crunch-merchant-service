package dtos

import "github.com/joelrose/crunch-merchant-service/db/models"

type GetMenuRequest struct {
	StoreId int `param:"id"`
}

type GetMenuResponse struct {
	Categories []models.MenuCategory
	Products   []models.MenuProduct
}
