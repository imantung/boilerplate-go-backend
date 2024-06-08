package service

import (
	"context"
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
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
	EmployeeSvcImpl struct{}
)

var _ = di.Provide(NewEmployeeSvc)

func NewEmployeeSvc() EmployeeSvc {
	return &EmployeeSvcImpl{}
}

func (e *EmployeeSvcImpl) ListEmployee(ctx context.Context, request oapi.ListEmployeeRequestObject) (oapi.ListEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
}

func (e *EmployeeSvcImpl) CreateEmployee(ctx context.Context, request oapi.CreateEmployeeRequestObject) (oapi.CreateEmployeeResponseObject, error) {
	return nil, &echo.HTTPError{Code: http.StatusNotImplemented, Message: "not implemented"}
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
