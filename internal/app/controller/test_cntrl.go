package controller

import (
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	TestCntrl interface {
		GetTest(ctx echo.Context) error
	}
	TestCntrlImpl struct{}
)

var _ = di.Provide(NewTestCntrl)

func NewTestCntrl() TestCntrl {
	return &TestCntrlImpl{}
}

func (s *TestCntrlImpl) GetTest(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "test")
}
