// Code generated by `tools/entity-gen`. DO NOT EDIT.
package entity

import "time"

var EmployeesTableName = "employees"

var EmployeesColumns = struct {
	ID             string
	EmployeeName   string
	JobTitle       string
	LastCheckInAt  string
	LastCheckOutAt string
	UpdateAt       string
	CreatedAt      string
}{
	ID:             "id",
	EmployeeName:   "employee_name",
	JobTitle:       "job_title",
	LastCheckInAt:  "last_check_in_at",
	LastCheckOutAt: "last_check_out_at",
	UpdateAt:       "update_at",
	CreatedAt:      "created_at",
}

type Employees struct {
	ID             int        `column:"id"`
	EmployeeName   string     `column:"employee_name"`
	JobTitle       string     `column:"job_title"`
	LastCheckInAt  *time.Time `column:"last_check_in_at"`
	LastCheckOutAt *time.Time `column:"last_check_out_at"`
	UpdateAt       time.Time  `column:"update_at"`
	CreatedAt      time.Time  `column:"created_at"`
}

var EmployeeClockHistoriesTableName = "employee_clock_histories"

var EmployeeClockHistoriesColumns = struct {
	ID                  string
	EmployeeID          string
	CheckInAt           string
	CheckOutAt          string
	WorkDuration        string
	WorkDurationMinutes string
	UpdateAt            string
	CreatedAt           string
}{
	ID:                  "id",
	EmployeeID:          "employee_id",
	CheckInAt:           "check_in_at",
	CheckOutAt:          "check_out_at",
	WorkDuration:        "work_duration",
	WorkDurationMinutes: "work_duration_minutes",
	UpdateAt:            "update_at",
	CreatedAt:           "created_at",
}

type EmployeeClockHistories struct {
	ID                  int        `column:"id"`
	EmployeeID          string     `column:"employee_id"`
	CheckInAt           *time.Time `column:"check_in_at"`
	CheckOutAt          *time.Time `column:"check_out_at"`
	WorkDuration        *string    `column:"work_duration"`
	WorkDurationMinutes *int       `column:"work_duration_minutes"`
	UpdateAt            time.Time  `column:"update_at"`
	CreatedAt           time.Time  `column:"created_at"`
}
