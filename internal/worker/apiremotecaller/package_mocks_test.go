// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/apiremotecaller (interfaces: RemoteServer)
//
// Generated by this command:
//
//	mockgen -typed -package apiremotecaller -destination package_mocks_test.go github.com/juju/juju/internal/worker/apiremotecaller RemoteServer
//

// Package apiremotecaller is a generated GoMock package.
package apiremotecaller

import (
	context "context"
	reflect "reflect"

	api "github.com/juju/juju/api"
	gomock "go.uber.org/mock/gomock"
)

// MockRemoteServer is a mock of RemoteServer interface.
type MockRemoteServer struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteServerMockRecorder
}

// MockRemoteServerMockRecorder is the mock recorder for MockRemoteServer.
type MockRemoteServerMockRecorder struct {
	mock *MockRemoteServer
}

// NewMockRemoteServer creates a new mock instance.
func NewMockRemoteServer(ctrl *gomock.Controller) *MockRemoteServer {
	mock := &MockRemoteServer{ctrl: ctrl}
	mock.recorder = &MockRemoteServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteServer) EXPECT() *MockRemoteServerMockRecorder {
	return m.recorder
}

// Connection mocks base method.
func (m *MockRemoteServer) Connection(arg0 context.Context) <-chan api.Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connection", arg0)
	ret0, _ := ret[0].(<-chan api.Connection)
	return ret0
}

// Connection indicates an expected call of Connection.
func (mr *MockRemoteServerMockRecorder) Connection(arg0 any) *MockRemoteServerConnectionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connection", reflect.TypeOf((*MockRemoteServer)(nil).Connection), arg0)
	return &MockRemoteServerConnectionCall{Call: call}
}

// MockRemoteServerConnectionCall wrap *gomock.Call
type MockRemoteServerConnectionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRemoteServerConnectionCall) Return(arg0 <-chan api.Connection) *MockRemoteServerConnectionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRemoteServerConnectionCall) Do(f func(context.Context) <-chan api.Connection) *MockRemoteServerConnectionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRemoteServerConnectionCall) DoAndReturn(f func(context.Context) <-chan api.Connection) *MockRemoteServerConnectionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Kill mocks base method.
func (m *MockRemoteServer) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockRemoteServerMockRecorder) Kill() *MockRemoteServerKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockRemoteServer)(nil).Kill))
	return &MockRemoteServerKillCall{Call: call}
}

// MockRemoteServerKillCall wrap *gomock.Call
type MockRemoteServerKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRemoteServerKillCall) Return() *MockRemoteServerKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRemoteServerKillCall) Do(f func()) *MockRemoteServerKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRemoteServerKillCall) DoAndReturn(f func()) *MockRemoteServerKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateAddresses mocks base method.
func (m *MockRemoteServer) UpdateAddresses(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAddresses", arg0)
}

// UpdateAddresses indicates an expected call of UpdateAddresses.
func (mr *MockRemoteServerMockRecorder) UpdateAddresses(arg0 any) *MockRemoteServerUpdateAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddresses", reflect.TypeOf((*MockRemoteServer)(nil).UpdateAddresses), arg0)
	return &MockRemoteServerUpdateAddressesCall{Call: call}
}

// MockRemoteServerUpdateAddressesCall wrap *gomock.Call
type MockRemoteServerUpdateAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRemoteServerUpdateAddressesCall) Return() *MockRemoteServerUpdateAddressesCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRemoteServerUpdateAddressesCall) Do(f func([]string)) *MockRemoteServerUpdateAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRemoteServerUpdateAddressesCall) DoAndReturn(f func([]string)) *MockRemoteServerUpdateAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockRemoteServer) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockRemoteServerMockRecorder) Wait() *MockRemoteServerWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockRemoteServer)(nil).Wait))
	return &MockRemoteServerWaitCall{Call: call}
}

// MockRemoteServerWaitCall wrap *gomock.Call
type MockRemoteServerWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRemoteServerWaitCall) Return(arg0 error) *MockRemoteServerWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRemoteServerWaitCall) Do(f func() error) *MockRemoteServerWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRemoteServerWaitCall) DoAndReturn(f func() error) *MockRemoteServerWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
