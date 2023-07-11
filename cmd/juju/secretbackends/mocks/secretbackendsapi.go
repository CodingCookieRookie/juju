// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/secretbackends (interfaces: ListSecretBackendsAPI,AddSecretBackendsAPI,RemoveSecretBackendsAPI,UpdateSecretBackendsAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	secretbackends "github.com/juju/juju/api/client/secretbackends"
)

// MockListSecretBackendsAPI is a mock of ListSecretBackendsAPI interface.
type MockListSecretBackendsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockListSecretBackendsAPIMockRecorder
}

// MockListSecretBackendsAPIMockRecorder is the mock recorder for MockListSecretBackendsAPI.
type MockListSecretBackendsAPIMockRecorder struct {
	mock *MockListSecretBackendsAPI
}

// NewMockListSecretBackendsAPI creates a new mock instance.
func NewMockListSecretBackendsAPI(ctrl *gomock.Controller) *MockListSecretBackendsAPI {
	mock := &MockListSecretBackendsAPI{ctrl: ctrl}
	mock.recorder = &MockListSecretBackendsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListSecretBackendsAPI) EXPECT() *MockListSecretBackendsAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockListSecretBackendsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockListSecretBackendsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockListSecretBackendsAPI)(nil).Close))
}

// ListSecretBackends mocks base method.
func (m *MockListSecretBackendsAPI) ListSecretBackends(arg0 []string, arg1 bool) ([]secretbackends.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretBackends", arg0, arg1)
	ret0, _ := ret[0].([]secretbackends.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretBackends indicates an expected call of ListSecretBackends.
func (mr *MockListSecretBackendsAPIMockRecorder) ListSecretBackends(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretBackends", reflect.TypeOf((*MockListSecretBackendsAPI)(nil).ListSecretBackends), arg0, arg1)
}

// MockAddSecretBackendsAPI is a mock of AddSecretBackendsAPI interface.
type MockAddSecretBackendsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAddSecretBackendsAPIMockRecorder
}

// MockAddSecretBackendsAPIMockRecorder is the mock recorder for MockAddSecretBackendsAPI.
type MockAddSecretBackendsAPIMockRecorder struct {
	mock *MockAddSecretBackendsAPI
}

// NewMockAddSecretBackendsAPI creates a new mock instance.
func NewMockAddSecretBackendsAPI(ctrl *gomock.Controller) *MockAddSecretBackendsAPI {
	mock := &MockAddSecretBackendsAPI{ctrl: ctrl}
	mock.recorder = &MockAddSecretBackendsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddSecretBackendsAPI) EXPECT() *MockAddSecretBackendsAPIMockRecorder {
	return m.recorder
}

// AddSecretBackend mocks base method.
func (m *MockAddSecretBackendsAPI) AddSecretBackend(arg0 secretbackends.CreateSecretBackend) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecretBackend", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSecretBackend indicates an expected call of AddSecretBackend.
func (mr *MockAddSecretBackendsAPIMockRecorder) AddSecretBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecretBackend", reflect.TypeOf((*MockAddSecretBackendsAPI)(nil).AddSecretBackend), arg0)
}

// Close mocks base method.
func (m *MockAddSecretBackendsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAddSecretBackendsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAddSecretBackendsAPI)(nil).Close))
}

// MockRemoveSecretBackendsAPI is a mock of RemoveSecretBackendsAPI interface.
type MockRemoveSecretBackendsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRemoveSecretBackendsAPIMockRecorder
}

// MockRemoveSecretBackendsAPIMockRecorder is the mock recorder for MockRemoveSecretBackendsAPI.
type MockRemoveSecretBackendsAPIMockRecorder struct {
	mock *MockRemoveSecretBackendsAPI
}

// NewMockRemoveSecretBackendsAPI creates a new mock instance.
func NewMockRemoveSecretBackendsAPI(ctrl *gomock.Controller) *MockRemoveSecretBackendsAPI {
	mock := &MockRemoveSecretBackendsAPI{ctrl: ctrl}
	mock.recorder = &MockRemoveSecretBackendsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoveSecretBackendsAPI) EXPECT() *MockRemoveSecretBackendsAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRemoveSecretBackendsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRemoveSecretBackendsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRemoveSecretBackendsAPI)(nil).Close))
}

// RemoveSecretBackend mocks base method.
func (m *MockRemoveSecretBackendsAPI) RemoveSecretBackend(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretBackend indicates an expected call of RemoveSecretBackend.
func (mr *MockRemoveSecretBackendsAPIMockRecorder) RemoveSecretBackend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretBackend", reflect.TypeOf((*MockRemoveSecretBackendsAPI)(nil).RemoveSecretBackend), arg0, arg1)
}

// MockUpdateSecretBackendsAPI is a mock of UpdateSecretBackendsAPI interface.
type MockUpdateSecretBackendsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateSecretBackendsAPIMockRecorder
}

// MockUpdateSecretBackendsAPIMockRecorder is the mock recorder for MockUpdateSecretBackendsAPI.
type MockUpdateSecretBackendsAPIMockRecorder struct {
	mock *MockUpdateSecretBackendsAPI
}

// NewMockUpdateSecretBackendsAPI creates a new mock instance.
func NewMockUpdateSecretBackendsAPI(ctrl *gomock.Controller) *MockUpdateSecretBackendsAPI {
	mock := &MockUpdateSecretBackendsAPI{ctrl: ctrl}
	mock.recorder = &MockUpdateSecretBackendsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateSecretBackendsAPI) EXPECT() *MockUpdateSecretBackendsAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockUpdateSecretBackendsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUpdateSecretBackendsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUpdateSecretBackendsAPI)(nil).Close))
}

// UpdateSecretBackend mocks base method.
func (m *MockUpdateSecretBackendsAPI) UpdateSecretBackend(arg0 secretbackends.UpdateSecretBackend, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretBackend indicates an expected call of UpdateSecretBackend.
func (mr *MockUpdateSecretBackendsAPIMockRecorder) UpdateSecretBackend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretBackend", reflect.TypeOf((*MockUpdateSecretBackendsAPI)(nil).UpdateSecretBackend), arg0, arg1)
}