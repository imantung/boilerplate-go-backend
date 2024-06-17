package service

import (
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
)

func convertEmployee(emp *entity.Employee) oapi.Employee {
	return oapi.Employee{
		Id:             emp.ID,
		EmployeeName:   emp.EmployeeName,
		JobTitle:       emp.JobTitle,
		LastClockInAt:  emp.LastClockInAt,
		LastClockOutAt: emp.LastClockOutAt,
	}
}
