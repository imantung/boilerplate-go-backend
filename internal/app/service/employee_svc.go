package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
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

func (e *EmployeeSvcImpl) ListEmployee(ctx context.Context, request oapi.ListEmployeeRequestObject) (oapi.ListEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (e *EmployeeSvcImpl) CreateEmployee(ctx context.Context, req oapi.CreateEmployeeRequestObject) (oapi.CreateEmployeeResponseObject, error) {
	if validationMsg := e.validateCreateEmployee(req); validationMsg != "" {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, validationMsg)
	}

	id, err := e.EmployeeRepo.Insert(ctx, &entity.Employee{
		EmployeeName: req.Body.EmployeeName,
		JobTitle:     req.Body.JobTitle,
	})
	if err != nil {
		return nil, err
	}
	return oapi.CreateEmployee201Response{
		Headers: oapi.CreateEmployee201ResponseHeaders{
			Location: fmt.Sprintf("/employees/%d", id),
		},
	}, nil
}

func (e *EmployeeSvcImpl) validateCreateEmployee(req oapi.CreateEmployeeRequestObject) string {
	if len(req.Body.EmployeeName) <= 0 {
		return "Employee Name can't be empty"
	}
	if len(req.Body.JobTitle) <= 0 {
		return "Job Title can't be empty"
	}
	return ""
}

func (e *EmployeeSvcImpl) DeleteEmployee(ctx context.Context, request oapi.DeleteEmployeeRequestObject) (oapi.DeleteEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (e *EmployeeSvcImpl) GetEmployee(ctx context.Context, request oapi.GetEmployeeRequestObject) (oapi.GetEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (e *EmployeeSvcImpl) PatchEmployee(ctx context.Context, request oapi.PatchEmployeeRequestObject) (oapi.PatchEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (e *EmployeeSvcImpl) UpdateEmployee(ctx context.Context, request oapi.UpdateEmployeeRequestObject) (oapi.UpdateEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}
