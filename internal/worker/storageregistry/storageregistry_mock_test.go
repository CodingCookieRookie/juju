// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/storageregistry (interfaces: StorageRegistryWorker)
//
// Generated by this command:
//
//	mockgen -typed -package storageregistry -destination storageregistry_mock_test.go github.com/juju/juju/internal/worker/storageregistry StorageRegistryWorker
//

// Package storageregistry is a generated GoMock package.
package storageregistry

import (
	reflect "reflect"

	storage "github.com/juju/juju/internal/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockStorageRegistryWorker is a mock of StorageRegistryWorker interface.
type MockStorageRegistryWorker struct {
	ctrl     *gomock.Controller
	recorder *MockStorageRegistryWorkerMockRecorder
}

// MockStorageRegistryWorkerMockRecorder is the mock recorder for MockStorageRegistryWorker.
type MockStorageRegistryWorkerMockRecorder struct {
	mock *MockStorageRegistryWorker
}

// NewMockStorageRegistryWorker creates a new mock instance.
func NewMockStorageRegistryWorker(ctrl *gomock.Controller) *MockStorageRegistryWorker {
	mock := &MockStorageRegistryWorker{ctrl: ctrl}
	mock.recorder = &MockStorageRegistryWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageRegistryWorker) EXPECT() *MockStorageRegistryWorkerMockRecorder {
	return m.recorder
}

// Kill mocks base method.
func (m *MockStorageRegistryWorker) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockStorageRegistryWorkerMockRecorder) Kill() *MockStorageRegistryWorkerKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockStorageRegistryWorker)(nil).Kill))
	return &MockStorageRegistryWorkerKillCall{Call: call}
}

// MockStorageRegistryWorkerKillCall wrap *gomock.Call
type MockStorageRegistryWorkerKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageRegistryWorkerKillCall) Return() *MockStorageRegistryWorkerKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageRegistryWorkerKillCall) Do(f func()) *MockStorageRegistryWorkerKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageRegistryWorkerKillCall) DoAndReturn(f func()) *MockStorageRegistryWorkerKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageProvider mocks base method.
func (m *MockStorageRegistryWorker) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockStorageRegistryWorkerMockRecorder) StorageProvider(arg0 any) *MockStorageRegistryWorkerStorageProviderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockStorageRegistryWorker)(nil).StorageProvider), arg0)
	return &MockStorageRegistryWorkerStorageProviderCall{Call: call}
}

// MockStorageRegistryWorkerStorageProviderCall wrap *gomock.Call
type MockStorageRegistryWorkerStorageProviderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageRegistryWorkerStorageProviderCall) Return(arg0 storage.Provider, arg1 error) *MockStorageRegistryWorkerStorageProviderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageRegistryWorkerStorageProviderCall) Do(f func(storage.ProviderType) (storage.Provider, error)) *MockStorageRegistryWorkerStorageProviderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageRegistryWorkerStorageProviderCall) DoAndReturn(f func(storage.ProviderType) (storage.Provider, error)) *MockStorageRegistryWorkerStorageProviderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageProviderTypes mocks base method.
func (m *MockStorageRegistryWorker) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockStorageRegistryWorkerMockRecorder) StorageProviderTypes() *MockStorageRegistryWorkerStorageProviderTypesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockStorageRegistryWorker)(nil).StorageProviderTypes))
	return &MockStorageRegistryWorkerStorageProviderTypesCall{Call: call}
}

// MockStorageRegistryWorkerStorageProviderTypesCall wrap *gomock.Call
type MockStorageRegistryWorkerStorageProviderTypesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageRegistryWorkerStorageProviderTypesCall) Return(arg0 []storage.ProviderType, arg1 error) *MockStorageRegistryWorkerStorageProviderTypesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageRegistryWorkerStorageProviderTypesCall) Do(f func() ([]storage.ProviderType, error)) *MockStorageRegistryWorkerStorageProviderTypesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageRegistryWorkerStorageProviderTypesCall) DoAndReturn(f func() ([]storage.ProviderType, error)) *MockStorageRegistryWorkerStorageProviderTypesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockStorageRegistryWorker) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockStorageRegistryWorkerMockRecorder) Wait() *MockStorageRegistryWorkerWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockStorageRegistryWorker)(nil).Wait))
	return &MockStorageRegistryWorkerWaitCall{Call: call}
}

// MockStorageRegistryWorkerWaitCall wrap *gomock.Call
type MockStorageRegistryWorkerWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageRegistryWorkerWaitCall) Return(arg0 error) *MockStorageRegistryWorkerWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageRegistryWorkerWaitCall) Do(f func() error) *MockStorageRegistryWorkerWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageRegistryWorkerWaitCall) DoAndReturn(f func() error) *MockStorageRegistryWorkerWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
