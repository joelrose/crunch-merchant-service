package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/test_helper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUserFound(t *testing.T) {
	rec, c, mockDB := test_helper.NewRequest(t, http.MethodGet, "")

	c.Set(middleware.FIREBASE_CONTEXT_KEY, &test_helper.MockToken)

	mockDB.
		EXPECT().
		GetUserByFirebaseId(gomock.Eq(test_helper.MockTokenUID)).
		Return(test_helper.MockUser, nil).
		Times(1)

	if assert.NoError(t, GetUser(c)) {
		responseJson, _ := json.Marshal(test_helper.MockUser)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(responseJson)+"\n", rec.Body.String())
	}
}

func TestGetUserNotFound(t *testing.T) {
	_, c, mockDB := test_helper.NewRequest(t, http.MethodGet, "")

	c.Set(middleware.FIREBASE_CONTEXT_KEY, &test_helper.MockToken)

	mockDB.
		EXPECT().
		GetUserByFirebaseId(gomock.Eq(test_helper.MockTokenUID)).
		Return(models.User{}, sql.ErrNoRows).
		Times(1)

	err := GetUser(c)
	if assert.NotNil(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusNotFound, he.Code)
		} else {
			t.Fatal("Error is not an HTTPError")
		}
	}
}
