package dashboard

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetProducts godoc
// @Summary      Get all products from a store
// @Tags         dashboard
// @Accept       json
// @Produce      json
// @Security 	 Auth0Token
// @Success      200  {object}  []dtos.GetStoreProduct
// @Success      400  {object} 	error
// @Failure      500  {object}  error
// @Router       /dashboard/products [get]
func GetProducts(c echo.Context) error {
	db := c.Get(middleware.DATABASE_CONTEXT_KEY).(db.DBInterface)

	/*userId := c.Get(middleware.AUTH0_USER_ID_CONTEXT_KEY)
	userIdString := userId.(string)

	storeId, err := db.GetStoreByMerchantUserId(userIdString)
	if err != nil {
		log.Errorf("Failed to get store by merchant user id: %v", err)
		return echo.NewHTTPError(http.StatusForbidden)
	}
	*/
	products, err := db.GetProducts(uuid.MustParse("9142ac52-e5a4-4ad8-8811-240c1f389ece"))
	if err != nil {
		log.Errorf("failed to get orders: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, products)
}
