// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/notification/v1/notification.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Level int32

const (
	Level_CRITICAL Level = 0
	Level_URGENT   Level = 1
	Level_NORMAL   Level = 2
	Level_LOW      Level = 3
	Level_NONE     Level = 4
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "CRITICAL",
		1: "URGENT",
		2: "NORMAL",
		3: "LOW",
		4: "NONE",
	}
	Level_value = map[string]int32{
		"CRITICAL": 0,
		"URGENT":   1,
		"NORMAL":   2,
		"LOW":      3,
		"NONE":     4,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_notification_v1_notification_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_proto_notification_v1_notification_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{0}
}

type Subject int32

const (
	Subject_LOW_CASH             Subject = 0
	Subject_HIGH_CASH            Subject = 1
	Subject_UPTO_2P_VARIATION    Subject = 2
	Subject_UPTO_5P_VARIATION    Subject = 3
	Subject_UPTO_10P_VARIATION   Subject = 4
	Subject_DOWNTO_2P_VARIATION  Subject = 5
	Subject_DOWNTO_5P_VARIATION  Subject = 6
	Subject_DOWNTO_10P_VARIATION Subject = 7
)

// Enum value maps for Subject.
var (
	Subject_name = map[int32]string{
		0: "LOW_CASH",
		1: "HIGH_CASH",
		2: "UPTO_2P_VARIATION",
		3: "UPTO_5P_VARIATION",
		4: "UPTO_10P_VARIATION",
		5: "DOWNTO_2P_VARIATION",
		6: "DOWNTO_5P_VARIATION",
		7: "DOWNTO_10P_VARIATION",
	}
	Subject_value = map[string]int32{
		"LOW_CASH":             0,
		"HIGH_CASH":            1,
		"UPTO_2P_VARIATION":    2,
		"UPTO_5P_VARIATION":    3,
		"UPTO_10P_VARIATION":   4,
		"DOWNTO_2P_VARIATION":  5,
		"DOWNTO_5P_VARIATION":  6,
		"DOWNTO_10P_VARIATION": 7,
	}
)

func (x Subject) Enum() *Subject {
	p := new(Subject)
	*p = x
	return p
}

func (x Subject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subject) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_notification_v1_notification_proto_enumTypes[1].Descriptor()
}

func (Subject) Type() protoreflect.EnumType {
	return &file_proto_notification_v1_notification_proto_enumTypes[1]
}

func (x Subject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subject.Descriptor instead.
func (Subject) EnumDescriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{1}
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_v1_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_v1_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Page) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type NotificationCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Level       Level   `protobuf:"varint,2,opt,name=level,proto3,enum=notification.v1.Level" json:"level,omitempty"`
	Subject     Subject `protobuf:"varint,4,opt,name=subject,proto3,enum=notification.v1.Subject" json:"subject,omitempty"`
	Title       string  `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *NotificationCreateRequest) Reset() {
	*x = NotificationCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_v1_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationCreateRequest) ProtoMessage() {}

func (x *NotificationCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_v1_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationCreateRequest.ProtoReflect.Descriptor instead.
func (*NotificationCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationCreateRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *NotificationCreateRequest) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_CRITICAL
}

func (x *NotificationCreateRequest) GetSubject() Subject {
	if x != nil {
		return x.Subject
	}
	return Subject_LOW_CASH
}

func (x *NotificationCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotificationCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type NotificationSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Level   Level  `protobuf:"varint,2,opt,name=level,proto3,enum=notification.v1.Level" json:"level,omitempty"`
	Date    string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Page    *Page  `protobuf:"bytes,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *NotificationSearchRequest) Reset() {
	*x = NotificationSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_v1_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationSearchRequest) ProtoMessage() {}

func (x *NotificationSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_v1_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationSearchRequest.ProtoReflect.Descriptor instead.
func (*NotificationSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationSearchRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *NotificationSearchRequest) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_CRITICAL
}

func (x *NotificationSearchRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *NotificationSearchRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type NotificationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject     Subject `protobuf:"varint,1,opt,name=subject,proto3,enum=notification.v1.Subject" json:"subject,omitempty"`
	Level       Level   `protobuf:"varint,2,opt,name=level,proto3,enum=notification.v1.Level" json:"level,omitempty"`
	Title       string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Date        int64   `protobuf:"varint,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *NotificationReply) Reset() {
	*x = NotificationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_v1_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationReply) ProtoMessage() {}

func (x *NotificationReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_v1_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationReply.ProtoReflect.Descriptor instead.
func (*NotificationReply) Descriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationReply) GetSubject() Subject {
	if x != nil {
		return x.Subject
	}
	return Subject_LOW_CASH
}

func (x *NotificationReply) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_CRITICAL
}

func (x *NotificationReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotificationReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotificationReply) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type NotificationReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*NotificationReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *NotificationReplies) Reset() {
	*x = NotificationReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_v1_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationReplies) ProtoMessage() {}

func (x *NotificationReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_v1_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationReplies.ProtoReflect.Descriptor instead.
func (*NotificationReplies) Descriptor() ([]byte, []int) {
	return file_proto_notification_v1_notification_proto_rawDescGZIP(), []int{4}
}

func (x *NotificationReplies) GetResults() []*NotificationReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *NotificationReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_notification_v1_notification_proto protoreflect.FileDescriptor

var file_proto_notification_v1_notification_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4c, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x1a, 0x05, 0x10, 0x90,
	0x4e, 0x20, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x1a,
	0x05, 0x10, 0x96, 0x01, 0x20, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xd9, 0x01,
	0x0a, 0x19, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x32, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd3, 0x01, 0x0a, 0x19, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x02, 0x33, 0x64, 0x52, 0x02, 0x31, 0x77, 0x52, 0x02, 0x32, 0x77,
	0x52, 0x02, 0x31, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0xc1, 0x01, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x69, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x40,
	0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49,
	0x43, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x4c, 0x4f, 0x57, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x04,
	0x2a, 0xb8, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x41, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x49,
	0x47, 0x48, 0x5f, 0x43, 0x41, 0x53, 0x48, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x50, 0x54,
	0x4f, 0x5f, 0x32, 0x50, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x55, 0x50, 0x54, 0x4f, 0x5f, 0x35, 0x50, 0x5f, 0x56, 0x41, 0x52, 0x49,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x50, 0x54, 0x4f, 0x5f,
	0x31, 0x30, 0x50, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12,
	0x17, 0x0a, 0x13, 0x44, 0x4f, 0x57, 0x4e, 0x54, 0x4f, 0x5f, 0x32, 0x50, 0x5f, 0x56, 0x41, 0x52,
	0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x4f, 0x57, 0x4e,
	0x54, 0x4f, 0x5f, 0x35, 0x50, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x06, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x4f, 0x57, 0x4e, 0x54, 0x4f, 0x5f, 0x31, 0x30, 0x50, 0x5f,
	0x56, 0x41, 0x52, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x32, 0xdc, 0x02, 0x0a, 0x0c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7f, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x58, 0x0a, 0x1e, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notification_v1_notification_proto_rawDescOnce sync.Once
	file_proto_notification_v1_notification_proto_rawDescData = file_proto_notification_v1_notification_proto_rawDesc
)

func file_proto_notification_v1_notification_proto_rawDescGZIP() []byte {
	file_proto_notification_v1_notification_proto_rawDescOnce.Do(func() {
		file_proto_notification_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notification_v1_notification_proto_rawDescData)
	})
	return file_proto_notification_v1_notification_proto_rawDescData
}

var file_proto_notification_v1_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_notification_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_notification_v1_notification_proto_goTypes = []interface{}{
	(Level)(0),                        // 0: notification.v1.Level
	(Subject)(0),                      // 1: notification.v1.Subject
	(*Page)(nil),                      // 2: notification.v1.Page
	(*NotificationCreateRequest)(nil), // 3: notification.v1.NotificationCreateRequest
	(*NotificationSearchRequest)(nil), // 4: notification.v1.NotificationSearchRequest
	(*NotificationReply)(nil),         // 5: notification.v1.NotificationReply
	(*NotificationReplies)(nil),       // 6: notification.v1.NotificationReplies
	(*emptypb.Empty)(nil),             // 7: google.protobuf.Empty
}
var file_proto_notification_v1_notification_proto_depIdxs = []int32{
	0,  // 0: notification.v1.NotificationCreateRequest.level:type_name -> notification.v1.Level
	1,  // 1: notification.v1.NotificationCreateRequest.subject:type_name -> notification.v1.Subject
	0,  // 2: notification.v1.NotificationSearchRequest.level:type_name -> notification.v1.Level
	2,  // 3: notification.v1.NotificationSearchRequest.page:type_name -> notification.v1.Page
	1,  // 4: notification.v1.NotificationReply.subject:type_name -> notification.v1.Subject
	0,  // 5: notification.v1.NotificationReply.level:type_name -> notification.v1.Level
	5,  // 6: notification.v1.NotificationReplies.results:type_name -> notification.v1.NotificationReply
	3,  // 7: notification.v1.Notification.Create:input_type -> notification.v1.NotificationCreateRequest
	4,  // 8: notification.v1.Notification.Search:input_type -> notification.v1.NotificationSearchRequest
	7,  // 9: notification.v1.Notification.Health:input_type -> google.protobuf.Empty
	5,  // 10: notification.v1.Notification.Create:output_type -> notification.v1.NotificationReply
	6,  // 11: notification.v1.Notification.Search:output_type -> notification.v1.NotificationReplies
	7,  // 12: notification.v1.Notification.Health:output_type -> google.protobuf.Empty
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_notification_v1_notification_proto_init() }
func file_proto_notification_v1_notification_proto_init() {
	if File_proto_notification_v1_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_notification_v1_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_proto_notification_v1_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationCreateRequest); i {
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
		file_proto_notification_v1_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationSearchRequest); i {
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
		file_proto_notification_v1_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationReply); i {
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
		file_proto_notification_v1_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationReplies); i {
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
			RawDescriptor: file_proto_notification_v1_notification_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notification_v1_notification_proto_goTypes,
		DependencyIndexes: file_proto_notification_v1_notification_proto_depIdxs,
		EnumInfos:         file_proto_notification_v1_notification_proto_enumTypes,
		MessageInfos:      file_proto_notification_v1_notification_proto_msgTypes,
	}.Build()
	File_proto_notification_v1_notification_proto = out.File
	file_proto_notification_v1_notification_proto_rawDesc = nil
	file_proto_notification_v1_notification_proto_goTypes = nil
	file_proto_notification_v1_notification_proto_depIdxs = nil
}
