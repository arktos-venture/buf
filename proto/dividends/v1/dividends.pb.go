// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.4
// source: proto/dividends/v1/dividends.proto

package v1Dividends

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

type DividendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval v1.Interval  `protobuf:"varint,1,opt,name=interval,proto3,enum=screener.v1.Interval" json:"interval,omitempty"`
	Filters  []*v1.Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *DividendRequest) Reset() {
	*x = DividendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DividendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DividendRequest) ProtoMessage() {}

func (x *DividendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DividendRequest.ProtoReflect.Descriptor instead.
func (*DividendRequest) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{0}
}

func (x *DividendRequest) GetInterval() v1.Interval {
	if x != nil {
		return x.Interval
	}
	return v1.Interval(0)
}

func (x *DividendRequest) GetFilters() []*v1.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type DividendLastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstrumentId int64 `protobuf:"varint,1,opt,name=instrument_id,json=instrumentId,proto3" json:"instrument_id,omitempty"`
}

func (x *DividendLastRequest) Reset() {
	*x = DividendLastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DividendLastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DividendLastRequest) ProtoMessage() {}

func (x *DividendLastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DividendLastRequest.ProtoReflect.Descriptor instead.
func (*DividendLastRequest) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{1}
}

func (x *DividendLastRequest) GetInstrumentId() int64 {
	if x != nil {
		return x.InstrumentId
	}
	return 0
}

type DividendLastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values          float32                `protobuf:"fixed32,1,opt,name=values,proto3" json:"values,omitempty"`
	Yield           float32                `protobuf:"fixed32,2,opt,name=yield,proto3" json:"yield,omitempty"`
	DeclarationDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=declaration_date,json=declarationDate,proto3" json:"declaration_date,omitempty"`
	RecordDate      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=record_date,json=recordDate,proto3" json:"record_date,omitempty"`
	PaymentDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *DividendLastReply) Reset() {
	*x = DividendLastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DividendLastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DividendLastReply) ProtoMessage() {}

func (x *DividendLastReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DividendLastReply.ProtoReflect.Descriptor instead.
func (*DividendLastReply) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{2}
}

func (x *DividendLastReply) GetValues() float32 {
	if x != nil {
		return x.Values
	}
	return 0
}

func (x *DividendLastReply) GetYield() float32 {
	if x != nil {
		return x.Yield
	}
	return 0
}

func (x *DividendLastReply) GetDeclarationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DeclarationDate
	}
	return nil
}

func (x *DividendLastReply) GetRecordDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordDate
	}
	return nil
}

func (x *DividendLastReply) GetPaymentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDate
	}
	return nil
}

func (x *DividendLastReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type DividendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values          []float32                `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	Yields          []float32                `protobuf:"fixed32,2,rep,packed,name=yields,proto3" json:"yields,omitempty"`
	DeclarationDate []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=declaration_date,json=declarationDate,proto3" json:"declaration_date,omitempty"`
	RecordDate      []*timestamppb.Timestamp `protobuf:"bytes,4,rep,name=record_date,json=recordDate,proto3" json:"record_date,omitempty"`
	PaymentDate     []*timestamppb.Timestamp `protobuf:"bytes,5,rep,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	CreatedAt       []*timestamppb.Timestamp `protobuf:"bytes,6,rep,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *DividendReply) Reset() {
	*x = DividendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dividends_v1_dividends_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DividendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DividendReply) ProtoMessage() {}

func (x *DividendReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dividends_v1_dividends_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DividendReply.ProtoReflect.Descriptor instead.
func (*DividendReply) Descriptor() ([]byte, []int) {
	return file_proto_dividends_v1_dividends_proto_rawDescGZIP(), []int{3}
}

func (x *DividendReply) GetValues() []float32 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *DividendReply) GetYields() []float32 {
	if x != nil {
		return x.Yields
	}
	return nil
}

func (x *DividendReply) GetDeclarationDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.DeclarationDate
	}
	return nil
}

func (x *DividendReply) GetRecordDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.RecordDate
	}
	return nil
}

func (x *DividendReply) GetPaymentDate() []*timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDate
	}
	return nil
}

func (x *DividendReply) GetCreatedAt() []*timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_proto_dividends_v1_dividends_proto protoreflect.FileDescriptor

var file_proto_dividends_v1_dividends_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a,
	0x0f, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3f, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x10, 0x14, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x43,
	0x0a, 0x13, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x10, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0xbf, 0x02, 0x0a, 0x11, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64,
	0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x79, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x79, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x45, 0x0a, 0x10, 0x64, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x64,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x79, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52,
	0x06, 0x79, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x64, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x64,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xdd, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x64, 0x73, 0x12, 0x70, 0x0a, 0x04, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x76,
	0x69, 0x64, 0x65, 0x6e, 0x64, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1d, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x64, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x8b, 0x03, 0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x41, 0x50, 0x49, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65,
	0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x44,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x92, 0x41, 0x97, 0x02, 0x12, 0x83, 0x01, 0x0a,
	0x0e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x12,
	0x0e, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x22,
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
	file_proto_dividends_v1_dividends_proto_rawDescOnce sync.Once
	file_proto_dividends_v1_dividends_proto_rawDescData = file_proto_dividends_v1_dividends_proto_rawDesc
)

func file_proto_dividends_v1_dividends_proto_rawDescGZIP() []byte {
	file_proto_dividends_v1_dividends_proto_rawDescOnce.Do(func() {
		file_proto_dividends_v1_dividends_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dividends_v1_dividends_proto_rawDescData)
	})
	return file_proto_dividends_v1_dividends_proto_rawDescData
}

var file_proto_dividends_v1_dividends_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_dividends_v1_dividends_proto_goTypes = []interface{}{
	(*DividendRequest)(nil),       // 0: dividends.v1.DividendRequest
	(*DividendLastRequest)(nil),   // 1: dividends.v1.DividendLastRequest
	(*DividendLastReply)(nil),     // 2: dividends.v1.DividendLastReply
	(*DividendReply)(nil),         // 3: dividends.v1.DividendReply
	(v1.Interval)(0),              // 4: screener.v1.Interval
	(*v1.Filter)(nil),             // 5: screener.v1.Filter
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_dividends_v1_dividends_proto_depIdxs = []int32{
	4,  // 0: dividends.v1.DividendRequest.interval:type_name -> screener.v1.Interval
	5,  // 1: dividends.v1.DividendRequest.filters:type_name -> screener.v1.Filter
	6,  // 2: dividends.v1.DividendLastReply.declaration_date:type_name -> google.protobuf.Timestamp
	6,  // 3: dividends.v1.DividendLastReply.record_date:type_name -> google.protobuf.Timestamp
	6,  // 4: dividends.v1.DividendLastReply.payment_date:type_name -> google.protobuf.Timestamp
	6,  // 5: dividends.v1.DividendLastReply.created_at:type_name -> google.protobuf.Timestamp
	6,  // 6: dividends.v1.DividendReply.declaration_date:type_name -> google.protobuf.Timestamp
	6,  // 7: dividends.v1.DividendReply.record_date:type_name -> google.protobuf.Timestamp
	6,  // 8: dividends.v1.DividendReply.payment_date:type_name -> google.protobuf.Timestamp
	6,  // 9: dividends.v1.DividendReply.created_at:type_name -> google.protobuf.Timestamp
	1,  // 10: dividends.v1.Dividends.Last:input_type -> dividends.v1.DividendLastRequest
	0,  // 11: dividends.v1.Dividends.Search:input_type -> dividends.v1.DividendRequest
	2,  // 12: dividends.v1.Dividends.Last:output_type -> dividends.v1.DividendLastReply
	3,  // 13: dividends.v1.Dividends.Search:output_type -> dividends.v1.DividendReply
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_dividends_v1_dividends_proto_init() }
func file_proto_dividends_v1_dividends_proto_init() {
	if File_proto_dividends_v1_dividends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dividends_v1_dividends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DividendRequest); i {
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
		file_proto_dividends_v1_dividends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DividendLastRequest); i {
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
		file_proto_dividends_v1_dividends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DividendLastReply); i {
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
		file_proto_dividends_v1_dividends_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DividendReply); i {
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
			RawDescriptor: file_proto_dividends_v1_dividends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dividends_v1_dividends_proto_goTypes,
		DependencyIndexes: file_proto_dividends_v1_dividends_proto_depIdxs,
		MessageInfos:      file_proto_dividends_v1_dividends_proto_msgTypes,
	}.Build()
	File_proto_dividends_v1_dividends_proto = out.File
	file_proto_dividends_v1_dividends_proto_rawDesc = nil
	file_proto_dividends_v1_dividends_proto_goTypes = nil
	file_proto_dividends_v1_dividends_proto_depIdxs = nil
}
