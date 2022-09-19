package menus

import "github.com/joelrose/crunch-merchant-service/models/dtos"

type (
	MenuRedisModel struct {
		Categories   []dtos.GetStoreCategory
		OpeningHours []dtos.GetStoreOpeningHour
	}
)
