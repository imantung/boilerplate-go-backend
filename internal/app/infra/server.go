package infra

import (
	"github.com/imantung/boilerplate-go-backend/internal/app/controller"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/dig"
)

type (
	Server struct {
		dig.In
		// add new controller here..
		controller.TestCntrl
	}
)

//go:generate mkdir -p ../../generated/openapi
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 --package openapi -generate types,server,spec --o ../../generated/openapi/openapi.go ../../../api/api-spec.yml

var _ openapi.ServerInterface = (*Server)(nil)
var _ = di.Provide(NewEcho)

func NewEcho(server Server) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // required for swagger ui
	openapi.RegisterHandlers(e, server)
	return e
}
