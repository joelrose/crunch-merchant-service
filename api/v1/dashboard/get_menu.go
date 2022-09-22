package dashboard

import (
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/domain/menus"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetMenu godoc
// @Summary      Get the menu for a store
// @Tags         dashboard
// @Accept       json
// @Produce      json
// @Security 	 Auth0Token
// @Success      200  {object}  []menus.MenuRedisModel
// @Success      400  {object} 	error
// @Failure      500  {object}  error
// @Router       /dashboard/menu [get]
func GetMenu(c echo.Context) error {
	db := c.Get(middleware.DATBASE_CONTEXT_KEY).(*db.DB)

	/*userId := c.Get(middleware.AUTH0_USER_ID_CONTEXT_KEY)
	userIdString := userId.(string)

	storeId, err := db.GetStoreByMerchantUserId(userIdString)
	if err != nil {
		log.Errorf("Failed to get store by merchant user id: %v", err)
		return echo.NewHTTPError(http.StatusForbidden)
	}*/

	rdb := c.Get(middleware.REDIS_CONTEXT_KEY).(*redis.Client)

	menuService := menus.NewMenuService(db, rdb, uuid.MustParse("69db0f24-7d29-4b95-9483-d3e1ab661c65"))
	menu, err := menuService.GetMenu()
	if err != nil {
		log.Debugf("failed to get menu: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, menu)
}
