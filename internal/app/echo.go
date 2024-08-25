package app

import (
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/logger"
	"github.com/labstack/echo/v4"
)

var _ = di.Provide(NewEcho)

func NewEcho(router Router, cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Debug = cfg.Debug

	logger.InitLogger(cfg, e)
	router.SetRoute(e)
	return e
}
