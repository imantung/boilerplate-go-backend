// Code generated by MockGen. DO NOT EDIT.
// Source: internal/generated/entity/employee_clock_histories.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	repokit "github.com/imantung/boilerplate-go-backend/pkg/repokit"
)

// MockEmployeeClockHistoryRepo is a mock of EmployeeClockHistoryRepo interface.
type MockEmployeeClockHistoryRepo struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeClockHistoryRepoMockRecorder
}

// MockEmployeeClockHistoryRepoMockRecorder is the mock recorder for MockEmployeeClockHistoryRepo.
type MockEmployeeClockHistoryRepoMockRecorder struct {
	mock *MockEmployeeClockHistoryRepo
}

// NewMockEmployeeClockHistoryRepo creates a new mock instance.
func NewMockEmployeeClockHistoryRepo(ctrl *gomock.Controller) *MockEmployeeClockHistoryRepo {
	mock := &MockEmployeeClockHistoryRepo{ctrl: ctrl}
	mock.recorder = &MockEmployeeClockHistoryRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeClockHistoryRepo) EXPECT() *MockEmployeeClockHistoryRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockEmployeeClockHistoryRepo) Count(arg0 context.Context, arg1 ...repokit.SelectOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockEmployeeClockHistoryRepoMockRecorder) Count(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockEmployeeClockHistoryRepo)(nil).Count), varargs...)
}

// Insert mocks base method.
func (m *MockEmployeeClockHistoryRepo) Insert(arg0 context.Context, arg1 *entity.EmployeeClockHistory) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockEmployeeClockHistoryRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEmployeeClockHistoryRepo)(nil).Insert), arg0, arg1)
}

// Patch mocks base method.
func (m *MockEmployeeClockHistoryRepo) Patch(arg0 context.Context, arg1 *entity.EmployeeClockHistory, arg2 ...repokit.UpdateOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockEmployeeClockHistoryRepoMockRecorder) Patch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockEmployeeClockHistoryRepo)(nil).Patch), varargs...)
}

// Select mocks base method.
func (m *MockEmployeeClockHistoryRepo) Select(arg0 context.Context, arg1 ...repokit.SelectOption) ([]*entity.EmployeeClockHistory, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].([]*entity.EmployeeClockHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockEmployeeClockHistoryRepoMockRecorder) Select(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockEmployeeClockHistoryRepo)(nil).Select), varargs...)
}

// SoftDelete mocks base method.
func (m *MockEmployeeClockHistoryRepo) SoftDelete(arg0 context.Context, arg1 int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDelete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SoftDelete indicates an expected call of SoftDelete.
func (mr *MockEmployeeClockHistoryRepoMockRecorder) SoftDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDelete", reflect.TypeOf((*MockEmployeeClockHistoryRepo)(nil).SoftDelete), arg0, arg1)
}

// Update mocks base method.
func (m *MockEmployeeClockHistoryRepo) Update(arg0 context.Context, arg1 *entity.EmployeeClockHistory, arg2 ...repokit.UpdateOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeClockHistoryRepoMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeClockHistoryRepo)(nil).Update), varargs...)
}
