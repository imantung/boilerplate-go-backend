package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/imantung/boilerplate-go-backend/internal/app/controller"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: trigger DI provide for database connection
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/oauth"
	"github.com/imantung/boilerplate-go-backend/internal/generated/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

//go:generate mkdir -p ../../generated/openapi
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 --package openapi -generate types,server,spec --o ../generated/openapi/openapi.go ../../api/api-spec.yaml

type (
	Server struct {
		dig.In

		// NOTE: add new controller below without variable name
		controller.HelloCntrl
	}
)

var _ openapi.ServerInterface = (*Server)(nil) // NOTE: server must be implemented `openapi.server`. The functions are defined in the controllers

var (
	e = echo.New()
)

func Start(
	server Server,
	cfg *config.Config,
	oauth *oauth.Handler,
	health HealthChecker,
) error {
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	group := e.Group("api", oauth.ValidateTokenMW())
	openapi.RegisterHandlers(group, server)

	e.Any("/oauth/authorize", oauth.HandleAuthorizeRequest)
	e.Any("/oauth/token", oauth.HandleTokenRequest)

	e.File("/swagger/api-spec.yaml", "api/api-spec.yaml")
	e.Static("/swagger/ui", "api/swagger-ui")

	e.Any("/health", health.Handle)

	return e.Start(cfg.Address)
}

func Stop(db *sql.DB) error {
	log.Printf("Gracefully shutdown the service")
	ctx := context.Background()

	var err error
	err = multierr.Append(err, e.Shutdown(ctx))
	err = multierr.Append(err, db.Close())

	return err
}
