// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/strategy-results/v1/strategy-results.proto

package v1

import (
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
	return file_proto_strategy_results_v1_strategy_results_proto_enumTypes[0].Descriptor()
}

func (StrategyReply_Action) Type() protoreflect.EnumType {
	return &file_proto_strategy_results_v1_strategy_results_proto_enumTypes[0]
}

func (x StrategyReply_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StrategyReply_Action.Descriptor instead.
func (StrategyReply_Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_strategy_results_v1_strategy_results_proto_rawDescGZIP(), []int{1, 0}
}

type StrategyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reversed bool                     `protobuf:"varint,1,opt,name=reversed,proto3" json:"reversed,omitempty"`
	Quotes   []float32                `protobuf:"fixed32,2,rep,packed,name=quotes,proto3" json:"quotes,omitempty"`
	Dates    []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=dates,proto3" json:"dates,omitempty"`
	Params   map[string]int32         `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *StrategyRequest) Reset() {
	*x = StrategyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategy_results_v1_strategy_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyRequest) ProtoMessage() {}

func (x *StrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategy_results_v1_strategy_results_proto_msgTypes[0]
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
	return file_proto_strategy_results_v1_strategy_results_proto_rawDescGZIP(), []int{0}
}

func (x *StrategyRequest) GetReversed() bool {
	if x != nil {
		return x.Reversed
	}
	return false
}

func (x *StrategyRequest) GetQuotes() []float32 {
	if x != nil {
		return x.Quotes
	}
	return nil
}

func (x *StrategyRequest) GetDates() []*timestamppb.Timestamp {
	if x != nil {
		return x.Dates
	}
	return nil
}

func (x *StrategyRequest) GetParams() map[string]int32 {
	if x != nil {
		return x.Params
	}
	return nil
}

type StrategyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Close  []float32                `protobuf:"fixed32,1,rep,packed,name=close,proto3" json:"close,omitempty"`
	Action []StrategyReply_Action   `protobuf:"varint,2,rep,packed,name=action,proto3,enum=strategy_results.v1.StrategyReply_Action" json:"action,omitempty"`
	Date   []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=date,proto3" json:"date,omitempty"`
}

func (x *StrategyReply) Reset() {
	*x = StrategyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategy_results_v1_strategy_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyReply) ProtoMessage() {}

func (x *StrategyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategy_results_v1_strategy_results_proto_msgTypes[1]
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
	return file_proto_strategy_results_v1_strategy_results_proto_rawDescGZIP(), []int{1}
}

func (x *StrategyReply) GetClose() []float32 {
	if x != nil {
		return x.Close
	}
	return nil
}

func (x *StrategyReply) GetAction() []StrategyReply_Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *StrategyReply) GetDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_proto_strategy_results_v1_strategy_results_proto protoreflect.FileDescriptor

var file_proto_strategy_results_v1_strategy_results_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x2d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x25, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4c, 0x44, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4c, 0x4c, 0x10, 0x02, 0x32, 0xcc, 0x01, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x74, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x7a, 0x42, 0x63, 0x0a, 0x22, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_strategy_results_v1_strategy_results_proto_rawDescOnce sync.Once
	file_proto_strategy_results_v1_strategy_results_proto_rawDescData = file_proto_strategy_results_v1_strategy_results_proto_rawDesc
)

func file_proto_strategy_results_v1_strategy_results_proto_rawDescGZIP() []byte {
	file_proto_strategy_results_v1_strategy_results_proto_rawDescOnce.Do(func() {
		file_proto_strategy_results_v1_strategy_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_strategy_results_v1_strategy_results_proto_rawDescData)
	})
	return file_proto_strategy_results_v1_strategy_results_proto_rawDescData
}

var file_proto_strategy_results_v1_strategy_results_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_strategy_results_v1_strategy_results_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_strategy_results_v1_strategy_results_proto_goTypes = []interface{}{
	(StrategyReply_Action)(0),     // 0: strategy_results.v1.StrategyReply.Action
	(*StrategyRequest)(nil),       // 1: strategy_results.v1.StrategyRequest
	(*StrategyReply)(nil),         // 2: strategy_results.v1.StrategyReply
	nil,                           // 3: strategy_results.v1.StrategyRequest.ParamsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_proto_strategy_results_v1_strategy_results_proto_depIdxs = []int32{
	4, // 0: strategy_results.v1.StrategyRequest.dates:type_name -> google.protobuf.Timestamp
	3, // 1: strategy_results.v1.StrategyRequest.params:type_name -> strategy_results.v1.StrategyRequest.ParamsEntry
	0, // 2: strategy_results.v1.StrategyReply.action:type_name -> strategy_results.v1.StrategyReply.Action
	4, // 3: strategy_results.v1.StrategyReply.date:type_name -> google.protobuf.Timestamp
	1, // 4: strategy_results.v1.Strategy.Execute:input_type -> strategy_results.v1.StrategyRequest
	5, // 5: strategy_results.v1.Strategy.Health:input_type -> google.protobuf.Empty
	2, // 6: strategy_results.v1.Strategy.Execute:output_type -> strategy_results.v1.StrategyReply
	5, // 7: strategy_results.v1.Strategy.Health:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_strategy_results_v1_strategy_results_proto_init() }
func file_proto_strategy_results_v1_strategy_results_proto_init() {
	if File_proto_strategy_results_v1_strategy_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_strategy_results_v1_strategy_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_strategy_results_v1_strategy_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_strategy_results_v1_strategy_results_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_strategy_results_v1_strategy_results_proto_goTypes,
		DependencyIndexes: file_proto_strategy_results_v1_strategy_results_proto_depIdxs,
		EnumInfos:         file_proto_strategy_results_v1_strategy_results_proto_enumTypes,
		MessageInfos:      file_proto_strategy_results_v1_strategy_results_proto_msgTypes,
	}.Build()
	File_proto_strategy_results_v1_strategy_results_proto = out.File
	file_proto_strategy_results_v1_strategy_results_proto_rawDesc = nil
	file_proto_strategy_results_v1_strategy_results_proto_goTypes = nil
	file_proto_strategy_results_v1_strategy_results_proto_depIdxs = nil
}
