package whitelist

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joelrose/crunch-merchant-service/test_helper"
	"github.com/stretchr/testify/assert"
)

func TestIsWhitelistedFound(t *testing.T) {
	id := "found"
	request := WhitelistRequest{
		Id: id,
	}

	requestJson, _ := json.Marshal(request)
	rec, c, mockDB := test_helper.NewRequest(t, http.MethodPost, string(requestJson))
	mockDB.
		EXPECT().
		IsWhitelisted(gomock.Eq("found")).
		Return(true).
		Times(1)

	if assert.NoError(t, IsWhitelisted(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "true\n", rec.Body.String())
	}
}

func TestIsWhitelistedNotFound(t *testing.T) {
	id := "not_found"
	request := WhitelistRequest{
		Id: id,
	}

	requestJson, _ := json.Marshal(request)
	rec, c, mockDB := test_helper.NewRequest(t, http.MethodPost, string(requestJson))

	mockDB.
		EXPECT().
		IsWhitelisted(gomock.Eq(id)).
		Return(false).
		Times(1)

	if assert.NoError(t, IsWhitelisted(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "false\n", rec.Body.String())
	}
}
