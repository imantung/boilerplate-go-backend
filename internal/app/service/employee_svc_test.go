package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/imantung/boilerplate-go-backend/internal/app/service"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/mock_entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/stretchr/testify/assert"
)

func TestListEmployee(t *testing.T) {
	testcases := []struct {
		TestName         string
		Request          oapi.ListEmployeeRequestObject
		OnMockRepo       func(*mock_entity.MockEmployeeRepo)
		ExpectedResponse oapi.ListEmployeeResponseObject
		ExpectedErr      string
	}{
		{
			TestName: "repo error",
			Request:  oapi.ListEmployeeRequestObject{},
			OnMockRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().Select(gomock.Any()).Return(nil, errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "success",
			Request:  oapi.ListEmployeeRequestObject{},
			OnMockRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().Select(gomock.Any()).
					Return([]*entity.Employee{
						{
							ID:             99,
							EmployeeName:   "some-name",
							JobTitle:       "some-title",
							LastClockInAt:  &time.Time{},
							LastClockOutAt: &time.Time{},
						},
					}, nil)
			},
			ExpectedResponse: oapi.ListEmployee200JSONResponse{
				{
					Id:             99,
					EmployeeName:   "some-name",
					JobTitle:       "some-title",
					LastClockInAt:  &time.Time{},
					LastClockOutAt: &time.Time{},
				},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_entity.NewMockEmployeeRepo(ctrl)
			if tt.OnMockRepo != nil {
				tt.OnMockRepo(repo)
			}

			svc := service.NewEmployeeSvc(repo)
			resp, err := svc.ListEmployee(context.Background(), tt.Request)
			if err != nil {
				assert.EqualError(t, err, tt.ExpectedErr)
			} else {
				assert.EqualValues(t, tt.ExpectedResponse, resp)
			}
		})
	}
}

func TestCreateEmployee(t *testing.T) {
	testcases := []struct {
		TestName         string
		Request          oapi.CreateEmployeeRequestObject
		OnMockRepo       func(*mock_entity.MockEmployeeRepo)
		ExpectedResponse oapi.CreateEmployeeResponseObject
		ExpectedErr      string
	}{
		{
			TestName:    "empty employee name",
			Request:     oapi.CreateEmployeeRequestObject{Body: &oapi.CreateEmployeeJSONRequestBody{}},
			ExpectedErr: "code=422, message=Employee Name can't be empty",
		},
		{
			TestName: "empty job title",
			Request: oapi.CreateEmployeeRequestObject{
				Body: &oapi.CreateEmployeeJSONRequestBody{
					EmployeeName: "some-name",
				},
			},
			ExpectedErr: "code=422, message=Job Title can't be empty",
		},
		{
			TestName: "repo error",
			Request: oapi.CreateEmployeeRequestObject{
				Body: &oapi.CreateEmployeeJSONRequestBody{
					EmployeeName: "some-name",
					JobTitle:     "some-title",
				},
			},
			OnMockRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().Insert(gomock.Any(), &entity.Employee{
					EmployeeName: "some-name",
					JobTitle:     "some-title",
				}).Return(-1, errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "success",
			Request: oapi.CreateEmployeeRequestObject{
				Body: &oapi.CreateEmployeeJSONRequestBody{
					EmployeeName: "some-name",
					JobTitle:     "some-title",
				},
			},
			OnMockRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().Insert(gomock.Any(), &entity.Employee{
					EmployeeName: "some-name",
					JobTitle:     "some-title",
				}).Return(99, nil)
			},
			ExpectedResponse: oapi.CreateEmployee201Response{
				Headers: oapi.CreateEmployee201ResponseHeaders{
					Location: "/employees/99",
				},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_entity.NewMockEmployeeRepo(ctrl)
			if tt.OnMockRepo != nil {
				tt.OnMockRepo(repo)
			}

			svc := service.NewEmployeeSvc(repo)
			resp, err := svc.CreateEmployee(context.Background(), tt.Request)
			if err != nil {
				assert.EqualError(t, err, tt.ExpectedErr)
			} else {
				assert.EqualValues(t, tt.ExpectedResponse, resp)
			}
		})
	}
}
