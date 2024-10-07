// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/controller (interfaces: Backend,Application,Relation)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/state_mock.go github.com/juju/juju/apiserver/facades/client/controller Backend,Application,Relation
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	controller "github.com/juju/juju/apiserver/facades/client/controller"
	charm "github.com/juju/juju/internal/charm"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// AllModelUUIDs mocks base method.
func (m *MockBackend) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs.
func (mr *MockBackendMockRecorder) AllModelUUIDs() *MockBackendAllModelUUIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockBackend)(nil).AllModelUUIDs))
	return &MockBackendAllModelUUIDsCall{Call: call}
}

// MockBackendAllModelUUIDsCall wrap *gomock.Call
type MockBackendAllModelUUIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendAllModelUUIDsCall) Return(arg0 []string, arg1 error) *MockBackendAllModelUUIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendAllModelUUIDsCall) Do(f func() ([]string, error)) *MockBackendAllModelUUIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendAllModelUUIDsCall) DoAndReturn(f func() ([]string, error)) *MockBackendAllModelUUIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Application mocks base method.
func (m *MockBackend) Application(arg0 string) (controller.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(controller.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockBackendMockRecorder) Application(arg0 any) *MockBackendApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBackend)(nil).Application), arg0)
	return &MockBackendApplicationCall{Call: call}
}

// MockBackendApplicationCall wrap *gomock.Call
type MockBackendApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendApplicationCall) Return(arg0 controller.Application, arg1 error) *MockBackendApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendApplicationCall) Do(f func(string) (controller.Application, error)) *MockBackendApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendApplicationCall) DoAndReturn(f func(string) (controller.Application, error)) *MockBackendApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerModelUUID mocks base method.
func (m *MockBackend) ControllerModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerModelUUID indicates an expected call of ControllerModelUUID.
func (mr *MockBackendMockRecorder) ControllerModelUUID() *MockBackendControllerModelUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerModelUUID", reflect.TypeOf((*MockBackend)(nil).ControllerModelUUID))
	return &MockBackendControllerModelUUIDCall{Call: call}
}

// MockBackendControllerModelUUIDCall wrap *gomock.Call
type MockBackendControllerModelUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendControllerModelUUIDCall) Return(arg0 string) *MockBackendControllerModelUUIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendControllerModelUUIDCall) Do(f func() string) *MockBackendControllerModelUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendControllerModelUUIDCall) DoAndReturn(f func() string) *MockBackendControllerModelUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockBackend) Model() (*state.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*state.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockBackendMockRecorder) Model() *MockBackendModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockBackend)(nil).Model))
	return &MockBackendModelCall{Call: call}
}

// MockBackendModelCall wrap *gomock.Call
type MockBackendModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendModelCall) Return(arg0 *state.Model, arg1 error) *MockBackendModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendModelCall) Do(f func() (*state.Model, error)) *MockBackendModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendModelCall) DoAndReturn(f func() (*state.Model, error)) *MockBackendModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelExists mocks base method.
func (m *MockBackend) ModelExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelExists indicates an expected call of ModelExists.
func (mr *MockBackendMockRecorder) ModelExists(arg0 any) *MockBackendModelExistsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelExists", reflect.TypeOf((*MockBackend)(nil).ModelExists), arg0)
	return &MockBackendModelExistsCall{Call: call}
}

// MockBackendModelExistsCall wrap *gomock.Call
type MockBackendModelExistsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendModelExistsCall) Return(arg0 bool, arg1 error) *MockBackendModelExistsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendModelExistsCall) Do(f func(string) (bool, error)) *MockBackendModelExistsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendModelExistsCall) DoAndReturn(f func(string) (bool, error)) *MockBackendModelExistsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MongoVersion mocks base method.
func (m *MockBackend) MongoVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoVersion indicates an expected call of MongoVersion.
func (mr *MockBackendMockRecorder) MongoVersion() *MockBackendMongoVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoVersion", reflect.TypeOf((*MockBackend)(nil).MongoVersion))
	return &MockBackendMongoVersionCall{Call: call}
}

// MockBackendMongoVersionCall wrap *gomock.Call
type MockBackendMongoVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendMongoVersionCall) Return(arg0 string, arg1 error) *MockBackendMongoVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendMongoVersionCall) Do(f func() (string, error)) *MockBackendMongoVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendMongoVersionCall) DoAndReturn(f func() (string, error)) *MockBackendMongoVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// CharmConfig mocks base method.
func (m *MockApplication) CharmConfig() (charm.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmConfig")
	ret0, _ := ret[0].(charm.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmConfig indicates an expected call of CharmConfig.
func (mr *MockApplicationMockRecorder) CharmConfig() *MockApplicationCharmConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmConfig", reflect.TypeOf((*MockApplication)(nil).CharmConfig))
	return &MockApplicationCharmConfigCall{Call: call}
}

// MockApplicationCharmConfigCall wrap *gomock.Call
type MockApplicationCharmConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationCharmConfigCall) Return(arg0 charm.Settings, arg1 error) *MockApplicationCharmConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationCharmConfigCall) Do(f func() (charm.Settings, error)) *MockApplicationCharmConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationCharmConfigCall) DoAndReturn(f func() (charm.Settings, error)) *MockApplicationCharmConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockApplication) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockApplicationMockRecorder) Name() *MockApplicationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockApplication)(nil).Name))
	return &MockApplicationNameCall{Call: call}
}

// MockApplicationNameCall wrap *gomock.Call
type MockApplicationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationNameCall) Return(arg0 string) *MockApplicationNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationNameCall) Do(f func() string) *MockApplicationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationNameCall) DoAndReturn(f func() string) *MockApplicationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Relations mocks base method.
func (m *MockApplication) Relations() ([]controller.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relations")
	ret0, _ := ret[0].([]controller.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relations indicates an expected call of Relations.
func (mr *MockApplicationMockRecorder) Relations() *MockApplicationRelationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relations", reflect.TypeOf((*MockApplication)(nil).Relations))
	return &MockApplicationRelationsCall{Call: call}
}

// MockApplicationRelationsCall wrap *gomock.Call
type MockApplicationRelationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationRelationsCall) Return(arg0 []controller.Relation, arg1 error) *MockApplicationRelationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationRelationsCall) Do(f func() ([]controller.Relation, error)) *MockApplicationRelationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationRelationsCall) DoAndReturn(f func() ([]controller.Relation, error)) *MockApplicationRelationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockRelation is a mock of Relation interface.
type MockRelation struct {
	ctrl     *gomock.Controller
	recorder *MockRelationMockRecorder
}

// MockRelationMockRecorder is the mock recorder for MockRelation.
type MockRelationMockRecorder struct {
	mock *MockRelation
}

// NewMockRelation creates a new mock instance.
func NewMockRelation(ctrl *gomock.Controller) *MockRelation {
	mock := &MockRelation{ctrl: ctrl}
	mock.recorder = &MockRelationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelation) EXPECT() *MockRelationMockRecorder {
	return m.recorder
}

// ApplicationSettings mocks base method.
func (m *MockRelation) ApplicationSettings(arg0 string) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationSettings", arg0)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationSettings indicates an expected call of ApplicationSettings.
func (mr *MockRelationMockRecorder) ApplicationSettings(arg0 any) *MockRelationApplicationSettingsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationSettings", reflect.TypeOf((*MockRelation)(nil).ApplicationSettings), arg0)
	return &MockRelationApplicationSettingsCall{Call: call}
}

// MockRelationApplicationSettingsCall wrap *gomock.Call
type MockRelationApplicationSettingsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationApplicationSettingsCall) Return(arg0 map[string]any, arg1 error) *MockRelationApplicationSettingsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationApplicationSettingsCall) Do(f func(string) (map[string]any, error)) *MockRelationApplicationSettingsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationApplicationSettingsCall) DoAndReturn(f func(string) (map[string]any, error)) *MockRelationApplicationSettingsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Endpoint mocks base method.
func (m *MockRelation) Endpoint(arg0 string) (state.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint", arg0)
	ret0, _ := ret[0].(state.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Endpoint indicates an expected call of Endpoint.
func (mr *MockRelationMockRecorder) Endpoint(arg0 any) *MockRelationEndpointCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockRelation)(nil).Endpoint), arg0)
	return &MockRelationEndpointCall{Call: call}
}

// MockRelationEndpointCall wrap *gomock.Call
type MockRelationEndpointCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationEndpointCall) Return(arg0 state.Endpoint, arg1 error) *MockRelationEndpointCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationEndpointCall) Do(f func(string) (state.Endpoint, error)) *MockRelationEndpointCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationEndpointCall) DoAndReturn(f func(string) (state.Endpoint, error)) *MockRelationEndpointCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelUUID mocks base method.
func (m *MockRelation) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockRelationMockRecorder) ModelUUID() *MockRelationModelUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockRelation)(nil).ModelUUID))
	return &MockRelationModelUUIDCall{Call: call}
}

// MockRelationModelUUIDCall wrap *gomock.Call
type MockRelationModelUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationModelUUIDCall) Return(arg0 string) *MockRelationModelUUIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationModelUUIDCall) Do(f func() string) *MockRelationModelUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationModelUUIDCall) DoAndReturn(f func() string) *MockRelationModelUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RelatedEndpoints mocks base method.
func (m *MockRelation) RelatedEndpoints(arg0 string) ([]state.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelatedEndpoints", arg0)
	ret0, _ := ret[0].([]state.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelatedEndpoints indicates an expected call of RelatedEndpoints.
func (mr *MockRelationMockRecorder) RelatedEndpoints(arg0 any) *MockRelationRelatedEndpointsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelatedEndpoints", reflect.TypeOf((*MockRelation)(nil).RelatedEndpoints), arg0)
	return &MockRelationRelatedEndpointsCall{Call: call}
}

// MockRelationRelatedEndpointsCall wrap *gomock.Call
type MockRelationRelatedEndpointsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationRelatedEndpointsCall) Return(arg0 []state.Endpoint, arg1 error) *MockRelationRelatedEndpointsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationRelatedEndpointsCall) Do(f func(string) ([]state.Endpoint, error)) *MockRelationRelatedEndpointsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationRelatedEndpointsCall) DoAndReturn(f func(string) ([]state.Endpoint, error)) *MockRelationRelatedEndpointsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
