package main

import (
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joelrose/crunch-merchant-service/api/channel/deliverect"
	"github.com/joelrose/crunch-merchant-service/api/menus"
	"github.com/joelrose/crunch-merchant-service/api/orders"
	"github.com/joelrose/crunch-merchant-service/api/stores"
	"github.com/joelrose/crunch-merchant-service/api/users"
	"github.com/joelrose/crunch-merchant-service/api/whitelist"
	"github.com/joelrose/crunch-merchant-service/auth_middleware"
	"github.com/stripe/stripe-go/v73"

	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func setupRoutes(e *echo.Echo, config config.Config) {
	okHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}

	e.GET("/", okHandler)

	apiGroup := e.Group("/api")

	apiGroup.GET("/whitelist", whitelist.IsWhitelisted)
	apiGroup.GET("/stores", stores.GetStores)
	apiGroup.GET("/menus/:id", menus.GetMenu)

	channelGroup := apiGroup.Group("/channel")

	deliverectGroup := channelGroup.Group("/deliverect")

	deliverectGroup.POST("/channel_status", deliverect.DeliverectChannelStatus)
	deliverectGroup.POST("/busy_mode", deliverect.DeliverectBusyMode)
	deliverectGroup.POST("/menu_push", deliverect.DeliverectMenuPush)

	dashboardGroup := apiGroup.Group("/dashboard", auth_middleware.Auth0Auth())

	dashboardGroup.GET("/status", okHandler)

	usersGroup := apiGroup.Group("/users", auth_middleware.FirebaseAuth(config.FirebaseConfig))

	usersGroup.GET("/status", okHandler)

	usersGroup.GET("/", users.GetUser)
	usersGroup.POST("/", users.CreateUser)

	ordersGroup := apiGroup.Group("/orders", auth_middleware.FirebaseAuth(config.FirebaseConfig))

	ordersGroup.POST("/", orders.CreateOrder)
}

func dbMiddleware(db *db.DB) echo.MiddlewareFunc {
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
	e.Use(dbMiddleware(&db.DB{Sqlx: *database}))

	setupRoutes(e, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
