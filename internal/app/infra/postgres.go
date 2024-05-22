package infra

import (
	"database/sql"
	"fmt"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
)

var _ = di.Provide(NewPostgres)

func NewPostgres(cfg *Config) (*sql.DB, error) {
	pg := cfg.Postgres
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pg.DBUser, pg.DBPass, pg.Host, pg.Port, pg.DBName,
	)

	// NOTE: Alternatively we can use github.com/jackc/pgx which is faster compared
	// with database/sql package. However the library didn't have same interface with
	// database/sql which cumbersome when use other library that require sql.DB
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	// NOTE: Set connection pool to prevent bottleneck in database side
	// Read more:
	// - https://go.dev/doc/database/manage-connections
	// - https://www.alexedwards.net/blog/configuring-sqldb
	db.SetConnMaxLifetime(pg.ConnMaxLifetime)
	db.SetMaxIdleConns(pg.MaxIdleConns)
	db.SetMaxOpenConns(pg.MaxOpenConns)

	return db, nil
}
