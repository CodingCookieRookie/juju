// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/application/modelmigration (interfaces: ImportService,ExportService)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/application/modelmigration ImportService,ExportService
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/juju/core/charm"
	config "github.com/juju/juju/core/config"
	constraints "github.com/juju/juju/core/constraints"
	status "github.com/juju/juju/core/status"
	unit "github.com/juju/juju/core/unit"
	application "github.com/juju/juju/domain/application"
	charm0 "github.com/juju/juju/domain/application/charm"
	service "github.com/juju/juju/domain/application/service"
	charm1 "github.com/juju/juju/internal/charm"
	gomock "go.uber.org/mock/gomock"
)

// MockImportService is a mock of ImportService interface.
type MockImportService struct {
	ctrl     *gomock.Controller
	recorder *MockImportServiceMockRecorder
}

// MockImportServiceMockRecorder is the mock recorder for MockImportService.
type MockImportServiceMockRecorder struct {
	mock *MockImportService
}

// NewMockImportService creates a new mock instance.
func NewMockImportService(ctrl *gomock.Controller) *MockImportService {
	mock := &MockImportService{ctrl: ctrl}
	mock.recorder = &MockImportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImportService) EXPECT() *MockImportServiceMockRecorder {
	return m.recorder
}

// ImportApplication mocks base method.
func (m *MockImportService) ImportApplication(arg0 context.Context, arg1 string, arg2 service.ImportApplicationArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportApplication indicates an expected call of ImportApplication.
func (mr *MockImportServiceMockRecorder) ImportApplication(arg0, arg1, arg2 any) *MockImportServiceImportApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportApplication", reflect.TypeOf((*MockImportService)(nil).ImportApplication), arg0, arg1, arg2)
	return &MockImportServiceImportApplicationCall{Call: call}
}

// MockImportServiceImportApplicationCall wrap *gomock.Call
type MockImportServiceImportApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceImportApplicationCall) Return(arg0 error) *MockImportServiceImportApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceImportApplicationCall) Do(f func(context.Context, string, service.ImportApplicationArgs) error) *MockImportServiceImportApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceImportApplicationCall) DoAndReturn(f func(context.Context, string, service.ImportApplicationArgs) error) *MockImportServiceImportApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveImportedApplication mocks base method.
func (m *MockImportService) RemoveImportedApplication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImportedApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImportedApplication indicates an expected call of RemoveImportedApplication.
func (mr *MockImportServiceMockRecorder) RemoveImportedApplication(arg0, arg1 any) *MockImportServiceRemoveImportedApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImportedApplication", reflect.TypeOf((*MockImportService)(nil).RemoveImportedApplication), arg0, arg1)
	return &MockImportServiceRemoveImportedApplicationCall{Call: call}
}

// MockImportServiceRemoveImportedApplicationCall wrap *gomock.Call
type MockImportServiceRemoveImportedApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceRemoveImportedApplicationCall) Return(arg0 error) *MockImportServiceRemoveImportedApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceRemoveImportedApplicationCall) Do(f func(context.Context, string) error) *MockImportServiceRemoveImportedApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceRemoveImportedApplicationCall) DoAndReturn(f func(context.Context, string) error) *MockImportServiceRemoveImportedApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockExportService is a mock of ExportService interface.
type MockExportService struct {
	ctrl     *gomock.Controller
	recorder *MockExportServiceMockRecorder
}

// MockExportServiceMockRecorder is the mock recorder for MockExportService.
type MockExportServiceMockRecorder struct {
	mock *MockExportService
}

// NewMockExportService creates a new mock instance.
func NewMockExportService(ctrl *gomock.Controller) *MockExportService {
	mock := &MockExportService{ctrl: ctrl}
	mock.recorder = &MockExportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportService) EXPECT() *MockExportServiceMockRecorder {
	return m.recorder
}

// GetApplicationConfigAndSettings mocks base method.
func (m *MockExportService) GetApplicationConfigAndSettings(arg0 context.Context, arg1 string) (config.ConfigAttributes, application.ApplicationSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationConfigAndSettings", arg0, arg1)
	ret0, _ := ret[0].(config.ConfigAttributes)
	ret1, _ := ret[1].(application.ApplicationSettings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApplicationConfigAndSettings indicates an expected call of GetApplicationConfigAndSettings.
func (mr *MockExportServiceMockRecorder) GetApplicationConfigAndSettings(arg0, arg1 any) *MockExportServiceGetApplicationConfigAndSettingsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationConfigAndSettings", reflect.TypeOf((*MockExportService)(nil).GetApplicationConfigAndSettings), arg0, arg1)
	return &MockExportServiceGetApplicationConfigAndSettingsCall{Call: call}
}

// MockExportServiceGetApplicationConfigAndSettingsCall wrap *gomock.Call
type MockExportServiceGetApplicationConfigAndSettingsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetApplicationConfigAndSettingsCall) Return(arg0 config.ConfigAttributes, arg1 application.ApplicationSettings, arg2 error) *MockExportServiceGetApplicationConfigAndSettingsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetApplicationConfigAndSettingsCall) Do(f func(context.Context, string) (config.ConfigAttributes, application.ApplicationSettings, error)) *MockExportServiceGetApplicationConfigAndSettingsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetApplicationConfigAndSettingsCall) DoAndReturn(f func(context.Context, string) (config.ConfigAttributes, application.ApplicationSettings, error)) *MockExportServiceGetApplicationConfigAndSettingsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationConstraints mocks base method.
func (m *MockExportService) GetApplicationConstraints(arg0 context.Context, arg1 string) (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationConstraints", arg0, arg1)
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationConstraints indicates an expected call of GetApplicationConstraints.
func (mr *MockExportServiceMockRecorder) GetApplicationConstraints(arg0, arg1 any) *MockExportServiceGetApplicationConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationConstraints", reflect.TypeOf((*MockExportService)(nil).GetApplicationConstraints), arg0, arg1)
	return &MockExportServiceGetApplicationConstraintsCall{Call: call}
}

// MockExportServiceGetApplicationConstraintsCall wrap *gomock.Call
type MockExportServiceGetApplicationConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetApplicationConstraintsCall) Return(arg0 constraints.Value, arg1 error) *MockExportServiceGetApplicationConstraintsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetApplicationConstraintsCall) Do(f func(context.Context, string) (constraints.Value, error)) *MockExportServiceGetApplicationConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetApplicationConstraintsCall) DoAndReturn(f func(context.Context, string) (constraints.Value, error)) *MockExportServiceGetApplicationConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationStatus mocks base method.
func (m *MockExportService) GetApplicationStatus(arg0 context.Context, arg1 string) (*status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationStatus", arg0, arg1)
	ret0, _ := ret[0].(*status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationStatus indicates an expected call of GetApplicationStatus.
func (mr *MockExportServiceMockRecorder) GetApplicationStatus(arg0, arg1 any) *MockExportServiceGetApplicationStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationStatus", reflect.TypeOf((*MockExportService)(nil).GetApplicationStatus), arg0, arg1)
	return &MockExportServiceGetApplicationStatusCall{Call: call}
}

// MockExportServiceGetApplicationStatusCall wrap *gomock.Call
type MockExportServiceGetApplicationStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetApplicationStatusCall) Return(arg0 *status.StatusInfo, arg1 error) *MockExportServiceGetApplicationStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetApplicationStatusCall) Do(f func(context.Context, string) (*status.StatusInfo, error)) *MockExportServiceGetApplicationStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetApplicationStatusCall) DoAndReturn(f func(context.Context, string) (*status.StatusInfo, error)) *MockExportServiceGetApplicationStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmByApplicationName mocks base method.
func (m *MockExportService) GetCharmByApplicationName(arg0 context.Context, arg1 string) (charm1.Charm, charm0.CharmLocator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmByApplicationName", arg0, arg1)
	ret0, _ := ret[0].(charm1.Charm)
	ret1, _ := ret[1].(charm0.CharmLocator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCharmByApplicationName indicates an expected call of GetCharmByApplicationName.
func (mr *MockExportServiceMockRecorder) GetCharmByApplicationName(arg0, arg1 any) *MockExportServiceGetCharmByApplicationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmByApplicationName", reflect.TypeOf((*MockExportService)(nil).GetCharmByApplicationName), arg0, arg1)
	return &MockExportServiceGetCharmByApplicationNameCall{Call: call}
}

// MockExportServiceGetCharmByApplicationNameCall wrap *gomock.Call
type MockExportServiceGetCharmByApplicationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetCharmByApplicationNameCall) Return(arg0 charm1.Charm, arg1 charm0.CharmLocator, arg2 error) *MockExportServiceGetCharmByApplicationNameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetCharmByApplicationNameCall) Do(f func(context.Context, string) (charm1.Charm, charm0.CharmLocator, error)) *MockExportServiceGetCharmByApplicationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetCharmByApplicationNameCall) DoAndReturn(f func(context.Context, string) (charm1.Charm, charm0.CharmLocator, error)) *MockExportServiceGetCharmByApplicationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmID mocks base method.
func (m *MockExportService) GetCharmID(arg0 context.Context, arg1 charm0.GetCharmArgs) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmID", arg0, arg1)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmID indicates an expected call of GetCharmID.
func (mr *MockExportServiceMockRecorder) GetCharmID(arg0, arg1 any) *MockExportServiceGetCharmIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmID", reflect.TypeOf((*MockExportService)(nil).GetCharmID), arg0, arg1)
	return &MockExportServiceGetCharmIDCall{Call: call}
}

// MockExportServiceGetCharmIDCall wrap *gomock.Call
type MockExportServiceGetCharmIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetCharmIDCall) Return(arg0 charm.ID, arg1 error) *MockExportServiceGetCharmIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetCharmIDCall) Do(f func(context.Context, charm0.GetCharmArgs) (charm.ID, error)) *MockExportServiceGetCharmIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetCharmIDCall) DoAndReturn(f func(context.Context, charm0.GetCharmArgs) (charm.ID, error)) *MockExportServiceGetCharmIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitWorkloadStatus mocks base method.
func (m *MockExportService) GetUnitWorkloadStatus(arg0 context.Context, arg1 unit.Name) (*status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitWorkloadStatus", arg0, arg1)
	ret0, _ := ret[0].(*status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitWorkloadStatus indicates an expected call of GetUnitWorkloadStatus.
func (mr *MockExportServiceMockRecorder) GetUnitWorkloadStatus(arg0, arg1 any) *MockExportServiceGetUnitWorkloadStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitWorkloadStatus", reflect.TypeOf((*MockExportService)(nil).GetUnitWorkloadStatus), arg0, arg1)
	return &MockExportServiceGetUnitWorkloadStatusCall{Call: call}
}

// MockExportServiceGetUnitWorkloadStatusCall wrap *gomock.Call
type MockExportServiceGetUnitWorkloadStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetUnitWorkloadStatusCall) Return(arg0 *status.StatusInfo, arg1 error) *MockExportServiceGetUnitWorkloadStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetUnitWorkloadStatusCall) Do(f func(context.Context, unit.Name) (*status.StatusInfo, error)) *MockExportServiceGetUnitWorkloadStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetUnitWorkloadStatusCall) DoAndReturn(f func(context.Context, unit.Name) (*status.StatusInfo, error)) *MockExportServiceGetUnitWorkloadStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
