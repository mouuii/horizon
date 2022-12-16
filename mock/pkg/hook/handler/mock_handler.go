// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/hook/handler/handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	hook "github.com/horizoncd/horizon/pkg/hook/hook"
	gomock "github.com/golang/mock/gomock"
)

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockEventHandler) Process(event *hook.EventCtx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockEventHandlerMockRecorder) Process(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockEventHandler)(nil).Process), event)
}
