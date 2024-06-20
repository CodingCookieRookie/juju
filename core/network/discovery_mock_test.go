// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/network (interfaces: ConfigSource,ConfigSourceNIC,ConfigSourceAddr)
//
// Generated by this command:
//
//	mockgen -typed -package network_test -destination discovery_mock_test.go github.com/juju/juju/core/network ConfigSource,ConfigSourceNIC,ConfigSourceAddr
//

// Package network_test is a generated GoMock package.
package network_test

import (
	net "net"
	reflect "reflect"

	set "github.com/juju/collections/set"
	network "github.com/juju/juju/core/network"
	gomock "go.uber.org/mock/gomock"
)

// MockConfigSource is a mock of ConfigSource interface.
type MockConfigSource struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSourceMockRecorder
}

// MockConfigSourceMockRecorder is the mock recorder for MockConfigSource.
type MockConfigSourceMockRecorder struct {
	mock *MockConfigSource
}

// NewMockConfigSource creates a new mock instance.
func NewMockConfigSource(ctrl *gomock.Controller) *MockConfigSource {
	mock := &MockConfigSource{ctrl: ctrl}
	mock.recorder = &MockConfigSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigSource) EXPECT() *MockConfigSourceMockRecorder {
	return m.recorder
}

// DefaultRoute mocks base method.
func (m *MockConfigSource) DefaultRoute() (net.IP, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultRoute")
	ret0, _ := ret[0].(net.IP)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DefaultRoute indicates an expected call of DefaultRoute.
func (mr *MockConfigSourceMockRecorder) DefaultRoute() *MockConfigSourceDefaultRouteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultRoute", reflect.TypeOf((*MockConfigSource)(nil).DefaultRoute))
	return &MockConfigSourceDefaultRouteCall{Call: call}
}

// MockConfigSourceDefaultRouteCall wrap *gomock.Call
type MockConfigSourceDefaultRouteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceDefaultRouteCall) Return(arg0 net.IP, arg1 string, arg2 error) *MockConfigSourceDefaultRouteCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceDefaultRouteCall) Do(f func() (net.IP, string, error)) *MockConfigSourceDefaultRouteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceDefaultRouteCall) DoAndReturn(f func() (net.IP, string, error)) *MockConfigSourceDefaultRouteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetBridgePorts mocks base method.
func (m *MockConfigSource) GetBridgePorts(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBridgePorts", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetBridgePorts indicates an expected call of GetBridgePorts.
func (mr *MockConfigSourceMockRecorder) GetBridgePorts(arg0 any) *MockConfigSourceGetBridgePortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBridgePorts", reflect.TypeOf((*MockConfigSource)(nil).GetBridgePorts), arg0)
	return &MockConfigSourceGetBridgePortsCall{Call: call}
}

// MockConfigSourceGetBridgePortsCall wrap *gomock.Call
type MockConfigSourceGetBridgePortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceGetBridgePortsCall) Return(arg0 []string) *MockConfigSourceGetBridgePortsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceGetBridgePortsCall) Do(f func(string) []string) *MockConfigSourceGetBridgePortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceGetBridgePortsCall) DoAndReturn(f func(string) []string) *MockConfigSourceGetBridgePortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Interfaces mocks base method.
func (m *MockConfigSource) Interfaces() ([]network.ConfigSourceNIC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Interfaces")
	ret0, _ := ret[0].([]network.ConfigSourceNIC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Interfaces indicates an expected call of Interfaces.
func (mr *MockConfigSourceMockRecorder) Interfaces() *MockConfigSourceInterfacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Interfaces", reflect.TypeOf((*MockConfigSource)(nil).Interfaces))
	return &MockConfigSourceInterfacesCall{Call: call}
}

// MockConfigSourceInterfacesCall wrap *gomock.Call
type MockConfigSourceInterfacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceInterfacesCall) Return(arg0 []network.ConfigSourceNIC, arg1 error) *MockConfigSourceInterfacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceInterfacesCall) Do(f func() ([]network.ConfigSourceNIC, error)) *MockConfigSourceInterfacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceInterfacesCall) DoAndReturn(f func() ([]network.ConfigSourceNIC, error)) *MockConfigSourceInterfacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OvsManagedBridges mocks base method.
func (m *MockConfigSource) OvsManagedBridges() (set.Strings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OvsManagedBridges")
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OvsManagedBridges indicates an expected call of OvsManagedBridges.
func (mr *MockConfigSourceMockRecorder) OvsManagedBridges() *MockConfigSourceOvsManagedBridgesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OvsManagedBridges", reflect.TypeOf((*MockConfigSource)(nil).OvsManagedBridges))
	return &MockConfigSourceOvsManagedBridgesCall{Call: call}
}

// MockConfigSourceOvsManagedBridgesCall wrap *gomock.Call
type MockConfigSourceOvsManagedBridgesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceOvsManagedBridgesCall) Return(arg0 set.Strings, arg1 error) *MockConfigSourceOvsManagedBridgesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceOvsManagedBridgesCall) Do(f func() (set.Strings, error)) *MockConfigSourceOvsManagedBridgesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceOvsManagedBridgesCall) DoAndReturn(f func() (set.Strings, error)) *MockConfigSourceOvsManagedBridgesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockConfigSourceNIC is a mock of ConfigSourceNIC interface.
type MockConfigSourceNIC struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSourceNICMockRecorder
}

// MockConfigSourceNICMockRecorder is the mock recorder for MockConfigSourceNIC.
type MockConfigSourceNICMockRecorder struct {
	mock *MockConfigSourceNIC
}

// NewMockConfigSourceNIC creates a new mock instance.
func NewMockConfigSourceNIC(ctrl *gomock.Controller) *MockConfigSourceNIC {
	mock := &MockConfigSourceNIC{ctrl: ctrl}
	mock.recorder = &MockConfigSourceNICMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigSourceNIC) EXPECT() *MockConfigSourceNICMockRecorder {
	return m.recorder
}

// Addresses mocks base method.
func (m *MockConfigSourceNIC) Addresses() ([]network.ConfigSourceAddr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].([]network.ConfigSourceAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses.
func (mr *MockConfigSourceNICMockRecorder) Addresses() *MockConfigSourceNICAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockConfigSourceNIC)(nil).Addresses))
	return &MockConfigSourceNICAddressesCall{Call: call}
}

// MockConfigSourceNICAddressesCall wrap *gomock.Call
type MockConfigSourceNICAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICAddressesCall) Return(arg0 []network.ConfigSourceAddr, arg1 error) *MockConfigSourceNICAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICAddressesCall) Do(f func() ([]network.ConfigSourceAddr, error)) *MockConfigSourceNICAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICAddressesCall) DoAndReturn(f func() ([]network.ConfigSourceAddr, error)) *MockConfigSourceNICAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HardwareAddr mocks base method.
func (m *MockConfigSourceNIC) HardwareAddr() net.HardwareAddr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareAddr")
	ret0, _ := ret[0].(net.HardwareAddr)
	return ret0
}

// HardwareAddr indicates an expected call of HardwareAddr.
func (mr *MockConfigSourceNICMockRecorder) HardwareAddr() *MockConfigSourceNICHardwareAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareAddr", reflect.TypeOf((*MockConfigSourceNIC)(nil).HardwareAddr))
	return &MockConfigSourceNICHardwareAddrCall{Call: call}
}

// MockConfigSourceNICHardwareAddrCall wrap *gomock.Call
type MockConfigSourceNICHardwareAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICHardwareAddrCall) Return(arg0 net.HardwareAddr) *MockConfigSourceNICHardwareAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICHardwareAddrCall) Do(f func() net.HardwareAddr) *MockConfigSourceNICHardwareAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICHardwareAddrCall) DoAndReturn(f func() net.HardwareAddr) *MockConfigSourceNICHardwareAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Index mocks base method.
func (m *MockConfigSourceNIC) Index() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(int)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MockConfigSourceNICMockRecorder) Index() *MockConfigSourceNICIndexCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockConfigSourceNIC)(nil).Index))
	return &MockConfigSourceNICIndexCall{Call: call}
}

// MockConfigSourceNICIndexCall wrap *gomock.Call
type MockConfigSourceNICIndexCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICIndexCall) Return(arg0 int) *MockConfigSourceNICIndexCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICIndexCall) Do(f func() int) *MockConfigSourceNICIndexCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICIndexCall) DoAndReturn(f func() int) *MockConfigSourceNICIndexCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsUp mocks base method.
func (m *MockConfigSourceNIC) IsUp() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUp")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUp indicates an expected call of IsUp.
func (mr *MockConfigSourceNICMockRecorder) IsUp() *MockConfigSourceNICIsUpCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUp", reflect.TypeOf((*MockConfigSourceNIC)(nil).IsUp))
	return &MockConfigSourceNICIsUpCall{Call: call}
}

// MockConfigSourceNICIsUpCall wrap *gomock.Call
type MockConfigSourceNICIsUpCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICIsUpCall) Return(arg0 bool) *MockConfigSourceNICIsUpCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICIsUpCall) Do(f func() bool) *MockConfigSourceNICIsUpCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICIsUpCall) DoAndReturn(f func() bool) *MockConfigSourceNICIsUpCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MTU mocks base method.
func (m *MockConfigSourceNIC) MTU() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MTU")
	ret0, _ := ret[0].(int)
	return ret0
}

// MTU indicates an expected call of MTU.
func (mr *MockConfigSourceNICMockRecorder) MTU() *MockConfigSourceNICMTUCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MTU", reflect.TypeOf((*MockConfigSourceNIC)(nil).MTU))
	return &MockConfigSourceNICMTUCall{Call: call}
}

// MockConfigSourceNICMTUCall wrap *gomock.Call
type MockConfigSourceNICMTUCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICMTUCall) Return(arg0 int) *MockConfigSourceNICMTUCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICMTUCall) Do(f func() int) *MockConfigSourceNICMTUCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICMTUCall) DoAndReturn(f func() int) *MockConfigSourceNICMTUCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockConfigSourceNIC) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockConfigSourceNICMockRecorder) Name() *MockConfigSourceNICNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockConfigSourceNIC)(nil).Name))
	return &MockConfigSourceNICNameCall{Call: call}
}

// MockConfigSourceNICNameCall wrap *gomock.Call
type MockConfigSourceNICNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICNameCall) Return(arg0 string) *MockConfigSourceNICNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICNameCall) Do(f func() string) *MockConfigSourceNICNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICNameCall) DoAndReturn(f func() string) *MockConfigSourceNICNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Type mocks base method.
func (m *MockConfigSourceNIC) Type() network.LinkLayerDeviceType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(network.LinkLayerDeviceType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockConfigSourceNICMockRecorder) Type() *MockConfigSourceNICTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockConfigSourceNIC)(nil).Type))
	return &MockConfigSourceNICTypeCall{Call: call}
}

// MockConfigSourceNICTypeCall wrap *gomock.Call
type MockConfigSourceNICTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceNICTypeCall) Return(arg0 network.LinkLayerDeviceType) *MockConfigSourceNICTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceNICTypeCall) Do(f func() network.LinkLayerDeviceType) *MockConfigSourceNICTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceNICTypeCall) DoAndReturn(f func() network.LinkLayerDeviceType) *MockConfigSourceNICTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockConfigSourceAddr is a mock of ConfigSourceAddr interface.
type MockConfigSourceAddr struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSourceAddrMockRecorder
}

// MockConfigSourceAddrMockRecorder is the mock recorder for MockConfigSourceAddr.
type MockConfigSourceAddrMockRecorder struct {
	mock *MockConfigSourceAddr
}

// NewMockConfigSourceAddr creates a new mock instance.
func NewMockConfigSourceAddr(ctrl *gomock.Controller) *MockConfigSourceAddr {
	mock := &MockConfigSourceAddr{ctrl: ctrl}
	mock.recorder = &MockConfigSourceAddrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigSourceAddr) EXPECT() *MockConfigSourceAddrMockRecorder {
	return m.recorder
}

// IP mocks base method.
func (m *MockConfigSourceAddr) IP() net.IP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IP")
	ret0, _ := ret[0].(net.IP)
	return ret0
}

// IP indicates an expected call of IP.
func (mr *MockConfigSourceAddrMockRecorder) IP() *MockConfigSourceAddrIPCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IP", reflect.TypeOf((*MockConfigSourceAddr)(nil).IP))
	return &MockConfigSourceAddrIPCall{Call: call}
}

// MockConfigSourceAddrIPCall wrap *gomock.Call
type MockConfigSourceAddrIPCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceAddrIPCall) Return(arg0 net.IP) *MockConfigSourceAddrIPCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceAddrIPCall) Do(f func() net.IP) *MockConfigSourceAddrIPCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceAddrIPCall) DoAndReturn(f func() net.IP) *MockConfigSourceAddrIPCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IPNet mocks base method.
func (m *MockConfigSourceAddr) IPNet() *net.IPNet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPNet")
	ret0, _ := ret[0].(*net.IPNet)
	return ret0
}

// IPNet indicates an expected call of IPNet.
func (mr *MockConfigSourceAddrMockRecorder) IPNet() *MockConfigSourceAddrIPNetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPNet", reflect.TypeOf((*MockConfigSourceAddr)(nil).IPNet))
	return &MockConfigSourceAddrIPNetCall{Call: call}
}

// MockConfigSourceAddrIPNetCall wrap *gomock.Call
type MockConfigSourceAddrIPNetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceAddrIPNetCall) Return(arg0 *net.IPNet) *MockConfigSourceAddrIPNetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceAddrIPNetCall) Do(f func() *net.IPNet) *MockConfigSourceAddrIPNetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceAddrIPNetCall) DoAndReturn(f func() *net.IPNet) *MockConfigSourceAddrIPNetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsSecondary mocks base method.
func (m *MockConfigSourceAddr) IsSecondary() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSecondary")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecondary indicates an expected call of IsSecondary.
func (mr *MockConfigSourceAddrMockRecorder) IsSecondary() *MockConfigSourceAddrIsSecondaryCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSecondary", reflect.TypeOf((*MockConfigSourceAddr)(nil).IsSecondary))
	return &MockConfigSourceAddrIsSecondaryCall{Call: call}
}

// MockConfigSourceAddrIsSecondaryCall wrap *gomock.Call
type MockConfigSourceAddrIsSecondaryCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceAddrIsSecondaryCall) Return(arg0 bool) *MockConfigSourceAddrIsSecondaryCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceAddrIsSecondaryCall) Do(f func() bool) *MockConfigSourceAddrIsSecondaryCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceAddrIsSecondaryCall) DoAndReturn(f func() bool) *MockConfigSourceAddrIsSecondaryCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// String mocks base method.
func (m *MockConfigSourceAddr) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockConfigSourceAddrMockRecorder) String() *MockConfigSourceAddrStringCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConfigSourceAddr)(nil).String))
	return &MockConfigSourceAddrStringCall{Call: call}
}

// MockConfigSourceAddrStringCall wrap *gomock.Call
type MockConfigSourceAddrStringCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConfigSourceAddrStringCall) Return(arg0 string) *MockConfigSourceAddrStringCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConfigSourceAddrStringCall) Do(f func() string) *MockConfigSourceAddrStringCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConfigSourceAddrStringCall) DoAndReturn(f func() string) *MockConfigSourceAddrStringCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}