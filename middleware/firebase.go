package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const (
	FIREBASE_CONTEXT_KEY = "token"
)

func FirebaseAuth(firebaseConfig string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			opt := option.WithCredentialsJSON([]byte(firebaseConfig))
			app, err := firebase.NewApp(context.Background(), nil, opt)
			if err != nil {
				log.Errorf("error initializing app: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			client, err := app.Auth(context.Background())
			if err != nil {
				log.Errorf("error getting Auth client: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			auth := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(auth, "Bearer ", "", 1)
			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			c.Set(FIREBASE_CONTEXT_KEY, token)
			return next(c)
		}
	}
}
