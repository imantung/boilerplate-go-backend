package service

import (
	"context"
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
)

type (
	ClockSvc interface {
		ClockIn(ctx context.Context, request oapi.ClockInRequestObject) (oapi.ClockInResponseObject, error)
		ClockOut(ctx context.Context, request oapi.ClockOutRequestObject) (oapi.ClockOutResponseObject, error)
		ListClock(ctx context.Context, request oapi.ListClockRequestObject) (oapi.ListClockResponseObject, error)
	}
	ClockSvcImpl struct{}
)

var _ = di.Provide(NewClockSvc)

func NewClockSvc() ClockSvc {
	return &ClockSvcImpl{}
}

func (*ClockSvcImpl) ClockIn(ctx context.Context, request oapi.ClockInRequestObject) (oapi.ClockInResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (*ClockSvcImpl) ClockOut(ctx context.Context, request oapi.ClockOutRequestObject) (oapi.ClockOutResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (*ClockSvcImpl) ListClock(ctx context.Context, request oapi.ListClockRequestObject) (oapi.ListClockResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}
