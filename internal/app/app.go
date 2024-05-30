package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/controller"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/auth"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: trigger DI provide for database connection
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

//go:generate mkdir -p ../../generated/openapi
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 --package openapi -generate types,server,spec --o ../generated/openapi/openapi.go ../../api/api-spec.yaml

type (
	App struct {
		dig.In

		Config *config.Config
		Oauth  *auth.OAuthHandler
		Health HealthMap

		// NOTE: add controller without variable name belows
		controller.HelloCntrl
	}
)

var _ openapi.ServerInterface = (*App)(nil) // NOTE: server must be implemented `openapi.server`. The functions are defined in the controllers
var _ = di.Provide(Health)

var (
	e = echo.New()
)

func Start(app App) error {
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	group := e.Group("api", app.Oauth.ValidateTokenMW())
	openapi.RegisterHandlers(group, app)

	e.Any("/oauth/authorize", app.Oauth.HandleAuthorizeRequest)
	e.Any("/oauth/token", app.Oauth.HandleTokenRequest)

	e.File("/swagger/api-spec.yaml", "api/api-spec.yaml")
	e.Static("/swagger/ui", "api/swagger-ui")

	e.Any("/health", app.Health.Handle)

	return e.Start(app.Config.Address)
}

func Stop(db *sql.DB) error {
	log.Printf("Gracefully stop the service")
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
