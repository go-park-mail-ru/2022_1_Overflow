// Code generated by MockGen. DO NOT EDIT.
// Source: proto/attach_proto/attach_grpc.pb.go

// Package mock_attach_proto is a generated GoMock package.
package attach_proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAttachClient is a mock of AttachClient interface.
type MockAttachClient struct {
	ctrl     *gomock.Controller
	recorder *MockAttachClientMockRecorder
}

// MockAttachClientMockRecorder is the mock recorder for MockAttachClient.
type MockAttachClientMockRecorder struct {
	mock *MockAttachClient
}

// NewMockAttachClient creates a new mock instance.
func NewMockAttachClient(ctrl *gomock.Controller) *MockAttachClient {
	mock := &MockAttachClient{ctrl: ctrl}
	mock.recorder = &MockAttachClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttachClient) EXPECT() *MockAttachClientMockRecorder {
	return m.recorder
}

// CheckAttachPermission mocks base method.
func (m *MockAttachClient) CheckAttachPermission(ctx context.Context, in *AttachPermissionRequest, opts ...grpc.CallOption) (*AttachPermissionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckAttachPermission", varargs...)
	ret0, _ := ret[0].(*AttachPermissionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAttachPermission indicates an expected call of CheckAttachPermission.
func (mr *MockAttachClientMockRecorder) CheckAttachPermission(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAttachPermission", reflect.TypeOf((*MockAttachClient)(nil).CheckAttachPermission), varargs...)
}

// GetAttach mocks base method.
func (m *MockAttachClient) GetAttach(ctx context.Context, in *GetAttachRequest, opts ...grpc.CallOption) (*AttachResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAttach", varargs...)
	ret0, _ := ret[0].(*AttachResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttach indicates an expected call of GetAttach.
func (mr *MockAttachClientMockRecorder) GetAttach(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttach", reflect.TypeOf((*MockAttachClient)(nil).GetAttach), varargs...)
}

// ListAttach mocks base method.
func (m *MockAttachClient) ListAttach(ctx context.Context, in *GetAttachRequest, opts ...grpc.CallOption) (*AttachListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttach", varargs...)
	ret0, _ := ret[0].(*AttachListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttach indicates an expected call of ListAttach.
func (mr *MockAttachClientMockRecorder) ListAttach(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttach", reflect.TypeOf((*MockAttachClient)(nil).ListAttach), varargs...)
}

// SaveAttach mocks base method.
func (m *MockAttachClient) SaveAttach(ctx context.Context, in *SaveAttachRequest, opts ...grpc.CallOption) (*Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveAttach", varargs...)
	ret0, _ := ret[0].(*Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveAttach indicates an expected call of SaveAttach.
func (mr *MockAttachClientMockRecorder) SaveAttach(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAttach", reflect.TypeOf((*MockAttachClient)(nil).SaveAttach), varargs...)
}

// MockAttachServer is a mock of AttachServer interface.
type MockAttachServer struct {
	ctrl     *gomock.Controller
	recorder *MockAttachServerMockRecorder
}

// MockAttachServerMockRecorder is the mock recorder for MockAttachServer.
type MockAttachServerMockRecorder struct {
	mock *MockAttachServer
}

// NewMockAttachServer creates a new mock instance.
func NewMockAttachServer(ctrl *gomock.Controller) *MockAttachServer {
	mock := &MockAttachServer{ctrl: ctrl}
	mock.recorder = &MockAttachServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttachServer) EXPECT() *MockAttachServerMockRecorder {
	return m.recorder
}

// CheckAttachPermission mocks base method.
func (m *MockAttachServer) CheckAttachPermission(arg0 context.Context, arg1 *AttachPermissionRequest) (*AttachPermissionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAttachPermission", arg0, arg1)
	ret0, _ := ret[0].(*AttachPermissionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAttachPermission indicates an expected call of CheckAttachPermission.
func (mr *MockAttachServerMockRecorder) CheckAttachPermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAttachPermission", reflect.TypeOf((*MockAttachServer)(nil).CheckAttachPermission), arg0, arg1)
}

// GetAttach mocks base method.
func (m *MockAttachServer) GetAttach(arg0 context.Context, arg1 *GetAttachRequest) (*AttachResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttach", arg0, arg1)
	ret0, _ := ret[0].(*AttachResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttach indicates an expected call of GetAttach.
func (mr *MockAttachServerMockRecorder) GetAttach(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttach", reflect.TypeOf((*MockAttachServer)(nil).GetAttach), arg0, arg1)
}

// ListAttach mocks base method.
func (m *MockAttachServer) ListAttach(arg0 context.Context, arg1 *GetAttachRequest) (*AttachListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttach", arg0, arg1)
	ret0, _ := ret[0].(*AttachListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttach indicates an expected call of ListAttach.
func (mr *MockAttachServerMockRecorder) ListAttach(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttach", reflect.TypeOf((*MockAttachServer)(nil).ListAttach), arg0, arg1)
}

// SaveAttach mocks base method.
func (m *MockAttachServer) SaveAttach(arg0 context.Context, arg1 *SaveAttachRequest) (*Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAttach", arg0, arg1)
	ret0, _ := ret[0].(*Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveAttach indicates an expected call of SaveAttach.
func (mr *MockAttachServerMockRecorder) SaveAttach(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAttach", reflect.TypeOf((*MockAttachServer)(nil).SaveAttach), arg0, arg1)
}

// mustEmbedUnimplementedAttachServer mocks base method.
func (m *MockAttachServer) mustEmbedUnimplementedAttachServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAttachServer")
}

// mustEmbedUnimplementedAttachServer indicates an expected call of mustEmbedUnimplementedAttachServer.
func (mr *MockAttachServerMockRecorder) mustEmbedUnimplementedAttachServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAttachServer", reflect.TypeOf((*MockAttachServer)(nil).mustEmbedUnimplementedAttachServer))
}

// MockUnsafeAttachServer is a mock of UnsafeAttachServer interface.
type MockUnsafeAttachServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAttachServerMockRecorder
}

// MockUnsafeAttachServerMockRecorder is the mock recorder for MockUnsafeAttachServer.
type MockUnsafeAttachServerMockRecorder struct {
	mock *MockUnsafeAttachServer
}

// NewMockUnsafeAttachServer creates a new mock instance.
func NewMockUnsafeAttachServer(ctrl *gomock.Controller) *MockUnsafeAttachServer {
	mock := &MockUnsafeAttachServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAttachServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAttachServer) EXPECT() *MockUnsafeAttachServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAttachServer mocks base method.
func (m *MockUnsafeAttachServer) mustEmbedUnimplementedAttachServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAttachServer")
}

// mustEmbedUnimplementedAttachServer indicates an expected call of mustEmbedUnimplementedAttachServer.
func (mr *MockUnsafeAttachServerMockRecorder) mustEmbedUnimplementedAttachServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAttachServer", reflect.TypeOf((*MockUnsafeAttachServer)(nil).mustEmbedUnimplementedAttachServer))
}
