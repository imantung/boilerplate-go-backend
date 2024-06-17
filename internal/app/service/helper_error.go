package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func validationError(errMsg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, errMsg)
}

func notFoundError(id int64) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%d not found", id))
}
