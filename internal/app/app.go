package app

import (
	"context"
	"database/sql"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/controller"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/auth"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/logger"
	"github.com/imantung/boilerplate-go-backend/internal/generated/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
	"go.uber.org/dig"
	"go.uber.org/multierr"

	_ "expvar" // enable `/debug/vars` endpoint
	"net/http"
	_ "net/http/pprof" // enable `/debug/pprof` endpoint
)

type (
	App struct {
		dig.In

		Config *config.Config
		Health HealthMap
		Logger *logger.Handler

		Oauth *auth.OAuthHandler
		Basic *auth.BasicAuthHandler

		// NOTE: add controller without variable name belows
		controller.EmployeeCntrl
	}
)

var _ openapi.ServerInterface = (*App)(nil) // NOTE: server must be implemented `openapi.server`. The functions are defined in the controllers
var _ = di.Provide(Health)

var (
	e = echo.New()
)

func Start(app App) error {
	e.HideBanner = true
	e.Debug = app.Config.Debug

	log.Logger = app.Logger.ZeroLogger
	e.Logger = app.Logger.LechoLogger

	e.Use(
		middleware.RequestID(), // NOTE: should be prior lecho middleware to append log field `request_id`
		lecho.Middleware(app.Logger.LechoConfig),
		middleware.CORS(),
		middleware.Recover(),
		middleware.Secure(),
	)

	group := e.Group("api", app.Oauth.ValidateToken)
	openapi.RegisterHandlers(group, app) // NOTE: register open-api endpoints

	e.Any("/oauth/authorize", app.Oauth.HandleAuthorizeRequest)
	e.Any("/oauth/token", app.Oauth.HandleTokenRequest)

	e.File("/swagger/api-spec.yaml", "api/api-spec.yaml")
	e.Static("/swagger/ui", "api/swagger-ui")

	basicAuth := middleware.BasicAuth(app.Basic.Validate)
	e.Any("/health", app.Health.Handle, basicAuth)
	e.GET("/debug/*/*", echo.WrapHandler(http.DefaultServeMux), basicAuth)

	return e.Start(app.Config.Address)
}

func Stop(db *sql.DB) error {
	log.Info().Msg("Gracefully stop the service")
	ctx := context.Background()

	var err error
	err = multierr.Append(err, e.Shutdown(ctx))
	err = multierr.Append(err, db.Close())

	return err
}

func Health(db *sql.DB) HealthMap {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return HealthMap{
		"postgres": db.PingContext(ctx),
	}

	// NOTE: Learn more about health check api at https://testfully.io/blog/api-health-check-monitoring/
}
