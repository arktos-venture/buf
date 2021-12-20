// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/account/v1/account.proto

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

type Broker int32

const (
	Broker_IBKR Broker = 0
)

// Enum value maps for Broker.
var (
	Broker_name = map[int32]string{
		0: "IBKR",
	}
	Broker_value = map[string]int32{
		"IBKR": 0,
	}
)

func (x Broker) Enum() *Broker {
	p := new(Broker)
	*p = x
	return p
}

func (x Broker) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Broker) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_account_v1_account_proto_enumTypes[0].Descriptor()
}

func (Broker) Type() protoreflect.EnumType {
	return &file_proto_account_v1_account_proto_enumTypes[0]
}

func (x Broker) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Broker.Descriptor instead.
func (Broker) EnumDescriptor() ([]byte, []int) {
	return file_proto_account_v1_account_proto_rawDescGZIP(), []int{0}
}

type AccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Broker   Broker `protobuf:"varint,1,opt,name=broker,proto3,enum=account.v1.Broker" json:"broker,omitempty"`
	Account  string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountRequest) GetBroker() Broker {
	if x != nil {
		return x.Broker
	}
	return Broker_IBKR
}

func (x *AccountRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type AccountCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Broker   Broker `protobuf:"varint,1,opt,name=broker,proto3,enum=account.v1.Broker" json:"broker,omitempty"`
	Account  string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *AccountCreateRequest) Reset() {
	*x = AccountCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreateRequest) ProtoMessage() {}

func (x *AccountCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreateRequest.ProtoReflect.Descriptor instead.
func (*AccountCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountCreateRequest) GetBroker() Broker {
	if x != nil {
		return x.Broker
	}
	return Broker_IBKR
}

func (x *AccountCreateRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountCreateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type AccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Active   bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Date     int64  `protobuf:"varint,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *AccountReply) Reset() {
	*x = AccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountReply) ProtoMessage() {}

func (x *AccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountReply.ProtoReflect.Descriptor instead.
func (*AccountReply) Descriptor() ([]byte, []int) {
	return file_proto_account_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountReply) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountReply) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AccountReply) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *AccountReply) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *AccountReply) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

var File_proto_account_v1_account_proto protoreflect.FileDescriptor

var file_proto_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x91, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x03, 0x18, 0x0f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x97, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x0f, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x98, 0x01, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x84,
	0x01, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x2a, 0x12, 0x0a, 0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x42, 0x4b, 0x52, 0x10, 0x00, 0x32, 0xa1, 0x02, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x12, 0x65, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x49, 0x0a,
	0x19, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_v1_account_proto_rawDescOnce sync.Once
	file_proto_account_v1_account_proto_rawDescData = file_proto_account_v1_account_proto_rawDesc
)

func file_proto_account_v1_account_proto_rawDescGZIP() []byte {
	file_proto_account_v1_account_proto_rawDescOnce.Do(func() {
		file_proto_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_v1_account_proto_rawDescData)
	})
	return file_proto_account_v1_account_proto_rawDescData
}

var file_proto_account_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_account_v1_account_proto_goTypes = []interface{}{
	(Broker)(0),                  // 0: account.v1.Broker
	(*AccountRequest)(nil),       // 1: account.v1.AccountRequest
	(*AccountCreateRequest)(nil), // 2: account.v1.AccountCreateRequest
	(*AccountReply)(nil),         // 3: account.v1.AccountReply
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_proto_account_v1_account_proto_depIdxs = []int32{
	0, // 0: account.v1.AccountRequest.broker:type_name -> account.v1.Broker
	0, // 1: account.v1.AccountCreateRequest.broker:type_name -> account.v1.Broker
	1, // 2: account.v1.Account.Get:input_type -> account.v1.AccountRequest
	2, // 3: account.v1.Account.Create:input_type -> account.v1.AccountCreateRequest
	4, // 4: account.v1.Account.Health:input_type -> google.protobuf.Empty
	3, // 5: account.v1.Account.Get:output_type -> account.v1.AccountReply
	3, // 6: account.v1.Account.Create:output_type -> account.v1.AccountReply
	4, // 7: account.v1.Account.Health:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_account_v1_account_proto_init() }
func file_proto_account_v1_account_proto_init() {
	if File_proto_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRequest); i {
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
		file_proto_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCreateRequest); i {
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
		file_proto_account_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountReply); i {
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
			RawDescriptor: file_proto_account_v1_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_account_v1_account_proto_goTypes,
		DependencyIndexes: file_proto_account_v1_account_proto_depIdxs,
		EnumInfos:         file_proto_account_v1_account_proto_enumTypes,
		MessageInfos:      file_proto_account_v1_account_proto_msgTypes,
	}.Build()
	File_proto_account_v1_account_proto = out.File
	file_proto_account_v1_account_proto_rawDesc = nil
	file_proto_account_v1_account_proto_goTypes = nil
	file_proto_account_v1_account_proto_depIdxs = nil
}
