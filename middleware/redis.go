package middleware

import (
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
)

const (
	REDIS_CONTEXT_KEY = "redis"
)

func RedisContext(redis *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(REDIS_CONTEXT_KEY, redis)
			return next(c)
		}
	}
}
