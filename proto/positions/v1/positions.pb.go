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

type PositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
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

func (x *PositionRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type PositionCompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker      string                            `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Exchange    string                            `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Size        float32                           `protobuf:"fixed32,3,opt,name=size,proto3" json:"size,omitempty"`
	Cost        *PositionCompanyReply_Cost        `protobuf:"bytes,4,opt,name=cost,proto3" json:"cost,omitempty"`
	Performance *PositionCompanyReply_Performance `protobuf:"bytes,5,opt,name=performance,proto3" json:"performance,omitempty"`
	CreatedAt   *timestamppb.Timestamp            `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp            `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PositionCompanyReply) Reset() {
	*x = PositionCompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCompanyReply) ProtoMessage() {}

func (x *PositionCompanyReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionCompanyReply.ProtoReflect.Descriptor instead.
func (*PositionCompanyReply) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{1}
}

func (x *PositionCompanyReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *PositionCompanyReply) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *PositionCompanyReply) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PositionCompanyReply) GetCost() *PositionCompanyReply_Cost {
	if x != nil {
		return x.Cost
	}
	return nil
}

func (x *PositionCompanyReply) GetPerformance() *PositionCompanyReply_Performance {
	if x != nil {
		return x.Performance
	}
	return nil
}

func (x *PositionCompanyReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PositionCompanyReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PositionCompanyReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PositionCompanyReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PositionCompanyReplies) Reset() {
	*x = PositionCompanyReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCompanyReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCompanyReplies) ProtoMessage() {}

func (x *PositionCompanyReplies) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionCompanyReplies.ProtoReflect.Descriptor instead.
func (*PositionCompanyReplies) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{2}
}

func (x *PositionCompanyReplies) GetResults() []*PositionCompanyReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PositionCompanyReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PositionCurrencyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency  string                 `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Size      float32                `protobuf:"fixed32,2,opt,name=size,proto3" json:"size,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PositionCurrencyReply) Reset() {
	*x = PositionCurrencyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCurrencyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCurrencyReply) ProtoMessage() {}

func (x *PositionCurrencyReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionCurrencyReply.ProtoReflect.Descriptor instead.
func (*PositionCurrencyReply) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{3}
}

func (x *PositionCurrencyReply) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PositionCurrencyReply) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PositionCurrencyReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PositionCurrencyReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PositionCurrencyReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PositionCurrencyReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PositionCurrencyReplies) Reset() {
	*x = PositionCurrencyReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCurrencyReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCurrencyReplies) ProtoMessage() {}

func (x *PositionCurrencyReplies) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionCurrencyReplies.ProtoReflect.Descriptor instead.
func (*PositionCurrencyReplies) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{4}
}

func (x *PositionCurrencyReplies) GetResults() []*PositionCurrencyReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PositionCurrencyReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PositionCompanyReply_Performance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percent float32 `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
	Money   float32 `protobuf:"fixed32,2,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *PositionCompanyReply_Performance) Reset() {
	*x = PositionCompanyReply_Performance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCompanyReply_Performance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCompanyReply_Performance) ProtoMessage() {}

func (x *PositionCompanyReply_Performance) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionCompanyReply_Performance.ProtoReflect.Descriptor instead.
func (*PositionCompanyReply_Performance) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PositionCompanyReply_Performance) GetPercent() float32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *PositionCompanyReply_Performance) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

type PositionCompanyReply_Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit  float32 `protobuf:"fixed32,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Total float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PositionCompanyReply_Cost) Reset() {
	*x = PositionCompanyReply_Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_positions_v1_positions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCompanyReply_Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCompanyReply_Cost) ProtoMessage() {}

func (x *PositionCompanyReply_Cost) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionCompanyReply_Cost.ProtoReflect.Descriptor instead.
func (*PositionCompanyReply_Cost) Descriptor() ([]byte, []int) {
	return file_proto_positions_v1_positions_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PositionCompanyReply_Cost) GetUnit() float32 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *PositionCompanyReply_Cost) GetTotal() float32 {
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
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0f, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xd4, 0x03, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x65, 0x72,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3d, 0x0a, 0x0b, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x1a, 0x30, 0x0a, 0x04, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x6c, 0x0a, 0x16,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xbd, 0x01, 0x0a, 0x15, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6e, 0x0a, 0x17, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x8c, 0x02, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x12, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x2c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x7d, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x42, 0x97, 0x03, 0x0a, 0x1b, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x41, 0x50, 0x49, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b,
	0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x76, 0x31, 0x92,
	0x41, 0xa0, 0x02, 0x12, 0x8c, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x12, 0x17, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65, 0x6e, 0x74, 0x75, 0x72,
	0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72,
	0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66,
	0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_positions_v1_positions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_positions_v1_positions_proto_goTypes = []interface{}{
	(*PositionRequest)(nil),                  // 0: positions.v1.PositionRequest
	(*PositionCompanyReply)(nil),             // 1: positions.v1.PositionCompanyReply
	(*PositionCompanyReplies)(nil),           // 2: positions.v1.PositionCompanyReplies
	(*PositionCurrencyReply)(nil),            // 3: positions.v1.PositionCurrencyReply
	(*PositionCurrencyReplies)(nil),          // 4: positions.v1.PositionCurrencyReplies
	(*PositionCompanyReply_Performance)(nil), // 5: positions.v1.PositionCompanyReply.Performance
	(*PositionCompanyReply_Cost)(nil),        // 6: positions.v1.PositionCompanyReply.Cost
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
}
var file_proto_positions_v1_positions_proto_depIdxs = []int32{
	6,  // 0: positions.v1.PositionCompanyReply.cost:type_name -> positions.v1.PositionCompanyReply.Cost
	5,  // 1: positions.v1.PositionCompanyReply.performance:type_name -> positions.v1.PositionCompanyReply.Performance
	7,  // 2: positions.v1.PositionCompanyReply.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: positions.v1.PositionCompanyReply.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 4: positions.v1.PositionCompanyReplies.results:type_name -> positions.v1.PositionCompanyReply
	7,  // 5: positions.v1.PositionCurrencyReply.created_at:type_name -> google.protobuf.Timestamp
	7,  // 6: positions.v1.PositionCurrencyReply.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 7: positions.v1.PositionCurrencyReplies.results:type_name -> positions.v1.PositionCurrencyReply
	0,  // 8: positions.v1.Position.Companies:input_type -> positions.v1.PositionRequest
	0,  // 9: positions.v1.Position.Currencies:input_type -> positions.v1.PositionRequest
	2,  // 10: positions.v1.Position.Companies:output_type -> positions.v1.PositionCompanyReplies
	4,  // 11: positions.v1.Position.Currencies:output_type -> positions.v1.PositionCurrencyReplies
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
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
			switch v := v.(*PositionCompanyReply); i {
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
			switch v := v.(*PositionCompanyReplies); i {
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
			switch v := v.(*PositionCurrencyReply); i {
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
			switch v := v.(*PositionCurrencyReplies); i {
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
			switch v := v.(*PositionCompanyReply_Performance); i {
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
			switch v := v.(*PositionCompanyReply_Cost); i {
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
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_positions_v1_positions_proto_goTypes,
		DependencyIndexes: file_proto_positions_v1_positions_proto_depIdxs,
		MessageInfos:      file_proto_positions_v1_positions_proto_msgTypes,
	}.Build()
	File_proto_positions_v1_positions_proto = out.File
	file_proto_positions_v1_positions_proto_rawDesc = nil
	file_proto_positions_v1_positions_proto_goTypes = nil
	file_proto_positions_v1_positions_proto_depIdxs = nil
}
