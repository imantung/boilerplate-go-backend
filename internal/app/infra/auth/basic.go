package auth

import (
	"crypto/subtle"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	BasicAuthHandler struct {
		Username string
		Secret   string
	}
)

var _ = di.Provide(NewBasicAuthHandler)

func NewBasicAuthHandler(cfg *config.Config) *BasicAuthHandler {
	return &BasicAuthHandler{
		Username: cfg.BasicAuth.Username,
		Secret:   cfg.BasicAuth.Secret,
	}
}

func (b *BasicAuthHandler) Validate(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare([]byte(username), []byte(b.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(b.Secret)) == 1 {
		return true, nil
	}
	return false, nil
}
