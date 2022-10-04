package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/routes"
	"github.com/joelrose/crunch-merchant-service/services/deliverect"
	"github.com/joelrose/crunch-merchant-service/services/http_client"
	redisService "github.com/joelrose/crunch-merchant-service/services/redis"
	"github.com/joelrose/crunch-merchant-service/services/tracing"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v73"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func run() error {
	log.SetLevel(log.DEBUG)

	c, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	ctx := context.Background()
	otelShutdown, err := tracing.InstallExportPipeline(ctx)
	if err != nil {
		return fmt.Errorf("error setting up OTel SDK - %e", err)
	}
	defer func() {
		if err := otelShutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	stripe.Key = c.Stripe.SecretKey
	stripe.SetHTTPClient(&http.Client{
		Timeout:   80 * time.Second, // defaultHTTPTimeout
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	})

	database, err := db.NewDatabase(c.DatabaseUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	redis := redisService.NewClient(c.RedisUrl)
	httpClient := http_client.NewClient()
	deliverect := deliverect.NewDeliverectService(c, redis, httpClient)

	e := echo.New()

	e.Use(otelecho.Middleware("crunch-backend-service"))
	e.Use(defaultMiddleware.RequestID())
	e.Use(defaultMiddleware.Logger())
	e.Use(middleware.ConfigContext(&c))
	e.Use(middleware.DatabaseContext(&db.DB{Sqlx: *database}))
	e.Use(middleware.RedisContext(redis))
	e.Use(middleware.DeliverectServiceContext(deliverect))

	routes.SetupRoutes(e, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return e.Start(":" + port)
}

// @title           Crunch Backend API
// @version         1.0
// @description     This is the Crunch Backend API
// @BasePath  /api/v1
// @host localhost:8080
// @securityDefinitions.apikey FirebaseToken
// @in header
// @name Authorization
// @securityDefinitions.apikey Auth0Token
// @in header
// @name Authorization
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
