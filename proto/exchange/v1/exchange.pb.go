// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/exchange/v1/exchange.proto

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

type ExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange string `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *ExchangeRequest) Reset() {
	*x = ExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRequest) ProtoMessage() {}

func (x *ExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eod         string `protobuf:"bytes,1,opt,name=eod,proto3" json:"eod,omitempty"`
	Ibkr        string `protobuf:"bytes,2,opt,name=ibkr,proto3" json:"ibkr,omitempty"`
	Tradingview string `protobuf:"bytes,3,opt,name=tradingview,proto3" json:"tradingview,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *Code) GetEod() string {
	if x != nil {
		return x.Eod
	}
	return ""
}

func (x *Code) GetIbkr() string {
	if x != nil {
		return x.Ibkr
	}
	return ""
}

func (x *Code) GetTradingview() string {
	if x != nil {
		return x.Tradingview
	}
	return ""
}

type TradingHours struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open        map[string]int32 `protobuf:"bytes,1,rep,name=open,proto3" json:"open,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Close       map[string]int32 `protobuf:"bytes,2,rep,name=close,proto3" json:"close,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	WorkingDays []int32          `protobuf:"varint,3,rep,packed,name=workingDays,proto3" json:"workingDays,omitempty"`
}

func (x *TradingHours) Reset() {
	*x = TradingHours{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradingHours) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradingHours) ProtoMessage() {}

func (x *TradingHours) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradingHours.ProtoReflect.Descriptor instead.
func (*TradingHours) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *TradingHours) GetOpen() map[string]int32 {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *TradingHours) GetClose() map[string]int32 {
	if x != nil {
		return x.Close
	}
	return nil
}

func (x *TradingHours) GetWorkingDays() []int32 {
	if x != nil {
		return x.WorkingDays
	}
	return nil
}

type ExchangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code         *Code                    `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Country      string                   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Currency     string                   `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Timezone     string                   `protobuf:"bytes,5,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Holidays     []*ExchangeReply_Holiday `protobuf:"bytes,6,rep,name=holidays,proto3" json:"holidays,omitempty"`
	TradingHours *TradingHours            `protobuf:"bytes,7,opt,name=tradingHours,proto3" json:"tradingHours,omitempty"`
	Companies    []*ExchangeReply_Company `protobuf:"bytes,8,rep,name=companies,proto3" json:"companies,omitempty"`
}

func (x *ExchangeReply) Reset() {
	*x = ExchangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeReply) ProtoMessage() {}

func (x *ExchangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeReply.ProtoReflect.Descriptor instead.
func (*ExchangeReply) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExchangeReply) GetCode() *Code {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *ExchangeReply) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ExchangeReply) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExchangeReply) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *ExchangeReply) GetHolidays() []*ExchangeReply_Holiday {
	if x != nil {
		return x.Holidays
	}
	return nil
}

func (x *ExchangeReply) GetTradingHours() *TradingHours {
	if x != nil {
		return x.TradingHours
	}
	return nil
}

func (x *ExchangeReply) GetCompanies() []*ExchangeReply_Company {
	if x != nil {
		return x.Companies
	}
	return nil
}

type ExchangeIsOpenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timezone          string        `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone,omitempty"`
	TradingHours      *TradingHours `protobuf:"bytes,4,opt,name=tradingHours,proto3" json:"tradingHours,omitempty"`
	Open              bool          `protobuf:"varint,5,opt,name=open,proto3" json:"open,omitempty"`
	TimeBeforeClosing int32         `protobuf:"varint,6,opt,name=timeBeforeClosing,proto3" json:"timeBeforeClosing,omitempty"`
}

func (x *ExchangeIsOpenReply) Reset() {
	*x = ExchangeIsOpenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeIsOpenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeIsOpenReply) ProtoMessage() {}

func (x *ExchangeIsOpenReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeIsOpenReply.ProtoReflect.Descriptor instead.
func (*ExchangeIsOpenReply) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *ExchangeIsOpenReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExchangeIsOpenReply) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *ExchangeIsOpenReply) GetTradingHours() *TradingHours {
	if x != nil {
		return x.TradingHours
	}
	return nil
}

func (x *ExchangeIsOpenReply) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

func (x *ExchangeIsOpenReply) GetTimeBeforeClosing() int32 {
	if x != nil {
		return x.TimeBeforeClosing
	}
	return 0
}

type ExchangeShortReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Country  string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Code     *Code  `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ExchangeShortReply) Reset() {
	*x = ExchangeShortReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeShortReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeShortReply) ProtoMessage() {}

func (x *ExchangeShortReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeShortReply.ProtoReflect.Descriptor instead.
func (*ExchangeShortReply) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{5}
}

func (x *ExchangeShortReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExchangeShortReply) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ExchangeShortReply) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExchangeShortReply) GetCode() *Code {
	if x != nil {
		return x.Code
	}
	return nil
}

type ExchangeReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ExchangeShortReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ExchangeReplies) Reset() {
	*x = ExchangeReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeReplies) ProtoMessage() {}

func (x *ExchangeReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeReplies.ProtoReflect.Descriptor instead.
func (*ExchangeReplies) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{6}
}

func (x *ExchangeReplies) GetResults() []*ExchangeShortReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ExchangeReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ExchangeReply_Holiday struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Dates    []float32 `protobuf:"fixed32,2,rep,packed,name=dates,proto3" json:"dates,omitempty"`
	Official bool      `protobuf:"varint,3,opt,name=official,proto3" json:"official,omitempty"`
}

func (x *ExchangeReply_Holiday) Reset() {
	*x = ExchangeReply_Holiday{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeReply_Holiday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeReply_Holiday) ProtoMessage() {}

func (x *ExchangeReply_Holiday) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeReply_Holiday.ProtoReflect.Descriptor instead.
func (*ExchangeReply_Holiday) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ExchangeReply_Holiday) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExchangeReply_Holiday) GetDates() []float32 {
	if x != nil {
		return x.Dates
	}
	return nil
}

func (x *ExchangeReply_Holiday) GetOfficial() bool {
	if x != nil {
		return x.Official
	}
	return false
}

type ExchangeReply_Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Isin     string `protobuf:"bytes,3,opt,name=isin,proto3" json:"isin,omitempty"`
	Activity int64  `protobuf:"varint,4,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *ExchangeReply_Company) Reset() {
	*x = ExchangeReply_Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeReply_Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeReply_Company) ProtoMessage() {}

func (x *ExchangeReply_Company) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeReply_Company.ProtoReflect.Descriptor instead.
func (*ExchangeReply_Company) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ExchangeReply_Company) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *ExchangeReply_Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExchangeReply_Company) GetIsin() string {
	if x != nil {
		return x.Isin
	}
	return ""
}

func (x *ExchangeReply_Company) GetActivity() int64 {
	if x != nil {
		return x.Activity
	}
	return 0
}

type ExchangeReplies_Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Isin   string `protobuf:"bytes,3,opt,name=isin,proto3" json:"isin,omitempty"`
}

func (x *ExchangeReplies_Company) Reset() {
	*x = ExchangeReplies_Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_v1_exchange_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeReplies_Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeReplies_Company) ProtoMessage() {}

func (x *ExchangeReplies_Company) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_v1_exchange_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeReplies_Company.ProtoReflect.Descriptor instead.
func (*ExchangeReplies_Company) Descriptor() ([]byte, []int) {
	return file_proto_exchange_v1_exchange_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ExchangeReplies_Company) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *ExchangeReplies_Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExchangeReplies_Company) GetIsin() string {
	if x != nil {
		return x.Isin
	}
	return ""
}

var File_proto_exchange_v1_exchange_proto protoreflect.FileDescriptor

var file_proto_exchange_v1_exchange_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x64, 0xfa, 0x42, 0x61, 0x72,
	0x5f, 0x52, 0x06, 0x4e, 0x41, 0x53, 0x44, 0x41, 0x51, 0x52, 0x04, 0x4e, 0x59, 0x53, 0x45, 0x52,
	0x02, 0x54, 0x4f, 0x52, 0x03, 0x4c, 0x53, 0x45, 0x52, 0x02, 0x50, 0x41, 0x52, 0x02, 0x42, 0x52,
	0x52, 0x02, 0x41, 0x53, 0x52, 0x02, 0x53, 0x47, 0x52, 0x03, 0x53, 0x48, 0x45, 0x52, 0x03, 0x53,
	0x48, 0x47, 0x52, 0x02, 0x48, 0x4b, 0x52, 0x02, 0x4c, 0x53, 0x52, 0x02, 0x4d, 0x43, 0x52, 0x01,
	0x46, 0x52, 0x02, 0x4d, 0x49, 0x52, 0x02, 0x4c, 0x55, 0x52, 0x04, 0x43, 0x4f, 0x4d, 0x4d, 0x52,
	0x05, 0x46, 0x4f, 0x52, 0x45, 0x58, 0x52, 0x04, 0x49, 0x4e, 0x44, 0x58, 0x52, 0x02, 0x43, 0x43,
	0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x62, 0x6b, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x6b, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x76, 0x69, 0x65, 0x77, 0x22, 0x98, 0x02, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x79, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x95, 0x04, 0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x08, 0x68, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x65, 0x73, 0x1a, 0x4f, 0x0a, 0x07, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x1a, 0x65, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x69,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0xc6, 0x01,
	0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x73, 0x4f, 0x70, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43,
	0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xad,
	0x01, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x1a, 0x49, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x69, 0x6e, 0x32, 0xff,
	0x02, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x70, 0x0a, 0x06, 0x49,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x73, 0x4f, 0x70, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x7b, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x7d, 0x2f, 0x69, 0x73, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x60, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2f, 0x7b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x7d, 0x12,
	0x53, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1c, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x42, 0x4c, 0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_exchange_v1_exchange_proto_rawDescOnce sync.Once
	file_proto_exchange_v1_exchange_proto_rawDescData = file_proto_exchange_v1_exchange_proto_rawDesc
)

func file_proto_exchange_v1_exchange_proto_rawDescGZIP() []byte {
	file_proto_exchange_v1_exchange_proto_rawDescOnce.Do(func() {
		file_proto_exchange_v1_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_exchange_v1_exchange_proto_rawDescData)
	})
	return file_proto_exchange_v1_exchange_proto_rawDescData
}

var file_proto_exchange_v1_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_exchange_v1_exchange_proto_goTypes = []interface{}{
	(*ExchangeRequest)(nil),         // 0: exchange.v1.ExchangeRequest
	(*Code)(nil),                    // 1: exchange.v1.Code
	(*TradingHours)(nil),            // 2: exchange.v1.TradingHours
	(*ExchangeReply)(nil),           // 3: exchange.v1.ExchangeReply
	(*ExchangeIsOpenReply)(nil),     // 4: exchange.v1.ExchangeIsOpenReply
	(*ExchangeShortReply)(nil),      // 5: exchange.v1.ExchangeShortReply
	(*ExchangeReplies)(nil),         // 6: exchange.v1.ExchangeReplies
	nil,                             // 7: exchange.v1.TradingHours.OpenEntry
	nil,                             // 8: exchange.v1.TradingHours.CloseEntry
	(*ExchangeReply_Holiday)(nil),   // 9: exchange.v1.ExchangeReply.Holiday
	(*ExchangeReply_Company)(nil),   // 10: exchange.v1.ExchangeReply.Company
	(*ExchangeReplies_Company)(nil), // 11: exchange.v1.ExchangeReplies.Company
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_proto_exchange_v1_exchange_proto_depIdxs = []int32{
	7,  // 0: exchange.v1.TradingHours.open:type_name -> exchange.v1.TradingHours.OpenEntry
	8,  // 1: exchange.v1.TradingHours.close:type_name -> exchange.v1.TradingHours.CloseEntry
	1,  // 2: exchange.v1.ExchangeReply.code:type_name -> exchange.v1.Code
	9,  // 3: exchange.v1.ExchangeReply.holidays:type_name -> exchange.v1.ExchangeReply.Holiday
	2,  // 4: exchange.v1.ExchangeReply.tradingHours:type_name -> exchange.v1.TradingHours
	10, // 5: exchange.v1.ExchangeReply.companies:type_name -> exchange.v1.ExchangeReply.Company
	2,  // 6: exchange.v1.ExchangeIsOpenReply.tradingHours:type_name -> exchange.v1.TradingHours
	1,  // 7: exchange.v1.ExchangeShortReply.code:type_name -> exchange.v1.Code
	5,  // 8: exchange.v1.ExchangeReplies.results:type_name -> exchange.v1.ExchangeShortReply
	0,  // 9: exchange.v1.Exchange.IsOpen:input_type -> exchange.v1.ExchangeRequest
	0,  // 10: exchange.v1.Exchange.Get:input_type -> exchange.v1.ExchangeRequest
	12, // 11: exchange.v1.Exchange.List:input_type -> google.protobuf.Empty
	12, // 12: exchange.v1.Exchange.Health:input_type -> google.protobuf.Empty
	4,  // 13: exchange.v1.Exchange.IsOpen:output_type -> exchange.v1.ExchangeIsOpenReply
	3,  // 14: exchange.v1.Exchange.Get:output_type -> exchange.v1.ExchangeReply
	6,  // 15: exchange.v1.Exchange.List:output_type -> exchange.v1.ExchangeReplies
	12, // 16: exchange.v1.Exchange.Health:output_type -> google.protobuf.Empty
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_exchange_v1_exchange_proto_init() }
func file_proto_exchange_v1_exchange_proto_init() {
	if File_proto_exchange_v1_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_exchange_v1_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRequest); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradingHours); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeReply); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeIsOpenReply); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeShortReply); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeReplies); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeReply_Holiday); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeReply_Company); i {
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
		file_proto_exchange_v1_exchange_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeReplies_Company); i {
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
			RawDescriptor: file_proto_exchange_v1_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_exchange_v1_exchange_proto_goTypes,
		DependencyIndexes: file_proto_exchange_v1_exchange_proto_depIdxs,
		MessageInfos:      file_proto_exchange_v1_exchange_proto_msgTypes,
	}.Build()
	File_proto_exchange_v1_exchange_proto = out.File
	file_proto_exchange_v1_exchange_proto_rawDesc = nil
	file_proto_exchange_v1_exchange_proto_goTypes = nil
	file_proto_exchange_v1_exchange_proto_depIdxs = nil
}
