package app

import (
	"context"
	"database/sql"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/logger"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/multierr"
)

var _ oapi.StrictServerInterface = (*Router)(nil)
var _ = di.Provide(func() *echo.Echo {
	return echo.New()
})

func Start(e *echo.Echo, router Router, cfg *config.Config) error {
	e.HideBanner = true
	e.Debug = cfg.Debug

	logger.InitLogger(cfg, e)
	router.SetRoute(e)

	return e.Start(cfg.Address)
}

func Stop(db *sql.DB, e *echo.Echo) error {
	log.Info().Msg("Gracefully stop the service")
	ctx := context.Background()

	return multierr.Combine(
		e.Shutdown(ctx),
		db.Close(),
	)
}
