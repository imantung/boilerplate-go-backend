// Code generated by MockGen. DO NOT EDIT.
// Source: internal/generated/entity/employees.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	sqkit "github.com/imantung/boilerplate-go-backend/pkg/sqkit"
)

// MockEmployeeRepo is a mock of EmployeeRepo interface.
type MockEmployeeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepoMockRecorder
}

// MockEmployeeRepoMockRecorder is the mock recorder for MockEmployeeRepo.
type MockEmployeeRepoMockRecorder struct {
	mock *MockEmployeeRepo
}

// NewMockEmployeeRepo creates a new mock instance.
func NewMockEmployeeRepo(ctrl *gomock.Controller) *MockEmployeeRepo {
	mock := &MockEmployeeRepo{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeRepo) EXPECT() *MockEmployeeRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockEmployeeRepo) Count(arg0 context.Context, arg1 ...sqkit.SelectOption) (int64, error) {
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
func (mr *MockEmployeeRepoMockRecorder) Count(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockEmployeeRepo)(nil).Count), varargs...)
}

// Find mocks base method.
func (m *MockEmployeeRepo) Find(arg0 context.Context, arg1 ...sqkit.SelectOption) ([]*entity.Employee, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*entity.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockEmployeeRepoMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockEmployeeRepo)(nil).Find), varargs...)
}

// Insert mocks base method.
func (m *MockEmployeeRepo) Insert(arg0 context.Context, arg1 *entity.Employee) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockEmployeeRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEmployeeRepo)(nil).Insert), arg0, arg1)
}

// Patch mocks base method.
func (m *MockEmployeeRepo) Patch(arg0 context.Context, arg1 *entity.Employee, arg2 ...sqkit.UpdateOption) (int64, error) {
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
func (mr *MockEmployeeRepoMockRecorder) Patch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockEmployeeRepo)(nil).Patch), varargs...)
}

// SoftDelete mocks base method.
func (m *MockEmployeeRepo) SoftDelete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SoftDelete indicates an expected call of SoftDelete.
func (mr *MockEmployeeRepoMockRecorder) SoftDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDelete", reflect.TypeOf((*MockEmployeeRepo)(nil).SoftDelete), arg0, arg1)
}

// Update mocks base method.
func (m *MockEmployeeRepo) Update(arg0 context.Context, arg1 *entity.Employee, arg2 ...sqkit.UpdateOption) (int64, error) {
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
func (mr *MockEmployeeRepoMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeRepo)(nil).Update), varargs...)
}
