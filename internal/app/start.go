package app

import (
	_ "expvar" // NOTE: enable `/debug/vars` endpoint
	"net/http"
	_ "net/http/pprof" //  NOTE: enable `/debug/pprof` endpoint

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/auth"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/logger"
	"github.com/imantung/boilerplate-go-backend/internal/app/service"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
	"go.uber.org/dig"
)

type App struct {
	dig.In

	Config *config.Config
	Health HealthMap
	Logger *logger.Handler

	Oauth *auth.OAuthHandler
	Basic *auth.BasicAuthHandler

	service.EmployeeSvc
	service.ClockSvc
}

var _ oapi.StrictServerInterface = (*App)(nil)

var e = echo.New()

func Start(app App) error {
	e.HideBanner = true
	e.Debug = app.Config.Debug

	log.Logger = app.Logger.ZeroLogger
	e.Logger = app.Logger.LechoLogger

	e.Pre(
		middleware.RequestID(),
	)

	e.Use(
		lecho.Middleware(app.Logger.LechoConfig),
		middleware.CORS(),
		middleware.Recover(),
		middleware.Secure(),
	)

	group := e.Group("api", app.Oauth.ValidateToken)
	server := oapi.NewStrictHandler(app, nil)
	oapi.RegisterHandlers(group, server) // NOTE: register open-api endpoints

	e.Any("/oauth/authorize", app.Oauth.HandleAuthorizeRequest)
	e.Any("/oauth/token", app.Oauth.HandleTokenRequest)

	e.File("/swagger/api-spec.yaml", "api/api-spec.yaml")
	e.Static("/swagger/ui", "api/swagger-ui")

	basicAuth := middleware.BasicAuth(app.Basic.Validate)
	e.Any("/health", app.Health.Handle, basicAuth)
	e.GET("/debug/*/*", echo.WrapHandler(http.DefaultServeMux), basicAuth)

	return e.Start(app.Config.Address)
}