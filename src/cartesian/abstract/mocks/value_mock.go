// Code generated by MockGen. DO NOT EDIT.
// Source: src/cartesian/abstract/value.go

// Package mock_abstract is a generated GoMock package.
package mock_abstract

import (
	gomock "github.com/golang/mock/gomock"
)

// MockValue is a mock of Value interface
type MockValue struct {
	ctrl     *gomock.Controller
	recorder *MockValueMockRecorder
}

// MockValueMockRecorder is the mock recorder for MockValue
type MockValueMockRecorder struct {
	mock *MockValue
}

// NewMockValue creates a new mock instance
func NewMockValue(ctrl *gomock.Controller) *MockValue {
	mock := &MockValue{ctrl: ctrl}
	mock.recorder = &MockValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValue) EXPECT() *MockValueMockRecorder {
	return m.recorder
}