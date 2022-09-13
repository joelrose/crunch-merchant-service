package main

import (
	"os"

	"github.com/go-redis/redis/v9"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/routes"
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v73"
)

func newRedis(redisUrl string) *redis.Client {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalf("failed to parse redis url: %v", err)
	}

	return redis.NewClient(opt)
}

func main() {
	log.SetLevel(log.DEBUG)

	c := config.LoadConfig()

	stripe.Key = c.StripeKey

	database := db.NewDatabase(c.DatabaseUrl)

	redis := newRedis(c.RedisUrl)

	e := echo.New()

	e.Use(defaultMiddleware.Logger())
	e.Use(middleware.DatabaseContext(&db.DB{Sqlx: *database}))
	e.Use(middleware.RedisContext(redis))

	routes.SetupRoutes(e, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
