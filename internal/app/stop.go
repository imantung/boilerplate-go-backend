package app

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"go.uber.org/multierr"
)

func Stop(db *sql.DB) error {
	log.Info().Msg("Gracefully stop the service")
	ctx := context.Background()

	var err error
	err = multierr.Append(err, e.Shutdown(ctx))
	err = multierr.Append(err, db.Close())

	return err
}
