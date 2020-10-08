// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: protobuf/jx3box.proto

package jx3box

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the user's name.
type UserQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserQueryParams) Reset() {
	*x = UserQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQueryParams) ProtoMessage() {}

func (x *UserQueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQueryParams.ProtoReflect.Descriptor instead.
func (*UserQueryParams) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{0}
}

func (x *UserQueryParams) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type UserRenameParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *UserRenameParams) Reset() {
	*x = UserRenameParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRenameParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRenameParams) ProtoMessage() {}

func (x *UserRenameParams) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRenameParams.ProtoReflect.Descriptor instead.
func (*UserRenameParams) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{1}
}

func (x *UserRenameParams) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserRenameParams) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// The response message containing the greetings
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status      int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Group       uint64 `protobuf:"varint,3,opt,name=group,proto3" json:"group,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[2]
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
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *User) GetGroup() uint64 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_protobuf_jx3box_proto protoreflect.FileDescriptor

var file_protobuf_jx3box_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6a, 0x78, 0x33, 0x62, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x22,
	0x23, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0x74, 0x0a, 0x06, 0x4a, 0x58, 0x33, 0x42, 0x6f, 0x78, 0x12, 0x32,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6a, 0x78, 0x33, 0x62,
	0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x6a, 0x78, 0x33,
	0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_jx3box_proto_rawDescOnce sync.Once
	file_protobuf_jx3box_proto_rawDescData = file_protobuf_jx3box_proto_rawDesc
)

func file_protobuf_jx3box_proto_rawDescGZIP() []byte {
	file_protobuf_jx3box_proto_rawDescOnce.Do(func() {
		file_protobuf_jx3box_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_jx3box_proto_rawDescData)
	})
	return file_protobuf_jx3box_proto_rawDescData
}

var file_protobuf_jx3box_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_jx3box_proto_goTypes = []interface{}{
	(*UserQueryParams)(nil),  // 0: jx3box.UserQueryParams
	(*UserRenameParams)(nil), // 1: jx3box.UserRenameParams
	(*User)(nil),             // 2: jx3box.User
}
var file_protobuf_jx3box_proto_depIdxs = []int32{
	0, // 0: jx3box.JX3Box.GetUser:input_type -> jx3box.UserQueryParams
	1, // 1: jx3box.JX3Box.UserRename:input_type -> jx3box.UserRenameParams
	2, // 2: jx3box.JX3Box.GetUser:output_type -> jx3box.User
	2, // 3: jx3box.JX3Box.UserRename:output_type -> jx3box.User
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_jx3box_proto_init() }
func file_protobuf_jx3box_proto_init() {
	if File_protobuf_jx3box_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_jx3box_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQueryParams); i {
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
		file_protobuf_jx3box_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRenameParams); i {
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
		file_protobuf_jx3box_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_jx3box_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_jx3box_proto_goTypes,
		DependencyIndexes: file_protobuf_jx3box_proto_depIdxs,
		MessageInfos:      file_protobuf_jx3box_proto_msgTypes,
	}.Build()
	File_protobuf_jx3box_proto = out.File
	file_protobuf_jx3box_proto_rawDesc = nil
	file_protobuf_jx3box_proto_goTypes = nil
	file_protobuf_jx3box_proto_depIdxs = nil
}
