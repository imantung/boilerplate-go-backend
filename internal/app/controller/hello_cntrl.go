package controller

import (
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	HelloCntrl interface {
		GetHello(ctx echo.Context) error
	}
	HelloCntrlImpl struct {
	}
)

var _ = di.Provide(NewHelloCntrl)

func NewHelloCntrl() HelloCntrl {
	return &HelloCntrlImpl{}
}

func (s *HelloCntrlImpl) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
