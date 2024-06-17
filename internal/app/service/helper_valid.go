package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func validationError(errMsg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, errMsg)
}

func isEmpty(s string) bool {
	return len(s) <= 0
}
