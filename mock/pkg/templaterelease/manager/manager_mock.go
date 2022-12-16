// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	reflect "reflect"

	models "github.com/horizoncd/horizon/pkg/application/models"
	models0 "github.com/horizoncd/horizon/pkg/cluster/models"
	models1 "github.com/horizoncd/horizon/pkg/templaterelease/models"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockManager) Create(ctx context.Context, templateRelease *models1.TemplateRelease) (*models1.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, templateRelease)
	ret0, _ := ret[0].(*models1.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockManagerMockRecorder) Create(ctx, templateRelease interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, templateRelease)
}

// DeleteByID mocks base method.
func (m *MockManager) DeleteByID(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockManagerMockRecorder) DeleteByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockManager)(nil).DeleteByID), ctx, id)
}

// GetByID mocks base method.
func (m *MockManager) GetByID(ctx context.Context, releaseID uint) (*models1.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, releaseID)
	ret0, _ := ret[0].(*models1.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockManagerMockRecorder) GetByID(ctx, releaseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockManager)(nil).GetByID), ctx, releaseID)
}

// GetByTemplateNameAndRelease mocks base method.
func (m *MockManager) GetByTemplateNameAndRelease(ctx context.Context, templateName, release string) (*models1.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTemplateNameAndRelease", ctx, templateName, release)
	ret0, _ := ret[0].(*models1.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTemplateNameAndRelease indicates an expected call of GetByTemplateNameAndRelease.
func (mr *MockManagerMockRecorder) GetByTemplateNameAndRelease(ctx, templateName, release interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTemplateNameAndRelease", reflect.TypeOf((*MockManager)(nil).GetByTemplateNameAndRelease), ctx, templateName, release)
}

// GetRefOfApplication mocks base method.
func (m *MockManager) GetRefOfApplication(ctx context.Context, id uint) ([]*models.Application, uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefOfApplication", ctx, id)
	ret0, _ := ret[0].([]*models.Application)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRefOfApplication indicates an expected call of GetRefOfApplication.
func (mr *MockManagerMockRecorder) GetRefOfApplication(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefOfApplication", reflect.TypeOf((*MockManager)(nil).GetRefOfApplication), ctx, id)
}

// GetRefOfCluster mocks base method.
func (m *MockManager) GetRefOfCluster(ctx context.Context, id uint) ([]*models0.Cluster, uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefOfCluster", ctx, id)
	ret0, _ := ret[0].([]*models0.Cluster)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRefOfCluster indicates an expected call of GetRefOfCluster.
func (mr *MockManagerMockRecorder) GetRefOfCluster(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefOfCluster", reflect.TypeOf((*MockManager)(nil).GetRefOfCluster), ctx, id)
}

// ListByTemplateID mocks base method.
func (m *MockManager) ListByTemplateID(ctx context.Context, id uint) ([]*models1.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByTemplateID", ctx, id)
	ret0, _ := ret[0].([]*models1.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByTemplateID indicates an expected call of ListByTemplateID.
func (mr *MockManagerMockRecorder) ListByTemplateID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByTemplateID", reflect.TypeOf((*MockManager)(nil).ListByTemplateID), ctx, id)
}

// ListByTemplateName mocks base method.
func (m *MockManager) ListByTemplateName(ctx context.Context, templateName string) ([]*models1.TemplateRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByTemplateName", ctx, templateName)
	ret0, _ := ret[0].([]*models1.TemplateRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByTemplateName indicates an expected call of ListByTemplateName.
func (mr *MockManagerMockRecorder) ListByTemplateName(ctx, templateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByTemplateName", reflect.TypeOf((*MockManager)(nil).ListByTemplateName), ctx, templateName)
}

// UpdateByID mocks base method.
func (m *MockManager) UpdateByID(ctx context.Context, releaseID uint, release *models1.TemplateRelease) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, releaseID, release)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockManagerMockRecorder) UpdateByID(ctx, releaseID, release interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockManager)(nil).UpdateByID), ctx, releaseID, release)
}
