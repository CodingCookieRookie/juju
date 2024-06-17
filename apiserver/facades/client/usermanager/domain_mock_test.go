// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/usermanager (interfaces: UserService)
//
// Generated by this command:
//
//	mockgen -typed -package usermanager_test -destination domain_mock_test.go github.com/juju/juju/apiserver/facades/client/usermanager UserService
//

// Package usermanager_test is a generated GoMock package.
package usermanager_test

import (
	context "context"
	reflect "reflect"

	user "github.com/juju/juju/core/user"
	service "github.com/juju/juju/domain/access/service"
	auth "github.com/juju/juju/internal/auth"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserService) AddUser(arg0 context.Context, arg1 service.AddUserArg) (user.UUID, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(user.UUID)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceMockRecorder) AddUser(arg0, arg1 any) *MockUserServiceAddUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserService)(nil).AddUser), arg0, arg1)
	return &MockUserServiceAddUserCall{Call: call}
}

// MockUserServiceAddUserCall wrap *gomock.Call
type MockUserServiceAddUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceAddUserCall) Return(arg0 user.UUID, arg1 []byte, arg2 error) *MockUserServiceAddUserCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceAddUserCall) Do(f func(context.Context, service.AddUserArg) (user.UUID, []byte, error)) *MockUserServiceAddUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceAddUserCall) DoAndReturn(f func(context.Context, service.AddUserArg) (user.UUID, []byte, error)) *MockUserServiceAddUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DisableUserAuthentication mocks base method.
func (m *MockUserService) DisableUserAuthentication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUserAuthentication indicates an expected call of DisableUserAuthentication.
func (mr *MockUserServiceMockRecorder) DisableUserAuthentication(arg0, arg1 any) *MockUserServiceDisableUserAuthenticationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUserAuthentication", reflect.TypeOf((*MockUserService)(nil).DisableUserAuthentication), arg0, arg1)
	return &MockUserServiceDisableUserAuthenticationCall{Call: call}
}

// MockUserServiceDisableUserAuthenticationCall wrap *gomock.Call
type MockUserServiceDisableUserAuthenticationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceDisableUserAuthenticationCall) Return(arg0 error) *MockUserServiceDisableUserAuthenticationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceDisableUserAuthenticationCall) Do(f func(context.Context, string) error) *MockUserServiceDisableUserAuthenticationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceDisableUserAuthenticationCall) DoAndReturn(f func(context.Context, string) error) *MockUserServiceDisableUserAuthenticationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnableUserAuthentication mocks base method.
func (m *MockUserService) EnableUserAuthentication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUserAuthentication indicates an expected call of EnableUserAuthentication.
func (mr *MockUserServiceMockRecorder) EnableUserAuthentication(arg0, arg1 any) *MockUserServiceEnableUserAuthenticationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUserAuthentication", reflect.TypeOf((*MockUserService)(nil).EnableUserAuthentication), arg0, arg1)
	return &MockUserServiceEnableUserAuthenticationCall{Call: call}
}

// MockUserServiceEnableUserAuthenticationCall wrap *gomock.Call
type MockUserServiceEnableUserAuthenticationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceEnableUserAuthenticationCall) Return(arg0 error) *MockUserServiceEnableUserAuthenticationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceEnableUserAuthenticationCall) Do(f func(context.Context, string) error) *MockUserServiceEnableUserAuthenticationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceEnableUserAuthenticationCall) DoAndReturn(f func(context.Context, string) error) *MockUserServiceEnableUserAuthenticationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllUsers mocks base method.
func (m *MockUserService) GetAllUsers(arg0 context.Context) ([]user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0)
	ret0, _ := ret[0].([]user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceMockRecorder) GetAllUsers(arg0 any) *MockUserServiceGetAllUsersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserService)(nil).GetAllUsers), arg0)
	return &MockUserServiceGetAllUsersCall{Call: call}
}

// MockUserServiceGetAllUsersCall wrap *gomock.Call
type MockUserServiceGetAllUsersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceGetAllUsersCall) Return(arg0 []user.User, arg1 error) *MockUserServiceGetAllUsersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceGetAllUsersCall) Do(f func(context.Context) ([]user.User, error)) *MockUserServiceGetAllUsersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceGetAllUsersCall) DoAndReturn(f func(context.Context) ([]user.User, error)) *MockUserServiceGetAllUsersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUserByName mocks base method.
func (m *MockUserService) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserServiceMockRecorder) GetUserByName(arg0, arg1 any) *MockUserServiceGetUserByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserService)(nil).GetUserByName), arg0, arg1)
	return &MockUserServiceGetUserByNameCall{Call: call}
}

// MockUserServiceGetUserByNameCall wrap *gomock.Call
type MockUserServiceGetUserByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceGetUserByNameCall) Return(arg0 user.User, arg1 error) *MockUserServiceGetUserByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceGetUserByNameCall) Do(f func(context.Context, string) (user.User, error)) *MockUserServiceGetUserByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceGetUserByNameCall) DoAndReturn(f func(context.Context, string) (user.User, error)) *MockUserServiceGetUserByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveUser mocks base method.
func (m *MockUserService) RemoveUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUserServiceMockRecorder) RemoveUser(arg0, arg1 any) *MockUserServiceRemoveUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUserService)(nil).RemoveUser), arg0, arg1)
	return &MockUserServiceRemoveUserCall{Call: call}
}

// MockUserServiceRemoveUserCall wrap *gomock.Call
type MockUserServiceRemoveUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceRemoveUserCall) Return(arg0 error) *MockUserServiceRemoveUserCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceRemoveUserCall) Do(f func(context.Context, string) error) *MockUserServiceRemoveUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceRemoveUserCall) DoAndReturn(f func(context.Context, string) error) *MockUserServiceRemoveUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResetPassword mocks base method.
func (m *MockUserService) ResetPassword(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserServiceMockRecorder) ResetPassword(arg0, arg1 any) *MockUserServiceResetPasswordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserService)(nil).ResetPassword), arg0, arg1)
	return &MockUserServiceResetPasswordCall{Call: call}
}

// MockUserServiceResetPasswordCall wrap *gomock.Call
type MockUserServiceResetPasswordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceResetPasswordCall) Return(arg0 []byte, arg1 error) *MockUserServiceResetPasswordCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceResetPasswordCall) Do(f func(context.Context, string) ([]byte, error)) *MockUserServiceResetPasswordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceResetPasswordCall) DoAndReturn(f func(context.Context, string) ([]byte, error)) *MockUserServiceResetPasswordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetPassword mocks base method.
func (m *MockUserService) SetPassword(arg0 context.Context, arg1 string, arg2 auth.Password) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPassword indicates an expected call of SetPassword.
func (mr *MockUserServiceMockRecorder) SetPassword(arg0, arg1, arg2 any) *MockUserServiceSetPasswordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockUserService)(nil).SetPassword), arg0, arg1, arg2)
	return &MockUserServiceSetPasswordCall{Call: call}
}

// MockUserServiceSetPasswordCall wrap *gomock.Call
type MockUserServiceSetPasswordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceSetPasswordCall) Return(arg0 error) *MockUserServiceSetPasswordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceSetPasswordCall) Do(f func(context.Context, string, auth.Password) error) *MockUserServiceSetPasswordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceSetPasswordCall) DoAndReturn(f func(context.Context, string, auth.Password) error) *MockUserServiceSetPasswordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}