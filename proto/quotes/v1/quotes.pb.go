// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.2
// source: proto/quotes/v1/quotes.proto

package v1Quotes

import (
	v1 "github.com/arktos-venture/buf/proto/screener/v1"
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

type QuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *QuoteRequest) Reset() {
	*x = QuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteRequest) ProtoMessage() {}

func (x *QuoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QuoteRequest.ProtoReflect.Descriptor instead.
func (*QuoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{0}
}

func (x *QuoteRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *QuoteRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type QuoteSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval v1.Interval  `protobuf:"varint,1,opt,name=interval,proto3,enum=screener.v1.Interval" json:"interval,omitempty"`
	Filters  []*v1.Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *QuoteSearchRequest) Reset() {
	*x = QuoteSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteSearchRequest) ProtoMessage() {}

func (x *QuoteSearchRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QuoteSearchRequest.ProtoReflect.Descriptor instead.
func (*QuoteSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{1}
}

func (x *QuoteSearchRequest) GetInterval() v1.Interval {
	if x != nil {
		return x.Interval
	}
	return v1.Interval(0)
}

func (x *QuoteSearchRequest) GetFilters() []*v1.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type QuoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open      float32                `protobuf:"fixed32,1,opt,name=open,proto3" json:"open,omitempty"`
	Close     float32                `protobuf:"fixed32,2,opt,name=close,proto3" json:"close,omitempty"`
	Low       float32                `protobuf:"fixed32,3,opt,name=low,proto3" json:"low,omitempty"`
	High      float32                `protobuf:"fixed32,4,opt,name=high,proto3" json:"high,omitempty"`
	Volume    int64                  `protobuf:"varint,5,opt,name=volume,proto3" json:"volume,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *QuoteReply) Reset() {
	*x = QuoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteReply) ProtoMessage() {}

func (x *QuoteReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QuoteReply.ProtoReflect.Descriptor instead.
func (*QuoteReply) Descriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{2}
}

func (x *QuoteReply) GetOpen() float32 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *QuoteReply) GetClose() float32 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *QuoteReply) GetLow() float32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *QuoteReply) GetHigh() float32 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *QuoteReply) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *QuoteReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type QuoteReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*QuoteReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int64         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *QuoteReplies) Reset() {
	*x = QuoteReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_v1_quotes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteReplies) ProtoMessage() {}

func (x *QuoteReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_v1_quotes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteReplies.ProtoReflect.Descriptor instead.
func (*QuoteReplies) Descriptor() ([]byte, []int) {
	return file_proto_quotes_v1_quotes_proto_rawDescGZIP(), []int{3}
}

func (x *QuoteReplies) GetResults() []*QuoteReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *QuoteReplies) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_quotes_v1_quotes_proto protoreflect.FileDescriptor

var file_proto_quotes_v1_quotes_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
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
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x10, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x01, 0x18, 0x10, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x92, 0x01,
	0x0a, 0x12, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x14, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0a, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x68, 0x69, 0x67,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x55, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc5, 0x01, 0x0a, 0x06,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x04, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x17,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x2f, 0x7b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x7d, 0x2f, 0x7b, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x7d, 0x2f, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x3a, 0x01, 0x2a, 0x42, 0x8f, 0x03, 0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x41, 0x50, 0x49, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x92, 0x41, 0xa5, 0x02,
	0x12, 0x91, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73,
	0x12, 0x1f, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65, 0x6e, 0x74,
	0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62,
	0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_quotes_v1_quotes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_quotes_v1_quotes_proto_goTypes = []interface{}{
	(*QuoteRequest)(nil),          // 0: quotes.v1.QuoteRequest
	(*QuoteSearchRequest)(nil),    // 1: quotes.v1.QuoteSearchRequest
	(*QuoteReply)(nil),            // 2: quotes.v1.QuoteReply
	(*QuoteReplies)(nil),          // 3: quotes.v1.QuoteReplies
	(v1.Interval)(0),              // 4: screener.v1.Interval
	(*v1.Filter)(nil),             // 5: screener.v1.Filter
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_quotes_v1_quotes_proto_depIdxs = []int32{
	4, // 0: quotes.v1.QuoteSearchRequest.interval:type_name -> screener.v1.Interval
	5, // 1: quotes.v1.QuoteSearchRequest.filters:type_name -> screener.v1.Filter
	6, // 2: quotes.v1.QuoteReply.created_at:type_name -> google.protobuf.Timestamp
	2, // 3: quotes.v1.QuoteReplies.results:type_name -> quotes.v1.QuoteReply
	0, // 4: quotes.v1.quotes.Last:input_type -> quotes.v1.QuoteRequest
	1, // 5: quotes.v1.quotes.Search:input_type -> quotes.v1.QuoteSearchRequest
	2, // 6: quotes.v1.quotes.Last:output_type -> quotes.v1.QuoteReply
	3, // 7: quotes.v1.quotes.Search:output_type -> quotes.v1.QuoteReplies
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_quotes_v1_quotes_proto_init() }
func file_proto_quotes_v1_quotes_proto_init() {
	if File_proto_quotes_v1_quotes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quotes_v1_quotes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteRequest); i {
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
			switch v := v.(*QuoteSearchRequest); i {
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
			switch v := v.(*QuoteReply); i {
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
		file_proto_quotes_v1_quotes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteReplies); i {
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
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quotes_v1_quotes_proto_goTypes,
		DependencyIndexes: file_proto_quotes_v1_quotes_proto_depIdxs,
		MessageInfos:      file_proto_quotes_v1_quotes_proto_msgTypes,
	}.Build()
	File_proto_quotes_v1_quotes_proto = out.File
	file_proto_quotes_v1_quotes_proto_rawDesc = nil
	file_proto_quotes_v1_quotes_proto_goTypes = nil
	file_proto_quotes_v1_quotes_proto_depIdxs = nil
}
