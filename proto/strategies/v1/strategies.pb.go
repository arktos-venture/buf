// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/strategies/v1/strategies.proto

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
	Interval_d Interval = 0
	Interval_w Interval = 1
	Interval_m Interval = 2
	Interval_y Interval = 3
)

// Enum value maps for Interval.
var (
	Interval_name = map[int32]string{
		0: "d",
		1: "w",
		2: "m",
		3: "y",
	}
	Interval_value = map[string]int32{
		"d": 0,
		"w": 1,
		"m": 2,
		"y": 3,
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
	return file_proto_strategies_v1_strategies_proto_enumTypes[0].Descriptor()
}

func (Interval) Type() protoreflect.EnumType {
	return &file_proto_strategies_v1_strategies_proto_enumTypes[0]
}

func (x Interval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interval.Descriptor instead.
func (Interval) EnumDescriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{0}
}

type TSDB int32

const (
	TSDB_Companies  TSDB = 0
	TSDB_Currencies TSDB = 1
	TSDB_Industries TSDB = 2
	TSDB_Exchanges  TSDB = 3
	TSDB_Countries  TSDB = 4
	TSDB_Indexes    TSDB = 5
	TSDB_Accounts   TSDB = 6
)

// Enum value maps for TSDB.
var (
	TSDB_name = map[int32]string{
		0: "Companies",
		1: "Currencies",
		2: "Industries",
		3: "Exchanges",
		4: "Countries",
		5: "Indexes",
		6: "Accounts",
	}
	TSDB_value = map[string]int32{
		"Companies":  0,
		"Currencies": 1,
		"Industries": 2,
		"Exchanges":  3,
		"Countries":  4,
		"Indexes":    5,
		"Accounts":   6,
	}
)

func (x TSDB) Enum() *TSDB {
	p := new(TSDB)
	*p = x
	return p
}

func (x TSDB) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TSDB) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_strategies_v1_strategies_proto_enumTypes[1].Descriptor()
}

func (TSDB) Type() protoreflect.EnumType {
	return &file_proto_strategies_v1_strategies_proto_enumTypes[1]
}

func (x TSDB) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TSDB.Descriptor instead.
func (TSDB) EnumDescriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{1}
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
	return file_proto_strategies_v1_strategies_proto_enumTypes[2].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_proto_strategies_v1_strategies_proto_enumTypes[2]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{2}
}

type Argument int32

const (
	Argument_Ticker   Argument = 0
	Argument_Exchange Argument = 1
	Argument_Industry Argument = 2
	Argument_Index    Argument = 3
	Argument_Country  Argument = 4
	Argument_Currency Argument = 5
	Argument_Time     Argument = 6
)

// Enum value maps for Argument.
var (
	Argument_name = map[int32]string{
		0: "Ticker",
		1: "Exchange",
		2: "Industry",
		3: "Index",
		4: "Country",
		5: "Currency",
		6: "Time",
	}
	Argument_value = map[string]int32{
		"Ticker":   0,
		"Exchange": 1,
		"Industry": 2,
		"Index":    3,
		"Country":  4,
		"Currency": 5,
		"Time":     6,
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
	return file_proto_strategies_v1_strategies_proto_enumTypes[3].Descriptor()
}

func (Argument) Type() protoreflect.EnumType {
	return &file_proto_strategies_v1_strategies_proto_enumTypes[3]
}

func (x Argument) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Argument.Descriptor instead.
func (Argument) EnumDescriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{3}
}

type StrategyReply_Action int32

const (
	StrategyReply_HOLD StrategyReply_Action = 0
	StrategyReply_BUY  StrategyReply_Action = 1
	StrategyReply_SELL StrategyReply_Action = 2
)

// Enum value maps for StrategyReply_Action.
var (
	StrategyReply_Action_name = map[int32]string{
		0: "HOLD",
		1: "BUY",
		2: "SELL",
	}
	StrategyReply_Action_value = map[string]int32{
		"HOLD": 0,
		"BUY":  1,
		"SELL": 2,
	}
)

func (x StrategyReply_Action) Enum() *StrategyReply_Action {
	p := new(StrategyReply_Action)
	*p = x
	return p
}

func (x StrategyReply_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StrategyReply_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_strategies_v1_strategies_proto_enumTypes[4].Descriptor()
}

func (StrategyReply_Action) Type() protoreflect.EnumType {
	return &file_proto_strategies_v1_strategies_proto_enumTypes[4]
}

func (x StrategyReply_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StrategyReply_Action.Descriptor instead.
func (StrategyReply_Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{3, 0}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=strategies.v1.Operator" json:"operator,omitempty"`
	Argument Argument `protobuf:"varint,2,opt,name=argument,proto3,enum=strategies.v1.Argument" json:"argument,omitempty"`
	Value    string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[0]
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
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{0}
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

type Strategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Params map[string]int32 `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Inject float32          `protobuf:"fixed32,3,opt,name=inject,proto3" json:"inject,omitempty"`
}

func (x *Strategy) Reset() {
	*x = Strategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strategy) ProtoMessage() {}

func (x *Strategy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strategy.ProtoReflect.Descriptor instead.
func (*Strategy) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{1}
}

func (x *Strategy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Strategy) GetParams() map[string]int32 {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Strategy) GetInject() float32 {
	if x != nil {
		return x.Inject
	}
	return 0
}

type StrategyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tsdb     TSDB      `protobuf:"varint,1,opt,name=tsdb,proto3,enum=strategies.v1.TSDB" json:"tsdb,omitempty"`
	Interval Interval  `protobuf:"varint,2,opt,name=interval,proto3,enum=strategies.v1.Interval" json:"interval,omitempty"`
	Strategy *Strategy `protobuf:"bytes,3,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Filters  []*Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *StrategyRequest) Reset() {
	*x = StrategyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyRequest) ProtoMessage() {}

func (x *StrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyRequest.ProtoReflect.Descriptor instead.
func (*StrategyRequest) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{2}
}

func (x *StrategyRequest) GetTsdb() TSDB {
	if x != nil {
		return x.Tsdb
	}
	return TSDB_Companies
}

func (x *StrategyRequest) GetInterval() Interval {
	if x != nil {
		return x.Interval
	}
	return Interval_d
}

func (x *StrategyRequest) GetStrategy() *Strategy {
	if x != nil {
		return x.Strategy
	}
	return nil
}

func (x *StrategyRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type StrategyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action        StrategyReply_Action   `protobuf:"varint,1,opt,name=action,proto3,enum=strategies.v1.StrategyReply_Action" json:"action,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	LastQuote     float32                `protobuf:"fixed32,3,opt,name=last_quote,json=lastQuote,proto3" json:"last_quote,omitempty"`
	LastQuoteDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_quote_date,json=lastQuoteDate,proto3" json:"last_quote_date,omitempty"`
}

func (x *StrategyReply) Reset() {
	*x = StrategyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyReply) ProtoMessage() {}

func (x *StrategyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyReply.ProtoReflect.Descriptor instead.
func (*StrategyReply) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{3}
}

func (x *StrategyReply) GetAction() StrategyReply_Action {
	if x != nil {
		return x.Action
	}
	return StrategyReply_HOLD
}

func (x *StrategyReply) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *StrategyReply) GetLastQuote() float32 {
	if x != nil {
		return x.LastQuote
	}
	return 0
}

func (x *StrategyReply) GetLastQuoteDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastQuoteDate
	}
	return nil
}

var File_proto_strategies_v1_strategies_proto protoreflect.FileDescriptor

var file_proto_strategies_v1_strategies_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x06, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x08, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x9a, 0x01, 0x04, 0x08, 0x02, 0x10, 0x0a, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfb, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x73,
	0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x53, 0x44, 0x42, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x73, 0x64, 0x62, 0x12, 0x3d, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x08,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x41, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x14, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55,
	0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x02, 0x2a, 0x26, 0x0a,
	0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x05, 0x0a, 0x01, 0x64, 0x10, 0x00,
	0x12, 0x05, 0x0a, 0x01, 0x77, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x6d, 0x10, 0x02, 0x12, 0x05,
	0x0a, 0x01, 0x79, 0x10, 0x03, 0x2a, 0x6e, 0x0a, 0x04, 0x54, 0x53, 0x44, 0x42, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x73, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x10, 0x06, 0x2a, 0x4d, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x72, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x65, 0x73, 0x73, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x65, 0x73, 0x73, 0x45, 0x71, 0x75,
	0x61, 0x6c, 0x10, 0x04, 0x2a, 0x62, 0x0a, 0x08, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0a, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x05, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x06, 0x32, 0xc2, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x52, 0x0a,
	0x1c, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31,
	0x50, 0x01, 0x5a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_strategies_v1_strategies_proto_rawDescOnce sync.Once
	file_proto_strategies_v1_strategies_proto_rawDescData = file_proto_strategies_v1_strategies_proto_rawDesc
)

func file_proto_strategies_v1_strategies_proto_rawDescGZIP() []byte {
	file_proto_strategies_v1_strategies_proto_rawDescOnce.Do(func() {
		file_proto_strategies_v1_strategies_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_strategies_v1_strategies_proto_rawDescData)
	})
	return file_proto_strategies_v1_strategies_proto_rawDescData
}

var file_proto_strategies_v1_strategies_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_proto_strategies_v1_strategies_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_strategies_v1_strategies_proto_goTypes = []interface{}{
	(Interval)(0),                 // 0: strategies.v1.Interval
	(TSDB)(0),                     // 1: strategies.v1.TSDB
	(Operator)(0),                 // 2: strategies.v1.Operator
	(Argument)(0),                 // 3: strategies.v1.Argument
	(StrategyReply_Action)(0),     // 4: strategies.v1.StrategyReply.Action
	(*Filter)(nil),                // 5: strategies.v1.Filter
	(*Strategy)(nil),              // 6: strategies.v1.Strategy
	(*StrategyRequest)(nil),       // 7: strategies.v1.StrategyRequest
	(*StrategyReply)(nil),         // 8: strategies.v1.StrategyReply
	nil,                           // 9: strategies.v1.Strategy.ParamsEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_proto_strategies_v1_strategies_proto_depIdxs = []int32{
	2,  // 0: strategies.v1.Filter.operator:type_name -> strategies.v1.Operator
	3,  // 1: strategies.v1.Filter.argument:type_name -> strategies.v1.Argument
	9,  // 2: strategies.v1.Strategy.params:type_name -> strategies.v1.Strategy.ParamsEntry
	1,  // 3: strategies.v1.StrategyRequest.tsdb:type_name -> strategies.v1.TSDB
	0,  // 4: strategies.v1.StrategyRequest.interval:type_name -> strategies.v1.Interval
	6,  // 5: strategies.v1.StrategyRequest.strategy:type_name -> strategies.v1.Strategy
	5,  // 6: strategies.v1.StrategyRequest.filters:type_name -> strategies.v1.Filter
	4,  // 7: strategies.v1.StrategyReply.action:type_name -> strategies.v1.StrategyReply.Action
	10, // 8: strategies.v1.StrategyReply.date:type_name -> google.protobuf.Timestamp
	10, // 9: strategies.v1.StrategyReply.last_quote_date:type_name -> google.protobuf.Timestamp
	7,  // 10: strategies.v1.Strategies.Execute:input_type -> strategies.v1.StrategyRequest
	11, // 11: strategies.v1.Strategies.Health:input_type -> google.protobuf.Empty
	8,  // 12: strategies.v1.Strategies.Execute:output_type -> strategies.v1.StrategyReply
	11, // 13: strategies.v1.Strategies.Health:output_type -> google.protobuf.Empty
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_strategies_v1_strategies_proto_init() }
func file_proto_strategies_v1_strategies_proto_init() {
	if File_proto_strategies_v1_strategies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_strategies_v1_strategies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strategy); i {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyRequest); i {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyReply); i {
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
			RawDescriptor: file_proto_strategies_v1_strategies_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_strategies_v1_strategies_proto_goTypes,
		DependencyIndexes: file_proto_strategies_v1_strategies_proto_depIdxs,
		EnumInfos:         file_proto_strategies_v1_strategies_proto_enumTypes,
		MessageInfos:      file_proto_strategies_v1_strategies_proto_msgTypes,
	}.Build()
	File_proto_strategies_v1_strategies_proto = out.File
	file_proto_strategies_v1_strategies_proto_rawDesc = nil
	file_proto_strategies_v1_strategies_proto_goTypes = nil
	file_proto_strategies_v1_strategies_proto_depIdxs = nil
}
