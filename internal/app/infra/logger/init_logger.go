package logger

import (
	"io"

	"os"
	"runtime/debug"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
)

func InitLogger(cfg *config.Config, e *echo.Echo) {

	var w io.Writer

	if cfg.Debug {
		w = zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		}
	} else {
		w = os.Stderr
	}

	buildInfo, _ := debug.ReadBuildInfo()

	// NOTE: update zerolog global logger
	log.Logger = zerolog.New(w).With().
		Timestamp().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	// NOTE: update echo server logger
	e.Logger = lecho.From(log.Logger)

}
