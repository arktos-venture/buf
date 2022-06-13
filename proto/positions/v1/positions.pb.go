// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/positions/v1/positions.proto

package positions_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Asset int32

const (
	Asset_STOCK  Asset = 0
	Asset_OPTION Asset = 1
	Asset_CASH   Asset = 2
)

// Enum value maps for Asset.
var (
	Asset_name = map[int32]string{
		0: "STOCK",
		1: "OPTION",
		2: "CASH",
	}
	Asset_value = map[string]int32{
		"STOCK":  0,
		"OPTION": 1,
		"CASH":   2,
	}
)

func (x Asset) Enum() *Asset {
	p := new(Asset)
	*p = x
	return p
}

func (x Asset) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Asset) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_positions_v1_positions_proto_enumTypes[0].Descriptor()
}

func (Asset) Type() protoreflect.EnumType {
	return &file_proto_positions_v1_positions_proto_enumTypes[0]
}

func (x Asset) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Asset.Descriptor instead.
func (Asset) EnumDescriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{0}
}

type PositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asset   []Asset `protobuf:"varint,1,rep,packed,name=asset,proto3,enum=positions.v1.Asset" json:"asset,omitempty"`
	Account string  `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *PositionRequest) Reset() {
	*x = PositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionRequest) ProtoMessage() {}

func (x *PositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionRequest.ProtoReflect.Descriptor instead.
func (*PositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{0}
}

func (x *PositionRequest) GetAsset() []Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

func (x *PositionRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type PositionDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *PositionDeleteRequest) Reset() {
	*x = PositionDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionDeleteRequest) ProtoMessage() {}

func (x *PositionDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionDeleteRequest.ProtoReflect.Descriptor instead.
func (*PositionDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{1}
}

func (x *PositionDeleteRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type PositionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker      string                     `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Exchange    string                     `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Size        float32                    `protobuf:"fixed32,3,opt,name=size,proto3" json:"size,omitempty"`
	Cost        *PositionReply_Cost        `protobuf:"bytes,4,opt,name=cost,proto3" json:"cost,omitempty"`
	Performance *PositionReply_Performance `protobuf:"bytes,5,opt,name=performance,proto3" json:"performance,omitempty"`
	CreatedAt   *timestamppb.Timestamp     `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp     `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PositionReply) Reset() {
	*x = PositionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionReply) ProtoMessage() {}

func (x *PositionReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionReply.ProtoReflect.Descriptor instead.
func (*PositionReply) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{2}
}

func (x *PositionReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *PositionReply) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *PositionReply) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PositionReply) GetCost() *PositionReply_Cost {
	if x != nil {
		return x.Cost
	}
	return nil
}

func (x *PositionReply) GetPerformance() *PositionReply_Performance {
	if x != nil {
		return x.Performance
	}
	return nil
}

func (x *PositionReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PositionReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PositionReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PositionReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PositionReplies) Reset() {
	*x = PositionReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionReplies) ProtoMessage() {}

func (x *PositionReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionReplies.ProtoReflect.Descriptor instead.
func (*PositionReplies) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{3}
}

func (x *PositionReplies) GetResults() []*PositionReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PositionReplies) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PositionDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PositionDelete) Reset() {
	*x = PositionDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionDelete) ProtoMessage() {}

func (x *PositionDelete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionDelete.ProtoReflect.Descriptor instead.
func (*PositionDelete) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{4}
}

func (x *PositionDelete) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PositionReply_Performance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percent float32 `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
	Money   float32 `protobuf:"fixed32,2,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *PositionReply_Performance) Reset() {
	*x = PositionReply_Performance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionReply_Performance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionReply_Performance) ProtoMessage() {}

func (x *PositionReply_Performance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionReply_Performance.ProtoReflect.Descriptor instead.
func (*PositionReply_Performance) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PositionReply_Performance) GetPercent() float32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *PositionReply_Performance) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

type PositionReply_Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit  float32 `protobuf:"fixed32,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Total float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PositionReply_Cost) Reset() {
	*x = PositionReply_Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionReply_Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionReply_Cost) ProtoMessage() {}

func (x *PositionReply_Cost) ProtoReflect() protoreflect.Message {
	mi := &file_proto_positions_v1_positions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionReply_Cost.ProtoReflect.Descriptor instead.
func (*PositionReply_Cost) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{2, 1}
}

func (x *PositionReply_Cost) GetUnit() float32 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *PositionReply_Cost) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_positions_v1_positions_proto protoreflect.FileDescriptor

var file_proto_positions_v1_positions_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x03, 0x18, 0x24, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a,
	0x15, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03,
	0x18, 0x24, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbf, 0x03, 0x0a, 0x0d,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3d, 0x0a, 0x0b,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x1a, 0x30, 0x0a, 0x04, 0x43,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5e, 0x0a,
	0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x26, 0x0a,
	0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x28, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x48, 0x10, 0x02, 0x32,
	0xe2, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x67, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x12, 0x6c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x23, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x7d, 0x42, 0x97, 0x03, 0x0a, 0x1b, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x41, 0x50, 0x49, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x76, 0x31, 0x92, 0x41, 0xa0, 0x02, 0x12, 0x8c,
	0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x73, 0x12, 0x17, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x73, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72,
	0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d,
	0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73,
	0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a,
	0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02,
	0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_positions_v1_positions_proto_rawDescOnce sync.Once
	file_proto_positions_v1_positions_proto_rawDescData = file_proto_positions_v1_positions_proto_rawDesc
)

func file_proto_positions_v1_positions_proto_rawDescGZIP() []byte {
	file_proto_positions_v1_positions_proto_rawDescOnce.Do(func() {
		file_proto_positions_v1_positions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_positions_v1_positions_proto_rawDescData)
	})
	return file_proto_positions_v1_positions_proto_rawDescData
}

var file_proto_positions_v1_positions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_positions_v1_positions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_positions_v1_positions_proto_goTypes = []interface{}{
	(Asset)(0),                        // 0: positions.v1.Asset
	(*PositionRequest)(nil),           // 1: positions.v1.PositionRequest
	(*PositionDeleteRequest)(nil),     // 2: positions.v1.PositionDeleteRequest
	(*PositionReply)(nil),             // 3: positions.v1.PositionReply
	(*PositionReplies)(nil),           // 4: positions.v1.PositionReplies
	(*PositionDelete)(nil),            // 5: positions.v1.PositionDelete
	(*PositionReply_Performance)(nil), // 6: positions.v1.PositionReply.Performance
	(*PositionReply_Cost)(nil),        // 7: positions.v1.PositionReply.Cost
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_proto_positions_v1_positions_proto_depIdxs = []int32{
	0, // 0: positions.v1.PositionRequest.asset:type_name -> positions.v1.Asset
	7, // 1: positions.v1.PositionReply.cost:type_name -> positions.v1.PositionReply.Cost
	6, // 2: positions.v1.PositionReply.performance:type_name -> positions.v1.PositionReply.Performance
	8, // 3: positions.v1.PositionReply.created_at:type_name -> google.protobuf.Timestamp
	8, // 4: positions.v1.PositionReply.updated_at:type_name -> google.protobuf.Timestamp
	3, // 5: positions.v1.PositionReplies.results:type_name -> positions.v1.PositionReply
	1, // 6: positions.v1.Positions.Search:input_type -> positions.v1.PositionRequest
	2, // 7: positions.v1.Positions.Delete:input_type -> positions.v1.PositionDeleteRequest
	4, // 8: positions.v1.Positions.Search:output_type -> positions.v1.PositionReplies
	5, // 9: positions.v1.Positions.Delete:output_type -> positions.v1.PositionDelete
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_positions_v1_positions_proto_init() }
func file_proto_positions_v1_positions_proto_init() {
	if File_proto_positions_v1_positions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_positions_v1_positions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionRequest); i {
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
		file_proto_positions_v1_positions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionDeleteRequest); i {
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
		file_proto_positions_v1_positions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionReply); i {
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
		file_proto_positions_v1_positions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionReplies); i {
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
		file_proto_positions_v1_positions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionDelete); i {
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
		file_proto_positions_v1_positions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionReply_Performance); i {
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
		file_proto_positions_v1_positions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionReply_Cost); i {
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
			RawDescriptor: file_proto_positions_v1_positions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_positions_v1_positions_proto_goTypes,
		DependencyIndexes: file_proto_positions_v1_positions_proto_depIdxs,
		EnumInfos:         file_proto_positions_v1_positions_proto_enumTypes,
		MessageInfos:      file_proto_positions_v1_positions_proto_msgTypes,
	}.Build()
	File_proto_positions_v1_positions_proto = out.File
	file_proto_positions_v1_positions_proto_rawDesc = nil
	file_proto_positions_v1_positions_proto_goTypes = nil
	file_proto_positions_v1_positions_proto_depIdxs = nil
}
