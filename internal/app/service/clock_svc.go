package service

import (
	"context"
	"net/http"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/imantung/dbtxn"
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
var Now = time.Now

func NewClockSvc(employeeRepo entity.EmployeeRepo, historyRepo entity.EmployeeClockHistoryRepo) ClockSvc {
	return &ClockSvcImpl{
		EmployeeRepo: employeeRepo,
		HistoryRepo:  historyRepo,
	}
}

func (c *ClockSvcImpl) ClockIn(ctx context.Context, req oapi.ClockInRequestObject) (resp oapi.ClockInResponseObject, err error) {
	if errMsg := validateClockRequest(req.Body); errMsg != "" {
		return nil, validationError(errMsg)
	}

	txn := dbtxn.Begin(&ctx)
	defer txn.CommitWithError(&err)

	employeeID := req.Body.EmployeeId
	clockInAt := Now()
	affectedRow, err := c.EmployeeRepo.Patch(ctx, &entity.Employee{LastClockInAt: &clockInAt}, repokit.Eq{"id": employeeID})
	if err != nil {
		return nil, err
	}
	if affectedRow < 1 {
		return nil, notFoundError(employeeID)
	}
	insertedID, err := c.HistoryRepo.Insert(ctx, &entity.EmployeeClockHistory{EmployeeID: employeeID, ClockInAt: &clockInAt})
	if err != nil {
		return nil, err
	}

	resp = oapi.ClockIn200JSONResponse{
		Id:         insertedID,
		EmployeeId: employeeID,
		ClockInAt:  &clockInAt,
	}
	return
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
