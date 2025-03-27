// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/upgrader (interfaces: UpgraderClient)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/upgrader_mocks.go github.com/juju/juju/internal/worker/upgrader UpgraderClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	semversion "github.com/juju/juju/core/semversion"
	watcher "github.com/juju/juju/core/watcher"
	tools "github.com/juju/juju/internal/tools"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgraderClient is a mock of UpgraderClient interface.
type MockUpgraderClient struct {
	ctrl     *gomock.Controller
	recorder *MockUpgraderClientMockRecorder
}

// MockUpgraderClientMockRecorder is the mock recorder for MockUpgraderClient.
type MockUpgraderClientMockRecorder struct {
	mock *MockUpgraderClient
}

// NewMockUpgraderClient creates a new mock instance.
func NewMockUpgraderClient(ctrl *gomock.Controller) *MockUpgraderClient {
	mock := &MockUpgraderClient{ctrl: ctrl}
	mock.recorder = &MockUpgraderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgraderClient) EXPECT() *MockUpgraderClientMockRecorder {
	return m.recorder
}

// DesiredVersion mocks base method.
func (m *MockUpgraderClient) DesiredVersion(arg0 context.Context, arg1 string) (semversion.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DesiredVersion", arg0, arg1)
	ret0, _ := ret[0].(semversion.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DesiredVersion indicates an expected call of DesiredVersion.
func (mr *MockUpgraderClientMockRecorder) DesiredVersion(arg0, arg1 any) *MockUpgraderClientDesiredVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DesiredVersion", reflect.TypeOf((*MockUpgraderClient)(nil).DesiredVersion), arg0, arg1)
	return &MockUpgraderClientDesiredVersionCall{Call: call}
}

// MockUpgraderClientDesiredVersionCall wrap *gomock.Call
type MockUpgraderClientDesiredVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgraderClientDesiredVersionCall) Return(arg0 semversion.Number, arg1 error) *MockUpgraderClientDesiredVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgraderClientDesiredVersionCall) Do(f func(context.Context, string) (semversion.Number, error)) *MockUpgraderClientDesiredVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgraderClientDesiredVersionCall) DoAndReturn(f func(context.Context, string) (semversion.Number, error)) *MockUpgraderClientDesiredVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetVersion mocks base method.
func (m *MockUpgraderClient) SetVersion(arg0 context.Context, arg1 string, arg2 semversion.Binary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVersion indicates an expected call of SetVersion.
func (mr *MockUpgraderClientMockRecorder) SetVersion(arg0, arg1, arg2 any) *MockUpgraderClientSetVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersion", reflect.TypeOf((*MockUpgraderClient)(nil).SetVersion), arg0, arg1, arg2)
	return &MockUpgraderClientSetVersionCall{Call: call}
}

// MockUpgraderClientSetVersionCall wrap *gomock.Call
type MockUpgraderClientSetVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgraderClientSetVersionCall) Return(arg0 error) *MockUpgraderClientSetVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgraderClientSetVersionCall) Do(f func(context.Context, string, semversion.Binary) error) *MockUpgraderClientSetVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgraderClientSetVersionCall) DoAndReturn(f func(context.Context, string, semversion.Binary) error) *MockUpgraderClientSetVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tools mocks base method.
func (m *MockUpgraderClient) Tools(arg0 context.Context, arg1 string) (tools.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tools", arg0, arg1)
	ret0, _ := ret[0].(tools.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tools indicates an expected call of Tools.
func (mr *MockUpgraderClientMockRecorder) Tools(arg0, arg1 any) *MockUpgraderClientToolsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tools", reflect.TypeOf((*MockUpgraderClient)(nil).Tools), arg0, arg1)
	return &MockUpgraderClientToolsCall{Call: call}
}

// MockUpgraderClientToolsCall wrap *gomock.Call
type MockUpgraderClientToolsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgraderClientToolsCall) Return(arg0 tools.List, arg1 error) *MockUpgraderClientToolsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgraderClientToolsCall) Do(f func(context.Context, string) (tools.List, error)) *MockUpgraderClientToolsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgraderClientToolsCall) DoAndReturn(f func(context.Context, string) (tools.List, error)) *MockUpgraderClientToolsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchAPIVersion mocks base method.
func (m *MockUpgraderClient) WatchAPIVersion(arg0 context.Context, arg1 string) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAPIVersion", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAPIVersion indicates an expected call of WatchAPIVersion.
func (mr *MockUpgraderClientMockRecorder) WatchAPIVersion(arg0, arg1 any) *MockUpgraderClientWatchAPIVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAPIVersion", reflect.TypeOf((*MockUpgraderClient)(nil).WatchAPIVersion), arg0, arg1)
	return &MockUpgraderClientWatchAPIVersionCall{Call: call}
}

// MockUpgraderClientWatchAPIVersionCall wrap *gomock.Call
type MockUpgraderClientWatchAPIVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgraderClientWatchAPIVersionCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockUpgraderClientWatchAPIVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgraderClientWatchAPIVersionCall) Do(f func(context.Context, string) (watcher.Watcher[struct{}], error)) *MockUpgraderClientWatchAPIVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgraderClientWatchAPIVersionCall) DoAndReturn(f func(context.Context, string) (watcher.Watcher[struct{}], error)) *MockUpgraderClientWatchAPIVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
