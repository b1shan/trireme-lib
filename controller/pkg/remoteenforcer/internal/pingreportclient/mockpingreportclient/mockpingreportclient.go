// Code generated by MockGen. DO NOT EDIT.
// Source: controller/pkg/remoteenforcer/internal/pingreportclient/interfaces.go

// Package mockpingreportclient is a generated GoMock package.
package mockpingreportclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPingReportClient is a mock of PingReportClient interface
// nolint
type MockPingReportClient struct {
	ctrl     *gomock.Controller
	recorder *MockPingReportClientMockRecorder
}

// MockPingReportClientMockRecorder is the mock recorder for MockPingReportClient
// nolint
type MockPingReportClientMockRecorder struct {
	mock *MockPingReportClient
}

// NewMockPingReportClient creates a new mock instance
// nolint
func NewMockPingReportClient(ctrl *gomock.Controller) *MockPingReportClient {
	mock := &MockPingReportClient{ctrl: ctrl}
	mock.recorder = &MockPingReportClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockPingReportClient) EXPECT() *MockPingReportClientMockRecorder {
	return m.recorder
}

// Run mocks base method
// nolint
func (m *MockPingReportClient) Run(ctx context.Context) error {
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
// nolint
func (mr *MockPingReportClientMockRecorder) Run(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPingReportClient)(nil).Run), ctx)
}
