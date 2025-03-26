// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api (interfaces: Connection)
//
// Generated by this command:
//
//	mockgen -typed -package remote -destination connection_mock_test.go github.com/juju/juju/api Connection
//

// Package remote is a generated GoMock package.
package remote

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	base "github.com/juju/juju/api/base"
	network "github.com/juju/juju/core/network"
	proxy "github.com/juju/juju/internal/proxy"
	version "github.com/juju/juju/internal/version"
	names "github.com/juju/names/v6"
	gomock "go.uber.org/mock/gomock"
	httprequest "gopkg.in/httprequest.v1"
	macaroon "gopkg.in/macaroon.v2"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// APICall mocks base method.
func (m *MockConnection) APICall(arg0 context.Context, arg1 string, arg2 int, arg3, arg4 string, arg5, arg6 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall.
func (mr *MockConnectionMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *MockConnectionAPICallCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockConnection)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	return &MockConnectionAPICallCall{Call: call}
}

// MockConnectionAPICallCall wrap *gomock.Call
type MockConnectionAPICallCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionAPICallCall) Return(arg0 error) *MockConnectionAPICallCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionAPICallCall) Do(f func(context.Context, string, int, string, string, any, any) error) *MockConnectionAPICallCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionAPICallCall) DoAndReturn(f func(context.Context, string, int, string, string, any, any) error) *MockConnectionAPICallCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// APIHostPorts mocks base method.
func (m *MockConnection) APIHostPorts() []network.MachineHostPorts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIHostPorts")
	ret0, _ := ret[0].([]network.MachineHostPorts)
	return ret0
}

// APIHostPorts indicates an expected call of APIHostPorts.
func (mr *MockConnectionMockRecorder) APIHostPorts() *MockConnectionAPIHostPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIHostPorts", reflect.TypeOf((*MockConnection)(nil).APIHostPorts))
	return &MockConnectionAPIHostPortsCall{Call: call}
}

// MockConnectionAPIHostPortsCall wrap *gomock.Call
type MockConnectionAPIHostPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionAPIHostPortsCall) Return(arg0 []network.MachineHostPorts) *MockConnectionAPIHostPortsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionAPIHostPortsCall) Do(f func() []network.MachineHostPorts) *MockConnectionAPIHostPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionAPIHostPortsCall) DoAndReturn(f func() []network.MachineHostPorts) *MockConnectionAPIHostPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Addr mocks base method.
func (m *MockConnection) Addr() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// Addr indicates an expected call of Addr.
func (mr *MockConnectionMockRecorder) Addr() *MockConnectionAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockConnection)(nil).Addr))
	return &MockConnectionAddrCall{Call: call}
}

// MockConnectionAddrCall wrap *gomock.Call
type MockConnectionAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionAddrCall) Return(arg0 *url.URL) *MockConnectionAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionAddrCall) Do(f func() *url.URL) *MockConnectionAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionAddrCall) DoAndReturn(f func() *url.URL) *MockConnectionAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AuthTag mocks base method.
func (m *MockConnection) AuthTag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// AuthTag indicates an expected call of AuthTag.
func (mr *MockConnectionMockRecorder) AuthTag() *MockConnectionAuthTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTag", reflect.TypeOf((*MockConnection)(nil).AuthTag))
	return &MockConnectionAuthTagCall{Call: call}
}

// MockConnectionAuthTagCall wrap *gomock.Call
type MockConnectionAuthTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionAuthTagCall) Return(arg0 names.Tag) *MockConnectionAuthTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionAuthTagCall) Do(f func() names.Tag) *MockConnectionAuthTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionAuthTagCall) DoAndReturn(f func() names.Tag) *MockConnectionAuthTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BakeryClient mocks base method.
func (m *MockConnection) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockConnectionMockRecorder) BakeryClient() *MockConnectionBakeryClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockConnection)(nil).BakeryClient))
	return &MockConnectionBakeryClientCall{Call: call}
}

// MockConnectionBakeryClientCall wrap *gomock.Call
type MockConnectionBakeryClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionBakeryClientCall) Return(arg0 base.MacaroonDischarger) *MockConnectionBakeryClientCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionBakeryClientCall) Do(f func() base.MacaroonDischarger) *MockConnectionBakeryClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionBakeryClientCall) DoAndReturn(f func() base.MacaroonDischarger) *MockConnectionBakeryClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BestFacadeVersion mocks base method.
func (m *MockConnection) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion.
func (mr *MockConnectionMockRecorder) BestFacadeVersion(arg0 any) *MockConnectionBestFacadeVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockConnection)(nil).BestFacadeVersion), arg0)
	return &MockConnectionBestFacadeVersionCall{Call: call}
}

// MockConnectionBestFacadeVersionCall wrap *gomock.Call
type MockConnectionBestFacadeVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionBestFacadeVersionCall) Return(arg0 int) *MockConnectionBestFacadeVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionBestFacadeVersionCall) Do(f func(string) int) *MockConnectionBestFacadeVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionBestFacadeVersionCall) DoAndReturn(f func(string) int) *MockConnectionBestFacadeVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Broken mocks base method.
func (m *MockConnection) Broken() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broken")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Broken indicates an expected call of Broken.
func (mr *MockConnectionMockRecorder) Broken() *MockConnectionBrokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broken", reflect.TypeOf((*MockConnection)(nil).Broken))
	return &MockConnectionBrokenCall{Call: call}
}

// MockConnectionBrokenCall wrap *gomock.Call
type MockConnectionBrokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionBrokenCall) Return(arg0 <-chan struct{}) *MockConnectionBrokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionBrokenCall) Do(f func() <-chan struct{}) *MockConnectionBrokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionBrokenCall) DoAndReturn(f func() <-chan struct{}) *MockConnectionBrokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *MockConnectionCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
	return &MockConnectionCloseCall{Call: call}
}

// MockConnectionCloseCall wrap *gomock.Call
type MockConnectionCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionCloseCall) Return(arg0 error) *MockConnectionCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionCloseCall) Do(f func() error) *MockConnectionCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionCloseCall) DoAndReturn(f func() error) *MockConnectionCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectControllerStream mocks base method.
func (m *MockConnection) ConnectControllerStream(arg0 context.Context, arg1 string, arg2 url.Values, arg3 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream.
func (mr *MockConnectionMockRecorder) ConnectControllerStream(arg0, arg1, arg2, arg3 any) *MockConnectionConnectControllerStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockConnection)(nil).ConnectControllerStream), arg0, arg1, arg2, arg3)
	return &MockConnectionConnectControllerStreamCall{Call: call}
}

// MockConnectionConnectControllerStreamCall wrap *gomock.Call
type MockConnectionConnectControllerStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionConnectControllerStreamCall) Return(arg0 base.Stream, arg1 error) *MockConnectionConnectControllerStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionConnectControllerStreamCall) Do(f func(context.Context, string, url.Values, http.Header) (base.Stream, error)) *MockConnectionConnectControllerStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionConnectControllerStreamCall) DoAndReturn(f func(context.Context, string, url.Values, http.Header) (base.Stream, error)) *MockConnectionConnectControllerStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectStream mocks base method.
func (m *MockConnection) ConnectStream(arg0 context.Context, arg1 string, arg2 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream.
func (mr *MockConnectionMockRecorder) ConnectStream(arg0, arg1, arg2 any) *MockConnectionConnectStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockConnection)(nil).ConnectStream), arg0, arg1, arg2)
	return &MockConnectionConnectStreamCall{Call: call}
}

// MockConnectionConnectStreamCall wrap *gomock.Call
type MockConnectionConnectStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionConnectStreamCall) Return(arg0 base.Stream, arg1 error) *MockConnectionConnectStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionConnectStreamCall) Do(f func(context.Context, string, url.Values) (base.Stream, error)) *MockConnectionConnectStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionConnectStreamCall) DoAndReturn(f func(context.Context, string, url.Values) (base.Stream, error)) *MockConnectionConnectStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerAccess mocks base method.
func (m *MockConnection) ControllerAccess() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAccess")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerAccess indicates an expected call of ControllerAccess.
func (mr *MockConnectionMockRecorder) ControllerAccess() *MockConnectionControllerAccessCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAccess", reflect.TypeOf((*MockConnection)(nil).ControllerAccess))
	return &MockConnectionControllerAccessCall{Call: call}
}

// MockConnectionControllerAccessCall wrap *gomock.Call
type MockConnectionControllerAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionControllerAccessCall) Return(arg0 string) *MockConnectionControllerAccessCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionControllerAccessCall) Do(f func() string) *MockConnectionControllerAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionControllerAccessCall) DoAndReturn(f func() string) *MockConnectionControllerAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerTag mocks base method.
func (m *MockConnection) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockConnectionMockRecorder) ControllerTag() *MockConnectionControllerTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockConnection)(nil).ControllerTag))
	return &MockConnectionControllerTagCall{Call: call}
}

// MockConnectionControllerTagCall wrap *gomock.Call
type MockConnectionControllerTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionControllerTagCall) Return(arg0 names.ControllerTag) *MockConnectionControllerTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionControllerTagCall) Do(f func() names.ControllerTag) *MockConnectionControllerTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionControllerTagCall) DoAndReturn(f func() names.ControllerTag) *MockConnectionControllerTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CookieURL mocks base method.
func (m *MockConnection) CookieURL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CookieURL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// CookieURL indicates an expected call of CookieURL.
func (mr *MockConnectionMockRecorder) CookieURL() *MockConnectionCookieURLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CookieURL", reflect.TypeOf((*MockConnection)(nil).CookieURL))
	return &MockConnectionCookieURLCall{Call: call}
}

// MockConnectionCookieURLCall wrap *gomock.Call
type MockConnectionCookieURLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionCookieURLCall) Return(arg0 *url.URL) *MockConnectionCookieURLCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionCookieURLCall) Do(f func() *url.URL) *MockConnectionCookieURLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionCookieURLCall) DoAndReturn(f func() *url.URL) *MockConnectionCookieURLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HTTPClient mocks base method.
func (m *MockConnection) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockConnectionMockRecorder) HTTPClient() *MockConnectionHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockConnection)(nil).HTTPClient))
	return &MockConnectionHTTPClientCall{Call: call}
}

// MockConnectionHTTPClientCall wrap *gomock.Call
type MockConnectionHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionHTTPClientCall) Return(arg0 *httprequest.Client, arg1 error) *MockConnectionHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionHTTPClientCall) Do(f func() (*httprequest.Client, error)) *MockConnectionHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionHTTPClientCall) DoAndReturn(f func() (*httprequest.Client, error)) *MockConnectionHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IPAddr mocks base method.
func (m *MockConnection) IPAddr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPAddr")
	ret0, _ := ret[0].(string)
	return ret0
}

// IPAddr indicates an expected call of IPAddr.
func (mr *MockConnectionMockRecorder) IPAddr() *MockConnectionIPAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPAddr", reflect.TypeOf((*MockConnection)(nil).IPAddr))
	return &MockConnectionIPAddrCall{Call: call}
}

// MockConnectionIPAddrCall wrap *gomock.Call
type MockConnectionIPAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionIPAddrCall) Return(arg0 string) *MockConnectionIPAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionIPAddrCall) Do(f func() string) *MockConnectionIPAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionIPAddrCall) DoAndReturn(f func() string) *MockConnectionIPAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsBroken mocks base method.
func (m *MockConnection) IsBroken(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBroken", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBroken indicates an expected call of IsBroken.
func (mr *MockConnectionMockRecorder) IsBroken(arg0 any) *MockConnectionIsBrokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBroken", reflect.TypeOf((*MockConnection)(nil).IsBroken), arg0)
	return &MockConnectionIsBrokenCall{Call: call}
}

// MockConnectionIsBrokenCall wrap *gomock.Call
type MockConnectionIsBrokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionIsBrokenCall) Return(arg0 bool) *MockConnectionIsBrokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionIsBrokenCall) Do(f func(context.Context) bool) *MockConnectionIsBrokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionIsBrokenCall) DoAndReturn(f func(context.Context) bool) *MockConnectionIsBrokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsProxied mocks base method.
func (m *MockConnection) IsProxied() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProxied")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsProxied indicates an expected call of IsProxied.
func (mr *MockConnectionMockRecorder) IsProxied() *MockConnectionIsProxiedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProxied", reflect.TypeOf((*MockConnection)(nil).IsProxied))
	return &MockConnectionIsProxiedCall{Call: call}
}

// MockConnectionIsProxiedCall wrap *gomock.Call
type MockConnectionIsProxiedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionIsProxiedCall) Return(arg0 bool) *MockConnectionIsProxiedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionIsProxiedCall) Do(f func() bool) *MockConnectionIsProxiedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionIsProxiedCall) DoAndReturn(f func() bool) *MockConnectionIsProxiedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Login mocks base method.
func (m *MockConnection) Login(arg0 context.Context, arg1 names.Tag, arg2, arg3 string, arg4 []macaroon.Slice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockConnectionMockRecorder) Login(arg0, arg1, arg2, arg3, arg4 any) *MockConnectionLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockConnection)(nil).Login), arg0, arg1, arg2, arg3, arg4)
	return &MockConnectionLoginCall{Call: call}
}

// MockConnectionLoginCall wrap *gomock.Call
type MockConnectionLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionLoginCall) Return(arg0 error) *MockConnectionLoginCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionLoginCall) Do(f func(context.Context, names.Tag, string, string, []macaroon.Slice) error) *MockConnectionLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionLoginCall) DoAndReturn(f func(context.Context, names.Tag, string, string, []macaroon.Slice) error) *MockConnectionLoginCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelTag mocks base method.
func (m *MockConnection) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockConnectionMockRecorder) ModelTag() *MockConnectionModelTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockConnection)(nil).ModelTag))
	return &MockConnectionModelTagCall{Call: call}
}

// MockConnectionModelTagCall wrap *gomock.Call
type MockConnectionModelTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionModelTagCall) Return(arg0 names.ModelTag, arg1 bool) *MockConnectionModelTagCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionModelTagCall) Do(f func() (names.ModelTag, bool)) *MockConnectionModelTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionModelTagCall) DoAndReturn(f func() (names.ModelTag, bool)) *MockConnectionModelTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Proxy mocks base method.
func (m *MockConnection) Proxy() proxy.Proxier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proxy")
	ret0, _ := ret[0].(proxy.Proxier)
	return ret0
}

// Proxy indicates an expected call of Proxy.
func (mr *MockConnectionMockRecorder) Proxy() *MockConnectionProxyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proxy", reflect.TypeOf((*MockConnection)(nil).Proxy))
	return &MockConnectionProxyCall{Call: call}
}

// MockConnectionProxyCall wrap *gomock.Call
type MockConnectionProxyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionProxyCall) Return(arg0 proxy.Proxier) *MockConnectionProxyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionProxyCall) Do(f func() proxy.Proxier) *MockConnectionProxyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionProxyCall) DoAndReturn(f func() proxy.Proxier) *MockConnectionProxyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PublicDNSName mocks base method.
func (m *MockConnection) PublicDNSName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicDNSName")
	ret0, _ := ret[0].(string)
	return ret0
}

// PublicDNSName indicates an expected call of PublicDNSName.
func (mr *MockConnectionMockRecorder) PublicDNSName() *MockConnectionPublicDNSNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicDNSName", reflect.TypeOf((*MockConnection)(nil).PublicDNSName))
	return &MockConnectionPublicDNSNameCall{Call: call}
}

// MockConnectionPublicDNSNameCall wrap *gomock.Call
type MockConnectionPublicDNSNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionPublicDNSNameCall) Return(arg0 string) *MockConnectionPublicDNSNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionPublicDNSNameCall) Do(f func() string) *MockConnectionPublicDNSNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionPublicDNSNameCall) DoAndReturn(f func() string) *MockConnectionPublicDNSNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RootHTTPClient mocks base method.
func (m *MockConnection) RootHTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootHTTPClient indicates an expected call of RootHTTPClient.
func (mr *MockConnectionMockRecorder) RootHTTPClient() *MockConnectionRootHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHTTPClient", reflect.TypeOf((*MockConnection)(nil).RootHTTPClient))
	return &MockConnectionRootHTTPClientCall{Call: call}
}

// MockConnectionRootHTTPClientCall wrap *gomock.Call
type MockConnectionRootHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionRootHTTPClientCall) Return(arg0 *httprequest.Client, arg1 error) *MockConnectionRootHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionRootHTTPClientCall) Do(f func() (*httprequest.Client, error)) *MockConnectionRootHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionRootHTTPClientCall) DoAndReturn(f func() (*httprequest.Client, error)) *MockConnectionRootHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ServerVersion mocks base method.
func (m *MockConnection) ServerVersion() (version.Number, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerVersion")
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ServerVersion indicates an expected call of ServerVersion.
func (mr *MockConnectionMockRecorder) ServerVersion() *MockConnectionServerVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerVersion", reflect.TypeOf((*MockConnection)(nil).ServerVersion))
	return &MockConnectionServerVersionCall{Call: call}
}

// MockConnectionServerVersionCall wrap *gomock.Call
type MockConnectionServerVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConnectionServerVersionCall) Return(arg0 version.Number, arg1 bool) *MockConnectionServerVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConnectionServerVersionCall) Do(f func() (version.Number, bool)) *MockConnectionServerVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConnectionServerVersionCall) DoAndReturn(f func() (version.Number, bool)) *MockConnectionServerVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
