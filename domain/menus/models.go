package menus

import (
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

type (
	MenuRedisModel struct {
		Categories   []dtos.GetStoreCategory    `json:"categories"`
		OpeningHours []dtos.GetStoreOpeningHour `json:"opening_hours"`
	} //@name Menu

	MenuService struct {
		db      db.DBInterface
		rdb     *redis.Client
		storeId uuid.UUID
	}
)
