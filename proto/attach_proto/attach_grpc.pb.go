// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/attach.proto

package attach_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AttachClient is the client API for Attach service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachClient interface {
	GetAttach(ctx context.Context, in *GetAttachRequest, opts ...grpc.CallOption) (*AttachResponse, error)
	SaveAttach(ctx context.Context, in *SaveAttachRequest, opts ...grpc.CallOption) (*Nothing, error)
	ListAttach(ctx context.Context, in *GetAttachRequest, opts ...grpc.CallOption) (*AttachListResponse, error)
	CheckAttachPermission(ctx context.Context, in *AttachPermissionRequest, opts ...grpc.CallOption) (*AttachPermissionResponse, error)
}

type attachClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachClient(cc grpc.ClientConnInterface) AttachClient {
	return &attachClient{cc}
}

func (c *attachClient) GetAttach(ctx context.Context, in *GetAttachRequest, opts ...grpc.CallOption) (*AttachResponse, error) {
	out := new(AttachResponse)
	err := c.cc.Invoke(ctx, "/attach_proto.Attach/GetAttach", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachClient) SaveAttach(ctx context.Context, in *SaveAttachRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/attach_proto.Attach/SaveAttach", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachClient) ListAttach(ctx context.Context, in *GetAttachRequest, opts ...grpc.CallOption) (*AttachListResponse, error) {
	out := new(AttachListResponse)
	err := c.cc.Invoke(ctx, "/attach_proto.Attach/ListAttach", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachClient) CheckAttachPermission(ctx context.Context, in *AttachPermissionRequest, opts ...grpc.CallOption) (*AttachPermissionResponse, error) {
	out := new(AttachPermissionResponse)
	err := c.cc.Invoke(ctx, "/attach_proto.Attach/CheckAttachPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachServer is the server API for Attach service.
// All implementations must embed UnimplementedAttachServer
// for forward compatibility
type AttachServer interface {
	GetAttach(context.Context, *GetAttachRequest) (*AttachResponse, error)
	SaveAttach(context.Context, *SaveAttachRequest) (*Nothing, error)
	ListAttach(context.Context, *GetAttachRequest) (*AttachListResponse, error)
	CheckAttachPermission(context.Context, *AttachPermissionRequest) (*AttachPermissionResponse, error)
	mustEmbedUnimplementedAttachServer()
}

// UnimplementedAttachServer must be embedded to have forward compatible implementations.
type UnimplementedAttachServer struct {
}

func (UnimplementedAttachServer) GetAttach(context.Context, *GetAttachRequest) (*AttachResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttach not implemented")
}
func (UnimplementedAttachServer) SaveAttach(context.Context, *SaveAttachRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAttach not implemented")
}
func (UnimplementedAttachServer) ListAttach(context.Context, *GetAttachRequest) (*AttachListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttach not implemented")
}
func (UnimplementedAttachServer) CheckAttachPermission(context.Context, *AttachPermissionRequest) (*AttachPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAttachPermission not implemented")
}
func (UnimplementedAttachServer) mustEmbedUnimplementedAttachServer() {}

// UnsafeAttachServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachServer will
// result in compilation errors.
type UnsafeAttachServer interface {
	mustEmbedUnimplementedAttachServer()
}

func RegisterAttachServer(s grpc.ServiceRegistrar, srv AttachServer) {
	s.RegisterService(&Attach_ServiceDesc, srv)
}

func _Attach_GetAttach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachServer).GetAttach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attach_proto.Attach/GetAttach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachServer).GetAttach(ctx, req.(*GetAttachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attach_SaveAttach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAttachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachServer).SaveAttach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attach_proto.Attach/SaveAttach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachServer).SaveAttach(ctx, req.(*SaveAttachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attach_ListAttach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachServer).ListAttach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attach_proto.Attach/ListAttach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachServer).ListAttach(ctx, req.(*GetAttachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attach_CheckAttachPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachServer).CheckAttachPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attach_proto.Attach/CheckAttachPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachServer).CheckAttachPermission(ctx, req.(*AttachPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Attach_ServiceDesc is the grpc.ServiceDesc for Attach service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attach_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attach_proto.Attach",
	HandlerType: (*AttachServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAttach",
			Handler:    _Attach_GetAttach_Handler,
		},
		{
			MethodName: "SaveAttach",
			Handler:    _Attach_SaveAttach_Handler,
		},
		{
			MethodName: "ListAttach",
			Handler:    _Attach_ListAttach_Handler,
		},
		{
			MethodName: "CheckAttachPermission",
			Handler:    _Attach_CheckAttachPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/attach.proto",
}
