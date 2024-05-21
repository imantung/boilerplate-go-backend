package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/imantung/empl-clocking-backend-go/internal/app/infra"
	"github.com/labstack/echo/v4"
	"go.uber.org/multierr"
)

func Start(e *echo.Echo, cfg *infra.Config) error {
	return e.Start(cfg.ServerAddress)
}

func Stop(e *echo.Echo, db *sql.DB) error {
	log.Printf("Gracefully shutdown the service")
	ctx := context.Background()

	var err error
	err = multierr.Append(err, e.Shutdown(ctx))
	err = multierr.Append(err, db.Close())

	return err
}
