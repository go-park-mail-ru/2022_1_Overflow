// Code generated by MockGen. DO NOT EDIT.
// Source: OverflowBackend/services/mailbox (interfaces: MailBoxServiceInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	config "OverflowBackend/internal/config"
	mailbox_proto "OverflowBackend/proto/mailbox_proto"
	profile_proto "OverflowBackend/proto/profile_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	utils_proto "OverflowBackend/proto/utils_proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMailBoxServiceInterface is a mock of MailBoxServiceInterface interface.
type MockMailBoxServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMailBoxServiceInterfaceMockRecorder
}

// MockMailBoxServiceInterfaceMockRecorder is the mock recorder for MockMailBoxServiceInterface.
type MockMailBoxServiceInterfaceMockRecorder struct {
	mock *MockMailBoxServiceInterface
}

// NewMockMailBoxServiceInterface creates a new mock instance.
func NewMockMailBoxServiceInterface(ctrl *gomock.Controller) *MockMailBoxServiceInterface {
	mock := &MockMailBoxServiceInterface{ctrl: ctrl}
	mock.recorder = &MockMailBoxServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailBoxServiceInterface) EXPECT() *MockMailBoxServiceInterfaceMockRecorder {
	return m.recorder
}

// DeleteMail mocks base method.
func (m *MockMailBoxServiceInterface) DeleteMail(arg0 *mailbox_proto.DeleteMailRequest) *utils_proto.JsonResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMail", arg0)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	return ret0
}

// DeleteMail indicates an expected call of DeleteMail.
func (mr *MockMailBoxServiceInterfaceMockRecorder) DeleteMail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMail", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).DeleteMail), arg0)
}

// GetMail mocks base method.
func (m *MockMailBoxServiceInterface) GetMail(arg0 *mailbox_proto.GetMailRequest) *mailbox_proto.ResponseMail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMail", arg0)
	ret0, _ := ret[0].(*mailbox_proto.ResponseMail)
	return ret0
}

// GetMail indicates an expected call of GetMail.
func (mr *MockMailBoxServiceInterfaceMockRecorder) GetMail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMail", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).GetMail), arg0)
}

// Income mocks base method.
func (m *MockMailBoxServiceInterface) Income(arg0 *utils_proto.Session) *mailbox_proto.ResponseMails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Income", arg0)
	ret0, _ := ret[0].(*mailbox_proto.ResponseMails)
	return ret0
}

// Income indicates an expected call of Income.
func (mr *MockMailBoxServiceInterfaceMockRecorder) Income(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Income", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).Income), arg0)
}

// Init mocks base method.
func (m *MockMailBoxServiceInterface) Init(arg0 *config.Config, arg1 repository_proto.DatabaseRepositoryClient, arg2 profile_proto.ProfileClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0, arg1, arg2)
}

// Init indicates an expected call of Init.
func (mr *MockMailBoxServiceInterfaceMockRecorder) Init(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).Init), arg0, arg1, arg2)
}

// Outcome mocks base method.
func (m *MockMailBoxServiceInterface) Outcome(arg0 *utils_proto.Session) *mailbox_proto.ResponseMails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Outcome", arg0)
	ret0, _ := ret[0].(*mailbox_proto.ResponseMails)
	return ret0
}

// Outcome indicates an expected call of Outcome.
func (mr *MockMailBoxServiceInterfaceMockRecorder) Outcome(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Outcome", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).Outcome), arg0)
}

// ReadMail mocks base method.
func (m *MockMailBoxServiceInterface) ReadMail(arg0 *mailbox_proto.ReadMailRequest) *utils_proto.JsonResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMail", arg0)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	return ret0
}

// ReadMail indicates an expected call of ReadMail.
func (mr *MockMailBoxServiceInterfaceMockRecorder) ReadMail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMail", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).ReadMail), arg0)
}

// SendMail mocks base method.
func (m *MockMailBoxServiceInterface) SendMail(arg0 *mailbox_proto.SendMailRequest) *utils_proto.JsonResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMail", arg0)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	return ret0
}

// SendMail indicates an expected call of SendMail.
func (mr *MockMailBoxServiceInterfaceMockRecorder) SendMail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMail", reflect.TypeOf((*MockMailBoxServiceInterface)(nil).SendMail), arg0)
}
