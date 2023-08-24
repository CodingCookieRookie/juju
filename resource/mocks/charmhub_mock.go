// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/resource (interfaces: CharmHub)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	url "net/url"
	reflect "reflect"

	charmhub "github.com/juju/juju/internal/charmhub"
	transport "github.com/juju/juju/internal/charmhub/transport"
	gomock "go.uber.org/mock/gomock"
)

// MockCharmHub is a mock of CharmHub interface.
type MockCharmHub struct {
	ctrl     *gomock.Controller
	recorder *MockCharmHubMockRecorder
}

// MockCharmHubMockRecorder is the mock recorder for MockCharmHub.
type MockCharmHubMockRecorder struct {
	mock *MockCharmHub
}

// NewMockCharmHub creates a new mock instance.
func NewMockCharmHub(ctrl *gomock.Controller) *MockCharmHub {
	mock := &MockCharmHub{ctrl: ctrl}
	mock.recorder = &MockCharmHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmHub) EXPECT() *MockCharmHubMockRecorder {
	return m.recorder
}

// DownloadResource mocks base method.
func (m *MockCharmHub) DownloadResource(arg0 context.Context, arg1 *url.URL) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadResource", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadResource indicates an expected call of DownloadResource.
func (mr *MockCharmHubMockRecorder) DownloadResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadResource", reflect.TypeOf((*MockCharmHub)(nil).DownloadResource), arg0, arg1)
}

// Refresh mocks base method.
func (m *MockCharmHub) Refresh(arg0 context.Context, arg1 charmhub.RefreshConfig) ([]transport.RefreshResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0, arg1)
	ret0, _ := ret[0].([]transport.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockCharmHubMockRecorder) Refresh(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockCharmHub)(nil).Refresh), arg0, arg1)
}
