// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/quotes/v1/quotes.proto

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
	return file_proto_quotes_v1_quotes_proto_enumTypes[0].Descriptor()
}

func (TSDB) Type() protoreflect.EnumType {
	return &file_proto_quotes_v1_quotes_proto_enumTypes[0]
}

func (x TSDB) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TSDB.Descriptor instead.
func (TSDB) EnumDescriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{0}
}

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
	return file_proto_quotes_v1_quotes_proto_enumTypes[1].Descriptor()
}

func (Interval) Type() protoreflect.EnumType {
	return &file_proto_quotes_v1_quotes_proto_enumTypes[1]
}

func (x Interval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interval.Descriptor instead.
func (Interval) EnumDescriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{1}
}

type Indicator int32

const (
	Indicator_Macd      Indicator = 0
	Indicator_Rsi       Indicator = 1
	Indicator_Obv       Indicator = 2
	Indicator_Adx       Indicator = 3
	Indicator_Beta      Indicator = 4
	Indicator_Atr       Indicator = 5
	Indicator_Bband     Indicator = 6
	Indicator_Ema       Indicator = 7
	Indicator_LinearReg Indicator = 8
	Indicator_WillR     Indicator = 9
	Indicator_StdDev    Indicator = 10
	Indicator_Sma10     Indicator = 11
	Indicator_Sma20     Indicator = 12
	Indicator_Sma50     Indicator = 13
	Indicator_Sma100    Indicator = 14
	Indicator_Sma200    Indicator = 15
)

// Enum value maps for Indicator.
var (
	Indicator_name = map[int32]string{
		0:  "Macd",
		1:  "Rsi",
		2:  "Obv",
		3:  "Adx",
		4:  "Beta",
		5:  "Atr",
		6:  "Bband",
		7:  "Ema",
		8:  "LinearReg",
		9:  "WillR",
		10: "StdDev",
		11: "Sma10",
		12: "Sma20",
		13: "Sma50",
		14: "Sma100",
		15: "Sma200",
	}
	Indicator_value = map[string]int32{
		"Macd":      0,
		"Rsi":       1,
		"Obv":       2,
		"Adx":       3,
		"Beta":      4,
		"Atr":       5,
		"Bband":     6,
		"Ema":       7,
		"LinearReg": 8,
		"WillR":     9,
		"StdDev":    10,
		"Sma10":     11,
		"Sma20":     12,
		"Sma50":     13,
		"Sma100":    14,
		"Sma200":    15,
	}
)

func (x Indicator) Enum() *Indicator {
	p := new(Indicator)
	*p = x
	return p
}

func (x Indicator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Indicator) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_quotes_v1_quotes_proto_enumTypes[2].Descriptor()
}

func (Indicator) Type() protoreflect.EnumType {
	return &file_proto_quotes_v1_quotes_proto_enumTypes[2]
}

func (x Indicator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Indicator.Descriptor instead.
func (Indicator) EnumDescriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{2}
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
	return file_proto_quotes_v1_quotes_proto_enumTypes[3].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_proto_quotes_v1_quotes_proto_enumTypes[3]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{3}
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
	return file_proto_quotes_v1_quotes_proto_enumTypes[4].Descriptor()
}

func (Argument) Type() protoreflect.EnumType {
	return &file_proto_quotes_v1_quotes_proto_enumTypes[4]
}

func (x Argument) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Argument.Descriptor instead.
func (Argument) EnumDescriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{4}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=quotes.v1.Operator" json:"operator,omitempty"`
	Argument Argument `protobuf:"varint,2,opt,name=argument,proto3,enum=quotes.v1.Argument" json:"argument,omitempty"`
	Value    string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_v1_quotes_proto_msgTypes[0]
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
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{0}
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

type QuotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tsdb     TSDB      `protobuf:"varint,1,opt,name=tsdb,proto3,enum=quotes.v1.TSDB" json:"tsdb,omitempty"`
	Interval Interval  `protobuf:"varint,2,opt,name=interval,proto3,enum=quotes.v1.Interval" json:"interval,omitempty"`
	Filters  []*Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *QuotesRequest) Reset() {
	*x = QuotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotesRequest) ProtoMessage() {}

func (x *QuotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_v1_quotes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotesRequest.ProtoReflect.Descriptor instead.
func (*QuotesRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{1}
}

func (x *QuotesRequest) GetTsdb() TSDB {
	if x != nil {
		return x.Tsdb
	}
	return TSDB_Companies
}

func (x *QuotesRequest) GetInterval() Interval {
	if x != nil {
		return x.Interval
	}
	return Interval_d
}

func (x *QuotesRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type QuotesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open      []float32                `protobuf:"fixed32,1,rep,packed,name=open,proto3" json:"open,omitempty"`
	Close     []float32                `protobuf:"fixed32,2,rep,packed,name=close,proto3" json:"close,omitempty"`
	Low       []float32                `protobuf:"fixed32,3,rep,packed,name=low,proto3" json:"low,omitempty"`
	High      []float32                `protobuf:"fixed32,4,rep,packed,name=high,proto3" json:"high,omitempty"`
	Volume    []float32                `protobuf:"fixed32,5,rep,packed,name=volume,proto3" json:"volume,omitempty"`
	CreatedAt []*timestamppb.Timestamp `protobuf:"bytes,7,rep,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *QuotesReply) Reset() {
	*x = QuotesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotesReply) ProtoMessage() {}

func (x *QuotesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_v1_quotes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotesReply.ProtoReflect.Descriptor instead.
func (*QuotesReply) Descriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{2}
}

func (x *QuotesReply) GetOpen() []float32 {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *QuotesReply) GetClose() []float32 {
	if x != nil {
		return x.Close
	}
	return nil
}

func (x *QuotesReply) GetLow() []float32 {
	if x != nil {
		return x.Low
	}
	return nil
}

func (x *QuotesReply) GetHigh() []float32 {
	if x != nil {
		return x.High
	}
	return nil
}

func (x *QuotesReply) GetVolume() []float32 {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *QuotesReply) GetCreatedAt() []*timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_proto_quotes_v1_quotes_proto protoreflect.FileDescriptor

var file_proto_quotes_v1_quotes_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x73, 0x64, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x53, 0x44, 0x42, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x74, 0x73, 0x64, 0x62, 0x12, 0x39, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x3d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa,
	0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x14, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x22, 0xb0, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x77, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x02, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x2a, 0x6e, 0x0a, 0x04, 0x54, 0x53, 0x44, 0x42, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x10, 0x06, 0x2a, 0x26, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x05, 0x0a, 0x01, 0x64, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x77, 0x10, 0x01, 0x12, 0x05, 0x0a,
	0x01, 0x6d, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x79, 0x10, 0x03, 0x2a, 0xb6, 0x01, 0x0a, 0x09,
	0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61, 0x63,
	0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x73, 0x69, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x62, 0x76, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x64, 0x78, 0x10, 0x03, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x65, 0x74, 0x61, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x74, 0x72, 0x10,
	0x05, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x62, 0x61, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03,
	0x45, 0x6d, 0x61, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x52,
	0x65, 0x67, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x69, 0x6c, 0x6c, 0x52, 0x10, 0x09, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x74, 0x64, 0x44, 0x65, 0x76, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x6d, 0x61, 0x31, 0x30, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x6d, 0x61, 0x32, 0x30, 0x10,
	0x0c, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x6d, 0x61, 0x35, 0x30, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x6d, 0x61, 0x31, 0x30, 0x30, 0x10, 0x0e, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6d, 0x61, 0x32,
	0x30, 0x30, 0x10, 0x0f, 0x2a, 0x4d, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x47,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x65,
	0x73, 0x73, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x65, 0x73, 0x73, 0x45, 0x71, 0x75, 0x61,
	0x6c, 0x10, 0x04, 0x2a, 0x62, 0x0a, 0x08, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x79, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x05, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x06, 0x32, 0xa7, 0x01, 0x0a, 0x06, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x51, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x7a, 0x42, 0x46, 0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x19,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_quotes_v1_quotes_proto_rawDescOnce sync.Once
	file_proto_quotes_v1_quotes_proto_rawDescData = file_proto_quotes_v1_quotes_proto_rawDesc
)

func file_proto_quotes_v1_quotes_proto_rawDescGZIP() []byte {
	file_proto_quotes_v1_quotes_proto_rawDescOnce.Do(func() {
		file_proto_quotes_v1_quotes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quotes_v1_quotes_proto_rawDescData)
	})
	return file_proto_quotes_v1_quotes_proto_rawDescData
}

var file_proto_quotes_v1_quotes_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_proto_quotes_v1_quotes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_quotes_v1_quotes_proto_goTypes = []interface{}{
	(TSDB)(0),                     // 0: quotes.v1.TSDB
	(Interval)(0),                 // 1: quotes.v1.Interval
	(Indicator)(0),                // 2: quotes.v1.Indicator
	(Operator)(0),                 // 3: quotes.v1.Operator
	(Argument)(0),                 // 4: quotes.v1.Argument
	(*Filter)(nil),                // 5: quotes.v1.Filter
	(*QuotesRequest)(nil),         // 6: quotes.v1.QuotesRequest
	(*QuotesReply)(nil),           // 7: quotes.v1.QuotesReply
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_proto_quotes_v1_quotes_proto_depIdxs = []int32{
	3, // 0: quotes.v1.Filter.operator:type_name -> quotes.v1.Operator
	4, // 1: quotes.v1.Filter.argument:type_name -> quotes.v1.Argument
	0, // 2: quotes.v1.QuotesRequest.tsdb:type_name -> quotes.v1.TSDB
	1, // 3: quotes.v1.QuotesRequest.interval:type_name -> quotes.v1.Interval
	5, // 4: quotes.v1.QuotesRequest.filters:type_name -> quotes.v1.Filter
	8, // 5: quotes.v1.QuotesReply.created_at:type_name -> google.protobuf.Timestamp
	6, // 6: quotes.v1.Quotes.Search:input_type -> quotes.v1.QuotesRequest
	9, // 7: quotes.v1.Quotes.Health:input_type -> google.protobuf.Empty
	7, // 8: quotes.v1.Quotes.Search:output_type -> quotes.v1.QuotesReply
	9, // 9: quotes.v1.Quotes.Health:output_type -> google.protobuf.Empty
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_quotes_v1_quotes_proto_init() }
func file_proto_quotes_v1_quotes_proto_init() {
	if File_proto_quotes_v1_quotes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quotes_v1_quotes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_quotes_v1_quotes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotesRequest); i {
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
		file_proto_quotes_v1_quotes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotesReply); i {
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
			RawDescriptor: file_proto_quotes_v1_quotes_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quotes_v1_quotes_proto_goTypes,
		DependencyIndexes: file_proto_quotes_v1_quotes_proto_depIdxs,
		EnumInfos:         file_proto_quotes_v1_quotes_proto_enumTypes,
		MessageInfos:      file_proto_quotes_v1_quotes_proto_msgTypes,
	}.Build()
	File_proto_quotes_v1_quotes_proto = out.File
	file_proto_quotes_v1_quotes_proto_rawDesc = nil
	file_proto_quotes_v1_quotes_proto_goTypes = nil
	file_proto_quotes_v1_quotes_proto_depIdxs = nil
}
