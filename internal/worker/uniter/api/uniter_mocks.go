// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface_generics.go
//
// Generated by this command:
//
//	mockgen -package api -destination uniter_mocks.go -source=./interface_generics.go
//

// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"
	time "time"

	uniter "github.com/juju/juju/api/agent/uniter"
	types "github.com/juju/juju/api/types"
	application "github.com/juju/juju/core/application"
	network "github.com/juju/juju/core/network"
	watcher "github.com/juju/juju/core/watcher"
	config "github.com/juju/juju/environs/config"
	params "github.com/juju/juju/rpc/params"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockUniterClient is a mock of UniterClient interface.
type MockUniterClient struct {
	ctrl     *gomock.Controller
	recorder *MockUniterClientMockRecorder
}

// MockUniterClientMockRecorder is the mock recorder for MockUniterClient.
type MockUniterClientMockRecorder struct {
	mock *MockUniterClient
}

// NewMockUniterClient creates a new mock instance.
func NewMockUniterClient(ctrl *gomock.Controller) *MockUniterClient {
	mock := &MockUniterClient{ctrl: ctrl}
	mock.recorder = &MockUniterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUniterClient) EXPECT() *MockUniterClientMockRecorder {
	return m.recorder
}

// APIAddresses mocks base method.
func (m *MockUniterClient) APIAddresses() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIAddresses")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIAddresses indicates an expected call of APIAddresses.
func (mr *MockUniterClientMockRecorder) APIAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIAddresses", reflect.TypeOf((*MockUniterClient)(nil).APIAddresses))
}

// Action mocks base method.
func (m *MockUniterClient) Action(ctx context.Context, tag names.ActionTag) (*uniter.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Action", ctx, tag)
	ret0, _ := ret[0].(*uniter.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Action indicates an expected call of Action.
func (mr *MockUniterClientMockRecorder) Action(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Action", reflect.TypeOf((*MockUniterClient)(nil).Action), ctx, tag)
}

// ActionBegin mocks base method.
func (m *MockUniterClient) ActionBegin(ctx context.Context, tag names.ActionTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionBegin", ctx, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionBegin indicates an expected call of ActionBegin.
func (mr *MockUniterClientMockRecorder) ActionBegin(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionBegin", reflect.TypeOf((*MockUniterClient)(nil).ActionBegin), ctx, tag)
}

// ActionFinish mocks base method.
func (m *MockUniterClient) ActionFinish(ctx context.Context, tag names.ActionTag, status string, results map[string]any, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionFinish", ctx, tag, status, results, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionFinish indicates an expected call of ActionFinish.
func (mr *MockUniterClientMockRecorder) ActionFinish(ctx, tag, status, results, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionFinish", reflect.TypeOf((*MockUniterClient)(nil).ActionFinish), ctx, tag, status, results, message)
}

// ActionStatus mocks base method.
func (m *MockUniterClient) ActionStatus(ctx context.Context, tag names.ActionTag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionStatus", ctx, tag)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionStatus indicates an expected call of ActionStatus.
func (mr *MockUniterClientMockRecorder) ActionStatus(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionStatus", reflect.TypeOf((*MockUniterClient)(nil).ActionStatus), ctx, tag)
}

// Application mocks base method.
func (m *MockUniterClient) Application(ctx context.Context, tag names.ApplicationTag) (Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", ctx, tag)
	ret0, _ := ret[0].(Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockUniterClientMockRecorder) Application(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUniterClient)(nil).Application), ctx, tag)
}

// Charm mocks base method.
func (m *MockUniterClient) Charm(curl string) (Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", curl)
	ret0, _ := ret[0].(Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockUniterClientMockRecorder) Charm(curl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockUniterClient)(nil).Charm), curl)
}

// CloudAPIVersion mocks base method.
func (m *MockUniterClient) CloudAPIVersion(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudAPIVersion", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudAPIVersion indicates an expected call of CloudAPIVersion.
func (mr *MockUniterClientMockRecorder) CloudAPIVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudAPIVersion", reflect.TypeOf((*MockUniterClient)(nil).CloudAPIVersion), arg0)
}

// CloudSpec mocks base method.
func (m *MockUniterClient) CloudSpec(arg0 context.Context) (*params.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudSpec", arg0)
	ret0, _ := ret[0].(*params.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudSpec indicates an expected call of CloudSpec.
func (mr *MockUniterClientMockRecorder) CloudSpec(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSpec", reflect.TypeOf((*MockUniterClient)(nil).CloudSpec), arg0)
}

// DestroyUnitStorageAttachments mocks base method.
func (m *MockUniterClient) DestroyUnitStorageAttachments(arg0 names.UnitTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyUnitStorageAttachments", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyUnitStorageAttachments indicates an expected call of DestroyUnitStorageAttachments.
func (mr *MockUniterClientMockRecorder) DestroyUnitStorageAttachments(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyUnitStorageAttachments", reflect.TypeOf((*MockUniterClient)(nil).DestroyUnitStorageAttachments), arg0)
}

// GoalState mocks base method.
func (m *MockUniterClient) GoalState(arg0 context.Context) (application.GoalState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoalState", arg0)
	ret0, _ := ret[0].(application.GoalState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GoalState indicates an expected call of GoalState.
func (mr *MockUniterClientMockRecorder) GoalState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoalState", reflect.TypeOf((*MockUniterClient)(nil).GoalState), arg0)
}

// LeadershipSettings mocks base method.
func (m *MockUniterClient) LeadershipSettings() uniter.LeadershipSettingsAccessor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipSettings")
	ret0, _ := ret[0].(uniter.LeadershipSettingsAccessor)
	return ret0
}

// LeadershipSettings indicates an expected call of LeadershipSettings.
func (mr *MockUniterClientMockRecorder) LeadershipSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipSettings", reflect.TypeOf((*MockUniterClient)(nil).LeadershipSettings))
}

// Model mocks base method.
func (m *MockUniterClient) Model(arg0 context.Context) (*types.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", arg0)
	ret0, _ := ret[0].(*types.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockUniterClientMockRecorder) Model(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockUniterClient)(nil).Model), arg0)
}

// ModelConfig mocks base method.
func (m *MockUniterClient) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockUniterClientMockRecorder) ModelConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockUniterClient)(nil).ModelConfig), arg0)
}

// OpenedMachinePortRangesByEndpoint mocks base method.
func (m *MockUniterClient) OpenedMachinePortRangesByEndpoint(ctx context.Context, machineTag names.MachineTag) (map[names.UnitTag]network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenedMachinePortRangesByEndpoint", ctx, machineTag)
	ret0, _ := ret[0].(map[names.UnitTag]network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenedMachinePortRangesByEndpoint indicates an expected call of OpenedMachinePortRangesByEndpoint.
func (mr *MockUniterClientMockRecorder) OpenedMachinePortRangesByEndpoint(ctx, machineTag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenedMachinePortRangesByEndpoint", reflect.TypeOf((*MockUniterClient)(nil).OpenedMachinePortRangesByEndpoint), ctx, machineTag)
}

// OpenedPortRangesByEndpoint mocks base method.
func (m *MockUniterClient) OpenedPortRangesByEndpoint(ctx context.Context) (map[names.UnitTag]network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenedPortRangesByEndpoint", ctx)
	ret0, _ := ret[0].(map[names.UnitTag]network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenedPortRangesByEndpoint indicates an expected call of OpenedPortRangesByEndpoint.
func (mr *MockUniterClientMockRecorder) OpenedPortRangesByEndpoint(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenedPortRangesByEndpoint", reflect.TypeOf((*MockUniterClient)(nil).OpenedPortRangesByEndpoint), ctx)
}

// Relation mocks base method.
func (m *MockUniterClient) Relation(ctx context.Context, tag names.RelationTag) (Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relation", ctx, tag)
	ret0, _ := ret[0].(Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation indicates an expected call of Relation.
func (mr *MockUniterClientMockRecorder) Relation(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation", reflect.TypeOf((*MockUniterClient)(nil).Relation), ctx, tag)
}

// RelationById mocks base method.
func (m *MockUniterClient) RelationById(arg0 context.Context, arg1 int) (Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationById", arg0, arg1)
	ret0, _ := ret[0].(Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelationById indicates an expected call of RelationById.
func (mr *MockUniterClientMockRecorder) RelationById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationById", reflect.TypeOf((*MockUniterClient)(nil).RelationById), arg0, arg1)
}

// RemoveStorageAttachment mocks base method.
func (m *MockUniterClient) RemoveStorageAttachment(arg0 names.StorageTag, arg1 names.UnitTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveStorageAttachment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveStorageAttachment indicates an expected call of RemoveStorageAttachment.
func (mr *MockUniterClientMockRecorder) RemoveStorageAttachment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStorageAttachment", reflect.TypeOf((*MockUniterClient)(nil).RemoveStorageAttachment), arg0, arg1)
}

// SetUnitWorkloadVersion mocks base method.
func (m *MockUniterClient) SetUnitWorkloadVersion(ctx context.Context, tag names.UnitTag, version string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitWorkloadVersion", ctx, tag, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitWorkloadVersion indicates an expected call of SetUnitWorkloadVersion.
func (mr *MockUniterClientMockRecorder) SetUnitWorkloadVersion(ctx, tag, version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitWorkloadVersion", reflect.TypeOf((*MockUniterClient)(nil).SetUnitWorkloadVersion), ctx, tag, version)
}

// StorageAttachment mocks base method.
func (m *MockUniterClient) StorageAttachment(arg0 names.StorageTag, arg1 names.UnitTag) (params.StorageAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAttachment", arg0, arg1)
	ret0, _ := ret[0].(params.StorageAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAttachment indicates an expected call of StorageAttachment.
func (mr *MockUniterClientMockRecorder) StorageAttachment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAttachment", reflect.TypeOf((*MockUniterClient)(nil).StorageAttachment), arg0, arg1)
}

// StorageAttachmentLife mocks base method.
func (m *MockUniterClient) StorageAttachmentLife(arg0 []params.StorageAttachmentId) ([]params.LifeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAttachmentLife", arg0)
	ret0, _ := ret[0].([]params.LifeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAttachmentLife indicates an expected call of StorageAttachmentLife.
func (mr *MockUniterClientMockRecorder) StorageAttachmentLife(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAttachmentLife", reflect.TypeOf((*MockUniterClient)(nil).StorageAttachmentLife), arg0)
}

// Unit mocks base method.
func (m *MockUniterClient) Unit(ctx context.Context, tag names.UnitTag) (Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", ctx, tag)
	ret0, _ := ret[0].(Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockUniterClientMockRecorder) Unit(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockUniterClient)(nil).Unit), ctx, tag)
}

// UnitStorageAttachments mocks base method.
func (m *MockUniterClient) UnitStorageAttachments(arg0 names.UnitTag) ([]params.StorageAttachmentId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitStorageAttachments", arg0)
	ret0, _ := ret[0].([]params.StorageAttachmentId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitStorageAttachments indicates an expected call of UnitStorageAttachments.
func (mr *MockUniterClientMockRecorder) UnitStorageAttachments(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitStorageAttachments", reflect.TypeOf((*MockUniterClient)(nil).UnitStorageAttachments), arg0)
}

// UnitWorkloadVersion mocks base method.
func (m *MockUniterClient) UnitWorkloadVersion(ctx context.Context, tag names.UnitTag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitWorkloadVersion", ctx, tag)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitWorkloadVersion indicates an expected call of UnitWorkloadVersion.
func (mr *MockUniterClientMockRecorder) UnitWorkloadVersion(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitWorkloadVersion", reflect.TypeOf((*MockUniterClient)(nil).UnitWorkloadVersion), ctx, tag)
}

// UpdateStatusHookInterval mocks base method.
func (m *MockUniterClient) UpdateStatusHookInterval() (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusHookInterval")
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatusHookInterval indicates an expected call of UpdateStatusHookInterval.
func (mr *MockUniterClientMockRecorder) UpdateStatusHookInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusHookInterval", reflect.TypeOf((*MockUniterClient)(nil).UpdateStatusHookInterval))
}

// WatchRelationUnits mocks base method.
func (m *MockUniterClient) WatchRelationUnits(arg0 context.Context, arg1 names.RelationTag, arg2 names.UnitTag) (watcher.RelationUnitsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRelationUnits", arg0, arg1, arg2)
	ret0, _ := ret[0].(watcher.RelationUnitsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRelationUnits indicates an expected call of WatchRelationUnits.
func (mr *MockUniterClientMockRecorder) WatchRelationUnits(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRelationUnits", reflect.TypeOf((*MockUniterClient)(nil).WatchRelationUnits), arg0, arg1, arg2)
}

// WatchStorageAttachment mocks base method.
func (m *MockUniterClient) WatchStorageAttachment(arg0 names.StorageTag, arg1 names.UnitTag) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchStorageAttachment", arg0, arg1)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchStorageAttachment indicates an expected call of WatchStorageAttachment.
func (mr *MockUniterClientMockRecorder) WatchStorageAttachment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchStorageAttachment", reflect.TypeOf((*MockUniterClient)(nil).WatchStorageAttachment), arg0, arg1)
}

// WatchUpdateStatusHookInterval mocks base method.
func (m *MockUniterClient) WatchUpdateStatusHookInterval() (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpdateStatusHookInterval")
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpdateStatusHookInterval indicates an expected call of WatchUpdateStatusHookInterval.
func (mr *MockUniterClientMockRecorder) WatchUpdateStatusHookInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpdateStatusHookInterval", reflect.TypeOf((*MockUniterClient)(nil).WatchUpdateStatusHookInterval))
}