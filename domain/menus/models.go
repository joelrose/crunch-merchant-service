package menus

import (
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

type (
	MenuRedisModel struct {
		Categories   []dtos.GetStoreCategory
		OpeningHours []dtos.GetStoreOpeningHour
	}

	MenuService struct {
		db      *db.DB
		rdb     *redis.Client
		storeId uuid.UUID
	}
)
