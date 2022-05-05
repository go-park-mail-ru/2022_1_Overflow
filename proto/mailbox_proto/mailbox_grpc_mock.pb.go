// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: proto/mailbox.proto

package mailbox_proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	utils_proto "OverflowBackend/proto/utils_proto"
)

// MockMailboxClient is a mock of MailboxClient interface.
type MockMailboxClient struct {
	ctrl     *gomock.Controller
	recorder *MockMailboxClientMockRecorder
}

// MockMailboxClientMockRecorder is the mock recorder for MockMailboxClient.
type MockMailboxClientMockRecorder struct {
	mock *MockMailboxClient
}

// NewMockMailboxClient creates a new mock instance.
func NewMockMailboxClient(ctrl *gomock.Controller) *MockMailboxClient {
	mock := &MockMailboxClient{ctrl: ctrl}
	mock.recorder = &MockMailboxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailboxClient) EXPECT() *MockMailboxClientMockRecorder {
	return m.recorder
}

// DeleteMail mocks base method.
func (m *MockMailboxClient) DeleteMail(ctx context.Context, in *DeleteMailRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMail", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMail indicates an expected call of DeleteMail.
func (mr *MockMailboxClientMockRecorder) DeleteMail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMail", reflect.TypeOf((*MockMailboxClient)(nil).DeleteMail), varargs...)
}

// GetMail mocks base method.
func (m *MockMailboxClient) GetMail(ctx context.Context, in *GetMailRequest, opts ...grpc.CallOption) (*ResponseMail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMail", varargs...)
	ret0, _ := ret[0].(*ResponseMail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMail indicates an expected call of GetMail.
func (mr *MockMailboxClientMockRecorder) GetMail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMail", reflect.TypeOf((*MockMailboxClient)(nil).GetMail), varargs...)
}

// Income mocks base method.
func (m *MockMailboxClient) Income(ctx context.Context, in *IncomeRequest, opts ...grpc.CallOption) (*ResponseMails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Income", varargs...)
	ret0, _ := ret[0].(*ResponseMails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Income indicates an expected call of Income.
func (mr *MockMailboxClientMockRecorder) Income(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Income", reflect.TypeOf((*MockMailboxClient)(nil).Income), varargs...)
}

// Outcome mocks base method.
func (m *MockMailboxClient) Outcome(ctx context.Context, in *OutcomeRequest, opts ...grpc.CallOption) (*ResponseMails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Outcome", varargs...)
	ret0, _ := ret[0].(*ResponseMails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Outcome indicates an expected call of Outcome.
func (mr *MockMailboxClientMockRecorder) Outcome(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Outcome", reflect.TypeOf((*MockMailboxClient)(nil).Outcome), varargs...)
}

// ReadMail mocks base method.
func (m *MockMailboxClient) ReadMail(ctx context.Context, in *ReadMailRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadMail", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMail indicates an expected call of ReadMail.
func (mr *MockMailboxClientMockRecorder) ReadMail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMail", reflect.TypeOf((*MockMailboxClient)(nil).ReadMail), varargs...)
}

// SendMail mocks base method.
func (m *MockMailboxClient) SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMail", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMail indicates an expected call of SendMail.
func (mr *MockMailboxClientMockRecorder) SendMail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMail", reflect.TypeOf((*MockMailboxClient)(nil).SendMail), varargs...)
}

// MockMailboxServer is a mock of MailboxServer interface.
type MockMailboxServer struct {
	ctrl     *gomock.Controller
	recorder *MockMailboxServerMockRecorder
}

// MockMailboxServerMockRecorder is the mock recorder for MockMailboxServer.
type MockMailboxServerMockRecorder struct {
	mock *MockMailboxServer
}

// NewMockMailboxServer creates a new mock instance.
func NewMockMailboxServer(ctrl *gomock.Controller) *MockMailboxServer {
	mock := &MockMailboxServer{ctrl: ctrl}
	mock.recorder = &MockMailboxServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailboxServer) EXPECT() *MockMailboxServerMockRecorder {
	return m.recorder
}

// DeleteMail mocks base method.
func (m *MockMailboxServer) DeleteMail(arg0 context.Context, arg1 *DeleteMailRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMail", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMail indicates an expected call of DeleteMail.
func (mr *MockMailboxServerMockRecorder) DeleteMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMail", reflect.TypeOf((*MockMailboxServer)(nil).DeleteMail), arg0, arg1)
}

// GetMail mocks base method.
func (m *MockMailboxServer) GetMail(arg0 context.Context, arg1 *GetMailRequest) (*ResponseMail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMail", arg0, arg1)
	ret0, _ := ret[0].(*ResponseMail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMail indicates an expected call of GetMail.
func (mr *MockMailboxServerMockRecorder) GetMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMail", reflect.TypeOf((*MockMailboxServer)(nil).GetMail), arg0, arg1)
}

// Income mocks base method.
func (m *MockMailboxServer) Income(arg0 context.Context, arg1 *IncomeRequest) (*ResponseMails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Income", arg0, arg1)
	ret0, _ := ret[0].(*ResponseMails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Income indicates an expected call of Income.
func (mr *MockMailboxServerMockRecorder) Income(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Income", reflect.TypeOf((*MockMailboxServer)(nil).Income), arg0, arg1)
}

// Outcome mocks base method.
func (m *MockMailboxServer) Outcome(arg0 context.Context, arg1 *OutcomeRequest) (*ResponseMails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Outcome", arg0, arg1)
	ret0, _ := ret[0].(*ResponseMails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Outcome indicates an expected call of Outcome.
func (mr *MockMailboxServerMockRecorder) Outcome(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Outcome", reflect.TypeOf((*MockMailboxServer)(nil).Outcome), arg0, arg1)
}

// ReadMail mocks base method.
func (m *MockMailboxServer) ReadMail(arg0 context.Context, arg1 *ReadMailRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMail", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMail indicates an expected call of ReadMail.
func (mr *MockMailboxServerMockRecorder) ReadMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMail", reflect.TypeOf((*MockMailboxServer)(nil).ReadMail), arg0, arg1)
}

// SendMail mocks base method.
func (m *MockMailboxServer) SendMail(arg0 context.Context, arg1 *SendMailRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMail", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMail indicates an expected call of SendMail.
func (mr *MockMailboxServerMockRecorder) SendMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMail", reflect.TypeOf((*MockMailboxServer)(nil).SendMail), arg0, arg1)
}

// mustEmbedUnimplementedMailboxServer mocks base method.
func (m *MockMailboxServer) mustEmbedUnimplementedMailboxServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMailboxServer")
}

// mustEmbedUnimplementedMailboxServer indicates an expected call of mustEmbedUnimplementedMailboxServer.
func (mr *MockMailboxServerMockRecorder) mustEmbedUnimplementedMailboxServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMailboxServer", reflect.TypeOf((*MockMailboxServer)(nil).mustEmbedUnimplementedMailboxServer))
}

// MockUnsafeMailboxServer is a mock of UnsafeMailboxServer interface.
type MockUnsafeMailboxServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeMailboxServerMockRecorder
}

// MockUnsafeMailboxServerMockRecorder is the mock recorder for MockUnsafeMailboxServer.
type MockUnsafeMailboxServerMockRecorder struct {
	mock *MockUnsafeMailboxServer
}

// NewMockUnsafeMailboxServer creates a new mock instance.
func NewMockUnsafeMailboxServer(ctrl *gomock.Controller) *MockUnsafeMailboxServer {
	mock := &MockUnsafeMailboxServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeMailboxServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeMailboxServer) EXPECT() *MockUnsafeMailboxServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedMailboxServer mocks base method.
func (m *MockUnsafeMailboxServer) mustEmbedUnimplementedMailboxServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMailboxServer")
}

// mustEmbedUnimplementedMailboxServer indicates an expected call of mustEmbedUnimplementedMailboxServer.
func (mr *MockUnsafeMailboxServerMockRecorder) mustEmbedUnimplementedMailboxServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMailboxServer", reflect.TypeOf((*MockUnsafeMailboxServer)(nil).mustEmbedUnimplementedMailboxServer))
}