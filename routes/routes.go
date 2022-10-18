package routes

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/api/app/v1/orders"
	"github.com/joelrose/crunch-merchant-service/api/app/v1/store"
	"github.com/joelrose/crunch-merchant-service/api/app/v1/stores"
	"github.com/joelrose/crunch-merchant-service/api/app/v1/users"
	"github.com/joelrose/crunch-merchant-service/api/app/v1/whitelist"
	"github.com/joelrose/crunch-merchant-service/api/channel/v1/deliverect"
	"github.com/joelrose/crunch-merchant-service/api/dashboard/v1"
	"github.com/joelrose/crunch-merchant-service/api/webhook/v1/stripe"
	"github.com/joelrose/crunch-merchant-service/config"
	_ "github.com/joelrose/crunch-merchant-service/docs"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	okHandler = func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}
)

func setupAppRoutes(e *echo.Echo, config config.Config) {
	appV1 := e.Group("/app/v1")

	appV1.Use(middleware.AppVersion(config.AllowedAppBuilds))

	appV1.POST("/whitelist", whitelist.IsWhitelisted)
	appV1.GET("/stores", stores.GetStoresOverview)
	appV1.GET("/store/:id", store.GetStore)

	usersGroup := appV1.Group("/users", middleware.FirebaseAuth(config.FirebaseConfig))
	usersGroup.GET("", users.GetUser)
	usersGroup.POST("", users.CreateUser)

	ordersGroup := appV1.Group("/orders", middleware.FirebaseAuth(config.FirebaseConfig))
	ordersGroup.POST("", orders.CreateOrder)
	ordersGroup.GET("", orders.GetOrders)
}

func setupIntegrationRoutes(e *echo.Echo, config config.Config) {
	e.POST("/webhook/v1/stripe", stripe.WebhookHandler)

	deliverectGroup := e.Group("/v1/channel/deliverect")

	deliverectGroup.POST("/channel_status", deliverect.ChannelStatus)
	deliverectGroup.POST("/menu_push", deliverect.MenuPush)
	deliverectGroup.POST("/snooze_unsnooze", deliverect.SnoozeUnsnooze)
	deliverectGroup.POST("/busy_mode", deliverect.BusyMode)
	deliverectGroup.POST("/prep_time", deliverect.PreparationTime)
	deliverectGroup.POST("/order_status", deliverect.OrderStatus)
}

func setupDashboardRoutes(e *echo.Echo, config config.Config) {
	dashboardGroup := e.Group("/dashboard/v1")

	// TODO: only allow specific origins and methods
	dashboardGroup.Use(defaultMiddleware.CORSWithConfig(defaultMiddleware.CORSConfig{
		Skipper:      defaultMiddleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	validator := middleware.NewValidator(config.Auth0.Audience, config.Auth0.Authority)
	dashboardGroup.Use(validator.Middleware())

	dashboardGroup.GET("/orders", dashboard.GetOrders)
	dashboardGroup.GET("/menu", dashboard.GetMenu)
	dashboardGroup.GET("/products", dashboard.GetProducts)
}

func SetupRoutes(e *echo.Echo, config config.Config) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", okHandler)

	setupAppRoutes(e, config)
	setupIntegrationRoutes(e, config)
	setupDashboardRoutes(e, config)
}
