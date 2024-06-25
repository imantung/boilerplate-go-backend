package service

import (
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
)

func convertToEmployeeOApi(emp *entity.Employee) oapi.Employee {
	return oapi.Employee{
		Id:             emp.ID,
		EmployeeName:   emp.EmployeeName,
		JobTitle:       emp.JobTitle,
		LastClockInAt:  emp.LastClockInAt,
		LastClockOutAt: emp.LastClockOutAt,
	}
}

func convertToEmployeeEntity(emp *oapi.Employee) entity.Employee {
	return entity.Employee{
		ID:             emp.Id,
		EmployeeName:   emp.EmployeeName,
		JobTitle:       emp.JobTitle,
		LastClockInAt:  emp.LastClockInAt,
		LastClockOutAt: emp.LastClockOutAt,
	}
}

func convertToClockHistoryOApi(history *entity.EmployeeClockHistory) oapi.EmployeeClockHistory {
	return oapi.EmployeeClockHistory{
		ClockInAt:           history.ClockInAt,
		ClockOutAt:          history.ClockOutAt,
		EmployeeId:          history.EmployeeID,
		Id:                  history.ID,
		WorkDuration:        history.WorkDuration,
		WorkDurationMinutes: history.WorkDurationMinutes,
	}
}
