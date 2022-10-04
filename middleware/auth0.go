package middleware

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	AUTH0_USER_ID_CONTEXT_KEY = "user_id"
)

type Auth0 struct {
	Audience  string
	Authority string
}

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func NewValidator(audience string, authority string) *Auth0 {
	return &Auth0{
		Audience:  audience,
		Authority: authority,
	}
}

func (auth0 *Auth0) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")

			if len(authHeader) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "No Authorization header")
			}

			issuerURL, err := url.Parse(auth0.Authority)
			if err != nil {
				log.Errorf("Failed to parse the issuer url: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse the issuer url")
			}

			provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

			jwtValidator, err := validator.New(
				provider.KeyFunc,
				validator.RS256,
				issuerURL.String(),
				[]string{auth0.Audience},
				validator.WithCustomClaims(
					func() validator.CustomClaims {
						return &CustomClaims{}
					},
				),
				validator.WithAllowedClockSkew(time.Minute),
			)
			if err != nil {
				log.Errorf("Failed to set up the jwt validator")
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to set up the jwt validator")
			}

			user, err := jwtValidator.ValidateToken(context.Background(), authHeader[1])
			if err != nil {
				log.Infof("Failed to validate the token: %v", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to validate the token")
			}

			claims := user.(*validator.ValidatedClaims)
			c.Set(AUTH0_USER_ID_CONTEXT_KEY, claims.RegisteredClaims.Subject)

			return next(c)
		}
	}
}
