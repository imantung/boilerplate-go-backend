package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/imantung/boilerplate-go-backend/pkg/sqkit"
	"github.com/labstack/echo/v4"
)

type (
	EmployeeSvc interface {
		ListEmployee(ctx context.Context, request oapi.ListEmployeeRequestObject) (oapi.ListEmployeeResponseObject, error)
		CreateEmployee(ctx context.Context, request oapi.CreateEmployeeRequestObject) (oapi.CreateEmployeeResponseObject, error)
		DeleteEmployee(ctx context.Context, request oapi.DeleteEmployeeRequestObject) (oapi.DeleteEmployeeResponseObject, error)
		GetEmployee(ctx context.Context, request oapi.GetEmployeeRequestObject) (oapi.GetEmployeeResponseObject, error)
		PatchEmployee(ctx context.Context, request oapi.PatchEmployeeRequestObject) (oapi.PatchEmployeeResponseObject, error)
		UpdateEmployee(ctx context.Context, request oapi.UpdateEmployeeRequestObject) (oapi.UpdateEmployeeResponseObject, error)
	}
	EmployeeSvcImpl struct {
		EmployeeRepo entity.EmployeeRepo
	}
)

var _ = di.Provide(NewEmployeeSvc)

func NewEmployeeSvc(employeeRepo entity.EmployeeRepo) EmployeeSvc {
	return &EmployeeSvcImpl{
		EmployeeRepo: employeeRepo,
	}
}

func (e *EmployeeSvcImpl) ListEmployee(ctx context.Context, req oapi.ListEmployeeRequestObject) (oapi.ListEmployeeResponseObject, error) {
	employees, err := e.EmployeeRepo.Select(ctx)
	if err != nil {
		return nil, err
	}

	resp := oapi.ListEmployee200JSONResponse{}
	for _, emp := range employees {
		resp = append(resp, convertEmployee(emp))
	}
	return resp, nil
}

func (e *EmployeeSvcImpl) CreateEmployee(ctx context.Context, req oapi.CreateEmployeeRequestObject) (oapi.CreateEmployeeResponseObject, error) {
	if errMsg := e.validateCreateEmployee(req); errMsg != "" {
		return nil, validationError(errMsg)
	}

	id, err := e.EmployeeRepo.Insert(ctx, &entity.Employee{
		EmployeeName: req.Body.EmployeeName,
		JobTitle:     req.Body.JobTitle,
	})
	if err != nil {
		return nil, err
	}

	resp := oapi.CreateEmployee201Response{
		Headers: oapi.CreateEmployee201ResponseHeaders{
			Location: fmt.Sprintf("/employees/%d", id),
		},
	}
	return resp, nil
}

func (e *EmployeeSvcImpl) validateCreateEmployee(req oapi.CreateEmployeeRequestObject) string {
	if isEmpty(req.Body.EmployeeName) {
		return "Employee Name can't be empty"
	}
	if isEmpty(req.Body.JobTitle) {
		return "Job Title can't be empty"
	}
	return ""
}

func (e *EmployeeSvcImpl) DeleteEmployee(ctx context.Context, req oapi.DeleteEmployeeRequestObject) (oapi.DeleteEmployeeResponseObject, error) {
	id := int(req.Id)
	affectedRow, err := e.EmployeeRepo.SoftDelete(ctx, id)
	if err != nil {
		return nil, err
	}

	if affectedRow < 1 {
		return nil, notFoundError(id)
	}

	return oapi.DeleteEmployee204Response{}, nil
}

func (e *EmployeeSvcImpl) GetEmployee(ctx context.Context, req oapi.GetEmployeeRequestObject) (oapi.GetEmployeeResponseObject, error) {
	id := int(req.Id)
	employees, err := e.EmployeeRepo.Select(ctx, sqkit.Eq{"id": id})
	if err != nil {
		return nil, err
	}

	if len(employees) < 1 {
		return nil, notFoundError(id)
	}

	resp := oapi.GetEmployee200JSONResponse(convertEmployee(employees[0]))
	return resp, nil
}

func (e *EmployeeSvcImpl) PatchEmployee(ctx context.Context, request oapi.PatchEmployeeRequestObject) (oapi.PatchEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (e *EmployeeSvcImpl) UpdateEmployee(ctx context.Context, request oapi.UpdateEmployeeRequestObject) (oapi.UpdateEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}
