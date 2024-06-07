package controller

import (
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	EmployeeCntrl interface {
		ListEmployee(ctx echo.Context) error
		CreateEmployee(ctx echo.Context) error
		DeleteEmployee(ctx echo.Context, id int64) error
		GetEmployee(ctx echo.Context, id int64) error
		PatchEmployee(ctx echo.Context, id int64) error
		UpdateEmployee(ctx echo.Context, id int64) error
		ClockIn(ctx echo.Context, id int64) error
		ClockOut(ctx echo.Context, id int64) error
		ListEmployeeClockHistory(ctx echo.Context, id int64) error
	}
	EmployeeCntrlImpl struct {
	}
)

var _ = di.Provide(NewEmployeeCntrl)

func NewEmployeeCntrl() EmployeeCntrl {
	return &EmployeeCntrlImpl{}
}

func (s *EmployeeCntrlImpl) ListEmployee(ec echo.Context) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) CreateEmployee(ec echo.Context) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) DeleteEmployee(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) GetEmployee(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) PatchEmployee(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) UpdateEmployee(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) ClockIn(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) ClockOut(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
func (s *EmployeeCntrlImpl) ListEmployeeClockHistory(ec echo.Context, id int64) error {
	return ec.JSON(http.StatusNotImplemented, "not implemented")
}
