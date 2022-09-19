package menus

import (
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
)

func ConvertToGetStoreResponse(store models.Store, menu *MenuRedisModel) dtos.GetStoreResponse {
	return dtos.GetStoreResponse{
		Id:                store.Id,
		Name:              store.Name,
		Description:       store.Description,
		Address:           store.Address,
		AveragePickupTime: store.AveragePickupTime,
		AverageReview:     store.AverageReview,
		ReviewCount:       store.ReviewCount,
		GoogleMapsLink:    store.GoogleMapsLink,
		PhoneNumber:       store.PhoneNumber,
		ImageUrl:          store.ImageUrl,
		Categories:        menu.Categories,
		OpeningHours:      menu.OpeningHours,
	}
}
