package entity

var EmployeesTableName = "employees"

var EmployeesColumns = struct {
	ID           string
	UpdateAt     string
	CreatedAt    string
	EmployeeName string
	JobTitle     string
}{
	ID:           "id",
	UpdateAt:     "update_at",
	CreatedAt:    "created_at",
	EmployeeName: "employee_name",
	JobTitle:     "job_title",
}

type Employees struct {
	ID           int `column:"id"`
	UpdateAt     int `column:"update_at"`
	CreatedAt    int `column:"created_at"`
	EmployeeName int `column:"employee_name"`
	JobTitle     int `column:"job_title"`
}

var EmployeeClockHistoriesTableName = "employee_clock_histories"

var EmployeeClockHistoriesColumns = struct {
	ID         string
	UpdateAt   string
	CreatedAt  string
	EmployeeID string
	JobTitle   string
}{
	ID:         "id",
	UpdateAt:   "update_at",
	CreatedAt:  "created_at",
	EmployeeID: "employee_id",
	JobTitle:   "job_title",
}

type EmployeeClockHistories struct {
	ID         int `column:"id"`
	UpdateAt   int `column:"update_at"`
	CreatedAt  int `column:"created_at"`
	EmployeeID int `column:"employee_id"`
	JobTitle   int `column:"job_title"`
}
