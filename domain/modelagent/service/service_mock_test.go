// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/modelagent/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination service_mock_test.go github.com/juju/juju/domain/modelagent/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	machine "github.com/juju/juju/core/machine"
	semversion "github.com/juju/juju/core/semversion"
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

// CheckMachineExists mocks base method.
func (m *MockState) CheckMachineExists(arg0 context.Context, arg1 machine.Name) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMachineExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckMachineExists indicates an expected call of CheckMachineExists.
func (mr *MockStateMockRecorder) CheckMachineExists(arg0, arg1 any) *MockStateCheckMachineExistsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMachineExists", reflect.TypeOf((*MockState)(nil).CheckMachineExists), arg0, arg1)
	return &MockStateCheckMachineExistsCall{Call: call}
}

// MockStateCheckMachineExistsCall wrap *gomock.Call
type MockStateCheckMachineExistsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateCheckMachineExistsCall) Return(arg0 error) *MockStateCheckMachineExistsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateCheckMachineExistsCall) Do(f func(context.Context, machine.Name) error) *MockStateCheckMachineExistsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateCheckMachineExistsCall) DoAndReturn(f func(context.Context, machine.Name) error) *MockStateCheckMachineExistsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CheckUnitExists mocks base method.
func (m *MockState) CheckUnitExists(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUnitExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUnitExists indicates an expected call of CheckUnitExists.
func (mr *MockStateMockRecorder) CheckUnitExists(arg0, arg1 any) *MockStateCheckUnitExistsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUnitExists", reflect.TypeOf((*MockState)(nil).CheckUnitExists), arg0, arg1)
	return &MockStateCheckUnitExistsCall{Call: call}
}

// MockStateCheckUnitExistsCall wrap *gomock.Call
type MockStateCheckUnitExistsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateCheckUnitExistsCall) Return(arg0 error) *MockStateCheckUnitExistsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateCheckUnitExistsCall) Do(f func(context.Context, string) error) *MockStateCheckUnitExistsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateCheckUnitExistsCall) DoAndReturn(f func(context.Context, string) error) *MockStateCheckUnitExistsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTargetAgentVersion mocks base method.
func (m *MockState) GetTargetAgentVersion(arg0 context.Context) (semversion.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetAgentVersion", arg0)
	ret0, _ := ret[0].(semversion.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetAgentVersion indicates an expected call of GetTargetAgentVersion.
func (mr *MockStateMockRecorder) GetTargetAgentVersion(arg0 any) *MockStateGetTargetAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetAgentVersion", reflect.TypeOf((*MockState)(nil).GetTargetAgentVersion), arg0)
	return &MockStateGetTargetAgentVersionCall{Call: call}
}

// MockStateGetTargetAgentVersionCall wrap *gomock.Call
type MockStateGetTargetAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetTargetAgentVersionCall) Return(arg0 semversion.Number, arg1 error) *MockStateGetTargetAgentVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetTargetAgentVersionCall) Do(f func(context.Context) (semversion.Number, error)) *MockStateGetTargetAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetTargetAgentVersionCall) DoAndReturn(f func(context.Context) (semversion.Number, error)) *MockStateGetTargetAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NamespaceForWatchAgentVersion mocks base method.
func (m *MockState) NamespaceForWatchAgentVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamespaceForWatchAgentVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// NamespaceForWatchAgentVersion indicates an expected call of NamespaceForWatchAgentVersion.
func (mr *MockStateMockRecorder) NamespaceForWatchAgentVersion() *MockStateNamespaceForWatchAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamespaceForWatchAgentVersion", reflect.TypeOf((*MockState)(nil).NamespaceForWatchAgentVersion))
	return &MockStateNamespaceForWatchAgentVersionCall{Call: call}
}

// MockStateNamespaceForWatchAgentVersionCall wrap *gomock.Call
type MockStateNamespaceForWatchAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateNamespaceForWatchAgentVersionCall) Return(arg0 string) *MockStateNamespaceForWatchAgentVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateNamespaceForWatchAgentVersionCall) Do(f func() string) *MockStateNamespaceForWatchAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateNamespaceForWatchAgentVersionCall) DoAndReturn(f func() string) *MockStateNamespaceForWatchAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
