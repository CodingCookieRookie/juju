// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/blockcommand/modelmigration (interfaces: Coordinator,ImportService,ExportService)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/blockcommand/modelmigration Coordinator,ImportService,ExportService
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"

	modelmigration "github.com/juju/juju/core/modelmigration"
	blockcommand "github.com/juju/juju/domain/blockcommand"
	gomock "go.uber.org/mock/gomock"
)

// MockCoordinator is a mock of Coordinator interface.
type MockCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockCoordinatorMockRecorder
}

// MockCoordinatorMockRecorder is the mock recorder for MockCoordinator.
type MockCoordinatorMockRecorder struct {
	mock *MockCoordinator
}

// NewMockCoordinator creates a new mock instance.
func NewMockCoordinator(ctrl *gomock.Controller) *MockCoordinator {
	mock := &MockCoordinator{ctrl: ctrl}
	mock.recorder = &MockCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoordinator) EXPECT() *MockCoordinatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCoordinator) Add(arg0 modelmigration.Operation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockCoordinatorMockRecorder) Add(arg0 any) *MockCoordinatorAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCoordinator)(nil).Add), arg0)
	return &MockCoordinatorAddCall{Call: call}
}

// MockCoordinatorAddCall wrap *gomock.Call
type MockCoordinatorAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCoordinatorAddCall) Return() *MockCoordinatorAddCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCoordinatorAddCall) Do(f func(modelmigration.Operation)) *MockCoordinatorAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCoordinatorAddCall) DoAndReturn(f func(modelmigration.Operation)) *MockCoordinatorAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

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

// SwitchBlockOn mocks base method.
func (m *MockImportService) SwitchBlockOn(arg0 context.Context, arg1 blockcommand.BlockType, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwitchBlockOn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwitchBlockOn indicates an expected call of SwitchBlockOn.
func (mr *MockImportServiceMockRecorder) SwitchBlockOn(arg0, arg1, arg2 any) *MockImportServiceSwitchBlockOnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchBlockOn", reflect.TypeOf((*MockImportService)(nil).SwitchBlockOn), arg0, arg1, arg2)
	return &MockImportServiceSwitchBlockOnCall{Call: call}
}

// MockImportServiceSwitchBlockOnCall wrap *gomock.Call
type MockImportServiceSwitchBlockOnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceSwitchBlockOnCall) Return(arg0 error) *MockImportServiceSwitchBlockOnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceSwitchBlockOnCall) Do(f func(context.Context, blockcommand.BlockType, string) error) *MockImportServiceSwitchBlockOnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceSwitchBlockOnCall) DoAndReturn(f func(context.Context, blockcommand.BlockType, string) error) *MockImportServiceSwitchBlockOnCall {
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

// GetBlocks mocks base method.
func (m *MockExportService) GetBlocks(arg0 context.Context) ([]blockcommand.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocks", arg0)
	ret0, _ := ret[0].([]blockcommand.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocks indicates an expected call of GetBlocks.
func (mr *MockExportServiceMockRecorder) GetBlocks(arg0 any) *MockExportServiceGetBlocksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocks", reflect.TypeOf((*MockExportService)(nil).GetBlocks), arg0)
	return &MockExportServiceGetBlocksCall{Call: call}
}

// MockExportServiceGetBlocksCall wrap *gomock.Call
type MockExportServiceGetBlocksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetBlocksCall) Return(arg0 []blockcommand.Block, arg1 error) *MockExportServiceGetBlocksCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetBlocksCall) Do(f func(context.Context) ([]blockcommand.Block, error)) *MockExportServiceGetBlocksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetBlocksCall) DoAndReturn(f func(context.Context) ([]blockcommand.Block, error)) *MockExportServiceGetBlocksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}