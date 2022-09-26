package middleware

import (
	"github.com/joelrose/crunch-merchant-service/services/deliverect"
	"github.com/labstack/echo/v4"
)

const (
	DELIVERECT_SERVICE_CONTEXT_KEY = "deliverect_service"
)

func DeliverectServiceContext(service deliverect.DeliverectInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DELIVERECT_SERVICE_CONTEXT_KEY, service)
			return next(c)
		}
	}
}
