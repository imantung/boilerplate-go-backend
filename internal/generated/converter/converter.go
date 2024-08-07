// Code generated by `tools/converter-gen`. DO NOT EDIT.
package converter

import (
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
)

func ConvertToEmployeeOApi(from entity.Employee) oapi.Employee {
	return oapi.Employee{
		EmployeeName:   from.EmployeeName,
		Id:             from.ID,
		JobTitle:       from.JobTitle,
		LastClockInAt:  from.LastClockInAt,
		LastClockOutAt: from.LastClockOutAt,
	}
}
func ConvertToEmployeeEntity(from oapi.Employee) entity.Employee {
	return entity.Employee{
		EmployeeName:   from.EmployeeName,
		ID:             from.Id,
		JobTitle:       from.JobTitle,
		LastClockInAt:  from.LastClockInAt,
		LastClockOutAt: from.LastClockOutAt,
	}
}
func ConvertToClockHistoryOApi(from entity.EmployeeClockHistory) oapi.EmployeeClockHistory {
	return oapi.EmployeeClockHistory{
		ClockInAt:           from.ClockInAt,
		ClockOutAt:          from.ClockOutAt,
		EmployeeId:          from.EmployeeID,
		Id:                  from.ID,
		WorkDuration:        from.WorkDuration,
		WorkDurationMinutes: from.WorkDurationMinutes,
	}
}
