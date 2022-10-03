package stores

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/config"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/joelrose/crunch-merchant-service/test_helper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	storeId1       = uuid.New()
	storeId2       = uuid.New()
	mockOpenStores = []dtos.GetStoresOverviewResponse{
		{
			Id:                storeId1,
			Name:              "Store 1",
			Description:       "Store 1 Description",
			Address:           "Store 1 Address",
			AveragePickupTime: 10,
			AverageReview:     4.5,
			ReviewCount:       10,
			GoogleMapsLink:    "https://google.com/maps",
			PhoneNumber:       "123-456-7890",
			ImageUrl:          "https://google.com/image/1",
		},
		{
			Id:                storeId2,
			Name:              "Store 2",
			Description:       "Store 2 Description",
			Address:           "Store 2 Address",
			AveragePickupTime: 15,
			AverageReview:     3.2,
			ReviewCount:       199,
			GoogleMapsLink:    "https://google.com/maps",
			PhoneNumber:       "+4917654321",
			ImageUrl:          "https://google.com/image/2",
		},
	}
	mockOpeningHoursAlwaysOpen = []dtos.GetStoreOpeningHour{
		{
			DayOfWeek:      0,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
		{
			DayOfWeek:      1,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
		{

			DayOfWeek:      2,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
		{

			DayOfWeek:      3,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
		{

			DayOfWeek:      4,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
		{

			DayOfWeek:      5,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
		{

			DayOfWeek:      6,
			StartTimestamp: 0,
			EndTimestamp:   1440,
		},
	}
	mockOpeningHoursNeverOpen = []dtos.GetStoreOpeningHour{
		{
			DayOfWeek:      0,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
		{
			DayOfWeek:      1,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
		{

			DayOfWeek:      2,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
		{

			DayOfWeek:      3,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
		{

			DayOfWeek:      4,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
		{

			DayOfWeek:      5,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
		{

			DayOfWeek:      6,
			StartTimestamp: 0,
			EndTimestamp:   0,
		},
	}
)

func TestGetStoresOverviewNotFound(t *testing.T) {
	_, c, mockDB := test_helper.NewRequest(t, http.MethodGet, "")

	timezone, _ := time.LoadLocation("Europe/Berlin")
	c.Set(middleware.CONFIG_CONTEXT_KEY, config.Config{Timezone: timezone})

	mockDB.
		EXPECT().
		GetOpenStores().
		Return(nil, sql.ErrNoRows).
		Times(1)

	err := GetStoresOverview(c)
	if assert.NotNil(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusNotFound, he.Code)
		} else {
			t.Fatal("Error is not an HTTPError")
		}
	}
}

func TestGetStoresOverviewFound(t *testing.T) {
	rec, c, mockDB := test_helper.NewRequest(t, http.MethodGet, "")

	timezone, _ := time.LoadLocation("Europe/Berlin")
	c.Set(middleware.CONFIG_CONTEXT_KEY, config.Config{Timezone: timezone})

	mockDB.
		EXPECT().
		GetOpenStores().
		Return(mockOpenStores, nil).
		Times(1)

	firstOpeningHours := mockDB.
		EXPECT().
		GetOpeningHours(gomock.Eq(storeId1)).
		Return(mockOpeningHoursAlwaysOpen, nil).
		Times(1)

	secondOpeningHours := mockDB.
		EXPECT().
		GetOpeningHours(gomock.Eq(storeId2)).
		Return(mockOpeningHoursNeverOpen, nil).
		Times(1)

	gomock.InOrder(
		firstOpeningHours,
		secondOpeningHours,
	)

	if assert.NoError(t, GetStoresOverview(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		mockOpenStores[0].IsAvailable = true
		mockOpenStores[1].IsAvailable = false

		var response []dtos.GetStoresOverviewResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		assert.Equal(t, mockOpenStores, response)
	}
}
