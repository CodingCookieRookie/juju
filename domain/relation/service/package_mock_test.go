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
	unit "github.com/juju/juju/core/unit"
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

// GetPeerRelationUUIDByEndpointIdentifiers mocks base method.
func (m *MockState) GetPeerRelationUUIDByEndpointIdentifiers(arg0 context.Context, arg1 relation0.EndpointIdentifier) (relation.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerRelationUUIDByEndpointIdentifiers", arg0, arg1)
	ret0, _ := ret[0].(relation.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerRelationUUIDByEndpointIdentifiers indicates an expected call of GetPeerRelationUUIDByEndpointIdentifiers.
func (mr *MockStateMockRecorder) GetPeerRelationUUIDByEndpointIdentifiers(arg0, arg1 any) *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerRelationUUIDByEndpointIdentifiers", reflect.TypeOf((*MockState)(nil).GetPeerRelationUUIDByEndpointIdentifiers), arg0, arg1)
	return &MockStateGetPeerRelationUUIDByEndpointIdentifiersCall{Call: call}
}

// MockStateGetPeerRelationUUIDByEndpointIdentifiersCall wrap *gomock.Call
type MockStateGetPeerRelationUUIDByEndpointIdentifiersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall) Return(arg0 relation.UUID, arg1 error) *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall) Do(f func(context.Context, relation0.EndpointIdentifier) (relation.UUID, error)) *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall) DoAndReturn(f func(context.Context, relation0.EndpointIdentifier) (relation.UUID, error)) *MockStateGetPeerRelationUUIDByEndpointIdentifiersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRegularRelationUUIDByEndpointIdentifiers mocks base method.
func (m *MockState) GetRegularRelationUUIDByEndpointIdentifiers(arg0 context.Context, arg1, arg2 relation0.EndpointIdentifier) (relation.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegularRelationUUIDByEndpointIdentifiers", arg0, arg1, arg2)
	ret0, _ := ret[0].(relation.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegularRelationUUIDByEndpointIdentifiers indicates an expected call of GetRegularRelationUUIDByEndpointIdentifiers.
func (mr *MockStateMockRecorder) GetRegularRelationUUIDByEndpointIdentifiers(arg0, arg1, arg2 any) *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegularRelationUUIDByEndpointIdentifiers", reflect.TypeOf((*MockState)(nil).GetRegularRelationUUIDByEndpointIdentifiers), arg0, arg1, arg2)
	return &MockStateGetRegularRelationUUIDByEndpointIdentifiersCall{Call: call}
}

// MockStateGetRegularRelationUUIDByEndpointIdentifiersCall wrap *gomock.Call
type MockStateGetRegularRelationUUIDByEndpointIdentifiersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall) Return(arg0 relation.UUID, arg1 error) *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall) Do(f func(context.Context, relation0.EndpointIdentifier, relation0.EndpointIdentifier) (relation.UUID, error)) *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall) DoAndReturn(f func(context.Context, relation0.EndpointIdentifier, relation0.EndpointIdentifier) (relation.UUID, error)) *MockStateGetRegularRelationUUIDByEndpointIdentifiersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRelationDetails mocks base method.
func (m *MockState) GetRelationDetails(arg0 context.Context, arg1 int) (relation0.RelationDetailsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationDetails", arg0, arg1)
	ret0, _ := ret[0].(relation0.RelationDetailsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationDetails indicates an expected call of GetRelationDetails.
func (mr *MockStateMockRecorder) GetRelationDetails(arg0, arg1 any) *MockStateGetRelationDetailsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationDetails", reflect.TypeOf((*MockState)(nil).GetRelationDetails), arg0, arg1)
	return &MockStateGetRelationDetailsCall{Call: call}
}

// MockStateGetRelationDetailsCall wrap *gomock.Call
type MockStateGetRelationDetailsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRelationDetailsCall) Return(arg0 relation0.RelationDetailsResult, arg1 error) *MockStateGetRelationDetailsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRelationDetailsCall) Do(f func(context.Context, int) (relation0.RelationDetailsResult, error)) *MockStateGetRelationDetailsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRelationDetailsCall) DoAndReturn(f func(context.Context, int) (relation0.RelationDetailsResult, error)) *MockStateGetRelationDetailsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
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

// GetRelationEndpoints mocks base method.
func (m *MockState) GetRelationEndpoints(arg0 context.Context, arg1 relation.UUID) ([]relation0.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]relation0.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationEndpoints indicates an expected call of GetRelationEndpoints.
func (mr *MockStateMockRecorder) GetRelationEndpoints(arg0, arg1 any) *MockStateGetRelationEndpointsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationEndpoints", reflect.TypeOf((*MockState)(nil).GetRelationEndpoints), arg0, arg1)
	return &MockStateGetRelationEndpointsCall{Call: call}
}

// MockStateGetRelationEndpointsCall wrap *gomock.Call
type MockStateGetRelationEndpointsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRelationEndpointsCall) Return(arg0 []relation0.Endpoint, arg1 error) *MockStateGetRelationEndpointsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRelationEndpointsCall) Do(f func(context.Context, relation.UUID) ([]relation0.Endpoint, error)) *MockStateGetRelationEndpointsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRelationEndpointsCall) DoAndReturn(f func(context.Context, relation.UUID) ([]relation0.Endpoint, error)) *MockStateGetRelationEndpointsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRelationID mocks base method.
func (m *MockState) GetRelationID(arg0 context.Context, arg1 relation.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationID", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationID indicates an expected call of GetRelationID.
func (mr *MockStateMockRecorder) GetRelationID(arg0, arg1 any) *MockStateGetRelationIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationID", reflect.TypeOf((*MockState)(nil).GetRelationID), arg0, arg1)
	return &MockStateGetRelationIDCall{Call: call}
}

// MockStateGetRelationIDCall wrap *gomock.Call
type MockStateGetRelationIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRelationIDCall) Return(arg0 int, arg1 error) *MockStateGetRelationIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRelationIDCall) Do(f func(context.Context, relation.UUID) (int, error)) *MockStateGetRelationIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRelationIDCall) DoAndReturn(f func(context.Context, relation.UUID) (int, error)) *MockStateGetRelationIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRelationUUIDByID mocks base method.
func (m *MockState) GetRelationUUIDByID(arg0 context.Context, arg1 int) (relation.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationUUIDByID", arg0, arg1)
	ret0, _ := ret[0].(relation.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationUUIDByID indicates an expected call of GetRelationUUIDByID.
func (mr *MockStateMockRecorder) GetRelationUUIDByID(arg0, arg1 any) *MockStateGetRelationUUIDByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationUUIDByID", reflect.TypeOf((*MockState)(nil).GetRelationUUIDByID), arg0, arg1)
	return &MockStateGetRelationUUIDByIDCall{Call: call}
}

// MockStateGetRelationUUIDByIDCall wrap *gomock.Call
type MockStateGetRelationUUIDByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRelationUUIDByIDCall) Return(arg0 relation.UUID, arg1 error) *MockStateGetRelationUUIDByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRelationUUIDByIDCall) Do(f func(context.Context, int) (relation.UUID, error)) *MockStateGetRelationUUIDByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRelationUUIDByIDCall) DoAndReturn(f func(context.Context, int) (relation.UUID, error)) *MockStateGetRelationUUIDByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRelationsStatusForUnit mocks base method.
func (m *MockState) GetRelationsStatusForUnit(arg0 context.Context, arg1 unit.UUID) ([]relation0.RelationUnitStatusResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationsStatusForUnit", arg0, arg1)
	ret0, _ := ret[0].([]relation0.RelationUnitStatusResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationsStatusForUnit indicates an expected call of GetRelationsStatusForUnit.
func (mr *MockStateMockRecorder) GetRelationsStatusForUnit(arg0, arg1 any) *MockStateGetRelationsStatusForUnitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationsStatusForUnit", reflect.TypeOf((*MockState)(nil).GetRelationsStatusForUnit), arg0, arg1)
	return &MockStateGetRelationsStatusForUnitCall{Call: call}
}

// MockStateGetRelationsStatusForUnitCall wrap *gomock.Call
type MockStateGetRelationsStatusForUnitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetRelationsStatusForUnitCall) Return(arg0 []relation0.RelationUnitStatusResult, arg1 error) *MockStateGetRelationsStatusForUnitCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetRelationsStatusForUnitCall) Do(f func(context.Context, unit.UUID) ([]relation0.RelationUnitStatusResult, error)) *MockStateGetRelationsStatusForUnitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetRelationsStatusForUnitCall) DoAndReturn(f func(context.Context, unit.UUID) ([]relation0.RelationUnitStatusResult, error)) *MockStateGetRelationsStatusForUnitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatcherApplicationSettingsNamespace mocks base method.
func (m *MockState) WatcherApplicationSettingsNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatcherApplicationSettingsNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// WatcherApplicationSettingsNamespace indicates an expected call of WatcherApplicationSettingsNamespace.
func (mr *MockStateMockRecorder) WatcherApplicationSettingsNamespace() *MockStateWatcherApplicationSettingsNamespaceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatcherApplicationSettingsNamespace", reflect.TypeOf((*MockState)(nil).WatcherApplicationSettingsNamespace))
	return &MockStateWatcherApplicationSettingsNamespaceCall{Call: call}
}

// MockStateWatcherApplicationSettingsNamespaceCall wrap *gomock.Call
type MockStateWatcherApplicationSettingsNamespaceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateWatcherApplicationSettingsNamespaceCall) Return(arg0 string) *MockStateWatcherApplicationSettingsNamespaceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateWatcherApplicationSettingsNamespaceCall) Do(f func() string) *MockStateWatcherApplicationSettingsNamespaceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateWatcherApplicationSettingsNamespaceCall) DoAndReturn(f func() string) *MockStateWatcherApplicationSettingsNamespaceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
