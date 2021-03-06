// Code generated by MockGen. DO NOT EDIT.
// Source: src/cartesian/configs/io.go

// Package mock_configs is a generated GoMock package.
package mock_configs

import (
	gomock "github.com/golang/mock/gomock"
	configs "github.com/zored/cartesian/src/cartesian/configs"
	reflect "reflect"
)

// MockIO is a mock of IO interface
type MockIO struct {
	ctrl     *gomock.Controller
	recorder *MockIOMockRecorder
}

// MockIOMockRecorder is the mock recorder for MockIO
type MockIOMockRecorder struct {
	mock *MockIO
}

// NewMockIO creates a new mock instance
func NewMockIO(ctrl *gomock.Controller) *MockIO {
	mock := &MockIO{ctrl: ctrl}
	mock.recorder = &MockIOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIO) EXPECT() *MockIOMockRecorder {
	return m.recorder
}

// GetInput mocks base method
func (m *MockIO) GetInput() *configs.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInput")
	ret0, _ := ret[0].(*configs.Config)
	return ret0
}

// GetInput indicates an expected call of GetInput
func (mr *MockIOMockRecorder) GetInput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInput", reflect.TypeOf((*MockIO)(nil).GetInput))
}

// SetParentIO mocks base method
func (m *MockIO) SetParentIO(arg0 configs.IO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParentIO", arg0)
}

// SetParentIO indicates an expected call of SetParentIO
func (mr *MockIOMockRecorder) SetParentIO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParentIO", reflect.TypeOf((*MockIO)(nil).SetParentIO), arg0)
}
