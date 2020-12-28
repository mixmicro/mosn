// Code generated by MockGen. DO NOT EDIT.
// Source: ./manager.go

// Package streamfilter is a generated GoMock package.
package streamfilter

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStreamFilterManager is a mock of StreamFilterManager interface.
type MockStreamFilterManager struct {
	ctrl     *gomock.Controller
	recorder *MockStreamFilterManagerMockRecorder
}

// MockStreamFilterManagerMockRecorder is the mock recorder for MockStreamFilterManager.
type MockStreamFilterManagerMockRecorder struct {
	mock *MockStreamFilterManager
}

// NewMockStreamFilterManager creates a new mock instance.
func NewMockStreamFilterManager(ctrl *gomock.Controller) *MockStreamFilterManager {
	mock := &MockStreamFilterManager{ctrl: ctrl}
	mock.recorder = &MockStreamFilterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamFilterManager) EXPECT() *MockStreamFilterManagerMockRecorder {
	return m.recorder
}

// AddOrUpdateStreamFilterConfig mocks base method.
func (m *MockStreamFilterManager) AddOrUpdateStreamFilterConfig(key string, config StreamFiltersConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateStreamFilterConfig", key, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateStreamFilterConfig indicates an expected call of AddOrUpdateStreamFilterConfig.
func (mr *MockStreamFilterManagerMockRecorder) AddOrUpdateStreamFilterConfig(key, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateStreamFilterConfig", reflect.TypeOf((*MockStreamFilterManager)(nil).AddOrUpdateStreamFilterConfig), key, config)
}

// GetStreamFilterFactory mocks base method.
func (m *MockStreamFilterManager) GetStreamFilterFactory(key string) StreamFilterFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamFilterFactory", key)
	ret0, _ := ret[0].(StreamFilterFactory)
	return ret0
}

// GetStreamFilterFactory indicates an expected call of GetStreamFilterFactory.
func (mr *MockStreamFilterManagerMockRecorder) GetStreamFilterFactory(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamFilterFactory", reflect.TypeOf((*MockStreamFilterManager)(nil).GetStreamFilterFactory), key)
}
