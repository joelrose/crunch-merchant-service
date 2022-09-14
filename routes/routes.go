package routes

import (
	"net/http"

	"github.com/joelrose/crunch-merchant-service/api/v1/channel/deliverect"
	"github.com/joelrose/crunch-merchant-service/api/v1/menus"
	"github.com/joelrose/crunch-merchant-service/api/v1/orders"
	"github.com/joelrose/crunch-merchant-service/api/v1/stores"
	"github.com/joelrose/crunch-merchant-service/api/v1/users"
	"github.com/joelrose/crunch-merchant-service/api/v1/whitelist"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, config config.Config) {
	okHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}

	e.GET("/", okHandler)

	apiV1 := e.Group("/api/v1")

	apiV1.GET("/whitelist", whitelist.IsWhitelisted)
	apiV1.GET("/stores", stores.GetStores)
	apiV1.GET("/menus/:id", menus.GetMenu)

	channelGroup := apiV1.Group("/channel")
	deliverectGroup := channelGroup.Group("/deliverect")

	deliverectGroup.POST("/channel_status", deliverect.DeliverectChannelStatus)
	deliverectGroup.POST("/busy_mode", deliverect.DeliverectBusyMode)
	deliverectGroup.POST("/menu_push", deliverect.DeliverectMenuPush)

	dashboardGroup := apiV1.Group("/dashboard", middleware.Auth0Auth())

	dashboardGroup.GET("/status", okHandler)

	usersGroup := apiV1.Group("/users", middleware.FirebaseAuth(config.FirebaseConfig))

	usersGroup.GET("/status", okHandler)

	usersGroup.GET("/", users.GetUser)
	usersGroup.POST("/", users.CreateUser)

	ordersGroup := apiV1.Group("/orders", middleware.FirebaseAuth(config.FirebaseConfig))

	ordersGroup.POST("/", orders.CreateOrder)
	ordersGroup.GET("/", orders.GetOrders)
}
