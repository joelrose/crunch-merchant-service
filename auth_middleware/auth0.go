package auth_middleware

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	certificateDirectory = "certificates/auth0.pem"
)

func getPemCert() []byte {
	pem, err := os.ReadFile(certificateDirectory)
	if err != nil {
		log.Fatal("cannot read pem file")
	}

	return pem
}

func Auth0Auth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: getPemCert(),
	})
}
