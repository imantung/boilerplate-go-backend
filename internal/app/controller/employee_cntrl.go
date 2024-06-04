package controller

import (
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	EmployeeCntrl interface {
		GetEmployees(ec echo.Context) error
		PostEmployees(ec echo.Context) error
		DeleteEmployeesId(ec echo.Context, id int64) error
		GetEmployeesId(ec echo.Context, id int64) error
		PatchEmployeesId(ec echo.Context, id int64) error
		PutEmployeesId(ec echo.Context, id int64) error
		PostEmployeesIdClockIn(ec echo.Context, id int64) error
		PostEmployeesIdClockOut(ec echo.Context, id int64) error
		GetEmployeesIdClocks(ec echo.Context, id int64) error
	}
	EmployeeCntrlImpl struct {
	}
)

var _ = di.Provide(NewEmployeeCntrl)

func NewEmployeeCntrl() EmployeeCntrl {
	return &EmployeeCntrlImpl{}
}

func (s *EmployeeCntrlImpl) GetEmployees(ec echo.Context) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) PostEmployees(ec echo.Context) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) DeleteEmployeesId(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) GetEmployeesId(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) PatchEmployeesId(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) PutEmployeesId(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) PostEmployeesIdClockIn(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) PostEmployeesIdClockOut(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) GetEmployeesIdClocks(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
