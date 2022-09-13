package main

import (
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stripe/stripe-go/v73"

	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func databaseMiddleware(db *db.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func main() {
	log.SetLevel(log.DEBUG)

	c := config.LoadConfig()

	stripe.Key = c.StripeKey

	database := db.NewDatabase(c.DatabaseUrl)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(databaseMiddleware(&db.DB{Sqlx: *database}))

	setupRoutes(e, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
