// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/utils.proto

package utils_proto

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type DatabaseExtendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status DatabaseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=utils.DatabaseStatus" json:"status,omitempty"`
	Param  string         `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *DatabaseExtendResponse) Reset() {
	*x = DatabaseExtendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseExtendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseExtendResponse) ProtoMessage() {}

func (x *DatabaseExtendResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DatabaseExtendResponse.ProtoReflect.Descriptor instead.
func (*DatabaseExtendResponse) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseExtendResponse) GetStatus() DatabaseStatus {
	if x != nil {
		return x.Status
	}
	return DatabaseStatus_OK
}

func (x *DatabaseExtendResponse) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string              `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Authenticated *wrappers.BoolValue `protobuf:"bytes,2,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Session) GetAuthenticated() *wrappers.BoolValue {
	if x != nil {
		return x.Authenticated
	}
	return nil
}

type JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *JsonResponse) Reset() {
	*x = JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonResponse) ProtoMessage() {}

func (x *JsonResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use JsonResponse.ProtoReflect.Descriptor instead.
func (*JsonResponse) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{3}
}

func (x *JsonResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

type JsonExtendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Param    string `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *JsonExtendResponse) Reset() {
	*x = JsonExtendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utils_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonExtendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonExtendResponse) ProtoMessage() {}

func (x *JsonExtendResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use JsonExtendResponse.ProtoReflect.Descriptor instead.
func (*JsonExtendResponse) Descriptor() ([]byte, []int) {
	return file_proto_utils_proto_rawDescGZIP(), []int{4}
}

func (x *JsonExtendResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *JsonExtendResponse) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

var File_proto_utils_proto protoreflect.FileDescriptor

var file_proto_utils_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a,
	0x16, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x67, 0x0a, 0x07,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x0c, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x46, 0x0a, 0x12, 0x4a, 0x73, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2a, 0x23, 0x0a, 0x0e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x15,
	0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_proto_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_utils_proto_goTypes = []interface{}{
	(DatabaseStatus)(0),            // 0: utils.DatabaseStatus
	(*DatabaseResponse)(nil),       // 1: utils.DatabaseResponse
	(*DatabaseExtendResponse)(nil), // 2: utils.DatabaseExtendResponse
	(*Session)(nil),                // 3: utils.Session
	(*JsonResponse)(nil),           // 4: utils.JsonResponse
	(*JsonExtendResponse)(nil),     // 5: utils.JsonExtendResponse
	(*wrappers.BoolValue)(nil),     // 6: google.protobuf.BoolValue
}
var file_proto_utils_proto_depIdxs = []int32{
	0, // 0: utils.DatabaseResponse.status:type_name -> utils.DatabaseStatus
	0, // 1: utils.DatabaseExtendResponse.status:type_name -> utils.DatabaseStatus
	6, // 2: utils.Session.authenticated:type_name -> google.protobuf.BoolValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
			switch v := v.(*DatabaseExtendResponse); i {
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
		file_proto_utils_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_utils_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonExtendResponse); i {
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
			NumMessages:   5,
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
