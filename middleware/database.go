package middleware

import (
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/echo/v4"
)

const (
	DATBASE_CONTEXT_KEY = "db"
)

func DatabaseContext(db *db.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DATBASE_CONTEXT_KEY, db)
			return next(c)
		}
	}
}
