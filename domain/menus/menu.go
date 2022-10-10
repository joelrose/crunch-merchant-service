package menus

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/gommon/log"
)

func (s *MenuService) cacheMenu(menu MenuRedisModel) error {
	ctx := context.Background()

	json, err := json.Marshal(menu)
	if err != nil {
		return fmt.Errorf("failed to marshal menu: %v", err)
	}

	// TODO: find a more profound way to handle this
	err = s.rdb.Set(ctx, s.storeId.String(), json, 1*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to save menu to redis: %v", err)
	}

	return nil
}

func (s *MenuService) getCachedMenu() (*MenuRedisModel, error) {
	ctx := context.Background()

	value, err := s.rdb.Get(ctx, s.storeId.String()).Result()
	if err == nil {
		var menu MenuRedisModel
		err := json.Unmarshal([]byte(value), &menu)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal menu model: %v", err)
		}

		return &menu, nil
	}

	return nil, nil
}

func (s *MenuService) GetMenu() (*MenuRedisModel, error) {
	cachedMenu, err := s.getCachedMenu()
	if err != nil {
		log.Errorf("failed to get cached menu: %v", err)
	}

	if cachedMenu != nil {
		log.Debugf("serving cached menu [%v]", s.storeId)
		return cachedMenu, nil
	}

	menu, err := s.build()
	if err != nil {
		return nil, fmt.Errorf("failed to build menu: %v", err)
	}

	err = s.cacheMenu(*menu)
	if err != nil {
		log.Errorf("failed to save menu in redis: %v", err)
	}

	return menu, nil
}

func NewMenuService(db db.DBInterface, rdb *redis.Client, storeId uuid.UUID) *MenuService {
	return &MenuService{
		db:      db,
		rdb:     rdb,
		storeId: storeId,
	}
}
