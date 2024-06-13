// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/machine/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/machine/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

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

// DeleteMachine mocks base method.
func (m *MockState) DeleteMachine(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMachine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMachine indicates an expected call of DeleteMachine.
func (mr *MockStateMockRecorder) DeleteMachine(arg0, arg1 any) *MockStateDeleteMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMachine", reflect.TypeOf((*MockState)(nil).DeleteMachine), arg0, arg1)
	return &MockStateDeleteMachineCall{Call: call}
}

// MockStateDeleteMachineCall wrap *gomock.Call
type MockStateDeleteMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateDeleteMachineCall) Return(arg0 error) *MockStateDeleteMachineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateDeleteMachineCall) Do(f func(context.Context, string) error) *MockStateDeleteMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateDeleteMachineCall) DoAndReturn(f func(context.Context, string) error) *MockStateDeleteMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InitialWatchStatement mocks base method.
func (m *MockState) InitialWatchStatement() (string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialWatchStatement")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// InitialWatchStatement indicates an expected call of InitialWatchStatement.
func (mr *MockStateMockRecorder) InitialWatchStatement() *MockStateInitialWatchStatementCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialWatchStatement", reflect.TypeOf((*MockState)(nil).InitialWatchStatement))
	return &MockStateInitialWatchStatementCall{Call: call}
}

// MockStateInitialWatchStatementCall wrap *gomock.Call
type MockStateInitialWatchStatementCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateInitialWatchStatementCall) Return(arg0, arg1 string) *MockStateInitialWatchStatementCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateInitialWatchStatementCall) Do(f func() (string, string)) *MockStateInitialWatchStatementCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateInitialWatchStatementCall) DoAndReturn(f func() (string, string)) *MockStateInitialWatchStatementCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpsertMachine mocks base method.
func (m *MockState) UpsertMachine(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMachine", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertMachine indicates an expected call of UpsertMachine.
func (mr *MockStateMockRecorder) UpsertMachine(arg0, arg1 any) *MockStateUpsertMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMachine", reflect.TypeOf((*MockState)(nil).UpsertMachine), arg0, arg1)
	return &MockStateUpsertMachineCall{Call: call}
}

// MockStateUpsertMachineCall wrap *gomock.Call
type MockStateUpsertMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpsertMachineCall) Return(arg0 string, arg1 error) *MockStateUpsertMachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpsertMachineCall) Do(f func(context.Context, string) (string, error)) *MockStateUpsertMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpsertMachineCall) DoAndReturn(f func(context.Context, string) (string, error)) *MockStateUpsertMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
