// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/index/v1/index.proto

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

type IndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IndexRequest) Reset() {
	*x = IndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_index_v1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRequest) ProtoMessage() {}

func (x *IndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_index_v1_index_proto_msgTypes[0]
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
	return file_proto_index_v1_index_proto_rawDescGZIP(), []int{0}
}

func (x *IndexRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IndexCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Companies   []*IndexCreateRequest_Company `protobuf:"bytes,3,rep,name=companies,proto3" json:"companies,omitempty"`
}

func (x *IndexCreateRequest) Reset() {
	*x = IndexCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_index_v1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexCreateRequest) ProtoMessage() {}

func (x *IndexCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_index_v1_index_proto_msgTypes[1]
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
	return file_proto_index_v1_index_proto_rawDescGZIP(), []int{1}
}

func (x *IndexCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndexCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IndexCreateRequest) GetCompanies() []*IndexCreateRequest_Company {
	if x != nil {
		return x.Companies
	}
	return nil
}

type IndexReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *IndexReply) Reset() {
	*x = IndexReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_index_v1_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReply) ProtoMessage() {}

func (x *IndexReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_index_v1_index_proto_msgTypes[2]
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
	return file_proto_index_v1_index_proto_rawDescGZIP(), []int{2}
}

func (x *IndexReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndexReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type IndexReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*IndexReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *IndexReplies) Reset() {
	*x = IndexReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_index_v1_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReplies) ProtoMessage() {}

func (x *IndexReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_index_v1_index_proto_msgTypes[3]
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
	return file_proto_index_v1_index_proto_rawDescGZIP(), []int{3}
}

func (x *IndexReplies) GetResults() []*IndexReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *IndexReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type IndexCreateRequest_Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency string  `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Exchange string  `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Ticker   string  `protobuf:"bytes,3,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Size     float32 `protobuf:"fixed32,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *IndexCreateRequest_Company) Reset() {
	*x = IndexCreateRequest_Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_index_v1_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexCreateRequest_Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexCreateRequest_Company) ProtoMessage() {}

func (x *IndexCreateRequest_Company) ProtoReflect() protoreflect.Message {
	mi := &file_proto_index_v1_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexCreateRequest_Company.ProtoReflect.Descriptor instead.
func (*IndexCreateRequest_Company) Descriptor() ([]byte, []int) {
	return file_proto_index_v1_index_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IndexCreateRequest_Company) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *IndexCreateRequest_Company) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *IndexCreateRequest_Company) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *IndexCreateRequest_Company) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_proto_index_v1_index_proto protoreflect.FileDescriptor

var file_proto_index_v1_index_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0c, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x03, 0x18, 0x18, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x12, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x18, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x10, 0x18, 0x80, 0x02,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x1a, 0x6d, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x42, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xe9, 0x03, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x52, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x1a, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x01, 0x2a, 0x12, 0x50, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x43, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_index_v1_index_proto_rawDescOnce sync.Once
	file_proto_index_v1_index_proto_rawDescData = file_proto_index_v1_index_proto_rawDesc
)

func file_proto_index_v1_index_proto_rawDescGZIP() []byte {
	file_proto_index_v1_index_proto_rawDescOnce.Do(func() {
		file_proto_index_v1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_index_v1_index_proto_rawDescData)
	})
	return file_proto_index_v1_index_proto_rawDescData
}

var file_proto_index_v1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_index_v1_index_proto_goTypes = []interface{}{
	(*IndexRequest)(nil),               // 0: index.v1.IndexRequest
	(*IndexCreateRequest)(nil),         // 1: index.v1.IndexCreateRequest
	(*IndexReply)(nil),                 // 2: index.v1.IndexReply
	(*IndexReplies)(nil),               // 3: index.v1.IndexReplies
	(*IndexCreateRequest_Company)(nil), // 4: index.v1.IndexCreateRequest.Company
	(*emptypb.Empty)(nil),              // 5: google.protobuf.Empty
}
var file_proto_index_v1_index_proto_depIdxs = []int32{
	4, // 0: index.v1.IndexCreateRequest.companies:type_name -> index.v1.IndexCreateRequest.Company
	2, // 1: index.v1.IndexReplies.results:type_name -> index.v1.IndexReply
	0, // 2: index.v1.Index.Get:input_type -> index.v1.IndexRequest
	0, // 3: index.v1.Index.Search:input_type -> index.v1.IndexRequest
	1, // 4: index.v1.Index.Create:input_type -> index.v1.IndexCreateRequest
	1, // 5: index.v1.Index.Update:input_type -> index.v1.IndexCreateRequest
	0, // 6: index.v1.Index.Delete:input_type -> index.v1.IndexRequest
	5, // 7: index.v1.Index.Health:input_type -> google.protobuf.Empty
	2, // 8: index.v1.Index.Get:output_type -> index.v1.IndexReply
	3, // 9: index.v1.Index.Search:output_type -> index.v1.IndexReplies
	2, // 10: index.v1.Index.Create:output_type -> index.v1.IndexReply
	2, // 11: index.v1.Index.Update:output_type -> index.v1.IndexReply
	2, // 12: index.v1.Index.Delete:output_type -> index.v1.IndexReply
	5, // 13: index.v1.Index.Health:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_index_v1_index_proto_init() }
func file_proto_index_v1_index_proto_init() {
	if File_proto_index_v1_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_index_v1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_index_v1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_index_v1_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_index_v1_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_index_v1_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexCreateRequest_Company); i {
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
			RawDescriptor: file_proto_index_v1_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_index_v1_index_proto_goTypes,
		DependencyIndexes: file_proto_index_v1_index_proto_depIdxs,
		MessageInfos:      file_proto_index_v1_index_proto_msgTypes,
	}.Build()
	File_proto_index_v1_index_proto = out.File
	file_proto_index_v1_index_proto_rawDesc = nil
	file_proto_index_v1_index_proto_goTypes = nil
	file_proto_index_v1_index_proto_depIdxs = nil
}
