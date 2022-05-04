// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: proto/folder_manager.proto

package folder_manager_proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	utils_proto "OverflowBackend/proto/utils_proto"
)

// MockFolderManagerClient is a mock of FolderManagerClient interface.
type MockFolderManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockFolderManagerClientMockRecorder
}

// MockFolderManagerClientMockRecorder is the mock recorder for MockFolderManagerClient.
type MockFolderManagerClientMockRecorder struct {
	mock *MockFolderManagerClient
}

// NewMockFolderManagerClient creates a new mock instance.
func NewMockFolderManagerClient(ctrl *gomock.Controller) *MockFolderManagerClient {
	mock := &MockFolderManagerClient{ctrl: ctrl}
	mock.recorder = &MockFolderManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFolderManagerClient) EXPECT() *MockFolderManagerClientMockRecorder {
	return m.recorder
}

// AddFolder mocks base method.
func (m *MockFolderManagerClient) AddFolder(ctx context.Context, in *AddFolderRequest, opts ...grpc.CallOption) (*ResponseFolder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFolder", varargs...)
	ret0, _ := ret[0].(*ResponseFolder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFolder indicates an expected call of AddFolder.
func (mr *MockFolderManagerClientMockRecorder) AddFolder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFolder", reflect.TypeOf((*MockFolderManagerClient)(nil).AddFolder), varargs...)
}

// AddMailToFolder mocks base method.
func (m *MockFolderManagerClient) AddMailToFolder(ctx context.Context, in *AddMailToFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMailToFolder", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMailToFolder indicates an expected call of AddMailToFolder.
func (mr *MockFolderManagerClientMockRecorder) AddMailToFolder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMailToFolder", reflect.TypeOf((*MockFolderManagerClient)(nil).AddMailToFolder), varargs...)
}

// ChangeFolder mocks base method.
func (m *MockFolderManagerClient) ChangeFolder(ctx context.Context, in *ChangeFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeFolder", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeFolder indicates an expected call of ChangeFolder.
func (mr *MockFolderManagerClientMockRecorder) ChangeFolder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeFolder", reflect.TypeOf((*MockFolderManagerClient)(nil).ChangeFolder), varargs...)
}

// DeleteFolder mocks base method.
func (m *MockFolderManagerClient) DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFolder", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFolder indicates an expected call of DeleteFolder.
func (mr *MockFolderManagerClientMockRecorder) DeleteFolder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolder", reflect.TypeOf((*MockFolderManagerClient)(nil).DeleteFolder), varargs...)
}

// DeleteFolderMail mocks base method.
func (m *MockFolderManagerClient) DeleteFolderMail(ctx context.Context, in *DeleteFolderMailRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFolderMail", varargs...)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFolderMail indicates an expected call of DeleteFolderMail.
func (mr *MockFolderManagerClientMockRecorder) DeleteFolderMail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolderMail", reflect.TypeOf((*MockFolderManagerClient)(nil).DeleteFolderMail), varargs...)
}

// ListFolder mocks base method.
func (m *MockFolderManagerClient) ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*ResponseMails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFolder", varargs...)
	ret0, _ := ret[0].(*ResponseMails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFolder indicates an expected call of ListFolder.
func (mr *MockFolderManagerClientMockRecorder) ListFolder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFolder", reflect.TypeOf((*MockFolderManagerClient)(nil).ListFolder), varargs...)
}

// ListFolders mocks base method.
func (m *MockFolderManagerClient) ListFolders(ctx context.Context, in *ListFoldersRequest, opts ...grpc.CallOption) (*ResponseFolders, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFolders", varargs...)
	ret0, _ := ret[0].(*ResponseFolders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFolders indicates an expected call of ListFolders.
func (mr *MockFolderManagerClientMockRecorder) ListFolders(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFolders", reflect.TypeOf((*MockFolderManagerClient)(nil).ListFolders), varargs...)
}

// MockFolderManagerServer is a mock of FolderManagerServer interface.
type MockFolderManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockFolderManagerServerMockRecorder
}

// MockFolderManagerServerMockRecorder is the mock recorder for MockFolderManagerServer.
type MockFolderManagerServerMockRecorder struct {
	mock *MockFolderManagerServer
}

// NewMockFolderManagerServer creates a new mock instance.
func NewMockFolderManagerServer(ctrl *gomock.Controller) *MockFolderManagerServer {
	mock := &MockFolderManagerServer{ctrl: ctrl}
	mock.recorder = &MockFolderManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFolderManagerServer) EXPECT() *MockFolderManagerServerMockRecorder {
	return m.recorder
}

// AddFolder mocks base method.
func (m *MockFolderManagerServer) AddFolder(arg0 context.Context, arg1 *AddFolderRequest) (*ResponseFolder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFolder", arg0, arg1)
	ret0, _ := ret[0].(*ResponseFolder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFolder indicates an expected call of AddFolder.
func (mr *MockFolderManagerServerMockRecorder) AddFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFolder", reflect.TypeOf((*MockFolderManagerServer)(nil).AddFolder), arg0, arg1)
}

// AddMailToFolder mocks base method.
func (m *MockFolderManagerServer) AddMailToFolder(arg0 context.Context, arg1 *AddMailToFolderRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMailToFolder", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMailToFolder indicates an expected call of AddMailToFolder.
func (mr *MockFolderManagerServerMockRecorder) AddMailToFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMailToFolder", reflect.TypeOf((*MockFolderManagerServer)(nil).AddMailToFolder), arg0, arg1)
}

// ChangeFolder mocks base method.
func (m *MockFolderManagerServer) ChangeFolder(arg0 context.Context, arg1 *ChangeFolderRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeFolder", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeFolder indicates an expected call of ChangeFolder.
func (mr *MockFolderManagerServerMockRecorder) ChangeFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeFolder", reflect.TypeOf((*MockFolderManagerServer)(nil).ChangeFolder), arg0, arg1)
}

// DeleteFolder mocks base method.
func (m *MockFolderManagerServer) DeleteFolder(arg0 context.Context, arg1 *DeleteFolderRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFolder", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFolder indicates an expected call of DeleteFolder.
func (mr *MockFolderManagerServerMockRecorder) DeleteFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolder", reflect.TypeOf((*MockFolderManagerServer)(nil).DeleteFolder), arg0, arg1)
}

// DeleteFolderMail mocks base method.
func (m *MockFolderManagerServer) DeleteFolderMail(arg0 context.Context, arg1 *DeleteFolderMailRequest) (*utils_proto.JsonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFolderMail", arg0, arg1)
	ret0, _ := ret[0].(*utils_proto.JsonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFolderMail indicates an expected call of DeleteFolderMail.
func (mr *MockFolderManagerServerMockRecorder) DeleteFolderMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolderMail", reflect.TypeOf((*MockFolderManagerServer)(nil).DeleteFolderMail), arg0, arg1)
}

// ListFolder mocks base method.
func (m *MockFolderManagerServer) ListFolder(arg0 context.Context, arg1 *ListFolderRequest) (*ResponseMails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFolder", arg0, arg1)
	ret0, _ := ret[0].(*ResponseMails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFolder indicates an expected call of ListFolder.
func (mr *MockFolderManagerServerMockRecorder) ListFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFolder", reflect.TypeOf((*MockFolderManagerServer)(nil).ListFolder), arg0, arg1)
}

// ListFolders mocks base method.
func (m *MockFolderManagerServer) ListFolders(arg0 context.Context, arg1 *ListFoldersRequest) (*ResponseFolders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFolders", arg0, arg1)
	ret0, _ := ret[0].(*ResponseFolders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFolders indicates an expected call of ListFolders.
func (mr *MockFolderManagerServerMockRecorder) ListFolders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFolders", reflect.TypeOf((*MockFolderManagerServer)(nil).ListFolders), arg0, arg1)
}

// mustEmbedUnimplementedFolderManagerServer mocks base method.
func (m *MockFolderManagerServer) mustEmbedUnimplementedFolderManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFolderManagerServer")
}

// mustEmbedUnimplementedFolderManagerServer indicates an expected call of mustEmbedUnimplementedFolderManagerServer.
func (mr *MockFolderManagerServerMockRecorder) mustEmbedUnimplementedFolderManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFolderManagerServer", reflect.TypeOf((*MockFolderManagerServer)(nil).mustEmbedUnimplementedFolderManagerServer))
}

// MockUnsafeFolderManagerServer is a mock of UnsafeFolderManagerServer interface.
type MockUnsafeFolderManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeFolderManagerServerMockRecorder
}

// MockUnsafeFolderManagerServerMockRecorder is the mock recorder for MockUnsafeFolderManagerServer.
type MockUnsafeFolderManagerServerMockRecorder struct {
	mock *MockUnsafeFolderManagerServer
}

// NewMockUnsafeFolderManagerServer creates a new mock instance.
func NewMockUnsafeFolderManagerServer(ctrl *gomock.Controller) *MockUnsafeFolderManagerServer {
	mock := &MockUnsafeFolderManagerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeFolderManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeFolderManagerServer) EXPECT() *MockUnsafeFolderManagerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedFolderManagerServer mocks base method.
func (m *MockUnsafeFolderManagerServer) mustEmbedUnimplementedFolderManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFolderManagerServer")
}

// mustEmbedUnimplementedFolderManagerServer indicates an expected call of mustEmbedUnimplementedFolderManagerServer.
func (mr *MockUnsafeFolderManagerServerMockRecorder) mustEmbedUnimplementedFolderManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFolderManagerServer", reflect.TypeOf((*MockUnsafeFolderManagerServer)(nil).mustEmbedUnimplementedFolderManagerServer))
}
