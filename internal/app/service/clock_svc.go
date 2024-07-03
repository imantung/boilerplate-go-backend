package service

import (
	"context"
	"fmt"
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/converter"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/imantung/dbtxn"
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
	clockInAt := Now().UTC()
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

func (c *ClockSvcImpl) ClockOut(ctx context.Context, req oapi.ClockOutRequestObject) (resp oapi.ClockOutResponseObject, err error) {
	if errMsg := validateClockRequest(req.Body); errMsg != "" {
		return nil, validationError(errMsg)
	}

	txn := dbtxn.Begin(&ctx)
	defer txn.CommitWithError(&err)

	employeeID := req.Body.EmployeeId

	employees, err := c.EmployeeRepo.Select(ctx, repokit.Eq{"id": employeeID})
	if err != nil {
		return nil, err
	}
	if len(employees) < 1 {
		return nil, notFoundError(employeeID)
	}

	histories, err := c.HistoryRepo.Select(ctx, repokit.Eq{"employee_id": employeeID}, repokit.Sorts{"-id"})
	if err != nil {
		return nil, err
	}
	if len(histories) < 1 {
		return nil, fmt.Errorf("no clock history for employee#%d", employeeID)
	}

	history := histories[0]
	if history.ClockOutAt != nil {
		return nil, fmt.Errorf("employee#%d already clock-out", employeeID)
	}

	clockOutAt := Now().UTC()
	history.ClockOutAt = &clockOutAt

	workDuration := clockOutAt.Sub(*history.ClockInAt)
	history.WorkDuration = workDuration.String()
	history.WorkDurationMinutes = int(workDuration.Minutes())

	_, err = c.HistoryRepo.Update(ctx, history, repokit.Eq{"id": history.ID})
	if err != nil {
		return nil, err
	}

	_, err = c.EmployeeRepo.Patch(ctx, &entity.Employee{LastClockOutAt: &clockOutAt}, repokit.Eq{"id": employeeID})
	if err != nil {
		return nil, err
	}

	resp = oapi.ClockOut200JSONResponse(converter.ConvertToClockHistoryOApi(*history))
	return
}

func (c *ClockSvcImpl) ListClock(ctx context.Context, req oapi.ListClockRequestObject) (oapi.ListClockResponseObject, error) {
	histories, err := c.HistoryRepo.Select(ctx)
	if err != nil {
		return nil, err
	}
	resp := oapi.ListClock200JSONResponse{}
	for _, history := range histories {
		resp = append(resp, converter.ConvertToClockHistoryOApi(*history))
	}
	return resp, nil
}
