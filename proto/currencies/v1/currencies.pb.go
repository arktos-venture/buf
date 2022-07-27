// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.4
// source: proto/currencies/v1/currencies.proto

package v1Currencies

import (
	v11 "github.com/arktos-venture/buf/proto/countries/v1"
	v1 "github.com/arktos-venture/buf/proto/exchanges/v1"
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

type CurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
}

func (x *CurrencyRequest) Reset() {
	*x = CurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_v1_currencies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRequest) ProtoMessage() {}

func (x *CurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_v1_currencies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRequest.ProtoReflect.Descriptor instead.
func (*CurrencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_currencies_v1_currencies_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

type CurrencyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker      string                       `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	CentralBank *CurrencyReply_CentralBank   `protobuf:"bytes,2,opt,name=central_bank,json=centralBank,proto3" json:"central_bank,omitempty"`
	Exchanges   []*v1.ExchangeReplies_Result `protobuf:"bytes,3,rep,name=exchanges,proto3" json:"exchanges,omitempty"`
	Countries   []*v11.CountryReplies_Result `protobuf:"bytes,4,rep,name=countries,proto3" json:"countries,omitempty"`
	CreatedAt   *timestamppb.Timestamp       `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp       `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CurrencyReply) Reset() {
	*x = CurrencyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_v1_currencies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyReply) ProtoMessage() {}

func (x *CurrencyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_v1_currencies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyReply.ProtoReflect.Descriptor instead.
func (*CurrencyReply) Descriptor() ([]byte, []int) {
	return file_proto_currencies_v1_currencies_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CurrencyReply) GetCentralBank() *CurrencyReply_CentralBank {
	if x != nil {
		return x.CentralBank
	}
	return nil
}

func (x *CurrencyReply) GetExchanges() []*v1.ExchangeReplies_Result {
	if x != nil {
		return x.Exchanges
	}
	return nil
}

func (x *CurrencyReply) GetCountries() []*v11.CountryReplies_Result {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *CurrencyReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CurrencyReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CurrencyReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []string `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CurrencyReplies) Reset() {
	*x = CurrencyReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_v1_currencies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyReplies) ProtoMessage() {}

func (x *CurrencyReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_v1_currencies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyReplies.ProtoReflect.Descriptor instead.
func (*CurrencyReplies) Descriptor() ([]byte, []int) {
	return file_proto_currencies_v1_currencies_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyReplies) GetResults() []string {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CurrencyReplies) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CurrencyReply_CentralBank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	BorrowingRate *CurrencyReply_CentralBank_BorrowingRate `protobuf:"bytes,3,opt,name=borrowing_rate,json=borrowingRate,proto3" json:"borrowing_rate,omitempty"`
}

func (x *CurrencyReply_CentralBank) Reset() {
	*x = CurrencyReply_CentralBank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_v1_currencies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyReply_CentralBank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyReply_CentralBank) ProtoMessage() {}

func (x *CurrencyReply_CentralBank) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_v1_currencies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyReply_CentralBank.ProtoReflect.Descriptor instead.
func (*CurrencyReply_CentralBank) Descriptor() ([]byte, []int) {
	return file_proto_currencies_v1_currencies_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CurrencyReply_CentralBank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CurrencyReply_CentralBank) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CurrencyReply_CentralBank) GetBorrowingRate() *CurrencyReply_CentralBank_BorrowingRate {
	if x != nil {
		return x.BorrowingRate
	}
	return nil
}

type CurrencyReply_CentralBank_BorrowingRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate  float32                `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	Since *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *CurrencyReply_CentralBank_BorrowingRate) Reset() {
	*x = CurrencyReply_CentralBank_BorrowingRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_v1_currencies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyReply_CentralBank_BorrowingRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyReply_CentralBank_BorrowingRate) ProtoMessage() {}

func (x *CurrencyReply_CentralBank_BorrowingRate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_v1_currencies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyReply_CentralBank_BorrowingRate.ProtoReflect.Descriptor instead.
func (*CurrencyReply_CentralBank_BorrowingRate) Descriptor() ([]byte, []int) {
	return file_proto_currencies_v1_currencies_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *CurrencyReply_CentralBank_BorrowingRate) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *CurrencyReply_CentralBank_BorrowingRate) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

var File_proto_currencies_v1_currencies_proto protoreflect.FileDescriptor

var file_proto_currencies_v1_currencies_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x03, 0x52,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x22, 0xed, 0x04, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x42, 0x61, 0x6e,
	0x6b, 0x52, 0x0b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x42,
	0x0a, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x12, 0x41, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0xf9, 0x01, 0x0a, 0x0b,
	0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5d, 0x0a, 0x0e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x42,
	0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x0d, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65,
	0x1a, 0x55, 0x0a, 0x0d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc8, 0x01, 0x0a, 0x0a, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x7d, 0x12, 0x56, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x42, 0x93, 0x03, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x41, 0x50, 0x49, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f,
	0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x92, 0x41,
	0x99, 0x02, 0x12, 0x85, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x12, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f,
	0x73, 0x20, 0x56, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e,
	0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72,
	0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d,
	0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a,
	0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_currencies_v1_currencies_proto_rawDescOnce sync.Once
	file_proto_currencies_v1_currencies_proto_rawDescData = file_proto_currencies_v1_currencies_proto_rawDesc
)

func file_proto_currencies_v1_currencies_proto_rawDescGZIP() []byte {
	file_proto_currencies_v1_currencies_proto_rawDescOnce.Do(func() {
		file_proto_currencies_v1_currencies_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currencies_v1_currencies_proto_rawDescData)
	})
	return file_proto_currencies_v1_currencies_proto_rawDescData
}

var file_proto_currencies_v1_currencies_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_currencies_v1_currencies_proto_goTypes = []interface{}{
	(*CurrencyRequest)(nil),                         // 0: currencies.v1.CurrencyRequest
	(*CurrencyReply)(nil),                           // 1: currencies.v1.CurrencyReply
	(*CurrencyReplies)(nil),                         // 2: currencies.v1.CurrencyReplies
	(*CurrencyReply_CentralBank)(nil),               // 3: currencies.v1.CurrencyReply.CentralBank
	(*CurrencyReply_CentralBank_BorrowingRate)(nil), // 4: currencies.v1.CurrencyReply.CentralBank.BorrowingRate
	(*v1.ExchangeReplies_Result)(nil),               // 5: exchanges.v1.ExchangeReplies.Result
	(*v11.CountryReplies_Result)(nil),               // 6: countries.v1.CountryReplies.Result
	(*timestamppb.Timestamp)(nil),                   // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                           // 8: google.protobuf.Empty
}
var file_proto_currencies_v1_currencies_proto_depIdxs = []int32{
	3, // 0: currencies.v1.CurrencyReply.central_bank:type_name -> currencies.v1.CurrencyReply.CentralBank
	5, // 1: currencies.v1.CurrencyReply.exchanges:type_name -> exchanges.v1.ExchangeReplies.Result
	6, // 2: currencies.v1.CurrencyReply.countries:type_name -> countries.v1.CountryReplies.Result
	7, // 3: currencies.v1.CurrencyReply.created_at:type_name -> google.protobuf.Timestamp
	7, // 4: currencies.v1.CurrencyReply.updated_at:type_name -> google.protobuf.Timestamp
	4, // 5: currencies.v1.CurrencyReply.CentralBank.borrowing_rate:type_name -> currencies.v1.CurrencyReply.CentralBank.BorrowingRate
	7, // 6: currencies.v1.CurrencyReply.CentralBank.BorrowingRate.since:type_name -> google.protobuf.Timestamp
	0, // 7: currencies.v1.Currencies.Get:input_type -> currencies.v1.CurrencyRequest
	8, // 8: currencies.v1.Currencies.List:input_type -> google.protobuf.Empty
	1, // 9: currencies.v1.Currencies.Get:output_type -> currencies.v1.CurrencyReply
	2, // 10: currencies.v1.Currencies.List:output_type -> currencies.v1.CurrencyReplies
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_currencies_v1_currencies_proto_init() }
func file_proto_currencies_v1_currencies_proto_init() {
	if File_proto_currencies_v1_currencies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currencies_v1_currencies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyRequest); i {
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
		file_proto_currencies_v1_currencies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyReply); i {
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
		file_proto_currencies_v1_currencies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyReplies); i {
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
		file_proto_currencies_v1_currencies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyReply_CentralBank); i {
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
		file_proto_currencies_v1_currencies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyReply_CentralBank_BorrowingRate); i {
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
			RawDescriptor: file_proto_currencies_v1_currencies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currencies_v1_currencies_proto_goTypes,
		DependencyIndexes: file_proto_currencies_v1_currencies_proto_depIdxs,
		MessageInfos:      file_proto_currencies_v1_currencies_proto_msgTypes,
	}.Build()
	File_proto_currencies_v1_currencies_proto = out.File
	file_proto_currencies_v1_currencies_proto_rawDesc = nil
	file_proto_currencies_v1_currencies_proto_goTypes = nil
	file_proto_currencies_v1_currencies_proto_depIdxs = nil
}
