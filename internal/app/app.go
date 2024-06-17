package app

import (
	"context"
	"database/sql"
	_ "expvar"         // NOTE: enable `/debug/vars` endpoint
	_ "net/http/pprof" //  NOTE: enable `/debug/pprof` endpoint

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/logger" // NOTE: provide logger constructor
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
	"go.uber.org/multierr"
)

var _ oapi.StrictServerInterface = (*Router)(nil)
var _ = di.Provide(func() *echo.Echo {
	return echo.New()
})

func Start(e *echo.Echo, router Router, logger zerolog.Logger, cfg *config.Config) error {
	log.Logger = logger

	e.HideBanner = true
	e.Debug = cfg.Debug
	e.Logger = lecho.From(logger)

	router.SetRoute(e)

	return e.Start(cfg.Address)
}

func Stop(db *sql.DB, e *echo.Echo) error {
	log.Info().Msg("Gracefully stop the service")
	ctx := context.Background()

	var err error
	err = multierr.Append(err, e.Shutdown(ctx))
	err = multierr.Append(err, db.Close())

	return err
}
