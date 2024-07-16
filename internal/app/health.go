package app

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type AppHealth map[string]error

var _ = di.Provide(Health)

func Health(pg *sql.DB, redis *redis.Client) AppHealth {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return AppHealth{
		"postgres": pg.PingContext(ctx),
		"redis":    redis.Ping(ctx).Err(),
	}
}

func (h AppHealth) Handle(ec echo.Context) error {
	// NOTE: disable cache
	ec.Response().Header().Set("Expires", "0")
	ec.Response().Header().Set("Pragma", "no-cache")
	ec.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	code := http.StatusOK
	msg := map[string]string{}
	for k, v := range h {
		if v != nil {
			msg[k] = v.Error()
			code = http.StatusServiceUnavailable
		} else {
			msg[k] = "ok"
		}
	}

	return ec.JSON(code, msg)
}
