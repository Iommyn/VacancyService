// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\iommy\Desktop\VacancyService\internal\usecases\interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "VacancyService/internal/entity"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVacancyService is a mock of VacancyService interface.
type MockVacancyService struct {
	ctrl     *gomock.Controller
	recorder *MockVacancyServiceMockRecorder
}

// MockVacancyServiceMockRecorder is the mock recorder for MockVacancyService.
type MockVacancyServiceMockRecorder struct {
	mock *MockVacancyService
}

// NewMockVacancyService creates a new mock instance.
func NewMockVacancyService(ctrl *gomock.Controller) *MockVacancyService {
	mock := &MockVacancyService{ctrl: ctrl}
	mock.recorder = &MockVacancyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVacancyService) EXPECT() *MockVacancyServiceMockRecorder {
	return m.recorder
}

// CreateVacancy mocks base method.
func (m *MockVacancyService) CreateVacancy(ctx context.Context, vacancy *entity.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVacancy", ctx, vacancy)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVacancy indicates an expected call of CreateVacancy.
func (mr *MockVacancyServiceMockRecorder) CreateVacancy(ctx, vacancy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVacancy", reflect.TypeOf((*MockVacancyService)(nil).CreateVacancy), ctx, vacancy)
}

// DeleteVacancy mocks base method.
func (m *MockVacancyService) DeleteVacancy(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVacancy", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVacancy indicates an expected call of DeleteVacancy.
func (mr *MockVacancyServiceMockRecorder) DeleteVacancy(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVacancy", reflect.TypeOf((*MockVacancyService)(nil).DeleteVacancy), ctx, id)
}

// GetAllVacancies mocks base method.
func (m *MockVacancyService) GetAllVacancies(ctx context.Context) ([]*entity.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVacancies", ctx)
	ret0, _ := ret[0].([]*entity.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVacancies indicates an expected call of GetAllVacancies.
func (mr *MockVacancyServiceMockRecorder) GetAllVacancies(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVacancies", reflect.TypeOf((*MockVacancyService)(nil).GetAllVacancies), ctx)
}

// GetVacancyByID mocks base method.
func (m *MockVacancyService) GetVacancyByID(ctx context.Context, id int64) (*entity.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVacancyByID", ctx, id)
	ret0, _ := ret[0].(*entity.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVacancyByID indicates an expected call of GetVacancyByID.
func (mr *MockVacancyServiceMockRecorder) GetVacancyByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVacancyByID", reflect.TypeOf((*MockVacancyService)(nil).GetVacancyByID), ctx, id)
}

// UpdateVacancy mocks base method.
func (m *MockVacancyService) UpdateVacancy(ctx context.Context, vacancy *entity.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVacancy", ctx, vacancy)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVacancy indicates an expected call of UpdateVacancy.
func (mr *MockVacancyServiceMockRecorder) UpdateVacancy(ctx, vacancy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVacancy", reflect.TypeOf((*MockVacancyService)(nil).UpdateVacancy), ctx, vacancy)
}
