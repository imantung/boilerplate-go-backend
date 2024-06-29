package auth

import (
	"crypto/subtle"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _ = di.Provide(NewBasicAuthValidator)

func NewBasicAuthValidator(cfg *config.Config) middleware.BasicAuthValidator {
	return func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(cfg.BasicAuth.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.BasicAuth.Secret)) == 1 {
			return true, nil
		}
		return false, nil
	}
}
