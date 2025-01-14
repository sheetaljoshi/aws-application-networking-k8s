// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/deploy/lattice/target_group_manager.go

// Package lattice is a generated GoMock package.
package lattice

import (
	context "context"
	reflect "reflect"

	lattice "github.com/aws/aws-application-networking-k8s/pkg/model/lattice"
	gomock "github.com/golang/mock/gomock"
)

// MockTargetGroupManager is a mock of TargetGroupManager interface.
type MockTargetGroupManager struct {
	ctrl     *gomock.Controller
	recorder *MockTargetGroupManagerMockRecorder
}

// MockTargetGroupManagerMockRecorder is the mock recorder for MockTargetGroupManager.
type MockTargetGroupManagerMockRecorder struct {
	mock *MockTargetGroupManager
}

// NewMockTargetGroupManager creates a new mock instance.
func NewMockTargetGroupManager(ctrl *gomock.Controller) *MockTargetGroupManager {
	mock := &MockTargetGroupManager{ctrl: ctrl}
	mock.recorder = &MockTargetGroupManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetGroupManager) EXPECT() *MockTargetGroupManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTargetGroupManager) Create(ctx context.Context, targetGroup *lattice.TargetGroup) (lattice.TargetGroupStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, targetGroup)
	ret0, _ := ret[0].(lattice.TargetGroupStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTargetGroupManagerMockRecorder) Create(ctx, targetGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTargetGroupManager)(nil).Create), ctx, targetGroup)
}

// Delete mocks base method.
func (m *MockTargetGroupManager) Delete(ctx context.Context, targetGroup *lattice.TargetGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, targetGroup)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTargetGroupManagerMockRecorder) Delete(ctx, targetGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTargetGroupManager)(nil).Delete), ctx, targetGroup)
}

// Get mocks base method.
func (m *MockTargetGroupManager) Get(tx context.Context, targetGroup *lattice.TargetGroup) (lattice.TargetGroupStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", tx, targetGroup)
	ret0, _ := ret[0].(lattice.TargetGroupStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTargetGroupManagerMockRecorder) Get(tx, targetGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTargetGroupManager)(nil).Get), tx, targetGroup)
}

// List mocks base method.
func (m *MockTargetGroupManager) List(ctx context.Context) ([]targetGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]targetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTargetGroupManagerMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTargetGroupManager)(nil).List), ctx)
}
