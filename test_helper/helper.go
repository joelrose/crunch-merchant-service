package test_helper

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/joelrose/crunch-merchant-service/middleware"
	"github.com/joelrose/crunch-merchant-service/test_helper/mock_db"
	"github.com/joelrose/crunch-merchant-service/utils"
	"github.com/labstack/echo/v4"
)

func NewEchoMock(t *testing.T, requestBody string) (*httptest.ResponseRecorder, echo.Context, *sqlx.DB, sqlmock.Sqlmock) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	return rec, c, sqlxDB, mock
}

func NewRequest(t *testing.T, requestMethod string, requestBody string) (*httptest.ResponseRecorder, echo.Context, *mock_db.MockDBInterface) {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	req := httptest.NewRequest(requestMethod, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mockDB := mock_db.NewMockDBInterface(ctrl)

	c.Set(middleware.DATABASE_CONTEXT_KEY, mockDB)

	return rec, c, mockDB
}
