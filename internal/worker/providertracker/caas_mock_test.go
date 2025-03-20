// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/caas (interfaces: Broker)
//
// Generated by this command:
//
//	mockgen -typed -package providertracker -destination caas_mock_test.go github.com/juju/juju/caas Broker
//

// Package providertracker is a generated GoMock package.
package providertracker

import (
	context "context"
	reflect "reflect"

	caas "github.com/juju/juju/caas"
	constraints "github.com/juju/juju/core/constraints"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	envcontext "github.com/juju/juju/environs/envcontext"
	docker "github.com/juju/juju/internal/docker"
	proxy "github.com/juju/juju/internal/proxy"
	storage "github.com/juju/juju/internal/storage"
	names "github.com/juju/names/v6"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// APIVersion mocks base method.
func (m *MockBroker) APIVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIVersion indicates an expected call of APIVersion.
func (mr *MockBrokerMockRecorder) APIVersion() *MockBrokerAPIVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockBroker)(nil).APIVersion))
	return &MockBrokerAPIVersionCall{Call: call}
}

// MockBrokerAPIVersionCall wrap *gomock.Call
type MockBrokerAPIVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerAPIVersionCall) Return(arg0 string, arg1 error) *MockBrokerAPIVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerAPIVersionCall) Do(f func() (string, error)) *MockBrokerAPIVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerAPIVersionCall) DoAndReturn(f func() (string, error)) *MockBrokerAPIVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AdoptResources mocks base method.
func (m *MockBroker) AdoptResources(arg0 envcontext.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources.
func (mr *MockBrokerMockRecorder) AdoptResources(arg0, arg1, arg2 any) *MockBrokerAdoptResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockBroker)(nil).AdoptResources), arg0, arg1, arg2)
	return &MockBrokerAdoptResourcesCall{Call: call}
}

// MockBrokerAdoptResourcesCall wrap *gomock.Call
type MockBrokerAdoptResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerAdoptResourcesCall) Return(arg0 error) *MockBrokerAdoptResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerAdoptResourcesCall) Do(f func(envcontext.ProviderCallContext, string, version.Number) error) *MockBrokerAdoptResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerAdoptResourcesCall) DoAndReturn(f func(envcontext.ProviderCallContext, string, version.Number) error) *MockBrokerAdoptResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AnnotateUnit mocks base method.
func (m *MockBroker) AnnotateUnit(arg0 context.Context, arg1, arg2 string, arg3 names.UnitTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnotateUnit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnotateUnit indicates an expected call of AnnotateUnit.
func (mr *MockBrokerMockRecorder) AnnotateUnit(arg0, arg1, arg2, arg3 any) *MockBrokerAnnotateUnitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotateUnit", reflect.TypeOf((*MockBroker)(nil).AnnotateUnit), arg0, arg1, arg2, arg3)
	return &MockBrokerAnnotateUnitCall{Call: call}
}

// MockBrokerAnnotateUnitCall wrap *gomock.Call
type MockBrokerAnnotateUnitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerAnnotateUnitCall) Return(arg0 error) *MockBrokerAnnotateUnitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerAnnotateUnitCall) Do(f func(context.Context, string, string, names.UnitTag) error) *MockBrokerAnnotateUnitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerAnnotateUnitCall) DoAndReturn(f func(context.Context, string, string, names.UnitTag) error) *MockBrokerAnnotateUnitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Application mocks base method.
func (m *MockBroker) Application(arg0 string, arg1 caas.DeploymentType) caas.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0, arg1)
	ret0, _ := ret[0].(caas.Application)
	return ret0
}

// Application indicates an expected call of Application.
func (mr *MockBrokerMockRecorder) Application(arg0, arg1 any) *MockBrokerApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBroker)(nil).Application), arg0, arg1)
	return &MockBrokerApplicationCall{Call: call}
}

// MockBrokerApplicationCall wrap *gomock.Call
type MockBrokerApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerApplicationCall) Return(arg0 caas.Application) *MockBrokerApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerApplicationCall) Do(f func(string, caas.DeploymentType) caas.Application) *MockBrokerApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerApplicationCall) DoAndReturn(f func(string, caas.DeploymentType) caas.Application) *MockBrokerApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Bootstrap mocks base method.
func (m *MockBroker) Bootstrap(arg0 environs.BootstrapContext, arg1 envcontext.ProviderCallContext, arg2 environs.BootstrapParams) (*environs.BootstrapResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*environs.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockBrokerMockRecorder) Bootstrap(arg0, arg1, arg2 any) *MockBrokerBootstrapCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockBroker)(nil).Bootstrap), arg0, arg1, arg2)
	return &MockBrokerBootstrapCall{Call: call}
}

// MockBrokerBootstrapCall wrap *gomock.Call
type MockBrokerBootstrapCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerBootstrapCall) Return(arg0 *environs.BootstrapResult, arg1 error) *MockBrokerBootstrapCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerBootstrapCall) Do(f func(environs.BootstrapContext, envcontext.ProviderCallContext, environs.BootstrapParams) (*environs.BootstrapResult, error)) *MockBrokerBootstrapCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerBootstrapCall) DoAndReturn(f func(environs.BootstrapContext, envcontext.ProviderCallContext, environs.BootstrapParams) (*environs.BootstrapResult, error)) *MockBrokerBootstrapCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CheckCloudCredentials mocks base method.
func (m *MockBroker) CheckCloudCredentials(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCloudCredentials", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCloudCredentials indicates an expected call of CheckCloudCredentials.
func (mr *MockBrokerMockRecorder) CheckCloudCredentials(arg0 any) *MockBrokerCheckCloudCredentialsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCloudCredentials", reflect.TypeOf((*MockBroker)(nil).CheckCloudCredentials), arg0)
	return &MockBrokerCheckCloudCredentialsCall{Call: call}
}

// MockBrokerCheckCloudCredentialsCall wrap *gomock.Call
type MockBrokerCheckCloudCredentialsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerCheckCloudCredentialsCall) Return(arg0 error) *MockBrokerCheckCloudCredentialsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerCheckCloudCredentialsCall) Do(f func(context.Context) error) *MockBrokerCheckCloudCredentialsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerCheckCloudCredentialsCall) DoAndReturn(f func(context.Context) error) *MockBrokerCheckCloudCredentialsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Config mocks base method.
func (m *MockBroker) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockBrokerMockRecorder) Config() *MockBrokerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBroker)(nil).Config))
	return &MockBrokerConfigCall{Call: call}
}

// MockBrokerConfigCall wrap *gomock.Call
type MockBrokerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerConfigCall) Return(arg0 *config.Config) *MockBrokerConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerConfigCall) Do(f func() *config.Config) *MockBrokerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerConfigCall) DoAndReturn(f func() *config.Config) *MockBrokerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConstraintsValidator mocks base method.
func (m *MockBroker) ConstraintsValidator(arg0 envcontext.ProviderCallContext) (constraints.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstraintsValidator", arg0)
	ret0, _ := ret[0].(constraints.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstraintsValidator indicates an expected call of ConstraintsValidator.
func (mr *MockBrokerMockRecorder) ConstraintsValidator(arg0 any) *MockBrokerConstraintsValidatorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstraintsValidator", reflect.TypeOf((*MockBroker)(nil).ConstraintsValidator), arg0)
	return &MockBrokerConstraintsValidatorCall{Call: call}
}

// MockBrokerConstraintsValidatorCall wrap *gomock.Call
type MockBrokerConstraintsValidatorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerConstraintsValidatorCall) Return(arg0 constraints.Validator, arg1 error) *MockBrokerConstraintsValidatorCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerConstraintsValidatorCall) Do(f func(envcontext.ProviderCallContext) (constraints.Validator, error)) *MockBrokerConstraintsValidatorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerConstraintsValidatorCall) DoAndReturn(f func(envcontext.ProviderCallContext) (constraints.Validator, error)) *MockBrokerConstraintsValidatorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Destroy mocks base method.
func (m *MockBroker) Destroy(arg0 envcontext.ProviderCallContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockBrokerMockRecorder) Destroy(arg0 any) *MockBrokerDestroyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockBroker)(nil).Destroy), arg0)
	return &MockBrokerDestroyCall{Call: call}
}

// MockBrokerDestroyCall wrap *gomock.Call
type MockBrokerDestroyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerDestroyCall) Return(arg0 error) *MockBrokerDestroyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerDestroyCall) Do(f func(envcontext.ProviderCallContext) error) *MockBrokerDestroyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerDestroyCall) DoAndReturn(f func(envcontext.ProviderCallContext) error) *MockBrokerDestroyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DestroyController mocks base method.
func (m *MockBroker) DestroyController(arg0 envcontext.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyController indicates an expected call of DestroyController.
func (mr *MockBrokerMockRecorder) DestroyController(arg0, arg1 any) *MockBrokerDestroyControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyController", reflect.TypeOf((*MockBroker)(nil).DestroyController), arg0, arg1)
	return &MockBrokerDestroyControllerCall{Call: call}
}

// MockBrokerDestroyControllerCall wrap *gomock.Call
type MockBrokerDestroyControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerDestroyControllerCall) Return(arg0 error) *MockBrokerDestroyControllerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerDestroyControllerCall) Do(f func(envcontext.ProviderCallContext, string) error) *MockBrokerDestroyControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerDestroyControllerCall) DoAndReturn(f func(envcontext.ProviderCallContext, string) error) *MockBrokerDestroyControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnsureImageRepoSecret mocks base method.
func (m *MockBroker) EnsureImageRepoSecret(arg0 context.Context, arg1 docker.ImageRepoDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureImageRepoSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureImageRepoSecret indicates an expected call of EnsureImageRepoSecret.
func (mr *MockBrokerMockRecorder) EnsureImageRepoSecret(arg0, arg1 any) *MockBrokerEnsureImageRepoSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureImageRepoSecret", reflect.TypeOf((*MockBroker)(nil).EnsureImageRepoSecret), arg0, arg1)
	return &MockBrokerEnsureImageRepoSecretCall{Call: call}
}

// MockBrokerEnsureImageRepoSecretCall wrap *gomock.Call
type MockBrokerEnsureImageRepoSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerEnsureImageRepoSecretCall) Return(arg0 error) *MockBrokerEnsureImageRepoSecretCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerEnsureImageRepoSecretCall) Do(f func(context.Context, docker.ImageRepoDetails) error) *MockBrokerEnsureImageRepoSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerEnsureImageRepoSecretCall) DoAndReturn(f func(context.Context, docker.ImageRepoDetails) error) *MockBrokerEnsureImageRepoSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnsureModelOperator mocks base method.
func (m *MockBroker) EnsureModelOperator(arg0 context.Context, arg1, arg2 string, arg3 *caas.ModelOperatorConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureModelOperator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureModelOperator indicates an expected call of EnsureModelOperator.
func (mr *MockBrokerMockRecorder) EnsureModelOperator(arg0, arg1, arg2, arg3 any) *MockBrokerEnsureModelOperatorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureModelOperator", reflect.TypeOf((*MockBroker)(nil).EnsureModelOperator), arg0, arg1, arg2, arg3)
	return &MockBrokerEnsureModelOperatorCall{Call: call}
}

// MockBrokerEnsureModelOperatorCall wrap *gomock.Call
type MockBrokerEnsureModelOperatorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerEnsureModelOperatorCall) Return(arg0 error) *MockBrokerEnsureModelOperatorCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerEnsureModelOperatorCall) Do(f func(context.Context, string, string, *caas.ModelOperatorConfig) error) *MockBrokerEnsureModelOperatorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerEnsureModelOperatorCall) DoAndReturn(f func(context.Context, string, string, *caas.ModelOperatorConfig) error) *MockBrokerEnsureModelOperatorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretToken mocks base method.
func (m *MockBroker) GetSecretToken(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretToken indicates an expected call of GetSecretToken.
func (mr *MockBrokerMockRecorder) GetSecretToken(arg0, arg1 any) *MockBrokerGetSecretTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretToken", reflect.TypeOf((*MockBroker)(nil).GetSecretToken), arg0, arg1)
	return &MockBrokerGetSecretTokenCall{Call: call}
}

// MockBrokerGetSecretTokenCall wrap *gomock.Call
type MockBrokerGetSecretTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerGetSecretTokenCall) Return(arg0 string, arg1 error) *MockBrokerGetSecretTokenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerGetSecretTokenCall) Do(f func(context.Context, string) (string, error)) *MockBrokerGetSecretTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerGetSecretTokenCall) DoAndReturn(f func(context.Context, string) (string, error)) *MockBrokerGetSecretTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetService mocks base method.
func (m *MockBroker) GetService(arg0 context.Context, arg1 string, arg2 bool) (*caas.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*caas.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockBrokerMockRecorder) GetService(arg0, arg1, arg2 any) *MockBrokerGetServiceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockBroker)(nil).GetService), arg0, arg1, arg2)
	return &MockBrokerGetServiceCall{Call: call}
}

// MockBrokerGetServiceCall wrap *gomock.Call
type MockBrokerGetServiceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerGetServiceCall) Return(arg0 *caas.Service, arg1 error) *MockBrokerGetServiceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerGetServiceCall) Do(f func(context.Context, string, bool) (*caas.Service, error)) *MockBrokerGetServiceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerGetServiceCall) DoAndReturn(f func(context.Context, string, bool) (*caas.Service, error)) *MockBrokerGetServiceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelOperator mocks base method.
func (m *MockBroker) ModelOperator(arg0 context.Context) (*caas.ModelOperatorConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOperator", arg0)
	ret0, _ := ret[0].(*caas.ModelOperatorConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOperator indicates an expected call of ModelOperator.
func (mr *MockBrokerMockRecorder) ModelOperator(arg0 any) *MockBrokerModelOperatorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOperator", reflect.TypeOf((*MockBroker)(nil).ModelOperator), arg0)
	return &MockBrokerModelOperatorCall{Call: call}
}

// MockBrokerModelOperatorCall wrap *gomock.Call
type MockBrokerModelOperatorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerModelOperatorCall) Return(arg0 *caas.ModelOperatorConfig, arg1 error) *MockBrokerModelOperatorCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerModelOperatorCall) Do(f func(context.Context) (*caas.ModelOperatorConfig, error)) *MockBrokerModelOperatorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerModelOperatorCall) DoAndReturn(f func(context.Context) (*caas.ModelOperatorConfig, error)) *MockBrokerModelOperatorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelOperatorExists mocks base method.
func (m *MockBroker) ModelOperatorExists(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOperatorExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOperatorExists indicates an expected call of ModelOperatorExists.
func (mr *MockBrokerMockRecorder) ModelOperatorExists(arg0 any) *MockBrokerModelOperatorExistsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOperatorExists", reflect.TypeOf((*MockBroker)(nil).ModelOperatorExists), arg0)
	return &MockBrokerModelOperatorExistsCall{Call: call}
}

// MockBrokerModelOperatorExistsCall wrap *gomock.Call
type MockBrokerModelOperatorExistsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerModelOperatorExistsCall) Return(arg0 bool, arg1 error) *MockBrokerModelOperatorExistsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerModelOperatorExistsCall) Do(f func(context.Context) (bool, error)) *MockBrokerModelOperatorExistsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerModelOperatorExistsCall) DoAndReturn(f func(context.Context) (bool, error)) *MockBrokerModelOperatorExistsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PrecheckInstance mocks base method.
func (m *MockBroker) PrecheckInstance(arg0 envcontext.ProviderCallContext, arg1 environs.PrecheckInstanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrecheckInstance indicates an expected call of PrecheckInstance.
func (mr *MockBrokerMockRecorder) PrecheckInstance(arg0, arg1 any) *MockBrokerPrecheckInstanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckInstance", reflect.TypeOf((*MockBroker)(nil).PrecheckInstance), arg0, arg1)
	return &MockBrokerPrecheckInstanceCall{Call: call}
}

// MockBrokerPrecheckInstanceCall wrap *gomock.Call
type MockBrokerPrecheckInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerPrecheckInstanceCall) Return(arg0 error) *MockBrokerPrecheckInstanceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerPrecheckInstanceCall) Do(f func(envcontext.ProviderCallContext, environs.PrecheckInstanceParams) error) *MockBrokerPrecheckInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerPrecheckInstanceCall) DoAndReturn(f func(envcontext.ProviderCallContext, environs.PrecheckInstanceParams) error) *MockBrokerPrecheckInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PrepareForBootstrap mocks base method.
func (m *MockBroker) PrepareForBootstrap(arg0 environs.BootstrapContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForBootstrap indicates an expected call of PrepareForBootstrap.
func (mr *MockBrokerMockRecorder) PrepareForBootstrap(arg0, arg1 any) *MockBrokerPrepareForBootstrapCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForBootstrap", reflect.TypeOf((*MockBroker)(nil).PrepareForBootstrap), arg0, arg1)
	return &MockBrokerPrepareForBootstrapCall{Call: call}
}

// MockBrokerPrepareForBootstrapCall wrap *gomock.Call
type MockBrokerPrepareForBootstrapCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerPrepareForBootstrapCall) Return(arg0 error) *MockBrokerPrepareForBootstrapCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerPrepareForBootstrapCall) Do(f func(environs.BootstrapContext, string) error) *MockBrokerPrepareForBootstrapCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerPrepareForBootstrapCall) DoAndReturn(f func(environs.BootstrapContext, string) error) *MockBrokerPrepareForBootstrapCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Provider mocks base method.
func (m *MockBroker) Provider() caas.ContainerEnvironProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(caas.ContainerEnvironProvider)
	return ret0
}

// Provider indicates an expected call of Provider.
func (mr *MockBrokerMockRecorder) Provider() *MockBrokerProviderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockBroker)(nil).Provider))
	return &MockBrokerProviderCall{Call: call}
}

// MockBrokerProviderCall wrap *gomock.Call
type MockBrokerProviderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerProviderCall) Return(arg0 caas.ContainerEnvironProvider) *MockBrokerProviderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerProviderCall) Do(f func() caas.ContainerEnvironProvider) *MockBrokerProviderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerProviderCall) DoAndReturn(f func() caas.ContainerEnvironProvider) *MockBrokerProviderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ProxyToApplication mocks base method.
func (m *MockBroker) ProxyToApplication(arg0 context.Context, arg1, arg2 string) (proxy.Proxier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxyToApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].(proxy.Proxier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProxyToApplication indicates an expected call of ProxyToApplication.
func (mr *MockBrokerMockRecorder) ProxyToApplication(arg0, arg1, arg2 any) *MockBrokerProxyToApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyToApplication", reflect.TypeOf((*MockBroker)(nil).ProxyToApplication), arg0, arg1, arg2)
	return &MockBrokerProxyToApplicationCall{Call: call}
}

// MockBrokerProxyToApplicationCall wrap *gomock.Call
type MockBrokerProxyToApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerProxyToApplicationCall) Return(arg0 proxy.Proxier, arg1 error) *MockBrokerProxyToApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerProxyToApplicationCall) Do(f func(context.Context, string, string) (proxy.Proxier, error)) *MockBrokerProxyToApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerProxyToApplicationCall) DoAndReturn(f func(context.Context, string, string) (proxy.Proxier, error)) *MockBrokerProxyToApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetConfig mocks base method.
func (m *MockBroker) SetConfig(arg0 context.Context, arg1 *config.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockBrokerMockRecorder) SetConfig(arg0, arg1 any) *MockBrokerSetConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockBroker)(nil).SetConfig), arg0, arg1)
	return &MockBrokerSetConfigCall{Call: call}
}

// MockBrokerSetConfigCall wrap *gomock.Call
type MockBrokerSetConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerSetConfigCall) Return(arg0 error) *MockBrokerSetConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerSetConfigCall) Do(f func(context.Context, *config.Config) error) *MockBrokerSetConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerSetConfigCall) DoAndReturn(f func(context.Context, *config.Config) error) *MockBrokerSetConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageProvider mocks base method.
func (m *MockBroker) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockBrokerMockRecorder) StorageProvider(arg0 any) *MockBrokerStorageProviderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockBroker)(nil).StorageProvider), arg0)
	return &MockBrokerStorageProviderCall{Call: call}
}

// MockBrokerStorageProviderCall wrap *gomock.Call
type MockBrokerStorageProviderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerStorageProviderCall) Return(arg0 storage.Provider, arg1 error) *MockBrokerStorageProviderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerStorageProviderCall) Do(f func(storage.ProviderType) (storage.Provider, error)) *MockBrokerStorageProviderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerStorageProviderCall) DoAndReturn(f func(storage.ProviderType) (storage.Provider, error)) *MockBrokerStorageProviderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageProviderTypes mocks base method.
func (m *MockBroker) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockBrokerMockRecorder) StorageProviderTypes() *MockBrokerStorageProviderTypesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockBroker)(nil).StorageProviderTypes))
	return &MockBrokerStorageProviderTypesCall{Call: call}
}

// MockBrokerStorageProviderTypesCall wrap *gomock.Call
type MockBrokerStorageProviderTypesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerStorageProviderTypesCall) Return(arg0 []storage.ProviderType, arg1 error) *MockBrokerStorageProviderTypesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerStorageProviderTypesCall) Do(f func() ([]storage.ProviderType, error)) *MockBrokerStorageProviderTypesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerStorageProviderTypesCall) DoAndReturn(f func() ([]storage.ProviderType, error)) *MockBrokerStorageProviderTypesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Units mocks base method.
func (m *MockBroker) Units(arg0 context.Context, arg1 string) ([]caas.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units", arg0, arg1)
	ret0, _ := ret[0].([]caas.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockBrokerMockRecorder) Units(arg0, arg1 any) *MockBrokerUnitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockBroker)(nil).Units), arg0, arg1)
	return &MockBrokerUnitsCall{Call: call}
}

// MockBrokerUnitsCall wrap *gomock.Call
type MockBrokerUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerUnitsCall) Return(arg0 []caas.Unit, arg1 error) *MockBrokerUnitsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerUnitsCall) Do(f func(context.Context, string) ([]caas.Unit, error)) *MockBrokerUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerUnitsCall) DoAndReturn(f func(context.Context, string) ([]caas.Unit, error)) *MockBrokerUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Upgrade mocks base method.
func (m *MockBroker) Upgrade(arg0 context.Context, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockBrokerMockRecorder) Upgrade(arg0, arg1, arg2 any) *MockBrokerUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockBroker)(nil).Upgrade), arg0, arg1, arg2)
	return &MockBrokerUpgradeCall{Call: call}
}

// MockBrokerUpgradeCall wrap *gomock.Call
type MockBrokerUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerUpgradeCall) Return(arg0 error) *MockBrokerUpgradeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerUpgradeCall) Do(f func(context.Context, string, version.Number) error) *MockBrokerUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerUpgradeCall) DoAndReturn(f func(context.Context, string, version.Number) error) *MockBrokerUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ValidateStorageClass mocks base method.
func (m *MockBroker) ValidateStorageClass(arg0 context.Context, arg1 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStorageClass", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStorageClass indicates an expected call of ValidateStorageClass.
func (mr *MockBrokerMockRecorder) ValidateStorageClass(arg0, arg1 any) *MockBrokerValidateStorageClassCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStorageClass", reflect.TypeOf((*MockBroker)(nil).ValidateStorageClass), arg0, arg1)
	return &MockBrokerValidateStorageClassCall{Call: call}
}

// MockBrokerValidateStorageClassCall wrap *gomock.Call
type MockBrokerValidateStorageClassCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerValidateStorageClassCall) Return(arg0 error) *MockBrokerValidateStorageClassCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerValidateStorageClassCall) Do(f func(context.Context, map[string]any) error) *MockBrokerValidateStorageClassCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerValidateStorageClassCall) DoAndReturn(f func(context.Context, map[string]any) error) *MockBrokerValidateStorageClassCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Version mocks base method.
func (m *MockBroker) Version() (*version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockBrokerMockRecorder) Version() *MockBrokerVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockBroker)(nil).Version))
	return &MockBrokerVersionCall{Call: call}
}

// MockBrokerVersionCall wrap *gomock.Call
type MockBrokerVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerVersionCall) Return(arg0 *version.Number, arg1 error) *MockBrokerVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerVersionCall) Do(f func() (*version.Number, error)) *MockBrokerVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerVersionCall) DoAndReturn(f func() (*version.Number, error)) *MockBrokerVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
