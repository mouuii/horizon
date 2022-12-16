// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cluster/cd/cd.go

// Package mock_cd is a generated GoMock package.
package mock_cd

import (
	context "context"
	reflect "reflect"

	cd "github.com/horizoncd/horizon/pkg/cluster/cd"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockCD is a mock of CD interface.
type MockCD struct {
	ctrl     *gomock.Controller
	recorder *MockCDMockRecorder
}

// MockCDMockRecorder is the mock recorder for MockCD.
type MockCDMockRecorder struct {
	mock *MockCD
}

// NewMockCD creates a new mock instance.
func NewMockCD(ctrl *gomock.Controller) *MockCD {
	mock := &MockCD{ctrl: ctrl}
	mock.recorder = &MockCDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCD) EXPECT() *MockCDMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method.
func (m *MockCD) CreateCluster(ctx context.Context, params *cd.CreateClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster.
func (mr *MockCDMockRecorder) CreateCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockCD)(nil).CreateCluster), ctx, params)
}

// DeleteCluster mocks base method.
func (m *MockCD) DeleteCluster(ctx context.Context, params *cd.DeleteClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockCDMockRecorder) DeleteCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockCD)(nil).DeleteCluster), ctx, params)
}

// DeletePods mocks base method.
func (m *MockCD) DeletePods(ctx context.Context, params *cd.DeletePodsParams) (map[string]cd.OperationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePods", ctx, params)
	ret0, _ := ret[0].(map[string]cd.OperationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePods indicates an expected call of DeletePods.
func (mr *MockCDMockRecorder) DeletePods(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePods", reflect.TypeOf((*MockCD)(nil).DeletePods), ctx, params)
}

// DeployCluster mocks base method.
func (m *MockCD) DeployCluster(ctx context.Context, params *cd.DeployClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployCluster indicates an expected call of DeployCluster.
func (mr *MockCDMockRecorder) DeployCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployCluster", reflect.TypeOf((*MockCD)(nil).DeployCluster), ctx, params)
}

// GetClusterState mocks base method.
func (m *MockCD) GetClusterState(ctx context.Context, params *cd.GetClusterStateParams) (*cd.ClusterState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterState", ctx, params)
	ret0, _ := ret[0].(*cd.ClusterState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterState indicates an expected call of GetClusterState.
func (mr *MockCDMockRecorder) GetClusterState(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterState", reflect.TypeOf((*MockCD)(nil).GetClusterState), ctx, params)
}

// GetContainerLog mocks base method.
func (m *MockCD) GetContainerLog(ctx context.Context, params *cd.GetContainerLogParams) (<-chan string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerLog", ctx, params)
	ret0, _ := ret[0].(<-chan string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerLog indicates an expected call of GetContainerLog.
func (mr *MockCDMockRecorder) GetContainerLog(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerLog", reflect.TypeOf((*MockCD)(nil).GetContainerLog), ctx, params)
}

// GetPod mocks base method.
func (m *MockCD) GetPod(ctx context.Context, params *cd.GetPodParams) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", ctx, params)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPod indicates an expected call of GetPod.
func (mr *MockCDMockRecorder) GetPod(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockCD)(nil).GetPod), ctx, params)
}

// GetPodContainers mocks base method.
func (m *MockCD) GetPodContainers(ctx context.Context, params *cd.GetPodParams) ([]cd.ContainerDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodContainers", ctx, params)
	ret0, _ := ret[0].([]cd.ContainerDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodContainers indicates an expected call of GetPodContainers.
func (mr *MockCDMockRecorder) GetPodContainers(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodContainers", reflect.TypeOf((*MockCD)(nil).GetPodContainers), ctx, params)
}

// GetPodEvents mocks base method.
func (m *MockCD) GetPodEvents(ctx context.Context, params *cd.GetPodEventsParams) ([]cd.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodEvents", ctx, params)
	ret0, _ := ret[0].([]cd.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodEvents indicates an expected call of GetPodEvents.
func (mr *MockCDMockRecorder) GetPodEvents(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodEvents", reflect.TypeOf((*MockCD)(nil).GetPodEvents), ctx, params)
}

// Next mocks base method.
func (m *MockCD) Next(ctx context.Context, params *cd.ClusterNextParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockCDMockRecorder) Next(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCD)(nil).Next), ctx, params)
}

// Offline mocks base method.
func (m *MockCD) Offline(ctx context.Context, params *cd.ExecParams) (map[string]cd.ExecResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offline", ctx, params)
	ret0, _ := ret[0].(map[string]cd.ExecResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offline indicates an expected call of Offline.
func (mr *MockCDMockRecorder) Offline(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offline", reflect.TypeOf((*MockCD)(nil).Offline), ctx, params)
}

// Online mocks base method.
func (m *MockCD) Online(ctx context.Context, params *cd.ExecParams) (map[string]cd.ExecResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Online", ctx, params)
	ret0, _ := ret[0].(map[string]cd.ExecResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Online indicates an expected call of Online.
func (mr *MockCDMockRecorder) Online(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Online", reflect.TypeOf((*MockCD)(nil).Online), ctx, params)
}

// Pause mocks base method.
func (m *MockCD) Pause(ctx context.Context, params *cd.ClusterPauseParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pause", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause.
func (mr *MockCDMockRecorder) Pause(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockCD)(nil).Pause), ctx, params)
}

// Promote mocks base method.
func (m *MockCD) Promote(ctx context.Context, params *cd.ClusterPromoteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Promote", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Promote indicates an expected call of Promote.
func (mr *MockCDMockRecorder) Promote(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Promote", reflect.TypeOf((*MockCD)(nil).Promote), ctx, params)
}

// Resume mocks base method.
func (m *MockCD) Resume(ctx context.Context, params *cd.ClusterResumeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resume", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resume indicates an expected call of Resume.
func (mr *MockCDMockRecorder) Resume(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockCD)(nil).Resume), ctx, params)
}
