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

// The response message containing the greetings
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Group         uint64 `protobuf:"varint,3,opt,name=group,proto3" json:"group,omitempty"`
	DisplayName   string `protobuf:"bytes,4,opt,name=displayName,proto3" json:"displayName,omitempty"`
	WechatUnionId string `protobuf:"bytes,5,opt,name=wechatUnionId,proto3" json:"wechatUnionId,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{1}
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

func (x *User) GetWechatUnionId() string {
	if x != nil {
		return x.WechatUnionId
	}
	return ""
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
		mi := &file_protobuf_jx3box_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRenameParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRenameParams) ProtoMessage() {}

func (x *UserRenameParams) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserRenameParams.ProtoReflect.Descriptor instead.
func (*UserRenameParams) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{2}
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

type UserRenameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Group       uint64 `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	DisplayName string `protobuf:"bytes,3,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Effect      int32  `protobuf:"varint,4,opt,name=effect,proto3" json:"effect,omitempty"` //受影响的条数
	Code        int32  `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`     //结果是否成功 0　成功,其他失败
	Msg         string `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`        // 失败原因
}

func (x *UserRenameResult) Reset() {
	*x = UserRenameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRenameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRenameResult) ProtoMessage() {}

func (x *UserRenameResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRenameResult.ProtoReflect.Descriptor instead.
func (*UserRenameResult) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{3}
}

func (x *UserRenameResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRenameResult) GetGroup() uint64 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *UserRenameResult) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserRenameResult) GetEffect() int32 {
	if x != nil {
		return x.Effect
	}
	return 0
}

func (x *UserRenameResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserRenameResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UserAwardParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Award string  `protobuf:"bytes,2,opt,name=award,proto3" json:"award,omitempty"`
	Uid   []int64 `protobuf:"varint,3,rep,packed,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserAwardParams) Reset() {
	*x = UserAwardParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAwardParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAwardParams) ProtoMessage() {}

func (x *UserAwardParams) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAwardParams.ProtoReflect.Descriptor instead.
func (*UserAwardParams) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{4}
}

func (x *UserAwardParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserAwardParams) GetAward() string {
	if x != nil {
		return x.Award
	}
	return ""
}

func (x *UserAwardParams) GetUid() []int64 {
	if x != nil {
		return x.Uid
	}
	return nil
}

type UserAwardResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Effect int64 `protobuf:"varint,1,opt,name=effect,proto3" json:"effect,omitempty"`
}

func (x *UserAwardResult) Reset() {
	*x = UserAwardResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAwardResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAwardResult) ProtoMessage() {}

func (x *UserAwardResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAwardResult.ProtoReflect.Descriptor instead.
func (*UserAwardResult) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{5}
}

func (x *UserAwardResult) GetEffect() int64 {
	if x != nil {
		return x.Effect
	}
	return 0
}

type PostsQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *PostsQueryParams) Reset() {
	*x = PostsQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsQueryParams) ProtoMessage() {}

func (x *PostsQueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsQueryParams.ProtoReflect.Descriptor instead.
func (*PostsQueryParams) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{6}
}

func (x *PostsQueryParams) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostsQueryParams) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type PostsQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId   uint64 `protobuf:"varint,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	AuthorName string `protobuf:"bytes,2,opt,name=authorName,proto3" json:"authorName,omitempty"`
}

func (x *PostsQueryResult) Reset() {
	*x = PostsQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsQueryResult) ProtoMessage() {}

func (x *PostsQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsQueryResult.ProtoReflect.Descriptor instead.
func (*PostsQueryResult) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{7}
}

func (x *PostsQueryResult) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PostsQueryResult) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

// 消息通知
type NotifyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessUserId uint64 `protobuf:"varint,1,opt,name=accessUserId,proto3" json:"accessUserId,omitempty"`
	SourceId     uint64 `protobuf:"varint,2,opt,name=sourceId,proto3" json:"sourceId,omitempty"`
	SourceType   string `protobuf:"bytes,3,opt,name=sourceType,proto3" json:"sourceType,omitempty"`
	Type         string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Subtype      string `protobuf:"bytes,5,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Content      string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *NotifyMessage) Reset() {
	*x = NotifyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMessage) ProtoMessage() {}

func (x *NotifyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMessage.ProtoReflect.Descriptor instead.
func (*NotifyMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{8}
}

func (x *NotifyMessage) GetAccessUserId() uint64 {
	if x != nil {
		return x.AccessUserId
	}
	return 0
}

func (x *NotifyMessage) GetSourceId() uint64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *NotifyMessage) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *NotifyMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NotifyMessage) GetSubtype() string {
	if x != nil {
		return x.Subtype
	}
	return ""
}

func (x *NotifyMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{9}
}

// 邮件通知
type EmailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessUserId            uint64 `protobuf:"varint,1,opt,name=accessUserId,proto3" json:"accessUserId,omitempty"` // 接收的用户id。 根据用户id获取
	ToMail                  string `protobuf:"bytes,2,opt,name=toMail,proto3" json:"toMail,omitempty"`
	Title                   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle                string `protobuf:"bytes,4,opt,name=subTitle,proto3" json:"subTitle,omitempty"`
	Content                 string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	AllowRepeatTimeInterval uint64 `protobuf:"varint,6,opt,name=allowRepeatTimeInterval,proto3" json:"allowRepeatTimeInterval,omitempty"` // 允许重复通知的时间间隔 单位秒
}

func (x *EmailMessage) Reset() {
	*x = EmailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jx3box_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailMessage) ProtoMessage() {}

func (x *EmailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jx3box_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailMessage.ProtoReflect.Descriptor instead.
func (*EmailMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_jx3box_proto_rawDescGZIP(), []int{10}
}

func (x *EmailMessage) GetAccessUserId() uint64 {
	if x != nil {
		return x.AccessUserId
	}
	return 0
}

func (x *EmailMessage) GetToMail() string {
	if x != nil {
		return x.ToMail
	}
	return ""
}

func (x *EmailMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EmailMessage) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *EmailMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *EmailMessage) GetAllowRepeatTimeInterval() uint64 {
	if x != nil {
		return x.AllowRepeatTimeInterval
	}
	return 0
}

var File_protobuf_jx3box_proto protoreflect.FileDescriptor

var file_protobuf_jx3box_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6a, 0x78, 0x33, 0x62, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x22,
	0x23, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x4d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x77,
	0x61, 0x72, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x77,
	0x61, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x22, 0x3e, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x4e, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0xd0, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x4d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x17,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x32, 0xb6, 0x03, 0x0a, 0x06, 0x4a, 0x58, 0x33, 0x42, 0x6f,
	0x78, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6a,
	0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e,
	0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x69, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x6a, 0x78, 0x33,
	0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x6a, 0x78, 0x33,
	0x62, 0x6f, 0x78, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15,
	0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x6a, 0x78,
	0x33, 0x62, 0x6f, 0x78, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0d, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54,
	0x6f, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x14, 0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f,
	0x78, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d,
	0x2e, 0x6a, 0x78, 0x33, 0x62, 0x6f, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6a, 0x78, 0x33, 0x62, 0x6f,
	0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protobuf_jx3box_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protobuf_jx3box_proto_goTypes = []interface{}{
	(*UserQueryParams)(nil),  // 0: jx3box.UserQueryParams
	(*User)(nil),             // 1: jx3box.User
	(*UserRenameParams)(nil), // 2: jx3box.UserRenameParams
	(*UserRenameResult)(nil), // 3: jx3box.UserRenameResult
	(*UserAwardParams)(nil),  // 4: jx3box.UserAwardParams
	(*UserAwardResult)(nil),  // 5: jx3box.UserAwardResult
	(*PostsQueryParams)(nil), // 6: jx3box.PostsQueryParams
	(*PostsQueryResult)(nil), // 7: jx3box.PostsQueryResult
	(*NotifyMessage)(nil),    // 8: jx3box.NotifyMessage
	(*Empty)(nil),            // 9: jx3box.Empty
	(*EmailMessage)(nil),     // 10: jx3box.EmailMessage
}
var file_protobuf_jx3box_proto_depIdxs = []int32{
	0,  // 0: jx3box.JX3Box.GetUser:input_type -> jx3box.UserQueryParams
	2,  // 1: jx3box.JX3Box.UserRename:input_type -> jx3box.UserRenameParams
	4,  // 2: jx3box.JX3Box.GiveUserAward:input_type -> jx3box.UserAwardParams
	6,  // 3: jx3box.JX3Box.GetPosts:input_type -> jx3box.PostsQueryParams
	8,  // 4: jx3box.JX3Box.SendNotify:input_type -> jx3box.NotifyMessage
	10, // 5: jx3box.JX3Box.SendEmailToUserID:input_type -> jx3box.EmailMessage
	10, // 6: jx3box.JX3Box.SendEmailToMailbox:input_type -> jx3box.EmailMessage
	1,  // 7: jx3box.JX3Box.GetUser:output_type -> jx3box.User
	3,  // 8: jx3box.JX3Box.UserRename:output_type -> jx3box.UserRenameResult
	5,  // 9: jx3box.JX3Box.GiveUserAward:output_type -> jx3box.UserAwardResult
	7,  // 10: jx3box.JX3Box.GetPosts:output_type -> jx3box.PostsQueryResult
	9,  // 11: jx3box.JX3Box.SendNotify:output_type -> jx3box.Empty
	9,  // 12: jx3box.JX3Box.SendEmailToUserID:output_type -> jx3box.Empty
	9,  // 13: jx3box.JX3Box.SendEmailToMailbox:output_type -> jx3box.Empty
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
		file_protobuf_jx3box_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protobuf_jx3box_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRenameResult); i {
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
		file_protobuf_jx3box_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAwardParams); i {
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
		file_protobuf_jx3box_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAwardResult); i {
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
		file_protobuf_jx3box_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsQueryParams); i {
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
		file_protobuf_jx3box_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsQueryResult); i {
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
		file_protobuf_jx3box_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyMessage); i {
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
		file_protobuf_jx3box_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protobuf_jx3box_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailMessage); i {
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
			NumMessages:   11,
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
