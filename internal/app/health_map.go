package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	HealthMap map[string]error
)

func (h HealthMap) httpCode() int {
	for _, v := range h {
		if v != nil {
			return http.StatusServiceUnavailable
		}
	}
	return http.StatusOK
}

func (h HealthMap) message() map[string]string {
	msg := map[string]string{}
	for k, v := range h {
		if v != nil {
			msg[k] = v.Error()
		} else {
			msg[k] = "ok"
		}
	}
	return msg
}

func (h HealthMap) Handle(ec echo.Context) error {

	// NOTE: disable cache
	ec.Response().Header().Set("Expires", "0")
	ec.Response().Header().Set("Pragma", "no-cache")
	ec.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	code := h.httpCode()
	msg := h.message()

	return ec.JSON(code, msg)
}
