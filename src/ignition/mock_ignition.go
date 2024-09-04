// Code generated by MockGen. DO NOT EDIT.
// Source: ignition.go
//
// Generated by this command:
//
//	mockgen -source=ignition.go -package=ignition -destination=mock_ignition.go
//
// Package ignition is a generated GoMock package.
package ignition

import (
	reflect "reflect"

	types "github.com/coreos/ignition/v2/config/v3_2/types"
	gomock "go.uber.org/mock/gomock"
)

// MockIgnition is a mock of Ignition interface.
type MockIgnition struct {
	ctrl     *gomock.Controller
	recorder *MockIgnitionMockRecorder
}

// MockIgnitionMockRecorder is the mock recorder for MockIgnition.
type MockIgnitionMockRecorder struct {
	mock *MockIgnition
}

// NewMockIgnition creates a new mock instance.
func NewMockIgnition(ctrl *gomock.Controller) *MockIgnition {
	mock := &MockIgnition{ctrl: ctrl}
	mock.recorder = &MockIgnitionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIgnition) EXPECT() *MockIgnitionMockRecorder {
	return m.recorder
}

// InjectKubeletTempPrivateKey mocks base method.
func (m *MockIgnition) InjectKubeletTempPrivateKey(pathToSourceIgnition string, privateKeyBytes []byte, certPathToInject string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectKubeletTempPrivateKey", pathToSourceIgnition, privateKeyBytes, certPathToInject)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectKubeletTempPrivateKey indicates an expected call of InjectKubeletTempPrivateKey.
func (mr *MockIgnitionMockRecorder) InjectKubeletTempPrivateKey(pathToSourceIgnition, privateKeyBytes, certPathToInject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectKubeletTempPrivateKey", reflect.TypeOf((*MockIgnition)(nil).InjectKubeletTempPrivateKey), pathToSourceIgnition, privateKeyBytes, certPathToInject)
}

// MergeIgnitionConfig mocks base method.
func (m *MockIgnition) MergeIgnitionConfig(base, overrides *types.Config) (*types.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeIgnitionConfig", base, overrides)
	ret0, _ := ret[0].(*types.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeIgnitionConfig indicates an expected call of MergeIgnitionConfig.
func (mr *MockIgnitionMockRecorder) MergeIgnitionConfig(base, overrides any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeIgnitionConfig", reflect.TypeOf((*MockIgnition)(nil).MergeIgnitionConfig), base, overrides)
}

// ParseIgnitionFile mocks base method.
func (m *MockIgnition) ParseIgnitionFile(path string) (*types.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseIgnitionFile", path)
	ret0, _ := ret[0].(*types.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseIgnitionFile indicates an expected call of ParseIgnitionFile.
func (mr *MockIgnitionMockRecorder) ParseIgnitionFile(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseIgnitionFile", reflect.TypeOf((*MockIgnition)(nil).ParseIgnitionFile), path)
}

// WriteIgnitionFile mocks base method.
func (m *MockIgnition) WriteIgnitionFile(path string, config *types.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteIgnitionFile", path, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteIgnitionFile indicates an expected call of WriteIgnitionFile.
func (mr *MockIgnitionMockRecorder) WriteIgnitionFile(path, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIgnitionFile", reflect.TypeOf((*MockIgnition)(nil).WriteIgnitionFile), path, config)
}
