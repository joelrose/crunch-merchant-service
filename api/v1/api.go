package api

import (
	"github.com/go-redis/redis/v9"
	"github.com/joelrose/crunch-merchant-service/db"
)

type (
	Handler struct {
		db    *db.DB
		redis *redis.Client
	}
)
