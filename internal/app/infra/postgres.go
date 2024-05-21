package infra

import (
	"database/sql"

	"github.com/imantung/empl-clocking-backend-go/internal/app/infra/di"
)

var _ = di.Provide(NewPostgres)

func NewPostgres(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// NOTE: set connection pool to prevent bottleneck in database side
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}
