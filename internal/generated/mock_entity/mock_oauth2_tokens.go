// Code generated by MockGen. DO NOT EDIT.
// Source: internal/generated/entity/oauth2_tokens.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	repokit "github.com/imantung/boilerplate-go-backend/pkg/repokit"
)

// MockOauth2TokenRepo is a mock of Oauth2TokenRepo interface.
type MockOauth2TokenRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOauth2TokenRepoMockRecorder
}

// MockOauth2TokenRepoMockRecorder is the mock recorder for MockOauth2TokenRepo.
type MockOauth2TokenRepoMockRecorder struct {
	mock *MockOauth2TokenRepo
}

// NewMockOauth2TokenRepo creates a new mock instance.
func NewMockOauth2TokenRepo(ctrl *gomock.Controller) *MockOauth2TokenRepo {
	mock := &MockOauth2TokenRepo{ctrl: ctrl}
	mock.recorder = &MockOauth2TokenRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauth2TokenRepo) EXPECT() *MockOauth2TokenRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockOauth2TokenRepo) Count(arg0 context.Context, arg1 ...repokit.SelectOption) (int64, error) {
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
func (mr *MockOauth2TokenRepoMockRecorder) Count(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOauth2TokenRepo)(nil).Count), varargs...)
}

// Insert mocks base method.
func (m *MockOauth2TokenRepo) Insert(arg0 context.Context, arg1 *entity.Oauth2Token) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockOauth2TokenRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOauth2TokenRepo)(nil).Insert), arg0, arg1)
}

// Patch mocks base method.
func (m *MockOauth2TokenRepo) Patch(arg0 context.Context, arg1 *entity.Oauth2Token, arg2 ...repokit.UpdateOption) (int64, error) {
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
func (mr *MockOauth2TokenRepoMockRecorder) Patch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockOauth2TokenRepo)(nil).Patch), varargs...)
}

// Select mocks base method.
func (m *MockOauth2TokenRepo) Select(arg0 context.Context, arg1 ...repokit.SelectOption) ([]*entity.Oauth2Token, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].([]*entity.Oauth2Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockOauth2TokenRepoMockRecorder) Select(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockOauth2TokenRepo)(nil).Select), varargs...)
}

// SoftDelete mocks base method.
func (m *MockOauth2TokenRepo) SoftDelete(arg0 context.Context, arg1 int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDelete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SoftDelete indicates an expected call of SoftDelete.
func (mr *MockOauth2TokenRepoMockRecorder) SoftDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDelete", reflect.TypeOf((*MockOauth2TokenRepo)(nil).SoftDelete), arg0, arg1)
}

// Update mocks base method.
func (m *MockOauth2TokenRepo) Update(arg0 context.Context, arg1 *entity.Oauth2Token, arg2 ...repokit.UpdateOption) (int64, error) {
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
func (mr *MockOauth2TokenRepoMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOauth2TokenRepo)(nil).Update), varargs...)
}
