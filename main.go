package main

import (
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joelrose/crunch-merchant-service/api/users"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

const (
	certificateDirectory = "certificates/auth0.pem"
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

	internal.GET("/users", users.GetUsers)
}

func getPemCert() []byte {
	pem, err := os.ReadFile(certificateDirectory)
	if err != nil {
		log.Fatal("cannot read pem file")
	}

	return pem
}

func dbMiddleware(db *db.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dbContextKey", db)
			return next(c)
		}
	}
}

func main() {
	c := config.LoadConfig()

	log.SetLevel(log.INFO)

	database := db.NewDatabase(c.DatabaseUrl)

	e := echo.New()

	e.Use(dbMiddleware(&db.DB{Sqlx: *database}))

	setupRoutes(e, c.InternalAuthToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
