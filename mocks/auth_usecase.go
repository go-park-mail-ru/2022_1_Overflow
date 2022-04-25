// Code generated by MockGen. DO NOT EDIT.
// Source: OverflowBackend/services/auth (interfaces: AuthServiceInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	config "OverflowBackend/internal/config"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	utils_proto "OverflowBackend/proto/utils_proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthServiceInterface is a mock of AuthServiceInterface interface.
type MockAuthServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceInterfaceMockRecorder
}

// MockAuthServiceInterfaceMockRecorder is the mock recorder for MockAuthServiceInterface.
type MockAuthServiceInterfaceMockRecorder struct {
	mock *MockAuthServiceInterface
}

// NewMockAuthServiceInterface creates a new mock instance.
func NewMockAuthServiceInterface(ctrl *gomock.Controller) *MockAuthServiceInterface {
	mock := &MockAuthServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAuthServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceInterface) EXPECT() *MockAuthServiceInterfaceMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockAuthServiceInterface) Init(arg0 *config.Config, arg1 repository_proto.DatabaseRepositoryClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0, arg1)
}

// Init indicates an expected call of Init.
func (mr *MockAuthServiceInterfaceMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAuthServiceInterface)(nil).Init), arg0, arg1)
}

// SignIn mocks base method.
func (m *MockAuthServiceInterface) SignIn(arg0 *auth_proto.SignInForm) *utils_proto.JsonResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	return ret0
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthServiceInterfaceMockRecorder) SignIn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthServiceInterface)(nil).SignIn), arg0)
}

// SignUp mocks base method.
func (m *MockAuthServiceInterface) SignUp(arg0 *auth_proto.SignUpForm) *utils_proto.JsonResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthServiceInterfaceMockRecorder) SignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthServiceInterface)(nil).SignUp), arg0)
}
