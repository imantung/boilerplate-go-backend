package service

import (
	"context"
	"fmt"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/imantung/boilerplate-go-backend/pkg/sqkit"
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
		resp = append(resp, convertToEmployeeOApi(emp))
	}
	return resp, nil
}

func (e *EmployeeSvcImpl) CreateEmployee(ctx context.Context, req oapi.CreateEmployeeRequestObject) (oapi.CreateEmployeeResponseObject, error) {
	employee := &entity.Employee{
		EmployeeName: req.Body.EmployeeName,
		JobTitle:     req.Body.JobTitle,
	}
	if errMsg := validateEmployee(employee); errMsg != "" {
		return nil, validationError(errMsg)
	}

	id, err := e.EmployeeRepo.Insert(ctx, employee)
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

	resp := oapi.GetEmployee200JSONResponse(convertToEmployeeOApi(employees[0]))
	return resp, nil
}

func (e *EmployeeSvcImpl) PatchEmployee(ctx context.Context, req oapi.PatchEmployeeRequestObject) (oapi.PatchEmployeeResponseObject, error) {
	id := int(req.Id)
	employee := convertToEmployeeEntity(req.Body)

	affectedRow, err := e.EmployeeRepo.Patch(ctx, &employee, sqkit.Eq{"id": id})
	if err != nil {
		return nil, err
	}
	if affectedRow < 1 {
		return nil, notFoundError(id)
	}
	return oapi.PatchEmployee204Response{}, nil
}

func (e *EmployeeSvcImpl) UpdateEmployee(ctx context.Context, req oapi.UpdateEmployeeRequestObject) (oapi.UpdateEmployeeResponseObject, error) {
	id := int(req.Id)
	employee := convertToEmployeeEntity(req.Body)
	if errMsg := validateEmployee(&employee); errMsg != "" {
		return nil, validationError(errMsg)
	}
	affectedRow, err := e.EmployeeRepo.Update(ctx, &employee, sqkit.Eq{"id": id})
	if err != nil {
		return nil, err
	}
	if affectedRow < 1 {
		return nil, notFoundError(id)
	}
	return oapi.UpdateEmployee204Response{}, nil
}
