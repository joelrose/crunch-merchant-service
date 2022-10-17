package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	APP_VERSION_KEY      = "X-App-Version"
)

func AppVersion(allowedAppVersions []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			appVersion := c.Request().Header.Get(APP_VERSION_KEY)

			for _, allowedVersion := range allowedAppVersions {
				if appVersion == allowedVersion {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusPreconditionFailed)
		}
	}
}
