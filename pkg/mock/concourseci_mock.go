// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/concourseci/concourseci.go

// Package mock_concourseci is a generated GoMock package.
package mock_concourseci

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConcourseCIClientInterface is a mock of ConcourseCIClientInterface interface
type MockConcourseCIClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConcourseCIClientInterfaceMockRecorder
}

// MockConcourseCIClientInterfaceMockRecorder is the mock recorder for MockConcourseCIClientInterface
type MockConcourseCIClientInterfaceMockRecorder struct {
	mock *MockConcourseCIClientInterface
}

// NewMockConcourseCIClientInterface creates a new mock instance
func NewMockConcourseCIClientInterface(ctrl *gomock.Controller) *MockConcourseCIClientInterface {
	mock := &MockConcourseCIClientInterface{ctrl: ctrl}
	mock.recorder = &MockConcourseCIClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConcourseCIClientInterface) EXPECT() *MockConcourseCIClientInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockConcourseCIClientInterface) Login() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login")
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockConcourseCIClientInterfaceMockRecorder) Login() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockConcourseCIClientInterface)(nil).Login))
}

// SetPipeline mocks base method
func (m *MockConcourseCIClientInterface) SetPipeline(target, pipeline, manifest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPipeline", target, pipeline, manifest)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPipeline indicates an expected call of SetPipeline
func (mr *MockConcourseCIClientInterfaceMockRecorder) SetPipeline(target, pipeline, manifest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPipeline", reflect.TypeOf((*MockConcourseCIClientInterface)(nil).SetPipeline), target, pipeline, manifest)
}

// DestroyPipeline mocks base method
func (m *MockConcourseCIClientInterface) DestroyPipeline(target, pipeline string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyPipeline", target, pipeline)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyPipeline indicates an expected call of DestroyPipeline
func (mr *MockConcourseCIClientInterfaceMockRecorder) DestroyPipeline(target, pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyPipeline", reflect.TypeOf((*MockConcourseCIClientInterface)(nil).DestroyPipeline), target, pipeline)
}

// UnpausePipeline mocks base method
func (m *MockConcourseCIClientInterface) UnpausePipeline(target, pipeline string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpausePipeline", target, pipeline)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpausePipeline indicates an expected call of UnpausePipeline
func (mr *MockConcourseCIClientInterfaceMockRecorder) UnpausePipeline(target, pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpausePipeline", reflect.TypeOf((*MockConcourseCIClientInterface)(nil).UnpausePipeline), target, pipeline)
}

// ExposePipeline mocks base method
func (m *MockConcourseCIClientInterface) ExposePipeline(target, pipeline string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExposePipeline", target, pipeline)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExposePipeline indicates an expected call of ExposePipeline
func (mr *MockConcourseCIClientInterfaceMockRecorder) ExposePipeline(target, pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExposePipeline", reflect.TypeOf((*MockConcourseCIClientInterface)(nil).ExposePipeline), target, pipeline)
}
