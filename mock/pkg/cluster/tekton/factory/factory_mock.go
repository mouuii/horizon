// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cluster/tekton/factory/factory.go

// Package mock_factory is a generated GoMock package.
package mock_factory

import (
	tekton "github.com/horizoncd/horizon/pkg/cluster/tekton"
	collector "github.com/horizoncd/horizon/pkg/cluster/tekton/collector"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// GetTekton mocks base method
func (m *MockFactory) GetTekton(environment string) (tekton.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTekton", environment)
	ret0, _ := ret[0].(tekton.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTekton indicates an expected call of GetTekton
func (mr *MockFactoryMockRecorder) GetTekton(environment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTekton", reflect.TypeOf((*MockFactory)(nil).GetTekton), environment)
}

// GetTektonCollector mocks base method
func (m *MockFactory) GetTektonCollector(environment string) (collector.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTektonCollector", environment)
	ret0, _ := ret[0].(collector.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTektonCollector indicates an expected call of GetTektonCollector
func (mr *MockFactoryMockRecorder) GetTektonCollector(environment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTektonCollector", reflect.TypeOf((*MockFactory)(nil).GetTektonCollector), environment)
}
