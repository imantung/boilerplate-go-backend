package database

import (
	"database/sql"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"

	_ "github.com/jackc/pgx/v5/stdlib" // NOTE: load postgres driver
)

var _ = di.Provide(NewPostgres)

func NewPostgres(cfg *config.Config) (*sql.DB, error) {
	pg := cfg.Postgres

	db, err := sql.Open("pgx", cfg.Postgres.Source)
	if err != nil {
		return nil, err
	}

	// NOTE: Set connection pool to prevent bottleneck in database side. Learn more: https://www.alexedwards.net/blog/configuring-sqldb
	db.SetConnMaxLifetime(pg.ConnMaxLifetime)
	db.SetMaxIdleConns(pg.MaxIdleConns)
	db.SetMaxOpenConns(pg.MaxOpenConns)

	return db, nil
}
