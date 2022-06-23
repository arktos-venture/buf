// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/strategies/v1/strategies.proto

package strategies_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// Transverse
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters map[string]int32 `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Period     map[string]int32 `protobuf:"bytes,2,rep,name=period,proto3" json:"period,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetParameters() map[string]int32 {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *Request) GetPeriod() map[string]int32 {
	if x != nil {
		return x.Period
	}
	return nil
}

type StrategyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
}

func (x *StrategyRequest) Reset() {
	*x = StrategyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyRequest) ProtoMessage() {}

func (x *StrategyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StrategyRequest.ProtoReflect.Descriptor instead.
func (*StrategyRequest) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{1}
}

func (x *StrategyRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

type StrategyModifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker      string   `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Account     string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Parameters  *Request `protobuf:"bytes,5,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *StrategyModifyRequest) Reset() {
	*x = StrategyModifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyModifyRequest) ProtoMessage() {}

func (x *StrategyModifyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StrategyModifyRequest.ProtoReflect.Descriptor instead.
func (*StrategyModifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{2}
}

func (x *StrategyModifyRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *StrategyModifyRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *StrategyModifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StrategyModifyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StrategyModifyRequest) GetParameters() *Request {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type StrategyDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickers []string `protobuf:"bytes,1,rep,name=tickers,proto3" json:"tickers,omitempty"`
}

func (x *StrategyDeleteRequest) Reset() {
	*x = StrategyDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyDeleteRequest) ProtoMessage() {}

func (x *StrategyDeleteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StrategyDeleteRequest.ProtoReflect.Descriptor instead.
func (*StrategyDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{3}
}

func (x *StrategyDeleteRequest) GetTickers() []string {
	if x != nil {
		return x.Tickers
	}
	return nil
}

type StrategyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker      string                 `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Account     string                 `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Request     *Request               `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *StrategyReply) Reset() {
	*x = StrategyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyReply) ProtoMessage() {}

func (x *StrategyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[4]
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
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{4}
}

func (x *StrategyReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *StrategyReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StrategyReply) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *StrategyReply) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *StrategyReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *StrategyReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type StrategyReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*StrategyReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *StrategyReplies) Reset() {
	*x = StrategyReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyReplies) ProtoMessage() {}

func (x *StrategyReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyReplies.ProtoReflect.Descriptor instead.
func (*StrategyReplies) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{5}
}

func (x *StrategyReplies) GetResults() []*StrategyReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *StrategyReplies) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type StrategyDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *StrategyDelete) Reset() {
	*x = StrategyDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_strategies_v1_strategies_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyDelete) ProtoMessage() {}

func (x *StrategyDelete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_strategies_v1_strategies_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyDelete.ProtoReflect.Descriptor instead.
func (*StrategyDelete) Descriptor() ([]byte, []int) {
	return file_proto_strategies_v1_strategies_proto_rawDescGZIP(), []int{6}
}

func (x *StrategyDelete) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
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
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3a,
	0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x22,
	0xb7, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x15, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x18, 0x01, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x8b,
	0x02, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5f, 0x0a, 0x0f,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12,
	0x36, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x26, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x88, 0x04, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f,
	0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x7d, 0x12, 0x56, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x12, 0x67, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a,
	0x15, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x7b, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x42, 0x94, 0x03, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x14, 0x41, 0x50, 0x49, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e,
	0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x92, 0x41, 0x99, 0x02, 0x12, 0x85,
	0x01, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x73, 0x12, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65,
	0x6e, 0x74, 0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73,
	0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59,
	0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62,
	0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_strategies_v1_strategies_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_strategies_v1_strategies_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: strategies.v1.Request
	(*StrategyRequest)(nil),       // 1: strategies.v1.StrategyRequest
	(*StrategyModifyRequest)(nil), // 2: strategies.v1.StrategyModifyRequest
	(*StrategyDeleteRequest)(nil), // 3: strategies.v1.StrategyDeleteRequest
	(*StrategyReply)(nil),         // 4: strategies.v1.StrategyReply
	(*StrategyReplies)(nil),       // 5: strategies.v1.StrategyReplies
	(*StrategyDelete)(nil),        // 6: strategies.v1.StrategyDelete
	nil,                           // 7: strategies.v1.Request.ParametersEntry
	nil,                           // 8: strategies.v1.Request.PeriodEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_proto_strategies_v1_strategies_proto_depIdxs = []int32{
	7,  // 0: strategies.v1.Request.parameters:type_name -> strategies.v1.Request.ParametersEntry
	8,  // 1: strategies.v1.Request.period:type_name -> strategies.v1.Request.PeriodEntry
	0,  // 2: strategies.v1.StrategyModifyRequest.parameters:type_name -> strategies.v1.Request
	0,  // 3: strategies.v1.StrategyReply.request:type_name -> strategies.v1.Request
	9,  // 4: strategies.v1.StrategyReply.created_at:type_name -> google.protobuf.Timestamp
	9,  // 5: strategies.v1.StrategyReply.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 6: strategies.v1.StrategyReplies.results:type_name -> strategies.v1.StrategyReply
	1,  // 7: strategies.v1.Strategies.Get:input_type -> strategies.v1.StrategyRequest
	10, // 8: strategies.v1.Strategies.List:input_type -> google.protobuf.Empty
	2,  // 9: strategies.v1.Strategies.Create:input_type -> strategies.v1.StrategyModifyRequest
	2,  // 10: strategies.v1.Strategies.Update:input_type -> strategies.v1.StrategyModifyRequest
	3,  // 11: strategies.v1.Strategies.Delete:input_type -> strategies.v1.StrategyDeleteRequest
	4,  // 12: strategies.v1.Strategies.Get:output_type -> strategies.v1.StrategyReply
	5,  // 13: strategies.v1.Strategies.List:output_type -> strategies.v1.StrategyReplies
	4,  // 14: strategies.v1.Strategies.Create:output_type -> strategies.v1.StrategyReply
	4,  // 15: strategies.v1.Strategies.Update:output_type -> strategies.v1.StrategyReply
	6,  // 16: strategies.v1.Strategies.Delete:output_type -> strategies.v1.StrategyDelete
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_strategies_v1_strategies_proto_init() }
func file_proto_strategies_v1_strategies_proto_init() {
	if File_proto_strategies_v1_strategies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_strategies_v1_strategies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyModifyRequest); i {
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
			switch v := v.(*StrategyDeleteRequest); i {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyReplies); i {
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
		file_proto_strategies_v1_strategies_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyDelete); i {
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
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_strategies_v1_strategies_proto_goTypes,
		DependencyIndexes: file_proto_strategies_v1_strategies_proto_depIdxs,
		MessageInfos:      file_proto_strategies_v1_strategies_proto_msgTypes,
	}.Build()
	File_proto_strategies_v1_strategies_proto = out.File
	file_proto_strategies_v1_strategies_proto_rawDesc = nil
	file_proto_strategies_v1_strategies_proto_goTypes = nil
	file_proto_strategies_v1_strategies_proto_depIdxs = nil
}
