package app

import (
	"context"
	"database/sql"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/multierr"
)

func Start(e *echo.Echo, cfg *config.Config) error {
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
