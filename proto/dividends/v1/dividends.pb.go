// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/dividends/v1/dividends.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Interval int32

const (
	Interval_m Interval = 0
	Interval_q Interval = 1
	Interval_y Interval = 2
)

// Enum value maps for Interval.
var (
	Interval_name = map[int32]string{
		0: "m",
		1: "q",
		2: "y",
	}
	Interval_value = map[string]int32{
		"m": 0,
		"q": 1,
		"y": 2,
	}
)

func (x Interval) Enum() *Interval {
	p := new(Interval)
	*p = x
	return p
}

func (x Interval) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interval) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dividends_v1_dividends_proto_enumTypes[0].Descriptor()
}

func (Interval) Type() protoreflect.EnumType {
	return &file_proto_dividends_v1_dividends_proto_enumTypes[0]
}

func (x Interval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interval.Descriptor instead.
func (Interval) EnumDescriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{0}
}

type Operator int32

const (
	Operator_Equal        Operator = 0
	Operator_Greater      Operator = 1
	Operator_GreaterEqual Operator = 2
	Operator_Less         Operator = 3
	Operator_LessEqual    Operator = 4
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "Equal",
		1: "Greater",
		2: "GreaterEqual",
		3: "Less",
		4: "LessEqual",
	}
	Operator_value = map[string]int32{
		"Equal":        0,
		"Greater":      1,
		"GreaterEqual": 2,
		"Less":         3,
		"LessEqual":    4,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dividends_v1_dividends_proto_enumTypes[1].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_proto_dividends_v1_dividends_proto_enumTypes[1]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{1}
}

type Argument int32

const (
	Argument_Ticker   Argument = 0
	Argument_Exchange Argument = 1
	Argument_Time     Argument = 2
)

// Enum value maps for Argument.
var (
	Argument_name = map[int32]string{
		0: "Ticker",
		1: "Exchange",
		2: "Time",
	}
	Argument_value = map[string]int32{
		"Ticker":   0,
		"Exchange": 1,
		"Time":     2,
	}
)

func (x Argument) Enum() *Argument {
	p := new(Argument)
	*p = x
	return p
}

func (x Argument) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Argument) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dividends_v1_dividends_proto_enumTypes[2].Descriptor()
}

func (Argument) Type() protoreflect.EnumType {
	return &file_proto_dividends_v1_dividends_proto_enumTypes[2]
}

func (x Argument) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Argument.Descriptor instead.
func (Argument) EnumDescriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{2}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=dividends.v1.Operator" json:"operator,omitempty"`
	Argument Argument `protobuf:"varint,2,opt,name=argument,proto3,enum=dividends.v1.Argument" json:"argument,omitempty"`
	Value    string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_Equal
}

func (x *Filter) GetArgument() Argument {
	if x != nil {
		return x.Argument
	}
	return Argument_Ticker
}

func (x *Filter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DividendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval Interval  `protobuf:"varint,1,opt,name=interval,proto3,enum=dividends.v1.Interval" json:"interval,omitempty"`
	Filters  []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *DividendsRequest) Reset() {
	*x = DividendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DividendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DividendsRequest) ProtoMessage() {}

func (x *DividendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DividendsRequest.ProtoReflect.Descriptor instead.
func (*DividendsRequest) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{1}
}

func (x *DividendsRequest) GetInterval() Interval {
	if x != nil {
		return x.Interval
	}
	return Interval_m
}

func (x *DividendsRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type DividendsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values          []float32                `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	DeclarationDate []*timestamppb.Timestamp `protobuf:"bytes,2,rep,name=declarationDate,proto3" json:"declarationDate,omitempty"`
	RecordDate      []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=recordDate,proto3" json:"recordDate,omitempty"`
	PaymentDate     []*timestamppb.Timestamp `protobuf:"bytes,4,rep,name=paymentDate,proto3" json:"paymentDate,omitempty"`
	CreatedAt       []*timestamppb.Timestamp `protobuf:"bytes,5,rep,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *DividendsReply) Reset() {
	*x = DividendsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DividendsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DividendsReply) ProtoMessage() {}

func (x *DividendsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DividendsReply.ProtoReflect.Descriptor instead.
func (*DividendsReply) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{2}
}

func (x *DividendsReply) GetValues() []float32 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *DividendsReply) GetDeclarationDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.DeclarationDate
	}
	return nil
}

func (x *DividendsReply) GetRecordDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.RecordDate
	}
	return nil
}

func (x *DividendsReply) GetPaymentDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDate
	}
	return nil
}

func (x *DividendsReply) GetCreatedAt() []*timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_proto_dividends_v1_dividends_proto protoreflect.FileDescriptor

var file_proto_dividends_v1_dividends_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x3c, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x3c, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x10,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x14,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xa3, 0x02, 0x0a, 0x0e, 0x44, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x64, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a,
	0x1f, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x05, 0x0a, 0x01, 0x6d,
	0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x71, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x79, 0x10, 0x02,
	0x2a, 0x4d, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x45,
	0x71, 0x75, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x65, 0x73, 0x73, 0x10, 0x03,
	0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x65, 0x73, 0x73, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x04, 0x2a,
	0x2e, 0x0a, 0x08, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x02, 0x32,
	0xb9, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x60, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x4c, 0x0a, 0x1a, 0x64,
	0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_dividends_v1_dividends_proto_rawDescOnce sync.Once
	file_proto_dividends_v1_dividends_proto_rawDescData = file_proto_dividends_v1_dividends_proto_rawDesc
)

func file_proto_dividends_v1_dividends_proto_rawDescGZIP() []byte {
	file_proto_dividends_v1_dividends_proto_rawDescOnce.Do(func() {
		file_proto_dividends_v1_dividends_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dividends_v1_dividends_proto_rawDescData)
	})
	return file_proto_dividends_v1_dividends_proto_rawDescData
}

var file_proto_dividends_v1_dividends_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_dividends_v1_dividends_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_dividends_v1_dividends_proto_goTypes = []interface{}{
	(Interval)(0),                 // 0: dividends.v1.Interval
	(Operator)(0),                 // 1: dividends.v1.Operator
	(Argument)(0),                 // 2: dividends.v1.Argument
	(*Filter)(nil),                // 3: dividends.v1.Filter
	(*DividendsRequest)(nil),      // 4: dividends.v1.DividendsRequest
	(*DividendsReply)(nil),        // 5: dividends.v1.DividendsReply
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_proto_dividends_v1_dividends_proto_depIdxs = []int32{
	1,  // 0: dividends.v1.Filter.operator:type_name -> dividends.v1.Operator
	2,  // 1: dividends.v1.Filter.argument:type_name -> dividends.v1.Argument
	0,  // 2: dividends.v1.DividendsRequest.interval:type_name -> dividends.v1.Interval
	3,  // 3: dividends.v1.DividendsRequest.filters:type_name -> dividends.v1.Filter
	6,  // 4: dividends.v1.DividendsReply.declarationDate:type_name -> google.protobuf.Timestamp
	6,  // 5: dividends.v1.DividendsReply.recordDate:type_name -> google.protobuf.Timestamp
	6,  // 6: dividends.v1.DividendsReply.paymentDate:type_name -> google.protobuf.Timestamp
	6,  // 7: dividends.v1.DividendsReply.created_at:type_name -> google.protobuf.Timestamp
	4,  // 8: dividends.v1.Dividends.Search:input_type -> dividends.v1.DividendsRequest
	7,  // 9: dividends.v1.Dividends.Health:input_type -> google.protobuf.Empty
	5,  // 10: dividends.v1.Dividends.Search:output_type -> dividends.v1.DividendsReply
	7,  // 11: dividends.v1.Dividends.Health:output_type -> google.protobuf.Empty
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_dividends_v1_dividends_proto_init() }
func file_proto_dividends_v1_dividends_proto_init() {
	if File_proto_dividends_v1_dividends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dividends_v1_dividends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_proto_dividends_v1_dividends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DividendsRequest); i {
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
		file_proto_dividends_v1_dividends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DividendsReply); i {
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
			RawDescriptor: file_proto_dividends_v1_dividends_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dividends_v1_dividends_proto_goTypes,
		DependencyIndexes: file_proto_dividends_v1_dividends_proto_depIdxs,
		EnumInfos:         file_proto_dividends_v1_dividends_proto_enumTypes,
		MessageInfos:      file_proto_dividends_v1_dividends_proto_msgTypes,
	}.Build()
	File_proto_dividends_v1_dividends_proto = out.File
	file_proto_dividends_v1_dividends_proto_rawDesc = nil
	file_proto_dividends_v1_dividends_proto_goTypes = nil
	file_proto_dividends_v1_dividends_proto_depIdxs = nil
}
