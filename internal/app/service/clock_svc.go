package service

import (
	"context"
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
)

type (
	ClockSvc interface {
		ClockIn(ctx context.Context, request oapi.ClockInRequestObject) (oapi.ClockInResponseObject, error)
		ClockOut(ctx context.Context, request oapi.ClockOutRequestObject) (oapi.ClockOutResponseObject, error)
		ListClock(ctx context.Context, request oapi.ListClockRequestObject) (oapi.ListClockResponseObject, error)
	}
	ClockSvcImpl struct {
		EmployeeRepo entity.EmployeeRepo
		HistoryRepo  entity.EmployeeClockHistoryRepo
	}
)

var _ = di.Provide(NewClockSvc)

func NewClockSvc(employeeRepo entity.EmployeeRepo, historyRepo entity.EmployeeClockHistoryRepo) ClockSvc {
	return &ClockSvcImpl{
		EmployeeRepo: employeeRepo,
		HistoryRepo:  historyRepo,
	}
}

func (c *ClockSvcImpl) ClockIn(ctx context.Context, req oapi.ClockInRequestObject) (oapi.ClockInResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (c *ClockSvcImpl) ClockOut(ctx context.Context, req oapi.ClockOutRequestObject) (oapi.ClockOutResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (c *ClockSvcImpl) ListClock(ctx context.Context, req oapi.ListClockRequestObject) (oapi.ListClockResponseObject, error) {
	histories, err := c.HistoryRepo.Select(ctx)
	if err != nil {
		return nil, err
	}
	resp := oapi.ListClock200JSONResponse{}
	for _, history := range histories {
		resp = append(resp, convertToClockHistoryOApi(history))
	}
	return resp, nil
}
