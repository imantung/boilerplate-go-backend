package infra

import (
	"github.com/imantung/boilerplate-go-backend/internal/app/controller"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/oauth"
	"github.com/imantung/boilerplate-go-backend/internal/generated/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/dig"
)

type (
	Server struct {
		dig.In
		controller.HelloCntrl
		// add new controller here..
	}
)

//go:generate mkdir -p ../../generated/openapi
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 --package openapi -generate types,server,spec --o ../../generated/openapi/openapi.go ../../../api/api-spec.yaml

var _ openapi.ServerInterface = (*Server)(nil)
var _ = di.Provide(NewEcho)

func NewEcho(server Server, oauthHandler *oauth.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	group := e.Group("api", oauthHandler.Middleware())
	openapi.RegisterHandlers(group, server)

	e.Any("/oauth/authorize", oauthHandler.HandleAuthorizeRequest)
	e.Any("/oauth/token", oauthHandler.HandleTokenRequest)

	e.File("/swagger/api-spec.yaml", "api/api-spec.yaml")
	e.Static("/swagger/ui", "api/swagger-ui")

	return e
}
