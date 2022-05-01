// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/utils.proto

package utils_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DatabaseStatus int32

const (
	DatabaseStatus_OK    DatabaseStatus = 0
	DatabaseStatus_ERROR DatabaseStatus = 1
)

// Enum value maps for DatabaseStatus.
var (
	DatabaseStatus_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	DatabaseStatus_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x DatabaseStatus) Enum() *DatabaseStatus {
	p := new(DatabaseStatus)
	*p = x
	return p
}

func (x DatabaseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatabaseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_utils_proto_enumTypes[0].Descriptor()
}

func (DatabaseStatus) Type() protoreflect.EnumType {
	return &file_proto_utils_proto_enumTypes[0]
}

func (x DatabaseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatabaseStatus.Descriptor instead.
func (DatabaseStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{0}
}

type DatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status DatabaseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=utils.DatabaseStatus" json:"status,omitempty"`
}

func (x *DatabaseResponse) Reset() {
	*x = DatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseResponse) ProtoMessage() {}

func (x *DatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseResponse.ProtoReflect.Descriptor instead.
func (*DatabaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseResponse) GetStatus() DatabaseStatus {
	if x != nil {
		return x.Status
	}
	return DatabaseStatus_OK
}

type JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JsonResponse) Reset() {
	*x = JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonResponse) ProtoMessage() {}

func (x *JsonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonResponse.ProtoReflect.Descriptor instead.
func (*JsonResponse) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{1}
}

func (x *JsonResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *JsonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MailForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addressee string `protobuf:"bytes,1,opt,name=addressee,proto3" json:"addressee,omitempty"`
	Theme     string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Files     string `protobuf:"bytes,4,opt,name=files,proto3" json:"files,omitempty"`
}

func (x *MailForm) Reset() {
	*x = MailForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailForm) ProtoMessage() {}

func (x *MailForm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailForm.ProtoReflect.Descriptor instead.
func (*MailForm) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{2}
}

func (x *MailForm) GetAddressee() string {
	if x != nil {
		return x.Addressee
	}
	return ""
}

func (x *MailForm) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *MailForm) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MailForm) GetFiles() string {
	if x != nil {
		return x.Files
	}
	return ""
}

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId  int32                  `protobuf:"varint,2,opt,name=client_id,proto3" json:"client_id,omitempty"`
	Sender    string                 `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Addressee string                 `protobuf:"bytes,4,opt,name=addressee,proto3" json:"addressee,omitempty"`
	Theme     string                 `protobuf:"bytes,5,opt,name=theme,proto3" json:"theme,omitempty"`
	Text      string                 `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Files     string                 `protobuf:"bytes,7,opt,name=files,proto3" json:"files,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Read      *wrapperspb.BoolValue  `protobuf:"bytes,9,opt,name=read,proto3" json:"read,omitempty"`
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{3}
}

func (x *Mail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Mail) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Mail) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Mail) GetAddressee() string {
	if x != nil {
		return x.Addressee
	}
	return ""
}

func (x *Mail) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *Mail) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Mail) GetFiles() string {
	if x != nil {
		return x.Files
	}
	return ""
}

func (x *Mail) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Mail) GetRead() *wrapperspb.BoolValue {
	if x != nil {
		return x.Read
	}
	return nil
}

type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId    int32                  `protobuf:"varint,3,opt,name=userId,json=user_id,proto3" json:"userId,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,json=created_at,proto3" json:"createdAt,omitempty"` //repeated utils.MailAdditional mails = 4 [json_name="mails"];
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{4}
}

func (x *Folder) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Folder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Folder) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Folder) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type MailAdditional struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail      *Mail  `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
	AvatarUrl string `protobuf:"bytes,2,opt,name=avatar_url,proto3" json:"avatar_url,omitempty"`
}

func (x *MailAdditional) Reset() {
	*x = MailAdditional{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailAdditional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailAdditional) ProtoMessage() {}

func (x *MailAdditional) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailAdditional.ProtoReflect.Descriptor instead.
func (*MailAdditional) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{5}
}

func (x *MailAdditional) GetMail() *Mail {
	if x != nil {
		return x.Mail
	}
	return nil
}

func (x *MailAdditional) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,proto3" json:"last_name,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ProfileSettingsForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,proto3" json:"last_name,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ProfileSettingsForm) Reset() {
	*x = ProfileSettingsForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSettingsForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSettingsForm) ProtoMessage() {}

func (x *ProfileSettingsForm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSettingsForm.ProtoReflect.Descriptor instead.
func (*ProfileSettingsForm) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{7}
}

func (x *ProfileSettingsForm) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ProfileSettingsForm) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ProfileSettingsForm) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string                `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Authenticated *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{8}
}

func (x *Session) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Session) GetAuthenticated() *wrapperspb.BoolValue {
	if x != nil {
		return x.Authenticated
	}
	return nil
}

type FolderSettingsForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewName string `protobuf:"bytes,1,opt,name=new_name,proto3" json:"new_name,omitempty"`
}

func (x *FolderSettingsForm) Reset() {
	*x = FolderSettingsForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderSettingsForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderSettingsForm) ProtoMessage() {}

func (x *FolderSettingsForm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utils_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderSettingsForm.ProtoReflect.Descriptor instead.
func (*FolderSettingsForm) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{9}
}

func (x *FolderSettingsForm) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

var File_proto_utils_proto protoreflect.FileDescriptor

var file_proto_utils_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40,
	0x0a, 0x0c, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x68, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8a, 0x02, 0x0a, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12,
	0x39, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x51, 0x0a, 0x0e, 0x4d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x04,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x8c, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6f, 0x0a, 0x13,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46,
	0x6f, 0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x67, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x30, 0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x23, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_utils_proto_rawDescOnce sync.Once
	file_proto_utils_proto_rawDescData = file_proto_utils_proto_rawDesc
)

func file_proto_utils_proto_rawDescGZIP() []byte {
	file_proto_utils_proto_rawDescOnce.Do(func() {
		file_proto_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_utils_proto_rawDescData)
	})
	return file_proto_utils_proto_rawDescData
}

var file_proto_utils_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_utils_proto_goTypes = []interface{}{
	(DatabaseStatus)(0),           // 0: utils.DatabaseStatus
	(*DatabaseResponse)(nil),      // 1: utils.DatabaseResponse
	(*JsonResponse)(nil),          // 2: utils.JsonResponse
	(*MailForm)(nil),              // 3: utils.MailForm
	(*Mail)(nil),                  // 4: utils.Mail
	(*Folder)(nil),                // 5: utils.Folder
	(*MailAdditional)(nil),        // 6: utils.MailAdditional
	(*User)(nil),                  // 7: utils.User
	(*ProfileSettingsForm)(nil),   // 8: utils.ProfileSettingsForm
	(*Session)(nil),               // 9: utils.Session
	(*FolderSettingsForm)(nil),    // 10: utils.FolderSettingsForm
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*wrapperspb.BoolValue)(nil),  // 12: google.protobuf.BoolValue
}
var file_proto_utils_proto_depIdxs = []int32{
	0,  // 0: utils.DatabaseResponse.status:type_name -> utils.DatabaseStatus
	11, // 1: utils.Mail.date:type_name -> google.protobuf.Timestamp
	12, // 2: utils.Mail.read:type_name -> google.protobuf.BoolValue
	11, // 3: utils.Folder.createdAt:type_name -> google.protobuf.Timestamp
	4,  // 4: utils.MailAdditional.mail:type_name -> utils.Mail
	12, // 5: utils.Session.authenticated:type_name -> google.protobuf.BoolValue
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_utils_proto_init() }
func file_proto_utils_proto_init() {
	if File_proto_utils_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_utils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseResponse); i {
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
		file_proto_utils_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonResponse); i {
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
		file_proto_utils_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailForm); i {
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
		file_proto_utils_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
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
		file_proto_utils_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
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
		file_proto_utils_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailAdditional); i {
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
		file_proto_utils_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_utils_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSettingsForm); i {
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
		file_proto_utils_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_proto_utils_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderSettingsForm); i {
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
			RawDescriptor: file_proto_utils_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_utils_proto_goTypes,
		DependencyIndexes: file_proto_utils_proto_depIdxs,
		EnumInfos:         file_proto_utils_proto_enumTypes,
		MessageInfos:      file_proto_utils_proto_msgTypes,
	}.Build()
	File_proto_utils_proto = out.File
	file_proto_utils_proto_rawDesc = nil
	file_proto_utils_proto_goTypes = nil
	file_proto_utils_proto_depIdxs = nil
}
