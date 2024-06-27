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
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/stretchr/testify/assert"
)

func TestListClock(t *testing.T) {
	testcases := []struct {
		TestName         string
		Request          oapi.ListClockRequestObject
		OnEmployeeRepo   func(*mock_entity.MockEmployeeRepo)
		OnHistoryRepo    func(*mock_entity.MockEmployeeClockHistoryRepo)
		ExpectedResponse oapi.ListClockResponseObject
		ExpectedErr      string
	}{
		{
			TestName: "repo error",
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().Select(gomock.Any()).Return(nil, errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "success",
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().Select(gomock.Any()).Return([]*entity.EmployeeClockHistory{
					{
						ID:                  11,
						EmployeeID:          22,
						WorkDuration:        "1h",
						WorkDurationMinutes: 33,
					},
				}, nil)
			},
			ExpectedResponse: oapi.ListClock200JSONResponse{
				{
					Id:                  11,
					EmployeeId:          22,
					WorkDuration:        "1h",
					WorkDurationMinutes: 33,
				},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			historyRepo := mock_entity.NewMockEmployeeClockHistoryRepo(ctrl)
			if tt.OnHistoryRepo != nil {
				tt.OnHistoryRepo(historyRepo)
			}

			employeeRepo := mock_entity.NewMockEmployeeRepo(ctrl)
			if tt.OnEmployeeRepo != nil {
				tt.OnEmployeeRepo(employeeRepo)
			}

			svc := service.NewClockSvc(employeeRepo, historyRepo)
			resp, err := svc.ListClock(context.Background(), tt.Request)
			if tt.ExpectedErr != "" {
				assert.EqualError(t, err, tt.ExpectedErr)
			} else {
				assert.EqualValues(t, tt.ExpectedResponse, resp)
			}
		})
	}
}

func TestClockIn(t *testing.T) {
	now := time.Now()
	clockInAtUTC := now.UTC()
	service.Now = func() time.Time {
		return now
	}

	testcases := []struct {
		TestName         string
		Request          oapi.ClockInRequestObject
		OnEmployeeRepo   func(*mock_entity.MockEmployeeRepo)
		OnHistoryRepo    func(*mock_entity.MockEmployeeClockHistoryRepo)
		ExpectedResponse oapi.ClockInResponseObject
		ExpectedErr      string
	}{
		{
			TestName: "validation error",
			Request: oapi.ClockInRequestObject{
				Body: &oapi.ClockInJSONRequestBody{
					EmployeeId: 0,
				},
			},
			ExpectedErr: "code=422, message=employee ID can't be zero",
		},
		{
			TestName: "patch employee error",
			Request: oapi.ClockInRequestObject{
				Body: &oapi.ClockInJSONRequestBody{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Patch(gomock.Any(), &entity.Employee{LastClockInAt: &clockInAtUTC}, repokit.Eq{"id": int(99)}).
					Return(int64(-1), errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "not found",
			Request: oapi.ClockInRequestObject{
				Body: &oapi.ClockInJSONRequestBody{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Patch(gomock.Any(), &entity.Employee{LastClockInAt: &clockInAtUTC}, repokit.Eq{"id": int(99)}).
					Return(int64(0), nil)
			},
			ExpectedErr: "code=404, message=ID #99 not found",
		},
		{
			TestName: "insert clock history error",
			Request: oapi.ClockInRequestObject{
				Body: &oapi.ClockInJSONRequestBody{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Patch(gomock.Any(), &entity.Employee{LastClockInAt: &clockInAtUTC}, repokit.Eq{"id": int(99)}).
					Return(int64(1), nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Insert(gomock.Any(), &entity.EmployeeClockHistory{EmployeeID: 99, ClockInAt: &clockInAtUTC}).
					Return(-1, errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "success",
			Request: oapi.ClockInRequestObject{
				Body: &oapi.ClockInJSONRequestBody{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Patch(gomock.Any(), &entity.Employee{LastClockInAt: &clockInAtUTC}, repokit.Eq{"id": int(99)}).
					Return(int64(1), nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Insert(gomock.Any(), &entity.EmployeeClockHistory{EmployeeID: 99, ClockInAt: &clockInAtUTC}).
					Return(88, nil)
			},
			ExpectedResponse: oapi.ClockIn200JSONResponse{
				Id:         88,
				EmployeeId: 99,
				ClockInAt:  &clockInAtUTC,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			historyRepo := mock_entity.NewMockEmployeeClockHistoryRepo(ctrl)
			if tt.OnHistoryRepo != nil {
				tt.OnHistoryRepo(historyRepo)
			}

			employeeRepo := mock_entity.NewMockEmployeeRepo(ctrl)
			if tt.OnEmployeeRepo != nil {
				tt.OnEmployeeRepo(employeeRepo)
			}

			svc := service.NewClockSvc(employeeRepo, historyRepo)
			resp, err := svc.ClockIn(context.Background(), tt.Request)
			if tt.ExpectedErr != "" {
				assert.EqualError(t, err, tt.ExpectedErr)
			} else {
				assert.EqualValues(t, tt.ExpectedResponse, resp)
			}
		})
	}
}

func TestClockOut(t *testing.T) {
	now := time.Now()
	clockInAtUTC := now.Add(-10 * time.Hour).UTC()
	clockOutAtUTC := now.UTC()
	service.Now = func() time.Time {
		return now
	}

	testcases := []struct {
		TestName         string
		Request          oapi.ClockOutRequestObject
		OnEmployeeRepo   func(*mock_entity.MockEmployeeRepo)
		OnHistoryRepo    func(*mock_entity.MockEmployeeClockHistoryRepo)
		ExpectedResponse oapi.ClockOutResponseObject
		ExpectedErr      string
	}{
		{
			TestName: "validation error",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 0,
				},
			},
			ExpectedErr: "code=422, message=employee ID can't be zero",
		},
		{
			TestName: "select employee error",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return(nil, errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "employee not found",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return(nil, nil)
			},
			ExpectedErr: "code=404, message=ID #99 not found",
		},
		{
			TestName: "history error",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return([]*entity.Employee{{}}, nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Select(gomock.Any(), repokit.Eq{"employee_id": int(99)}, repokit.Sorts{"-id"}).
					Return(nil, errors.New("some-error"))
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "history not found",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return([]*entity.Employee{{}}, nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Select(gomock.Any(), repokit.Eq{"employee_id": int(99)}, repokit.Sorts{"-id"}).
					Return(nil, nil)
			},
			ExpectedErr: "no clock history for employee#99",
		},
		{
			TestName: "already clock out",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return([]*entity.Employee{{}}, nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {

				mechr.EXPECT().
					Select(gomock.Any(), repokit.Eq{"employee_id": int(99)}, repokit.Sorts{"-id"}).
					Return([]*entity.EmployeeClockHistory{
						{
							ID:         77,
							ClockInAt:  &clockInAtUTC,
							ClockOutAt: &clockOutAtUTC,
						},
					}, nil)

			},
			ExpectedErr: "employee#99 already clock-out",
		},
		{
			TestName: "update history error",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return([]*entity.Employee{{}}, nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Select(gomock.Any(), repokit.Eq{"employee_id": int(99)}, repokit.Sorts{"-id"}).
					Return([]*entity.EmployeeClockHistory{
						{
							ID:         77,
							EmployeeID: 99,
							ClockInAt:  &clockInAtUTC,
						},
					}, nil)

				mechr.EXPECT().
					Update(gomock.Any(),
						&entity.EmployeeClockHistory{
							ID:                  77,
							EmployeeID:          99,
							ClockInAt:           &clockInAtUTC,
							ClockOutAt:          &clockOutAtUTC,
							WorkDuration:        "10h0m0s",
							WorkDurationMinutes: 600,
						},
						repokit.Eq{"id": int(77)}).
					Return(int64(-1), errors.New("some-error"))

			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "patch employee error",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return([]*entity.Employee{{}}, nil)

				mer.EXPECT().
					Patch(gomock.Any(), &entity.Employee{LastClockOutAt: &clockOutAtUTC}, repokit.Eq{"id": int(99)}).
					Return(int64(-1), errors.New("some-error"))
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Select(gomock.Any(), repokit.Eq{"employee_id": int(99)}, repokit.Sorts{"-id"}).
					Return([]*entity.EmployeeClockHistory{
						{
							ID:         77,
							EmployeeID: 99,
							ClockInAt:  &clockInAtUTC,
						},
					}, nil)

				mechr.EXPECT().
					Update(gomock.Any(),
						&entity.EmployeeClockHistory{
							ID:                  77,
							EmployeeID:          99,
							ClockInAt:           &clockInAtUTC,
							ClockOutAt:          &clockOutAtUTC,
							WorkDuration:        "10h0m0s",
							WorkDurationMinutes: 600,
						},
						repokit.Eq{"id": int(77)}).
					Return(int64(1), nil)

			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "success",
			Request: oapi.ClockOutRequestObject{
				Body: &oapi.ClockRequest{
					EmployeeId: 99,
				},
			},
			OnEmployeeRepo: func(mer *mock_entity.MockEmployeeRepo) {
				mer.EXPECT().
					Select(gomock.Any(), repokit.Eq{"id": int(99)}).
					Return([]*entity.Employee{{}}, nil)

				mer.EXPECT().
					Patch(gomock.Any(), &entity.Employee{LastClockOutAt: &clockOutAtUTC}, repokit.Eq{"id": int(99)}).
					Return(int64(1), nil)
			},
			OnHistoryRepo: func(mechr *mock_entity.MockEmployeeClockHistoryRepo) {
				mechr.EXPECT().
					Select(gomock.Any(), repokit.Eq{"employee_id": int(99)}, repokit.Sorts{"-id"}).
					Return([]*entity.EmployeeClockHistory{
						{
							ID:         77,
							EmployeeID: 99,
							ClockInAt:  &clockInAtUTC,
						},
					}, nil)

				mechr.EXPECT().
					Update(gomock.Any(),
						&entity.EmployeeClockHistory{
							ID:                  77,
							EmployeeID:          99,
							ClockInAt:           &clockInAtUTC,
							ClockOutAt:          &clockOutAtUTC,
							WorkDuration:        "10h0m0s",
							WorkDurationMinutes: 600,
						},
						repokit.Eq{"id": int(77)}).
					Return(int64(1), nil)

			},
			ExpectedResponse: oapi.ClockOut200JSONResponse{
				Id:                  77,
				EmployeeId:          99,
				ClockInAt:           &clockInAtUTC,
				ClockOutAt:          &clockOutAtUTC,
				WorkDuration:        "10h0m0s",
				WorkDurationMinutes: 600,
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			historyRepo := mock_entity.NewMockEmployeeClockHistoryRepo(ctrl)
			if tt.OnHistoryRepo != nil {
				tt.OnHistoryRepo(historyRepo)
			}

			employeeRepo := mock_entity.NewMockEmployeeRepo(ctrl)
			if tt.OnEmployeeRepo != nil {
				tt.OnEmployeeRepo(employeeRepo)
			}

			svc := service.NewClockSvc(employeeRepo, historyRepo)
			resp, err := svc.ClockOut(context.Background(), tt.Request)
			if tt.ExpectedErr != "" {
				assert.EqualError(t, err, tt.ExpectedErr)
			} else {
				assert.EqualValues(t, tt.ExpectedResponse, resp)
			}
		})
	}
}
