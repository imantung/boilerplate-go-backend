package logger

import (
	"io"
	"os"
	"runtime/debug"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/config"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/rs/zerolog"
)

var _ = di.Provide(NewZeroLogger)

func NewZeroLogger(cfg *config.Config) zerolog.Logger {
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

	return zerolog.New(w).With().
		// Caller(). // NOTE: uncomment to append caller to the log
		Timestamp().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}
