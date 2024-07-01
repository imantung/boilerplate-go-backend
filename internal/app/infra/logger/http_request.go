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

// if config.Skipper(c) {
// 	return next(c)
// }

// var err error
// req := c.Request()
// res := c.Response()
// start := time.Now()

// id := req.Header.Get(config.RequestIDHeader)

// if id == "" {
// 	id = res.Header().Get(config.RequestIDHeader)
// }

// cloned := false
// logger := config.Logger

// if id != "" {
// 	logger = From(logger.log, WithField(config.RequestIDKey, id))
// 	cloned = true
// }

// if config.Enricher != nil {
// 	// to avoid mutation of shared instance
// 	if !cloned {
// 		logger = From(logger.log)
// 		cloned = true
// 	}

// 	logger.log = config.Enricher(c, logger.log.With()).Logger()
// }

// ctx := req.Context()

// if ctx == nil {
// 	ctx = context.Background()
// }

// // Pass logger down to request context
// c.SetRequest(req.WithContext(logger.WithContext(ctx)))
// c = NewContext(c, logger)

// if config.BeforeNext != nil {
// 	config.BeforeNext(c)
// }

// if err = next(c); err != nil {
// 	if config.HandleError {
// 		c.Error(err)
// 	}
// }

// stop := time.Now()
// latency := stop.Sub(start)
// var mainEvt *zerolog.Event
// if err != nil {
// 	mainEvt = logger.log.Err(err)
// } else if config.RequestLatencyLimit != 0 && latency > config.RequestLatencyLimit {
// 	mainEvt = logger.log.WithLevel(config.RequestLatencyLevel)
// } else {
// 	mainEvt = logger.log.WithLevel(logger.log.GetLevel())
// }

// var evt *zerolog.Event
// if config.NestKey != "" { // Start a new event (dict) if there's a nest key.
// 	evt = zerolog.Dict()
// } else {
// 	evt = mainEvt
// }

// evt.Str("remote_ip", c.RealIP())
// evt.Str("host", req.Host)
// evt.Str("method", req.Method)
// evt.Str("uri", req.RequestURI)
// evt.Str("user_agent", req.UserAgent())
// evt.Int("status", res.Status)
// evt.Str("referer", req.Referer())
// evt.Dur("latency", latency)
// evt.Str("latency_human", latency.String())

// cl := req.Header.Get(echo.HeaderContentLength)
// if cl == "" {
// 	cl = "0"
// }

// evt.Str("bytes_in", cl)
// evt.Str("bytes_out", strconv.FormatInt(res.Size, 10))

// if config.NestKey != "" { // Nest the new event (dict) under the nest key.
// 	mainEvt.Dict(config.NestKey, evt)
// }
// mainEvt.Send()
