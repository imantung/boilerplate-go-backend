package app

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

// NOTE: Learn more about health check api at
// - https://testfully.io/blog/api-health-check-monitoring/

type (
	HealthChecker struct {
		PG *sql.DB
	}
)

var _ = di.Provide(NewHealthChecker)

func NewHealthChecker(pg *sql.DB) *HealthChecker {
	return &HealthChecker{
		PG: pg,
	}
}

func (h *HealthChecker) healthMap() map[string]error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return map[string]error{
		"postgres": h.PG.PingContext(ctx),
	}
}

func (h *HealthChecker) httpCode(m map[string]error) int {
	for _, v := range m {
		if v != nil {
			return http.StatusServiceUnavailable
		}
	}
	return http.StatusOK
}

func (h *HealthChecker) message(m map[string]error) map[string]string {
	msg := map[string]string{}
	for k, v := range m {
		if v != nil {
			msg[k] = v.Error()
		} else {
			msg[k] = "ok"
		}
	}
	return msg
}

func (h *HealthChecker) Handle(ec echo.Context) error {
	m := h.healthMap()

	// NOTE: disable cache
	ec.Response().Header().Set("Expires", "0")
	ec.Response().Header().Set("Pragma", "no-cache")
	ec.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	code := h.httpCode(m)
	msg := h.message(m)

	return ec.JSON(code, msg)
}
