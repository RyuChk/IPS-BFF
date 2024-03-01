// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_usermanagerclient is a generated GoMock package.
package mock_usermanagerclient

import (
	context "context"
	reflect "reflect"

	userv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/user/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetCoordinate mocks base method.
func (m *MockService) GetCoordinate(ctx context.Context, body *userv1.GetCoordinateRequest) (*userv1.GetCoordinateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoordinate", ctx, body)
	ret0, _ := ret[0].(*userv1.GetCoordinateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoordinate indicates an expected call of GetCoordinate.
func (mr *MockServiceMockRecorder) GetCoordinate(ctx, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoordinate", reflect.TypeOf((*MockService)(nil).GetCoordinate), ctx, body)
}