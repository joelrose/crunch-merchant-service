package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/joelrose/crunch-merchant-service/test_helper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userRequest = dtos.CreateUserRequest{
		LanguageCode: "en",
		Firstname:    "John",
		Lastname:     "Doe",
	}
)

func TestCreateExisitingUser(t *testing.T) {
	_, c, mockDB := test_helper.NewRequest(t, http.MethodPost, "")
	c.Set(middleware.FIREBASE_CONTEXT_KEY, &test_helper.MockToken)

	mockDB.
		EXPECT().
		GetUserByFirebaseId(gomock.Eq(test_helper.MockTokenUID)).
		Return(test_helper.MockUser, nil).
		Times(1)

	err := CreateUser(c)
	if assert.NotNil(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusConflict, he.Code)
		} else {
			t.Fatal("Error is not an HTTPError")
		}
	}
}

func TestCreateUserSuccessful(t *testing.T) {
	requestJson, _ := json.Marshal(userRequest)
	rec, c, mockDB := test_helper.NewRequest(t, http.MethodPost, string(requestJson))
	c.Set(middleware.FIREBASE_CONTEXT_KEY, &test_helper.MockToken)

	mockDB.
		EXPECT().
		GetUserByFirebaseId(gomock.Eq(test_helper.MockTokenUID)).
		Return(models.User{}, sql.ErrNoRows).
		Times(1)

	mockDB.
		EXPECT().
		CreateUser(gomock.Eq(test_helper.MockTokenUID), gomock.Eq(userRequest)).
		Return(nil).
		Times(1)

	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
