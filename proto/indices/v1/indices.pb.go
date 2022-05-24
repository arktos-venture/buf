// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/indices/v1/indices.proto

package indices_v1

import (
	v1 "github.com/arktos-venture/buf/proto/quotes/v1"
	v11 "github.com/arktos-venture/buf/proto/screener/v1"
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

type IndicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indice string `protobuf:"bytes,1,opt,name=indice,proto3" json:"indice,omitempty"`
}

func (x *IndicesRequest) Reset() {
	*x = IndicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indices_v1_indices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicesRequest) ProtoMessage() {}

func (x *IndicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indices_v1_indices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicesRequest.ProtoReflect.Descriptor instead.
func (*IndicesRequest) Descriptor() ([]byte, []int) {
	return file_proto_indices_v1_indices_proto_rawDescGZIP(), []int{0}
}

func (x *IndicesRequest) GetIndice() string {
	if x != nil {
		return x.Indice
	}
	return ""
}

type IndicesExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange string `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *IndicesExchangeRequest) Reset() {
	*x = IndicesExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indices_v1_indices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndicesExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicesExchangeRequest) ProtoMessage() {}

func (x *IndicesExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indices_v1_indices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicesExchangeRequest.ProtoReflect.Descriptor instead.
func (*IndicesExchangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_indices_v1_indices_proto_rawDescGZIP(), []int{1}
}

func (x *IndicesExchangeRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type IndicesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ticker      string                 `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Quote       *v1.QuotesLastReply    `protobuf:"bytes,5,opt,name=quote,proto3" json:"quote,omitempty"`
	Companies   []*v11.ScreenerReply   `protobuf:"bytes,6,rep,name=companies,proto3" json:"companies,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *IndicesReply) Reset() {
	*x = IndicesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indices_v1_indices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndicesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicesReply) ProtoMessage() {}

func (x *IndicesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indices_v1_indices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicesReply.ProtoReflect.Descriptor instead.
func (*IndicesReply) Descriptor() ([]byte, []int) {
	return file_proto_indices_v1_indices_proto_rawDescGZIP(), []int{2}
}

func (x *IndicesReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IndicesReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *IndicesReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndicesReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IndicesReply) GetQuote() *v1.QuotesLastReply {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *IndicesReply) GetCompanies() []*v11.ScreenerReply {
	if x != nil {
		return x.Companies
	}
	return nil
}

func (x *IndicesReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IndicesReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type IndicesShortReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ticker      string                 `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *IndicesShortReply) Reset() {
	*x = IndicesShortReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indices_v1_indices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndicesShortReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicesShortReply) ProtoMessage() {}

func (x *IndicesShortReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indices_v1_indices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicesShortReply.ProtoReflect.Descriptor instead.
func (*IndicesShortReply) Descriptor() ([]byte, []int) {
	return file_proto_indices_v1_indices_proto_rawDescGZIP(), []int{3}
}

func (x *IndicesShortReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IndicesShortReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *IndicesShortReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndicesShortReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IndicesShortReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IndicesShortReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type IndicesShortReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*IndicesShortReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *IndicesShortReplies) Reset() {
	*x = IndicesShortReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indices_v1_indices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndicesShortReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicesShortReplies) ProtoMessage() {}

func (x *IndicesShortReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indices_v1_indices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicesShortReplies.ProtoReflect.Descriptor instead.
func (*IndicesShortReplies) Descriptor() ([]byte, []int) {
	return file_proto_indices_v1_indices_proto_rawDescGZIP(), []int{4}
}

func (x *IndicesShortReplies) GetResults() []*IndicesShortReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *IndicesShortReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_indices_v1_indices_proto protoreflect.FileDescriptor

var file_proto_indices_v1_indices_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x33, 0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x10, 0x52, 0x06,
	0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x16, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x10, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xce, 0x02, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe7, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x64, 0x0a, 0x13, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xaa, 0x02, 0x0a, 0x07, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x69, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x7d, 0x12, 0x74,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x7b, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x7d, 0x12, 0x4f, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x92, 0x41, 0x02, 0x62, 0x00, 0x42, 0x81, 0x03, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x11, 0x41, 0x50, 0x49, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x5f, 0x76, 0x31, 0x92, 0x41, 0x92, 0x02, 0x12, 0x7f, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x12, 0x0c, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20,
	0x56, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x25, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74,
	0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02,
	0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a,
	0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_indices_v1_indices_proto_rawDescOnce sync.Once
	file_proto_indices_v1_indices_proto_rawDescData = file_proto_indices_v1_indices_proto_rawDesc
)

func file_proto_indices_v1_indices_proto_rawDescGZIP() []byte {
	file_proto_indices_v1_indices_proto_rawDescOnce.Do(func() {
		file_proto_indices_v1_indices_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_indices_v1_indices_proto_rawDescData)
	})
	return file_proto_indices_v1_indices_proto_rawDescData
}

var file_proto_indices_v1_indices_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_indices_v1_indices_proto_goTypes = []interface{}{
	(*IndicesRequest)(nil),         // 0: indices.v1.IndicesRequest
	(*IndicesExchangeRequest)(nil), // 1: indices.v1.IndicesExchangeRequest
	(*IndicesReply)(nil),           // 2: indices.v1.IndicesReply
	(*IndicesShortReply)(nil),      // 3: indices.v1.IndicesShortReply
	(*IndicesShortReplies)(nil),    // 4: indices.v1.IndicesShortReplies
	(*v1.QuotesLastReply)(nil),     // 5: quotes.v1.QuotesLastReply
	(*v11.ScreenerReply)(nil),      // 6: screener.v1.ScreenerReply
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_proto_indices_v1_indices_proto_depIdxs = []int32{
	5,  // 0: indices.v1.IndicesReply.quote:type_name -> quotes.v1.QuotesLastReply
	6,  // 1: indices.v1.IndicesReply.companies:type_name -> screener.v1.ScreenerReply
	7,  // 2: indices.v1.IndicesReply.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: indices.v1.IndicesReply.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 4: indices.v1.IndicesShortReply.created_at:type_name -> google.protobuf.Timestamp
	7,  // 5: indices.v1.IndicesShortReply.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 6: indices.v1.IndicesShortReplies.results:type_name -> indices.v1.IndicesShortReply
	0,  // 7: indices.v1.Indices.Get:input_type -> indices.v1.IndicesRequest
	1,  // 8: indices.v1.Indices.List:input_type -> indices.v1.IndicesExchangeRequest
	8,  // 9: indices.v1.Indices.Health:input_type -> google.protobuf.Empty
	2,  // 10: indices.v1.Indices.Get:output_type -> indices.v1.IndicesReply
	4,  // 11: indices.v1.Indices.List:output_type -> indices.v1.IndicesShortReplies
	8,  // 12: indices.v1.Indices.Health:output_type -> google.protobuf.Empty
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_indices_v1_indices_proto_init() }
func file_proto_indices_v1_indices_proto_init() {
	if File_proto_indices_v1_indices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_indices_v1_indices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndicesRequest); i {
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
		file_proto_indices_v1_indices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndicesExchangeRequest); i {
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
		file_proto_indices_v1_indices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndicesReply); i {
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
		file_proto_indices_v1_indices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndicesShortReply); i {
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
		file_proto_indices_v1_indices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndicesShortReplies); i {
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
			RawDescriptor: file_proto_indices_v1_indices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_indices_v1_indices_proto_goTypes,
		DependencyIndexes: file_proto_indices_v1_indices_proto_depIdxs,
		MessageInfos:      file_proto_indices_v1_indices_proto_msgTypes,
	}.Build()
	File_proto_indices_v1_indices_proto = out.File
	file_proto_indices_v1_indices_proto_rawDesc = nil
	file_proto_indices_v1_indices_proto_goTypes = nil
	file_proto_indices_v1_indices_proto_depIdxs = nil
}
