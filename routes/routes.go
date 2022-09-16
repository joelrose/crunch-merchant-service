package routes

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/api/v1/channel/deliverect"
	"github.com/joelrose/crunch-merchant-service/api/v1/orders"
	"github.com/joelrose/crunch-merchant-service/api/v1/store"
	"github.com/joelrose/crunch-merchant-service/api/v1/stores"
	"github.com/joelrose/crunch-merchant-service/api/v1/users"
	"github.com/joelrose/crunch-merchant-service/api/v1/webhook"
	"github.com/joelrose/crunch-merchant-service/api/v1/whitelist"
	"github.com/joelrose/crunch-merchant-service/config"
	_ "github.com/joelrose/crunch-merchant-service/docs"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRoutes(e *echo.Echo, config config.Config) {
	okHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", okHandler)

	apiV1 := e.Group("/api/v1")

	apiV1.POST("/whitelist", whitelist.IsWhitelisted)
	apiV1.GET("/stores", stores.GetStoresOverview)
	apiV1.GET("/store/:id", store.GetStore)

	apiV1.POST("/webhook/stripe", webhook.HandleStripe)

	channelGroup := apiV1.Group("/channel")
	deliverectGroup := channelGroup.Group("/deliverect")

	deliverectGroup.POST("/channel_status", deliverect.ChannelStatus)

	deliverectGroup.POST("/menu_push", deliverect.MenuPush)
	deliverectGroup.POST("/snooze_unsnooze", deliverect.SnoozeUnsnooze)
	deliverectGroup.POST("/busy_mode", deliverect.BusyMode)
	deliverectGroup.POST("/prep_time", deliverect.PreparationTime)
	deliverectGroup.POST("/order_status", deliverect.OrderStatus)

	dashboardGroup := apiV1.Group("/dashboard", middleware.Auth0Auth())

	dashboardGroup.GET("/status", okHandler)

	usersGroup := apiV1.Group("/users", middleware.FirebaseAuth(config.FirebaseConfig))

	usersGroup.GET("/status", okHandler)

	usersGroup.GET("", users.GetUser)
	usersGroup.POST("", users.CreateUser)

	ordersGroup := apiV1.Group("/orders", middleware.FirebaseAuth(config.FirebaseConfig))

	ordersGroup.POST("", orders.CreateOrder)
	ordersGroup.GET("", orders.GetOrders)
}
