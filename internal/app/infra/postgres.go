package infra

import (
	"database/sql"
	"fmt"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var _ = di.Provide(NewPostgres)

func NewPostgres(cfg *Config) (*sql.DB, error) {
	pg := cfg.Postgres
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pg.DBUser, pg.DBPass, pg.Host, pg.Port, pg.DBName,
	)

	db, err := sql.Open("pgx", conn)
	if err != nil {
		return nil, err
	}

	// NOTE: Set connection pool to prevent bottleneck in database side
	// - https://go.dev/doc/database/manage-connections
	// - https://www.alexedwards.net/blog/configuring-sqldb
	db.SetConnMaxLifetime(pg.ConnMaxLifetime)
	db.SetMaxIdleConns(pg.MaxIdleConns)
	db.SetMaxOpenConns(pg.MaxOpenConns)

	return db, nil
}
