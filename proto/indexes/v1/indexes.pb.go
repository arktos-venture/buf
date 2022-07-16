// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/indexes/v1/indexes.proto

package v1Indexes

import (
	v1 "github.com/arktos-venture/buf/proto/instruments/v1"
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

type IndexSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []string `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	Sort    []string `protobuf:"bytes,2,rep,name=sort,proto3" json:"sort,omitempty"`
	Limit   int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *IndexSearch) Reset() {
	*x = IndexSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSearch) ProtoMessage() {}

func (x *IndexSearch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSearch.ProtoReflect.Descriptor instead.
func (*IndexSearch) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{0}
}

func (x *IndexSearch) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *IndexSearch) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *IndexSearch) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type IndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Ticker  string `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
}

func (x *IndexRequest) Reset() {
	*x = IndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRequest) ProtoMessage() {}

func (x *IndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRequest.ProtoReflect.Descriptor instead.
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{1}
}

func (x *IndexRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *IndexRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

type IndexSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *IndexSearchRequest) Reset() {
	*x = IndexSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSearchRequest) ProtoMessage() {}

func (x *IndexSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSearchRequest.ProtoReflect.Descriptor instead.
func (*IndexSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{2}
}

func (x *IndexSearchRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type IndexCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string       `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Ticker  string       `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Request *IndexSearch `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Public  bool         `protobuf:"varint,4,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *IndexCreateRequest) Reset() {
	*x = IndexCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexCreateRequest) ProtoMessage() {}

func (x *IndexCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexCreateRequest.ProtoReflect.Descriptor instead.
func (*IndexCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{3}
}

func (x *IndexCreateRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *IndexCreateRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *IndexCreateRequest) GetRequest() *IndexSearch {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *IndexCreateRequest) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

type IndexDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickers []string `protobuf:"bytes,1,rep,name=tickers,proto3" json:"tickers,omitempty"`
}

func (x *IndexDeleteRequest) Reset() {
	*x = IndexDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDeleteRequest) ProtoMessage() {}

func (x *IndexDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDeleteRequest.ProtoReflect.Descriptor instead.
func (*IndexDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{4}
}

func (x *IndexDeleteRequest) GetTickers() []string {
	if x != nil {
		return x.Tickers
	}
	return nil
}

type IndexReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker    string                         `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Public    bool                           `protobuf:"varint,2,opt,name=public,proto3" json:"public,omitempty"`
	Companies []*v1.InstrumentReplies_Result `protobuf:"bytes,3,rep,name=companies,proto3" json:"companies,omitempty"`
	CreatedAt *timestamppb.Timestamp         `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp         `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *IndexReply) Reset() {
	*x = IndexReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReply) ProtoMessage() {}

func (x *IndexReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexReply.ProtoReflect.Descriptor instead.
func (*IndexReply) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{5}
}

func (x *IndexReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *IndexReply) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *IndexReply) GetCompanies() []*v1.InstrumentReplies_Result {
	if x != nil {
		return x.Companies
	}
	return nil
}

func (x *IndexReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IndexReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type IndexReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*IndexReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int64         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *IndexReplies) Reset() {
	*x = IndexReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReplies) ProtoMessage() {}

func (x *IndexReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexReplies.ProtoReflect.Descriptor instead.
func (*IndexReplies) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{6}
}

func (x *IndexReplies) GetResults() []*IndexReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *IndexReplies) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type IndexSearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker    string                 `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Public    bool                   `protobuf:"varint,2,opt,name=public,proto3" json:"public,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *IndexSearchReply) Reset() {
	*x = IndexSearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSearchReply) ProtoMessage() {}

func (x *IndexSearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSearchReply.ProtoReflect.Descriptor instead.
func (*IndexSearchReply) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{7}
}

func (x *IndexSearchReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *IndexSearchReply) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *IndexSearchReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IndexSearchReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type IndexSearchReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*IndexSearchReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int64               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *IndexSearchReplies) Reset() {
	*x = IndexSearchReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSearchReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSearchReplies) ProtoMessage() {}

func (x *IndexSearchReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSearchReplies.ProtoReflect.Descriptor instead.
func (*IndexSearchReplies) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{8}
}

func (x *IndexSearchReplies) GetResults() []*IndexSearchReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *IndexSearchReplies) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type IndexDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *IndexDelete) Reset() {
	*x = IndexDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indexes_v1_indexes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDelete) ProtoMessage() {}

func (x *IndexDelete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indexes_v1_indexes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDelete.ProtoReflect.Descriptor instead.
func (*IndexDelete) Descriptor() ([]byte, []int) {
	return file_proto_indexes_v1_indexes_proto_rawDescGZIP(), []int{9}
}

func (x *IndexDelete) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_indexes_v1_indexes_proto protoreflect.FileDescriptor

var file_proto_indexes_v1_indexes_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x56, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x24, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x24, 0x52,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x12, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x24, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x03, 0x18, 0x24, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x24, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x40, 0x0a, 0x12,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x18, 0x01, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x22, 0xfa,
	0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x46, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x56, 0x0a, 0x0c, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0xb8, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x62,
	0x0a, 0x12, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x23, 0x0a, 0x0b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xe5, 0x03, 0x0a, 0x07, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x7d, 0x12, 0x6e, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x58, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x42,
	0x80, 0x03, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x41,
	0x50, 0x49, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x92, 0x41, 0x92, 0x02,
	0x12, 0x7f, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73,
	0x12, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x22, 0x5c,
	0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65,
	0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b,
	0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x1a,
	0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c,
	0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_indexes_v1_indexes_proto_rawDescOnce sync.Once
	file_proto_indexes_v1_indexes_proto_rawDescData = file_proto_indexes_v1_indexes_proto_rawDesc
)

func file_proto_indexes_v1_indexes_proto_rawDescGZIP() []byte {
	file_proto_indexes_v1_indexes_proto_rawDescOnce.Do(func() {
		file_proto_indexes_v1_indexes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_indexes_v1_indexes_proto_rawDescData)
	})
	return file_proto_indexes_v1_indexes_proto_rawDescData
}

var file_proto_indexes_v1_indexes_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_indexes_v1_indexes_proto_goTypes = []interface{}{
	(*IndexSearch)(nil),                 // 0: indexes.v1.IndexSearch
	(*IndexRequest)(nil),                // 1: indexes.v1.IndexRequest
	(*IndexSearchRequest)(nil),          // 2: indexes.v1.IndexSearchRequest
	(*IndexCreateRequest)(nil),          // 3: indexes.v1.IndexCreateRequest
	(*IndexDeleteRequest)(nil),          // 4: indexes.v1.IndexDeleteRequest
	(*IndexReply)(nil),                  // 5: indexes.v1.IndexReply
	(*IndexReplies)(nil),                // 6: indexes.v1.IndexReplies
	(*IndexSearchReply)(nil),            // 7: indexes.v1.IndexSearchReply
	(*IndexSearchReplies)(nil),          // 8: indexes.v1.IndexSearchReplies
	(*IndexDelete)(nil),                 // 9: indexes.v1.IndexDelete
	(*v1.InstrumentReplies_Result)(nil), // 10: instruments.v1.InstrumentReplies.Result
	(*timestamppb.Timestamp)(nil),       // 11: google.protobuf.Timestamp
}
var file_proto_indexes_v1_indexes_proto_depIdxs = []int32{
	0,  // 0: indexes.v1.IndexCreateRequest.request:type_name -> indexes.v1.IndexSearch
	10, // 1: indexes.v1.IndexReply.companies:type_name -> instruments.v1.InstrumentReplies.Result
	11, // 2: indexes.v1.IndexReply.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: indexes.v1.IndexReply.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 4: indexes.v1.IndexReplies.results:type_name -> indexes.v1.IndexReply
	11, // 5: indexes.v1.IndexSearchReply.created_at:type_name -> google.protobuf.Timestamp
	11, // 6: indexes.v1.IndexSearchReply.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 7: indexes.v1.IndexSearchReplies.results:type_name -> indexes.v1.IndexSearchReply
	1,  // 8: indexes.v1.Indexes.Get:input_type -> indexes.v1.IndexRequest
	2,  // 9: indexes.v1.Indexes.Search:input_type -> indexes.v1.IndexSearchRequest
	3,  // 10: indexes.v1.Indexes.Create:input_type -> indexes.v1.IndexCreateRequest
	3,  // 11: indexes.v1.Indexes.Update:input_type -> indexes.v1.IndexCreateRequest
	4,  // 12: indexes.v1.Indexes.Delete:input_type -> indexes.v1.IndexDeleteRequest
	5,  // 13: indexes.v1.Indexes.Get:output_type -> indexes.v1.IndexReply
	8,  // 14: indexes.v1.Indexes.Search:output_type -> indexes.v1.IndexSearchReplies
	5,  // 15: indexes.v1.Indexes.Create:output_type -> indexes.v1.IndexReply
	5,  // 16: indexes.v1.Indexes.Update:output_type -> indexes.v1.IndexReply
	9,  // 17: indexes.v1.Indexes.Delete:output_type -> indexes.v1.IndexDelete
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_indexes_v1_indexes_proto_init() }
func file_proto_indexes_v1_indexes_proto_init() {
	if File_proto_indexes_v1_indexes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_indexes_v1_indexes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSearch); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRequest); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSearchRequest); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexCreateRequest); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDeleteRequest); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexReply); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexReplies); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSearchReply); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSearchReplies); i {
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
		file_proto_indexes_v1_indexes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDelete); i {
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
			RawDescriptor: file_proto_indexes_v1_indexes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_indexes_v1_indexes_proto_goTypes,
		DependencyIndexes: file_proto_indexes_v1_indexes_proto_depIdxs,
		MessageInfos:      file_proto_indexes_v1_indexes_proto_msgTypes,
	}.Build()
	File_proto_indexes_v1_indexes_proto = out.File
	file_proto_indexes_v1_indexes_proto_rawDesc = nil
	file_proto_indexes_v1_indexes_proto_goTypes = nil
	file_proto_indexes_v1_indexes_proto_depIdxs = nil
}
