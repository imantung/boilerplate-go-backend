package logger

import (
	"net/http"
	"strconv"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/auth"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var SlowLatencyLimit = 500 * time.Millisecond

func HTTPRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			req := c.Request()
			res := c.Response()

			latency := time.Since(start)

			var evt *zerolog.Event
			if err != nil {
				evt = log.Error()
			} else if latency > SlowLatencyLimit {
				evt = log.Warn()
			} else {
				evt = log.Info()
			}

			// NOTE: check middleware.RequestID()
			if reqID := res.Header().Get(echo.HeaderXRequestID); reqID != "" {
				evt.Str("req_id", reqID)
			}

			// NOTE: check auth.ValidateToken()
			if userID := req.Header.Get(auth.HeaderXUserID); userID != "" {
				evt.Str("user_id", userID)
			}

			evt.Str("remote_ip", c.RealIP())
			evt.Str("host", req.Host)
			evt.Str("method", req.Method)
			evt.Str("uri", req.RequestURI)
			evt.Str("user_agent", req.UserAgent())
			evt.Int("status", res.Status)
			evt.Str("referer", req.Referer())
			evt.Dur("latency", latency)
			evt.Str("latency_human", latency.String())
			evt.Str("bytes_in", requestSize(req))
			evt.Str("bytes_out", strconv.FormatInt(res.Size, 10))

			evt.Send()

			return err
		}
	}
}

func requestSize(req *http.Request) string {
	cl := req.Header.Get(echo.HeaderContentLength)
	if cl == "" {
		return "0"
	}
	return cl
}
