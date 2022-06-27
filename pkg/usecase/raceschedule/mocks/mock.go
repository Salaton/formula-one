// Code generated by MockGen. DO NOT EDIT.
// Source: schedule.go

// Package mock_raceschedule is a generated GoMock package.
package mock_raceschedule

import (
	context "context"
	reflect "reflect"

	domain "github.com/Salaton/formula-one/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockRaceSchedule is a mock of RaceSchedule interface.
type MockRaceSchedule struct {
	ctrl     *gomock.Controller
	recorder *MockRaceScheduleMockRecorder
}

// MockRaceScheduleMockRecorder is the mock recorder for MockRaceSchedule.
type MockRaceScheduleMockRecorder struct {
	mock *MockRaceSchedule
}

// NewMockRaceSchedule creates a new mock instance.
func NewMockRaceSchedule(ctrl *gomock.Controller) *MockRaceSchedule {
	mock := &MockRaceSchedule{ctrl: ctrl}
	mock.recorder = &MockRaceScheduleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRaceSchedule) EXPECT() *MockRaceScheduleMockRecorder {
	return m.recorder
}

// GetSeasonRaceSchedules mocks base method.
func (m *MockRaceSchedule) GetSeasonRaceSchedules(ctx context.Context, year int) (*domain.DataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeasonRaceSchedules", ctx, year)
	ret0, _ := ret[0].(*domain.DataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeasonRaceSchedules indicates an expected call of GetSeasonRaceSchedules.
func (mr *MockRaceScheduleMockRecorder) GetSeasonRaceSchedules(ctx, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeasonRaceSchedules", reflect.TypeOf((*MockRaceSchedule)(nil).GetSeasonRaceSchedules), ctx, year)
}
