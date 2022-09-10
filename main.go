package main

import (
	"net/http"
	"os"

	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const (
	certificateDirectory = "certificates/crunch-app.pem"
)

func setupRoutes(e *echo.Echo, internalAuthKey string) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	apiGroup := e.Group("/api")

	// dashboard routes api/dashboard

	dashboard := apiGroup.Group("/dashboard")

	dashboard.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: getPemCert(),
	}))

	dashboard.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	dashboard.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	// internal routes: api/internal

	internal := apiGroup.Group("/internal")

	internal.Use(middleware.KeyAuth(
		func(auth string, c echo.Context) (bool, error) {
			return auth == internalAuthKey, nil
		},
	))

	internal.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
}

func getPemCert() []byte {
	pem, err := os.ReadFile(certificateDirectory)
	if err != nil {
		log.Fatal("cannot read pem file")
	}

	return pem
}

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	c := config.LoadConfig()

	setupRoutes(e, c.InternalAuthToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
