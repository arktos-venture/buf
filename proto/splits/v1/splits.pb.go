// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/splits/v1/splits.proto

package splits_v1

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

type SplitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval v1.Interval        `protobuf:"varint,1,opt,name=interval,proto3,enum=screener.v1.Interval" json:"interval,omitempty"`
	Filters  []*v1.FilterSimple `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *SplitsRequest) Reset() {
	*x = SplitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_splits_v1_splits_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitsRequest) ProtoMessage() {}

func (x *SplitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_splits_v1_splits_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitsRequest.ProtoReflect.Descriptor instead.
func (*SplitsRequest) Descriptor() ([]byte, []int) {
	return file_proto_splits_v1_splits_proto_rawDescGZIP(), []int{0}
}

func (x *SplitsRequest) GetInterval() v1.Interval {
	if x != nil {
		return x.Interval
	}
	return v1.Interval(0)
}

func (x *SplitsRequest) GetFilters() []*v1.FilterSimple {
	if x != nil {
		return x.Filters
	}
	return nil
}

type SplitsLastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange string `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Ticker   string `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
}

func (x *SplitsLastRequest) Reset() {
	*x = SplitsLastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_splits_v1_splits_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitsLastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitsLastRequest) ProtoMessage() {}

func (x *SplitsLastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_splits_v1_splits_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitsLastRequest.ProtoReflect.Descriptor instead.
func (*SplitsLastRequest) Descriptor() ([]byte, []int) {
	return file_proto_splits_v1_splits_proto_rawDescGZIP(), []int{1}
}

func (x *SplitsLastRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *SplitsLastRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

type SplitsLastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      float32                `protobuf:"fixed32,1,opt,name=from,proto3" json:"from,omitempty"`
	To        float32                `protobuf:"fixed32,2,opt,name=to,proto3" json:"to,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SplitsLastReply) Reset() {
	*x = SplitsLastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_splits_v1_splits_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitsLastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitsLastReply) ProtoMessage() {}

func (x *SplitsLastReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_splits_v1_splits_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitsLastReply.ProtoReflect.Descriptor instead.
func (*SplitsLastReply) Descriptor() ([]byte, []int) {
	return file_proto_splits_v1_splits_proto_rawDescGZIP(), []int{2}
}

func (x *SplitsLastReply) GetFrom() float32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *SplitsLastReply) GetTo() float32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *SplitsLastReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SplitsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      []float32                `protobuf:"fixed32,1,rep,packed,name=from,proto3" json:"from,omitempty"`
	To        []float32                `protobuf:"fixed32,2,rep,packed,name=to,proto3" json:"to,omitempty"`
	CreatedAt []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SplitsReply) Reset() {
	*x = SplitsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_splits_v1_splits_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitsReply) ProtoMessage() {}

func (x *SplitsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_splits_v1_splits_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitsReply.ProtoReflect.Descriptor instead.
func (*SplitsReply) Descriptor() ([]byte, []int) {
	return file_proto_splits_v1_splits_proto_rawDescGZIP(), []int{3}
}

func (x *SplitsReply) GetFrom() []float32 {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *SplitsReply) GetTo() []float32 {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SplitsReply) GetCreatedAt() []*timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_proto_splits_v1_splits_proto protoreflect.FileDescriptor

var file_proto_splits_v1_splits_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
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
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x10,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x14,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x11, 0x53, 0x70, 0x6c,
	0x69, 0x74, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x22, 0x70, 0x0a, 0x0f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x4c, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x6c, 0x0a, 0x0b, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x32, 0xc5, 0x01, 0x0a, 0x06, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x12, 0x68, 0x0a,
	0x04, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x7d, 0x2f, 0x7b,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x7d, 0x12, 0x51, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x18, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70,
	0x6c, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0xfb, 0x02, 0x0a, 0x18, 0x64,
	0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x50, 0x49, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x73, 0x5f, 0x76, 0x31, 0x92, 0x41, 0x90, 0x02, 0x12, 0x7d, 0x0a, 0x0b, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x12, 0x0b, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72,
	0x65, 0x2f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f,
	0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a,
	0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12,
	0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20,
	0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_splits_v1_splits_proto_rawDescOnce sync.Once
	file_proto_splits_v1_splits_proto_rawDescData = file_proto_splits_v1_splits_proto_rawDesc
)

func file_proto_splits_v1_splits_proto_rawDescGZIP() []byte {
	file_proto_splits_v1_splits_proto_rawDescOnce.Do(func() {
		file_proto_splits_v1_splits_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_splits_v1_splits_proto_rawDescData)
	})
	return file_proto_splits_v1_splits_proto_rawDescData
}

var file_proto_splits_v1_splits_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_splits_v1_splits_proto_goTypes = []interface{}{
	(*SplitsRequest)(nil),         // 0: splits.v1.SplitsRequest
	(*SplitsLastRequest)(nil),     // 1: splits.v1.SplitsLastRequest
	(*SplitsLastReply)(nil),       // 2: splits.v1.SplitsLastReply
	(*SplitsReply)(nil),           // 3: splits.v1.SplitsReply
	(v1.Interval)(0),              // 4: screener.v1.Interval
	(*v1.FilterSimple)(nil),       // 5: screener.v1.FilterSimple
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_splits_v1_splits_proto_depIdxs = []int32{
	4, // 0: splits.v1.SplitsRequest.interval:type_name -> screener.v1.Interval
	5, // 1: splits.v1.SplitsRequest.filters:type_name -> screener.v1.FilterSimple
	6, // 2: splits.v1.SplitsLastReply.created_at:type_name -> google.protobuf.Timestamp
	6, // 3: splits.v1.SplitsReply.created_at:type_name -> google.protobuf.Timestamp
	1, // 4: splits.v1.Splits.Last:input_type -> splits.v1.SplitsLastRequest
	0, // 5: splits.v1.Splits.Search:input_type -> splits.v1.SplitsRequest
	2, // 6: splits.v1.Splits.Last:output_type -> splits.v1.SplitsLastReply
	3, // 7: splits.v1.Splits.Search:output_type -> splits.v1.SplitsReply
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_splits_v1_splits_proto_init() }
func file_proto_splits_v1_splits_proto_init() {
	if File_proto_splits_v1_splits_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_splits_v1_splits_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitsRequest); i {
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
		file_proto_splits_v1_splits_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitsLastRequest); i {
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
		file_proto_splits_v1_splits_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitsLastReply); i {
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
		file_proto_splits_v1_splits_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitsReply); i {
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
			RawDescriptor: file_proto_splits_v1_splits_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_splits_v1_splits_proto_goTypes,
		DependencyIndexes: file_proto_splits_v1_splits_proto_depIdxs,
		MessageInfos:      file_proto_splits_v1_splits_proto_msgTypes,
	}.Build()
	File_proto_splits_v1_splits_proto = out.File
	file_proto_splits_v1_splits_proto_rawDesc = nil
	file_proto_splits_v1_splits_proto_goTypes = nil
	file_proto_splits_v1_splits_proto_depIdxs = nil
}
