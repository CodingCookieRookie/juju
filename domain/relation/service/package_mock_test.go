// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/relation/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/relation/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	relation "github.com/juju/juju/core/relation"
	relation0 "github.com/juju/juju/domain/relation"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// GetRelationEndpointUUID mocks base method.
func (m *MockState) GetRelationEndpointUUID(arg0 context.Context, arg1 relation0.GetRelationEndpointUUIDArgs) (relation.EndpointUUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationEndpointUUID", arg0, arg1)
	ret0, _ := ret[0].(relation.EndpointUUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationEndpointUUID indicates an expected call of GetRelationEndpointUUID.
func (mr *MockStateMockRecorder) GetRelationEndpointUUID(arg0, arg1 any) *MockStateGetRelationEndpointUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationEndpointUUID", reflect.TypeOf((*MockState)(nil).GetRelationEndpointUUID), arg0, arg1)
	return &MockStateGetRelationEndpointUUIDCall{Call: call}
}

// MockStateGetRelationEndpointUUIDCall wrap *gomock.Call
type MockStateGetRelationEndpointUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRelationEndpointUUIDCall) Return(arg0 relation.EndpointUUID, arg1 error) *MockStateGetRelationEndpointUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRelationEndpointUUIDCall) Do(f func(context.Context, relation0.GetRelationEndpointUUIDArgs) (relation.EndpointUUID, error)) *MockStateGetRelationEndpointUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRelationEndpointUUIDCall) DoAndReturn(f func(context.Context, relation0.GetRelationEndpointUUIDArgs) (relation.EndpointUUID, error)) *MockStateGetRelationEndpointUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchRelationApplicationSettingTable mocks base method.
func (m *MockState) WatchRelationApplicationSettingTable() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRelationApplicationSettingTable")
	ret0, _ := ret[0].(string)
	return ret0
}

// WatchRelationApplicationSettingTable indicates an expected call of WatchRelationApplicationSettingTable.
func (mr *MockStateMockRecorder) WatchRelationApplicationSettingTable() *MockStateWatchRelationApplicationSettingTableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRelationApplicationSettingTable", reflect.TypeOf((*MockState)(nil).WatchRelationApplicationSettingTable))
	return &MockStateWatchRelationApplicationSettingTableCall{Call: call}
}

// MockStateWatchRelationApplicationSettingTableCall wrap *gomock.Call
type MockStateWatchRelationApplicationSettingTableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateWatchRelationApplicationSettingTableCall) Return(arg0 string) *MockStateWatchRelationApplicationSettingTableCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateWatchRelationApplicationSettingTableCall) Do(f func() string) *MockStateWatchRelationApplicationSettingTableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateWatchRelationApplicationSettingTableCall) DoAndReturn(f func() string) *MockStateWatchRelationApplicationSettingTableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
