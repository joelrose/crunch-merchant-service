package middleware

import (
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/labstack/echo/v4"
)

const (
	CONFIG_CONTEXT_KEY = "config"
)

func ConfigContext(config *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(CONFIG_CONTEXT_KEY, config)
			return next(c)
		}
	}
}
