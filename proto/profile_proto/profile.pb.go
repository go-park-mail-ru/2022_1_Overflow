// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/profile.proto

package profile_proto

import (
	utils_proto "OverflowBackend/proto/utils_proto"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Avatar []byte               `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *SetAvatarRequest) Reset() {
	*x = SetAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarRequest) ProtoMessage() {}

func (x *SetAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarRequest.ProtoReflect.Descriptor instead.
func (*SetAvatarRequest) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{0}
}

func (x *SetAvatarRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetAvatarRequest) GetAvatar() []byte {
	if x != nil {
		return x.Avatar
	}
	return nil
}

type SetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Form []byte               `protobuf:"bytes,2,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *SetInfoRequest) Reset() {
	*x = SetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetInfoRequest) ProtoMessage() {}

func (x *SetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetInfoRequest.ProtoReflect.Descriptor instead.
func (*SetInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{1}
}

func (x *SetInfoRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetInfoRequest) GetForm() []byte {
	if x != nil {
		return x.Form
	}
	return nil
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	DummyName string `protobuf:"bytes,2,opt,name=dummyName,proto3" json:"dummyName,omitempty"`
}

func (x *GetAvatarRequest) Reset() {
	*x = GetAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarRequest) ProtoMessage() {}

func (x *GetAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarRequest.ProtoReflect.Descriptor instead.
func (*GetAvatarRequest) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{3}
}

func (x *GetAvatarRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetAvatarRequest) GetDummyName() string {
	if x != nil {
		return x.DummyName
	}
	return ""
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PasswordOld string               `protobuf:"bytes,2,opt,name=passwordOld,proto3" json:"passwordOld,omitempty"`
	PasswordNew string               `protobuf:"bytes,3,opt,name=passwordNew,proto3" json:"passwordNew,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePasswordRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChangePasswordRequest) GetPasswordOld() string {
	if x != nil {
		return x.PasswordOld
	}
	return ""
}

func (x *ChangePasswordRequest) GetPasswordNew() string {
	if x != nil {
		return x.PasswordNew
	}
	return ""
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *utils_proto.JsonResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Data     []byte                    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{5}
}

func (x *GetInfoResponse) GetResponse() *utils_proto.JsonResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GetInfoResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *utils_proto.JsonResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Url      string                    `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetAvatarResponse) Reset() {
	*x = GetAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarResponse) ProtoMessage() {}

func (x *GetAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarResponse.ProtoReflect.Descriptor instead.
func (*GetAvatarResponse) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{6}
}

func (x *GetAvatarResponse) GetResponse() *utils_proto.JsonResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GetAvatarResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_proto_profile_proto protoreflect.FileDescriptor

var file_proto_profile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x6f, 0x72,
	0x6d, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75,
	0x74, 0x69, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4f, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x4f, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4e, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4e, 0x65, 0x77, 0x22, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xfc, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x07, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73,
	0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_profile_proto_rawDescOnce sync.Once
	file_proto_profile_proto_rawDescData = file_proto_profile_proto_rawDesc
)

func file_proto_profile_proto_rawDescGZIP() []byte {
	file_proto_profile_proto_rawDescOnce.Do(func() {
		file_proto_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_profile_proto_rawDescData)
	})
	return file_proto_profile_proto_rawDescData
}

var file_proto_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_profile_proto_goTypes = []interface{}{
	(*SetAvatarRequest)(nil),         // 0: profile_proto.SetAvatarRequest
	(*SetInfoRequest)(nil),           // 1: profile_proto.SetInfoRequest
	(*GetInfoRequest)(nil),           // 2: profile_proto.GetInfoRequest
	(*GetAvatarRequest)(nil),         // 3: profile_proto.GetAvatarRequest
	(*ChangePasswordRequest)(nil),    // 4: profile_proto.ChangePasswordRequest
	(*GetInfoResponse)(nil),          // 5: profile_proto.GetInfoResponse
	(*GetAvatarResponse)(nil),        // 6: profile_proto.GetAvatarResponse
	(*utils_proto.Session)(nil),      // 7: utils.Session
	(*utils_proto.JsonResponse)(nil), // 8: utils.JsonResponse
}
var file_proto_profile_proto_depIdxs = []int32{
	7,  // 0: profile_proto.SetAvatarRequest.data:type_name -> utils.Session
	7,  // 1: profile_proto.SetInfoRequest.data:type_name -> utils.Session
	7,  // 2: profile_proto.GetInfoRequest.data:type_name -> utils.Session
	7,  // 3: profile_proto.ChangePasswordRequest.data:type_name -> utils.Session
	8,  // 4: profile_proto.GetInfoResponse.response:type_name -> utils.JsonResponse
	8,  // 5: profile_proto.GetAvatarResponse.response:type_name -> utils.JsonResponse
	2,  // 6: profile_proto.Profile.GetInfo:input_type -> profile_proto.GetInfoRequest
	1,  // 7: profile_proto.Profile.SetInfo:input_type -> profile_proto.SetInfoRequest
	3,  // 8: profile_proto.Profile.GetAvatar:input_type -> profile_proto.GetAvatarRequest
	0,  // 9: profile_proto.Profile.SetAvatar:input_type -> profile_proto.SetAvatarRequest
	4,  // 10: profile_proto.Profile.ChangePassword:input_type -> profile_proto.ChangePasswordRequest
	5,  // 11: profile_proto.Profile.GetInfo:output_type -> profile_proto.GetInfoResponse
	8,  // 12: profile_proto.Profile.SetInfo:output_type -> utils.JsonResponse
	6,  // 13: profile_proto.Profile.GetAvatar:output_type -> profile_proto.GetAvatarResponse
	8,  // 14: profile_proto.Profile.SetAvatar:output_type -> utils.JsonResponse
	8,  // 15: profile_proto.Profile.ChangePassword:output_type -> utils.JsonResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_profile_proto_init() }
func file_proto_profile_proto_init() {
	if File_proto_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvatarRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvatarRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvatarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_profile_proto_goTypes,
		DependencyIndexes: file_proto_profile_proto_depIdxs,
		MessageInfos:      file_proto_profile_proto_msgTypes,
	}.Build()
	File_proto_profile_proto = out.File
	file_proto_profile_proto_rawDesc = nil
	file_proto_profile_proto_goTypes = nil
	file_proto_profile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	SetInfo(ctx context.Context, in *SetInfoRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
	GetAvatar(ctx context.Context, in *GetAvatarRequest, opts ...grpc.CallOption) (*GetAvatarResponse, error)
	SetAvatar(ctx context.Context, in *SetAvatarRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/profile_proto.Profile/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) SetInfo(ctx context.Context, in *SetInfoRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/profile_proto.Profile/SetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetAvatar(ctx context.Context, in *GetAvatarRequest, opts ...grpc.CallOption) (*GetAvatarResponse, error) {
	out := new(GetAvatarResponse)
	err := c.cc.Invoke(ctx, "/profile_proto.Profile/GetAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) SetAvatar(ctx context.Context, in *SetAvatarRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/profile_proto.Profile/SetAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/profile_proto.Profile/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
type ProfileServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	SetInfo(context.Context, *SetInfoRequest) (*utils_proto.JsonResponse, error)
	GetAvatar(context.Context, *GetAvatarRequest) (*GetAvatarResponse, error)
	SetAvatar(context.Context, *SetAvatarRequest) (*utils_proto.JsonResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*utils_proto.JsonResponse, error)
}

// UnimplementedProfileServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (*UnimplementedProfileServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedProfileServer) SetInfo(context.Context, *SetInfoRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (*UnimplementedProfileServer) GetAvatar(context.Context, *GetAvatarRequest) (*GetAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatar not implemented")
}
func (*UnimplementedProfileServer) SetAvatar(context.Context, *SetAvatarRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAvatar not implemented")
}
func (*UnimplementedProfileServer) ChangePassword(context.Context, *ChangePasswordRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_proto.Profile/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_SetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).SetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_proto.Profile/SetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).SetInfo(ctx, req.(*SetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_proto.Profile/GetAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetAvatar(ctx, req.(*GetAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_SetAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).SetAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_proto.Profile/SetAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).SetAvatar(ctx, req.(*SetAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile_proto.Profile/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile_proto.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Profile_GetInfo_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _Profile_SetInfo_Handler,
		},
		{
			MethodName: "GetAvatar",
			Handler:    _Profile_GetAvatar_Handler,
		},
		{
			MethodName: "SetAvatar",
			Handler:    _Profile_SetAvatar_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Profile_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/profile.proto",
}
