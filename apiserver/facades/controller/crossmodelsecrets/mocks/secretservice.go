// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/crossmodelsecrets (interfaces: SecretService,SecretBackendService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/secretservice.go github.com/juju/juju/apiserver/facades/controller/crossmodelsecrets SecretService,SecretBackendService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	unit "github.com/juju/juju/core/unit"
	service "github.com/juju/juju/domain/secret/service"
	service0 "github.com/juju/juju/domain/secretbackend/service"
	provider "github.com/juju/juju/internal/secrets/provider"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretService is a mock of SecretService interface.
type MockSecretService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceMockRecorder
}

// MockSecretServiceMockRecorder is the mock recorder for MockSecretService.
type MockSecretServiceMockRecorder struct {
	mock *MockSecretService
}

// NewMockSecretService creates a new mock instance.
func NewMockSecretService(ctrl *gomock.Controller) *MockSecretService {
	mock := &MockSecretService{ctrl: ctrl}
	mock.recorder = &MockSecretServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretService) EXPECT() *MockSecretServiceMockRecorder {
	return m.recorder
}

// GetSecret mocks base method.
func (m *MockSecretService) GetSecret(arg0 context.Context, arg1 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretServiceMockRecorder) GetSecret(arg0, arg1 any) *MockSecretServiceGetSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretService)(nil).GetSecret), arg0, arg1)
	return &MockSecretServiceGetSecretCall{Call: call}
}

// MockSecretServiceGetSecretCall wrap *gomock.Call
type MockSecretServiceGetSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetSecretCall) Return(arg0 *secrets.SecretMetadata, arg1 error) *MockSecretServiceGetSecretCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetSecretCall) Do(f func(context.Context, *secrets.URI) (*secrets.SecretMetadata, error)) *MockSecretServiceGetSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetSecretCall) DoAndReturn(f func(context.Context, *secrets.URI) (*secrets.SecretMetadata, error)) *MockSecretServiceGetSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretAccessScope mocks base method.
func (m *MockSecretService) GetSecretAccessScope(arg0 context.Context, arg1 *secrets.URI, arg2 service.SecretAccessor) (service.SecretAccessScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretAccessScope", arg0, arg1, arg2)
	ret0, _ := ret[0].(service.SecretAccessScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretAccessScope indicates an expected call of GetSecretAccessScope.
func (mr *MockSecretServiceMockRecorder) GetSecretAccessScope(arg0, arg1, arg2 any) *MockSecretServiceGetSecretAccessScopeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretAccessScope", reflect.TypeOf((*MockSecretService)(nil).GetSecretAccessScope), arg0, arg1, arg2)
	return &MockSecretServiceGetSecretAccessScopeCall{Call: call}
}

// MockSecretServiceGetSecretAccessScopeCall wrap *gomock.Call
type MockSecretServiceGetSecretAccessScopeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetSecretAccessScopeCall) Return(arg0 service.SecretAccessScope, arg1 error) *MockSecretServiceGetSecretAccessScopeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetSecretAccessScopeCall) Do(f func(context.Context, *secrets.URI, service.SecretAccessor) (service.SecretAccessScope, error)) *MockSecretServiceGetSecretAccessScopeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetSecretAccessScopeCall) DoAndReturn(f func(context.Context, *secrets.URI, service.SecretAccessor) (service.SecretAccessScope, error)) *MockSecretServiceGetSecretAccessScopeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretValue mocks base method.
func (m *MockSecretService) GetSecretValue(arg0 context.Context, arg1 *secrets.URI, arg2 int, arg3 service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretServiceMockRecorder) GetSecretValue(arg0, arg1, arg2, arg3 any) *MockSecretServiceGetSecretValueCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretService)(nil).GetSecretValue), arg0, arg1, arg2, arg3)
	return &MockSecretServiceGetSecretValueCall{Call: call}
}

// MockSecretServiceGetSecretValueCall wrap *gomock.Call
type MockSecretServiceGetSecretValueCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetSecretValueCall) Return(arg0 secrets.SecretValue, arg1 *secrets.ValueRef, arg2 error) *MockSecretServiceGetSecretValueCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetSecretValueCall) Do(f func(context.Context, *secrets.URI, int, service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error)) *MockSecretServiceGetSecretValueCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetSecretValueCall) DoAndReturn(f func(context.Context, *secrets.URI, int, service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error)) *MockSecretServiceGetSecretValueCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListGrantedSecretsForBackend mocks base method.
func (m *MockSecretService) ListGrantedSecretsForBackend(arg0 context.Context, arg1 string, arg2 secrets.SecretRole, arg3 ...service.SecretAccessor) ([]*secrets.SecretRevisionRef, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGrantedSecretsForBackend", varargs...)
	ret0, _ := ret[0].([]*secrets.SecretRevisionRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGrantedSecretsForBackend indicates an expected call of ListGrantedSecretsForBackend.
func (mr *MockSecretServiceMockRecorder) ListGrantedSecretsForBackend(arg0, arg1, arg2 any, arg3 ...any) *MockSecretServiceListGrantedSecretsForBackendCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGrantedSecretsForBackend", reflect.TypeOf((*MockSecretService)(nil).ListGrantedSecretsForBackend), varargs...)
	return &MockSecretServiceListGrantedSecretsForBackendCall{Call: call}
}

// MockSecretServiceListGrantedSecretsForBackendCall wrap *gomock.Call
type MockSecretServiceListGrantedSecretsForBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceListGrantedSecretsForBackendCall) Return(arg0 []*secrets.SecretRevisionRef, arg1 error) *MockSecretServiceListGrantedSecretsForBackendCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceListGrantedSecretsForBackendCall) Do(f func(context.Context, string, secrets.SecretRole, ...service.SecretAccessor) ([]*secrets.SecretRevisionRef, error)) *MockSecretServiceListGrantedSecretsForBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceListGrantedSecretsForBackendCall) DoAndReturn(f func(context.Context, string, secrets.SecretRole, ...service.SecretAccessor) ([]*secrets.SecretRevisionRef, error)) *MockSecretServiceListGrantedSecretsForBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateRemoteConsumedRevision mocks base method.
func (m *MockSecretService) UpdateRemoteConsumedRevision(arg0 context.Context, arg1 *secrets.URI, arg2 unit.Name, arg3 bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteConsumedRevision", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRemoteConsumedRevision indicates an expected call of UpdateRemoteConsumedRevision.
func (mr *MockSecretServiceMockRecorder) UpdateRemoteConsumedRevision(arg0, arg1, arg2, arg3 any) *MockSecretServiceUpdateRemoteConsumedRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteConsumedRevision", reflect.TypeOf((*MockSecretService)(nil).UpdateRemoteConsumedRevision), arg0, arg1, arg2, arg3)
	return &MockSecretServiceUpdateRemoteConsumedRevisionCall{Call: call}
}

// MockSecretServiceUpdateRemoteConsumedRevisionCall wrap *gomock.Call
type MockSecretServiceUpdateRemoteConsumedRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceUpdateRemoteConsumedRevisionCall) Return(arg0 int, arg1 error) *MockSecretServiceUpdateRemoteConsumedRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceUpdateRemoteConsumedRevisionCall) Do(f func(context.Context, *secrets.URI, unit.Name, bool) (int, error)) *MockSecretServiceUpdateRemoteConsumedRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceUpdateRemoteConsumedRevisionCall) DoAndReturn(f func(context.Context, *secrets.URI, unit.Name, bool) (int, error)) *MockSecretServiceUpdateRemoteConsumedRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockSecretBackendService is a mock of SecretBackendService interface.
type MockSecretBackendService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendServiceMockRecorder
}

// MockSecretBackendServiceMockRecorder is the mock recorder for MockSecretBackendService.
type MockSecretBackendServiceMockRecorder struct {
	mock *MockSecretBackendService
}

// NewMockSecretBackendService creates a new mock instance.
func NewMockSecretBackendService(ctrl *gomock.Controller) *MockSecretBackendService {
	mock := &MockSecretBackendService{ctrl: ctrl}
	mock.recorder = &MockSecretBackendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendService) EXPECT() *MockSecretBackendServiceMockRecorder {
	return m.recorder
}

// BackendConfigInfo mocks base method.
func (m *MockSecretBackendService) BackendConfigInfo(arg0 context.Context, arg1 service0.BackendConfigParams) (*provider.ModelBackendConfigInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackendConfigInfo", arg0, arg1)
	ret0, _ := ret[0].(*provider.ModelBackendConfigInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackendConfigInfo indicates an expected call of BackendConfigInfo.
func (mr *MockSecretBackendServiceMockRecorder) BackendConfigInfo(arg0, arg1 any) *MockSecretBackendServiceBackendConfigInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackendConfigInfo", reflect.TypeOf((*MockSecretBackendService)(nil).BackendConfigInfo), arg0, arg1)
	return &MockSecretBackendServiceBackendConfigInfoCall{Call: call}
}

// MockSecretBackendServiceBackendConfigInfoCall wrap *gomock.Call
type MockSecretBackendServiceBackendConfigInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretBackendServiceBackendConfigInfoCall) Return(arg0 *provider.ModelBackendConfigInfo, arg1 error) *MockSecretBackendServiceBackendConfigInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretBackendServiceBackendConfigInfoCall) Do(f func(context.Context, service0.BackendConfigParams) (*provider.ModelBackendConfigInfo, error)) *MockSecretBackendServiceBackendConfigInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretBackendServiceBackendConfigInfoCall) DoAndReturn(f func(context.Context, service0.BackendConfigParams) (*provider.ModelBackendConfigInfo, error)) *MockSecretBackendServiceBackendConfigInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
