// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/folder_manager.proto

package folder_manager_proto

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

type AddFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Name string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddFolderRequest) Reset() {
	*x = AddFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFolderRequest) ProtoMessage() {}

func (x *AddFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFolderRequest.ProtoReflect.Descriptor instead.
func (*AddFolderRequest) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{0}
}

func (x *AddFolderRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AddFolderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddMailToFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	FolderId int32                `protobuf:"varint,2,opt,name=folderId,proto3" json:"folderId,omitempty"`
	MailId   int32                `protobuf:"varint,3,opt,name=mailId,proto3" json:"mailId,omitempty"`
}

func (x *AddMailToFolderRequest) Reset() {
	*x = AddMailToFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMailToFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMailToFolderRequest) ProtoMessage() {}

func (x *AddMailToFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMailToFolderRequest.ProtoReflect.Descriptor instead.
func (*AddMailToFolderRequest) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{1}
}

func (x *AddMailToFolderRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AddMailToFolderRequest) GetFolderId() int32 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *AddMailToFolderRequest) GetMailId() int32 {
	if x != nil {
		return x.MailId
	}
	return 0
}

type ChangeFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data          *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	FolderId      int32                `protobuf:"varint,2,opt,name=folderId,proto3" json:"folderId,omitempty"`
	FolderNewName string               `protobuf:"bytes,3,opt,name=folderNewName,proto3" json:"folderNewName,omitempty"`
}

func (x *ChangeFolderRequest) Reset() {
	*x = ChangeFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeFolderRequest) ProtoMessage() {}

func (x *ChangeFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeFolderRequest.ProtoReflect.Descriptor instead.
func (*ChangeFolderRequest) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeFolderRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChangeFolderRequest) GetFolderId() int32 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *ChangeFolderRequest) GetFolderNewName() string {
	if x != nil {
		return x.FolderNewName
	}
	return ""
}

type ResponseFolders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *utils_proto.JsonResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Folders  []byte                    `protobuf:"bytes,2,opt,name=folders,proto3" json:"folders,omitempty"`
}

func (x *ResponseFolders) Reset() {
	*x = ResponseFolders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFolders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFolders) ProtoMessage() {}

func (x *ResponseFolders) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFolders.ProtoReflect.Descriptor instead.
func (*ResponseFolders) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseFolders) GetResponse() *utils_proto.JsonResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ResponseFolders) GetFolders() []byte {
	if x != nil {
		return x.Folders
	}
	return nil
}

type ResponseMails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *utils_proto.JsonResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Mails    []byte                    `protobuf:"bytes,2,opt,name=mails,proto3" json:"mails,omitempty"`
}

func (x *ResponseMails) Reset() {
	*x = ResponseMails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMails) ProtoMessage() {}

func (x *ResponseMails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMails.ProtoReflect.Descriptor instead.
func (*ResponseMails) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseMails) GetResponse() *utils_proto.JsonResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ResponseMails) GetMails() []byte {
	if x != nil {
		return x.Mails
	}
	return nil
}

type DeleteFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Name string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteFolderRequest) Reset() {
	*x = DeleteFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFolderRequest) ProtoMessage() {}

func (x *DeleteFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFolderRequest.ProtoReflect.Descriptor instead.
func (*DeleteFolderRequest) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFolderRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DeleteFolderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListFoldersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ListFoldersRequest) Reset() {
	*x = ListFoldersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoldersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoldersRequest) ProtoMessage() {}

func (x *ListFoldersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoldersRequest.ProtoReflect.Descriptor instead.
func (*ListFoldersRequest) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{6}
}

func (x *ListFoldersRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *utils_proto.Session `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	FolderId int32                `protobuf:"varint,2,opt,name=folderId,proto3" json:"folderId,omitempty"`
}

func (x *ListFolderRequest) Reset() {
	*x = ListFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_folder_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFolderRequest) ProtoMessage() {}

func (x *ListFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_folder_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFolderRequest.ProtoReflect.Descriptor instead.
func (*ListFolderRequest) Descriptor() ([]byte, []int) {
	return file_proto_folder_manager_proto_rawDescGZIP(), []int{7}
}

func (x *ListFolderRequest) GetData() *utils_proto.Session {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListFolderRequest) GetFolderId() int32 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

var File_proto_folder_manager_proto protoreflect.FileDescriptor

var file_proto_folder_manager_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x70, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x69,
	0x6c, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x56,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x4d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x53, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x32, 0x97, 0x04, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00,
	0x12, 0x5c, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x27,
	0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x16, 0x2e, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_folder_manager_proto_rawDescOnce sync.Once
	file_proto_folder_manager_proto_rawDescData = file_proto_folder_manager_proto_rawDesc
)

func file_proto_folder_manager_proto_rawDescGZIP() []byte {
	file_proto_folder_manager_proto_rawDescOnce.Do(func() {
		file_proto_folder_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_folder_manager_proto_rawDescData)
	})
	return file_proto_folder_manager_proto_rawDescData
}

var file_proto_folder_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_folder_manager_proto_goTypes = []interface{}{
	(*AddFolderRequest)(nil),         // 0: folder_manager_proto.AddFolderRequest
	(*AddMailToFolderRequest)(nil),   // 1: folder_manager_proto.AddMailToFolderRequest
	(*ChangeFolderRequest)(nil),      // 2: folder_manager_proto.ChangeFolderRequest
	(*ResponseFolders)(nil),          // 3: folder_manager_proto.ResponseFolders
	(*ResponseMails)(nil),            // 4: folder_manager_proto.ResponseMails
	(*DeleteFolderRequest)(nil),      // 5: folder_manager_proto.DeleteFolderRequest
	(*ListFoldersRequest)(nil),       // 6: folder_manager_proto.ListFoldersRequest
	(*ListFolderRequest)(nil),        // 7: folder_manager_proto.ListFolderRequest
	(*utils_proto.Session)(nil),      // 8: utils.Session
	(*utils_proto.JsonResponse)(nil), // 9: utils.JsonResponse
}
var file_proto_folder_manager_proto_depIdxs = []int32{
	8,  // 0: folder_manager_proto.AddFolderRequest.data:type_name -> utils.Session
	8,  // 1: folder_manager_proto.AddMailToFolderRequest.data:type_name -> utils.Session
	8,  // 2: folder_manager_proto.ChangeFolderRequest.data:type_name -> utils.Session
	9,  // 3: folder_manager_proto.ResponseFolders.response:type_name -> utils.JsonResponse
	9,  // 4: folder_manager_proto.ResponseMails.response:type_name -> utils.JsonResponse
	8,  // 5: folder_manager_proto.DeleteFolderRequest.data:type_name -> utils.Session
	8,  // 6: folder_manager_proto.ListFoldersRequest.data:type_name -> utils.Session
	8,  // 7: folder_manager_proto.ListFolderRequest.data:type_name -> utils.Session
	0,  // 8: folder_manager_proto.FolderManager.AddFolder:input_type -> folder_manager_proto.AddFolderRequest
	1,  // 9: folder_manager_proto.FolderManager.AddMailToFolder:input_type -> folder_manager_proto.AddMailToFolderRequest
	2,  // 10: folder_manager_proto.FolderManager.ChangeFolder:input_type -> folder_manager_proto.ChangeFolderRequest
	5,  // 11: folder_manager_proto.FolderManager.DeleteFolder:input_type -> folder_manager_proto.DeleteFolderRequest
	6,  // 12: folder_manager_proto.FolderManager.ListFolders:input_type -> folder_manager_proto.ListFoldersRequest
	7,  // 13: folder_manager_proto.FolderManager.ListFolder:input_type -> folder_manager_proto.ListFolderRequest
	9,  // 14: folder_manager_proto.FolderManager.AddFolder:output_type -> utils.JsonResponse
	9,  // 15: folder_manager_proto.FolderManager.AddMailToFolder:output_type -> utils.JsonResponse
	9,  // 16: folder_manager_proto.FolderManager.ChangeFolder:output_type -> utils.JsonResponse
	9,  // 17: folder_manager_proto.FolderManager.DeleteFolder:output_type -> utils.JsonResponse
	3,  // 18: folder_manager_proto.FolderManager.ListFolders:output_type -> folder_manager_proto.ResponseFolders
	4,  // 19: folder_manager_proto.FolderManager.ListFolder:output_type -> folder_manager_proto.ResponseMails
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_folder_manager_proto_init() }
func file_proto_folder_manager_proto_init() {
	if File_proto_folder_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_folder_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFolderRequest); i {
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
		file_proto_folder_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMailToFolderRequest); i {
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
		file_proto_folder_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeFolderRequest); i {
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
		file_proto_folder_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFolders); i {
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
		file_proto_folder_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMails); i {
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
		file_proto_folder_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFolderRequest); i {
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
		file_proto_folder_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoldersRequest); i {
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
		file_proto_folder_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFolderRequest); i {
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
			RawDescriptor: file_proto_folder_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_folder_manager_proto_goTypes,
		DependencyIndexes: file_proto_folder_manager_proto_depIdxs,
		MessageInfos:      file_proto_folder_manager_proto_msgTypes,
	}.Build()
	File_proto_folder_manager_proto = out.File
	file_proto_folder_manager_proto_rawDesc = nil
	file_proto_folder_manager_proto_goTypes = nil
	file_proto_folder_manager_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FolderManagerClient is the client API for FolderManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FolderManagerClient interface {
	AddFolder(ctx context.Context, in *AddFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
	AddMailToFolder(ctx context.Context, in *AddMailToFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
	ChangeFolder(ctx context.Context, in *ChangeFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
	//rpc GetFolderByName(GetFolderByNameRequest) returns (ResponseFolder) {}
	//rpc GetFolderById(GetFolderByIdRequest) returns (ResponseFolder) {}
	DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error)
	ListFolders(ctx context.Context, in *ListFoldersRequest, opts ...grpc.CallOption) (*ResponseFolders, error)
	ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*ResponseMails, error)
}

type folderManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewFolderManagerClient(cc grpc.ClientConnInterface) FolderManagerClient {
	return &folderManagerClient{cc}
}

func (c *folderManagerClient) AddFolder(ctx context.Context, in *AddFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/folder_manager_proto.FolderManager/AddFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderManagerClient) AddMailToFolder(ctx context.Context, in *AddMailToFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/folder_manager_proto.FolderManager/AddMailToFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderManagerClient) ChangeFolder(ctx context.Context, in *ChangeFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/folder_manager_proto.FolderManager/ChangeFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderManagerClient) DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*utils_proto.JsonResponse, error) {
	out := new(utils_proto.JsonResponse)
	err := c.cc.Invoke(ctx, "/folder_manager_proto.FolderManager/DeleteFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderManagerClient) ListFolders(ctx context.Context, in *ListFoldersRequest, opts ...grpc.CallOption) (*ResponseFolders, error) {
	out := new(ResponseFolders)
	err := c.cc.Invoke(ctx, "/folder_manager_proto.FolderManager/ListFolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderManagerClient) ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*ResponseMails, error) {
	out := new(ResponseMails)
	err := c.cc.Invoke(ctx, "/folder_manager_proto.FolderManager/ListFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FolderManagerServer is the server API for FolderManager service.
type FolderManagerServer interface {
	AddFolder(context.Context, *AddFolderRequest) (*utils_proto.JsonResponse, error)
	AddMailToFolder(context.Context, *AddMailToFolderRequest) (*utils_proto.JsonResponse, error)
	ChangeFolder(context.Context, *ChangeFolderRequest) (*utils_proto.JsonResponse, error)
	//rpc GetFolderByName(GetFolderByNameRequest) returns (ResponseFolder) {}
	//rpc GetFolderById(GetFolderByIdRequest) returns (ResponseFolder) {}
	DeleteFolder(context.Context, *DeleteFolderRequest) (*utils_proto.JsonResponse, error)
	ListFolders(context.Context, *ListFoldersRequest) (*ResponseFolders, error)
	ListFolder(context.Context, *ListFolderRequest) (*ResponseMails, error)
}

// UnimplementedFolderManagerServer can be embedded to have forward compatible implementations.
type UnimplementedFolderManagerServer struct {
}

func (*UnimplementedFolderManagerServer) AddFolder(context.Context, *AddFolderRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFolder not implemented")
}
func (*UnimplementedFolderManagerServer) AddMailToFolder(context.Context, *AddMailToFolderRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMailToFolder not implemented")
}
func (*UnimplementedFolderManagerServer) ChangeFolder(context.Context, *ChangeFolderRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeFolder not implemented")
}
func (*UnimplementedFolderManagerServer) DeleteFolder(context.Context, *DeleteFolderRequest) (*utils_proto.JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}
func (*UnimplementedFolderManagerServer) ListFolders(context.Context, *ListFoldersRequest) (*ResponseFolders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFolders not implemented")
}
func (*UnimplementedFolderManagerServer) ListFolder(context.Context, *ListFolderRequest) (*ResponseMails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFolder not implemented")
}

func RegisterFolderManagerServer(s *grpc.Server, srv FolderManagerServer) {
	s.RegisterService(&_FolderManager_serviceDesc, srv)
}

func _FolderManager_AddFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderManagerServer).AddFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folder_manager_proto.FolderManager/AddFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderManagerServer).AddFolder(ctx, req.(*AddFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderManager_AddMailToFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMailToFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderManagerServer).AddMailToFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folder_manager_proto.FolderManager/AddMailToFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderManagerServer).AddMailToFolder(ctx, req.(*AddMailToFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderManager_ChangeFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderManagerServer).ChangeFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folder_manager_proto.FolderManager/ChangeFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderManagerServer).ChangeFolder(ctx, req.(*ChangeFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderManager_DeleteFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderManagerServer).DeleteFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folder_manager_proto.FolderManager/DeleteFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderManagerServer).DeleteFolder(ctx, req.(*DeleteFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderManager_ListFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderManagerServer).ListFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folder_manager_proto.FolderManager/ListFolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderManagerServer).ListFolders(ctx, req.(*ListFoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderManager_ListFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderManagerServer).ListFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folder_manager_proto.FolderManager/ListFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderManagerServer).ListFolder(ctx, req.(*ListFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FolderManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "folder_manager_proto.FolderManager",
	HandlerType: (*FolderManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFolder",
			Handler:    _FolderManager_AddFolder_Handler,
		},
		{
			MethodName: "AddMailToFolder",
			Handler:    _FolderManager_AddMailToFolder_Handler,
		},
		{
			MethodName: "ChangeFolder",
			Handler:    _FolderManager_ChangeFolder_Handler,
		},
		{
			MethodName: "DeleteFolder",
			Handler:    _FolderManager_DeleteFolder_Handler,
		},
		{
			MethodName: "ListFolders",
			Handler:    _FolderManager_ListFolders_Handler,
		},
		{
			MethodName: "ListFolder",
			Handler:    _FolderManager_ListFolder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/folder_manager.proto",
}
