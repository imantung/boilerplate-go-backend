package logger

import (
	"io"
	"os"
	"runtime/debug"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

type (
	Handler struct {
		ZeroLogger  zerolog.Logger
		LechoLogger *lecho.Logger
		LechoConfig lecho.Config
	}
)

var _ = di.Provide(NewHandler)

func NewHandler(cfg *config.Config) *Handler {
	buildInfo, _ := debug.ReadBuildInfo()

	var w io.Writer

	if cfg.Debug {
		w = zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		}
	} else {
		w = os.Stderr
	}

	zeroLogger := zerolog.New(w).With().
		// Caller(). // NOTE: uncomment to put caller to the log
		Timestamp().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	lechoLogger := lecho.From(zeroLogger)
	lechoConfig := lecho.Config{
		Logger:              lechoLogger,
		RequestIDKey:        "request_id",
		RequestLatencyLevel: zerolog.WarnLevel,
		RequestLatencyLimit: 500 * time.Millisecond,
	}
	return &Handler{
		ZeroLogger:  zeroLogger,
		LechoLogger: lechoLogger,
		LechoConfig: lechoConfig,
	}
}
