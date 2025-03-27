// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/modelupgrader (interfaces: StatePool,State,Model,UpgradeService,ControllerConfigService,ModelAgentService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/state_mock.go github.com/juju/juju/apiserver/facades/client/modelupgrader StatePool,State,Model,UpgradeService,ControllerConfigService,ModelAgentService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	modelupgrader "github.com/juju/juju/apiserver/facades/client/modelupgrader"
	controller "github.com/juju/juju/controller"
	semversion "github.com/juju/juju/core/semversion"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v6"
	replicaset "github.com/juju/replicaset/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockStatePool is a mock of StatePool interface.
type MockStatePool struct {
	ctrl     *gomock.Controller
	recorder *MockStatePoolMockRecorder
}

// MockStatePoolMockRecorder is the mock recorder for MockStatePool.
type MockStatePoolMockRecorder struct {
	mock *MockStatePool
}

// NewMockStatePool creates a new mock instance.
func NewMockStatePool(ctrl *gomock.Controller) *MockStatePool {
	mock := &MockStatePool{ctrl: ctrl}
	mock.recorder = &MockStatePoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatePool) EXPECT() *MockStatePoolMockRecorder {
	return m.recorder
}

// ControllerModel mocks base method.
func (m *MockStatePool) ControllerModel() (modelupgrader.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerModel")
	ret0, _ := ret[0].(modelupgrader.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerModel indicates an expected call of ControllerModel.
func (mr *MockStatePoolMockRecorder) ControllerModel() *MockStatePoolControllerModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerModel", reflect.TypeOf((*MockStatePool)(nil).ControllerModel))
	return &MockStatePoolControllerModelCall{Call: call}
}

// MockStatePoolControllerModelCall wrap *gomock.Call
type MockStatePoolControllerModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStatePoolControllerModelCall) Return(arg0 modelupgrader.Model, arg1 error) *MockStatePoolControllerModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStatePoolControllerModelCall) Do(f func() (modelupgrader.Model, error)) *MockStatePoolControllerModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStatePoolControllerModelCall) DoAndReturn(f func() (modelupgrader.Model, error)) *MockStatePoolControllerModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockStatePool) Get(arg0 string) (modelupgrader.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(modelupgrader.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStatePoolMockRecorder) Get(arg0 any) *MockStatePoolGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStatePool)(nil).Get), arg0)
	return &MockStatePoolGetCall{Call: call}
}

// MockStatePoolGetCall wrap *gomock.Call
type MockStatePoolGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStatePoolGetCall) Return(arg0 modelupgrader.State, arg1 error) *MockStatePoolGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStatePoolGetCall) Do(f func(string) (modelupgrader.State, error)) *MockStatePoolGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStatePoolGetCall) DoAndReturn(f func(string) (modelupgrader.State, error)) *MockStatePoolGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MongoVersion mocks base method.
func (m *MockStatePool) MongoVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoVersion indicates an expected call of MongoVersion.
func (mr *MockStatePoolMockRecorder) MongoVersion() *MockStatePoolMongoVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoVersion", reflect.TypeOf((*MockStatePool)(nil).MongoVersion))
	return &MockStatePoolMongoVersionCall{Call: call}
}

// MockStatePoolMongoVersionCall wrap *gomock.Call
type MockStatePoolMongoVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStatePoolMongoVersionCall) Return(arg0 string, arg1 error) *MockStatePoolMongoVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStatePoolMongoVersionCall) Do(f func() (string, error)) *MockStatePoolMongoVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStatePoolMongoVersionCall) DoAndReturn(f func() (string, error)) *MockStatePoolMongoVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

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

// AllMachinesCount mocks base method.
func (m *MockState) AllMachinesCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachinesCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachinesCount indicates an expected call of AllMachinesCount.
func (mr *MockStateMockRecorder) AllMachinesCount() *MockStateAllMachinesCountCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachinesCount", reflect.TypeOf((*MockState)(nil).AllMachinesCount))
	return &MockStateAllMachinesCountCall{Call: call}
}

// MockStateAllMachinesCountCall wrap *gomock.Call
type MockStateAllMachinesCountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateAllMachinesCountCall) Return(arg0 int, arg1 error) *MockStateAllMachinesCountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateAllMachinesCountCall) Do(f func() (int, error)) *MockStateAllMachinesCountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateAllMachinesCountCall) DoAndReturn(f func() (int, error)) *MockStateAllMachinesCountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllModelUUIDs mocks base method.
func (m *MockState) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs.
func (mr *MockStateMockRecorder) AllModelUUIDs() *MockStateAllModelUUIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockState)(nil).AllModelUUIDs))
	return &MockStateAllModelUUIDsCall{Call: call}
}

// MockStateAllModelUUIDsCall wrap *gomock.Call
type MockStateAllModelUUIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateAllModelUUIDsCall) Return(arg0 []string, arg1 error) *MockStateAllModelUUIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateAllModelUUIDsCall) Do(f func() ([]string, error)) *MockStateAllModelUUIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateAllModelUUIDsCall) DoAndReturn(f func() ([]string, error)) *MockStateAllModelUUIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MachineCountForBase mocks base method.
func (m *MockState) MachineCountForBase(arg0 ...state.Base) (map[string]int, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MachineCountForBase", varargs...)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineCountForBase indicates an expected call of MachineCountForBase.
func (mr *MockStateMockRecorder) MachineCountForBase(arg0 ...any) *MockStateMachineCountForBaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineCountForBase", reflect.TypeOf((*MockState)(nil).MachineCountForBase), arg0...)
	return &MockStateMachineCountForBaseCall{Call: call}
}

// MockStateMachineCountForBaseCall wrap *gomock.Call
type MockStateMachineCountForBaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateMachineCountForBaseCall) Return(arg0 map[string]int, arg1 error) *MockStateMachineCountForBaseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateMachineCountForBaseCall) Do(f func(...state.Base) (map[string]int, error)) *MockStateMachineCountForBaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateMachineCountForBaseCall) DoAndReturn(f func(...state.Base) (map[string]int, error)) *MockStateMachineCountForBaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockState) Model() (modelupgrader.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(modelupgrader.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateMockRecorder) Model() *MockStateModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
	return &MockStateModelCall{Call: call}
}

// MockStateModelCall wrap *gomock.Call
type MockStateModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateModelCall) Return(arg0 modelupgrader.Model, arg1 error) *MockStateModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateModelCall) Do(f func() (modelupgrader.Model, error)) *MockStateModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateModelCall) DoAndReturn(f func() (modelupgrader.Model, error)) *MockStateModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MongoCurrentStatus mocks base method.
func (m *MockState) MongoCurrentStatus() (*replicaset.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoCurrentStatus")
	ret0, _ := ret[0].(*replicaset.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoCurrentStatus indicates an expected call of MongoCurrentStatus.
func (mr *MockStateMockRecorder) MongoCurrentStatus() *MockStateMongoCurrentStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoCurrentStatus", reflect.TypeOf((*MockState)(nil).MongoCurrentStatus))
	return &MockStateMongoCurrentStatusCall{Call: call}
}

// MockStateMongoCurrentStatusCall wrap *gomock.Call
type MockStateMongoCurrentStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateMongoCurrentStatusCall) Return(arg0 *replicaset.Status, arg1 error) *MockStateMongoCurrentStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateMongoCurrentStatusCall) Do(f func() (*replicaset.Status, error)) *MockStateMongoCurrentStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateMongoCurrentStatusCall) DoAndReturn(f func() (*replicaset.Status, error)) *MockStateMongoCurrentStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Release mocks base method.
func (m *MockState) Release() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Release indicates an expected call of Release.
func (mr *MockStateMockRecorder) Release() *MockStateReleaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockState)(nil).Release))
	return &MockStateReleaseCall{Call: call}
}

// MockStateReleaseCall wrap *gomock.Call
type MockStateReleaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateReleaseCall) Return(arg0 bool) *MockStateReleaseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateReleaseCall) Do(f func() bool) *MockStateReleaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateReleaseCall) DoAndReturn(f func() bool) *MockStateReleaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetModelAgentVersion mocks base method.
func (m *MockState) SetModelAgentVersion(arg0 semversion.Number, arg1 *string, arg2 bool, arg3 state.Upgrader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModelAgentVersion", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModelAgentVersion indicates an expected call of SetModelAgentVersion.
func (mr *MockStateMockRecorder) SetModelAgentVersion(arg0, arg1, arg2, arg3 any) *MockStateSetModelAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModelAgentVersion", reflect.TypeOf((*MockState)(nil).SetModelAgentVersion), arg0, arg1, arg2, arg3)
	return &MockStateSetModelAgentVersionCall{Call: call}
}

// MockStateSetModelAgentVersionCall wrap *gomock.Call
type MockStateSetModelAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetModelAgentVersionCall) Return(arg0 error) *MockStateSetModelAgentVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetModelAgentVersionCall) Do(f func(semversion.Number, *string, bool, state.Upgrader) error) *MockStateSetModelAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetModelAgentVersionCall) DoAndReturn(f func(semversion.Number, *string, bool, state.Upgrader) error) *MockStateSetModelAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// IsControllerModel mocks base method.
func (m *MockModel) IsControllerModel() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsControllerModel")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsControllerModel indicates an expected call of IsControllerModel.
func (mr *MockModelMockRecorder) IsControllerModel() *MockModelIsControllerModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsControllerModel", reflect.TypeOf((*MockModel)(nil).IsControllerModel))
	return &MockModelIsControllerModelCall{Call: call}
}

// MockModelIsControllerModelCall wrap *gomock.Call
type MockModelIsControllerModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelIsControllerModelCall) Return(arg0 bool) *MockModelIsControllerModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelIsControllerModelCall) Do(f func() bool) *MockModelIsControllerModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelIsControllerModelCall) DoAndReturn(f func() bool) *MockModelIsControllerModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Life mocks base method.
func (m *MockModel) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockModelMockRecorder) Life() *MockModelLifeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockModel)(nil).Life))
	return &MockModelLifeCall{Call: call}
}

// MockModelLifeCall wrap *gomock.Call
type MockModelLifeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelLifeCall) Return(arg0 state.Life) *MockModelLifeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelLifeCall) Do(f func() state.Life) *MockModelLifeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelLifeCall) DoAndReturn(f func() state.Life) *MockModelLifeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MigrationMode mocks base method.
func (m *MockModel) MigrationMode() state.MigrationMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationMode")
	ret0, _ := ret[0].(state.MigrationMode)
	return ret0
}

// MigrationMode indicates an expected call of MigrationMode.
func (mr *MockModelMockRecorder) MigrationMode() *MockModelMigrationModeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationMode", reflect.TypeOf((*MockModel)(nil).MigrationMode))
	return &MockModelMigrationModeCall{Call: call}
}

// MockModelMigrationModeCall wrap *gomock.Call
type MockModelMigrationModeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationModeCall) Return(arg0 state.MigrationMode) *MockModelMigrationModeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationModeCall) Do(f func() state.MigrationMode) *MockModelMigrationModeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationModeCall) DoAndReturn(f func() state.MigrationMode) *MockModelMigrationModeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelMockRecorder) Name() *MockModelNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModel)(nil).Name))
	return &MockModelNameCall{Call: call}
}

// MockModelNameCall wrap *gomock.Call
type MockModelNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelNameCall) Return(arg0 string) *MockModelNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelNameCall) Do(f func() string) *MockModelNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelNameCall) DoAndReturn(f func() string) *MockModelNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Owner mocks base method.
func (m *MockModel) Owner() names.UserTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(names.UserTag)
	return ret0
}

// Owner indicates an expected call of Owner.
func (mr *MockModelMockRecorder) Owner() *MockModelOwnerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockModel)(nil).Owner))
	return &MockModelOwnerCall{Call: call}
}

// MockModelOwnerCall wrap *gomock.Call
type MockModelOwnerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelOwnerCall) Return(arg0 names.UserTag) *MockModelOwnerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelOwnerCall) Do(f func() names.UserTag) *MockModelOwnerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelOwnerCall) DoAndReturn(f func() names.UserTag) *MockModelOwnerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *MockModelTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
	return &MockModelTypeCall{Call: call}
}

// MockModelTypeCall wrap *gomock.Call
type MockModelTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelTypeCall) Return(arg0 state.ModelType) *MockModelTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelTypeCall) Do(f func() state.ModelType) *MockModelTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelTypeCall) DoAndReturn(f func() state.ModelType) *MockModelTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUpgradeService is a mock of UpgradeService interface.
type MockUpgradeService struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeServiceMockRecorder
}

// MockUpgradeServiceMockRecorder is the mock recorder for MockUpgradeService.
type MockUpgradeServiceMockRecorder struct {
	mock *MockUpgradeService
}

// NewMockUpgradeService creates a new mock instance.
func NewMockUpgradeService(ctrl *gomock.Controller) *MockUpgradeService {
	mock := &MockUpgradeService{ctrl: ctrl}
	mock.recorder = &MockUpgradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeService) EXPECT() *MockUpgradeServiceMockRecorder {
	return m.recorder
}

// IsUpgrading mocks base method.
func (m *MockUpgradeService) IsUpgrading(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpgrading", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpgrading indicates an expected call of IsUpgrading.
func (mr *MockUpgradeServiceMockRecorder) IsUpgrading(arg0 any) *MockUpgradeServiceIsUpgradingCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpgrading", reflect.TypeOf((*MockUpgradeService)(nil).IsUpgrading), arg0)
	return &MockUpgradeServiceIsUpgradingCall{Call: call}
}

// MockUpgradeServiceIsUpgradingCall wrap *gomock.Call
type MockUpgradeServiceIsUpgradingCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceIsUpgradingCall) Return(arg0 bool, arg1 error) *MockUpgradeServiceIsUpgradingCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceIsUpgradingCall) Do(f func(context.Context) (bool, error)) *MockUpgradeServiceIsUpgradingCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceIsUpgradingCall) DoAndReturn(f func(context.Context) (bool, error)) *MockUpgradeServiceIsUpgradingCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *MockControllerConfigServiceControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
	return &MockControllerConfigServiceControllerConfigCall{Call: call}
}

// MockControllerConfigServiceControllerConfigCall wrap *gomock.Call
type MockControllerConfigServiceControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerConfigServiceControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerConfigServiceControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerConfigServiceControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelAgentService is a mock of ModelAgentService interface.
type MockModelAgentService struct {
	ctrl     *gomock.Controller
	recorder *MockModelAgentServiceMockRecorder
}

// MockModelAgentServiceMockRecorder is the mock recorder for MockModelAgentService.
type MockModelAgentServiceMockRecorder struct {
	mock *MockModelAgentService
}

// NewMockModelAgentService creates a new mock instance.
func NewMockModelAgentService(ctrl *gomock.Controller) *MockModelAgentService {
	mock := &MockModelAgentService{ctrl: ctrl}
	mock.recorder = &MockModelAgentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelAgentService) EXPECT() *MockModelAgentServiceMockRecorder {
	return m.recorder
}

// GetModelTargetAgentVersion mocks base method.
func (m *MockModelAgentService) GetModelTargetAgentVersion(arg0 context.Context) (semversion.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelTargetAgentVersion", arg0)
	ret0, _ := ret[0].(semversion.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelTargetAgentVersion indicates an expected call of GetModelTargetAgentVersion.
func (mr *MockModelAgentServiceMockRecorder) GetModelTargetAgentVersion(arg0 any) *MockModelAgentServiceGetModelTargetAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelTargetAgentVersion", reflect.TypeOf((*MockModelAgentService)(nil).GetModelTargetAgentVersion), arg0)
	return &MockModelAgentServiceGetModelTargetAgentVersionCall{Call: call}
}

// MockModelAgentServiceGetModelTargetAgentVersionCall wrap *gomock.Call
type MockModelAgentServiceGetModelTargetAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelAgentServiceGetModelTargetAgentVersionCall) Return(arg0 semversion.Number, arg1 error) *MockModelAgentServiceGetModelTargetAgentVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelAgentServiceGetModelTargetAgentVersionCall) Do(f func(context.Context) (semversion.Number, error)) *MockModelAgentServiceGetModelTargetAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelAgentServiceGetModelTargetAgentVersionCall) DoAndReturn(f func(context.Context) (semversion.Number, error)) *MockModelAgentServiceGetModelTargetAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
