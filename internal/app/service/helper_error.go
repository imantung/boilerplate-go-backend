package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func validationError(errMsg string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, errMsg)
}

func notFoundError(id int) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("ID #%d not found", id))
}
