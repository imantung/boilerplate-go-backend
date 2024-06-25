package service

import (
	"strings"

	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}

func validateEmployee(emp *entity.Employee) string {
	if isEmpty(emp.EmployeeName) {
		return "Employee Name can't be empty"
	}
	if isEmpty(emp.JobTitle) {
		return "Job Title can't be empty"
	}
	return ""
}

func validateClockRequest(req *oapi.ClockRequest) string {
	if req.EmployeeId <= 0 {
		return "employee ID can't be zero"
	}
	return ""
}
