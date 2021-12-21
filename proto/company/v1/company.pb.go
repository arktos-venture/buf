// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/company/v1/company.proto

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

type CompanySearchByIDsRequest_Broker int32

const (
	CompanySearchByIDsRequest_IBKR CompanySearchByIDsRequest_Broker = 0
)

// Enum value maps for CompanySearchByIDsRequest_Broker.
var (
	CompanySearchByIDsRequest_Broker_name = map[int32]string{
		0: "IBKR",
	}
	CompanySearchByIDsRequest_Broker_value = map[string]int32{
		"IBKR": 0,
	}
)

func (x CompanySearchByIDsRequest_Broker) Enum() *CompanySearchByIDsRequest_Broker {
	p := new(CompanySearchByIDsRequest_Broker)
	*p = x
	return p
}

func (x CompanySearchByIDsRequest_Broker) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompanySearchByIDsRequest_Broker) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_company_v1_company_proto_enumTypes[0].Descriptor()
}

func (CompanySearchByIDsRequest_Broker) Type() protoreflect.EnumType {
	return &file_proto_company_v1_company_proto_enumTypes[0]
}

func (x CompanySearchByIDsRequest_Broker) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompanySearchByIDsRequest_Broker.Descriptor instead.
func (CompanySearchByIDsRequest_Broker) EnumDescriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{1, 0}
}

type CompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Exchange string `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *CompanyRequest) Reset() {
	*x = CompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyRequest) ProtoMessage() {}

func (x *CompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyRequest.ProtoReflect.Descriptor instead.
func (*CompanyRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{0}
}

func (x *CompanyRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CompanyRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CompanyRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type CompanySearchByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Broker   CompanySearchByIDsRequest_Broker `protobuf:"varint,1,opt,name=broker,proto3,enum=company.v1.CompanySearchByIDsRequest_Broker" json:"broker,omitempty"`
	Currency string                           `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Ids      []int64                          `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CompanySearchByIDsRequest) Reset() {
	*x = CompanySearchByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanySearchByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanySearchByIDsRequest) ProtoMessage() {}

func (x *CompanySearchByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanySearchByIDsRequest.ProtoReflect.Descriptor instead.
func (*CompanySearchByIDsRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{1}
}

func (x *CompanySearchByIDsRequest) GetBroker() CompanySearchByIDsRequest_Broker {
	if x != nil {
		return x.Broker
	}
	return CompanySearchByIDsRequest_IBKR
}

func (x *CompanySearchByIDsRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CompanySearchByIDsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CompanySearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency     string            `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	FilterString map[string]string `protobuf:"bytes,2,rep,name=filterString,proto3" json:"filterString,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FilterInt    map[string]int64  `protobuf:"bytes,3,rep,name=filterInt,proto3" json:"filterInt,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Limit        int32             `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CompanySearchRequest) Reset() {
	*x = CompanySearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanySearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanySearchRequest) ProtoMessage() {}

func (x *CompanySearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanySearchRequest.ProtoReflect.Descriptor instead.
func (*CompanySearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{2}
}

func (x *CompanySearchRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CompanySearchRequest) GetFilterString() map[string]string {
	if x != nil {
		return x.FilterString
	}
	return nil
}

func (x *CompanySearchRequest) GetFilterInt() map[string]int64 {
	if x != nil {
		return x.FilterInt
	}
	return nil
}

func (x *CompanySearchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type CompanyAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *CompanyAddress) Reset() {
	*x = CompanyAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyAddress) ProtoMessage() {}

func (x *CompanyAddress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyAddress.ProtoReflect.Descriptor instead.
func (*CompanyAddress) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{3}
}

func (x *CompanyAddress) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CompanyAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CompanyAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CompanyAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type CompanyContact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Website string `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CompanyContact) Reset() {
	*x = CompanyContact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyContact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyContact) ProtoMessage() {}

func (x *CompanyContact) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyContact.ProtoReflect.Descriptor instead.
func (*CompanyContact) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{4}
}

func (x *CompanyContact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CompanyContact) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *CompanyContact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker         string                          `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name           string                          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Exchange       string                          `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Isin           string                          `protobuf:"bytes,5,opt,name=isin,proto3" json:"isin,omitempty"`
	Activity       int64                           `protobuf:"varint,6,opt,name=activity,proto3" json:"activity,omitempty"`
	CurrencyReport string                          `protobuf:"bytes,7,opt,name=currencyReport,proto3" json:"currencyReport,omitempty"`
	FiscalYearEnd  int64                           `protobuf:"varint,8,opt,name=fiscalYearEnd,proto3" json:"fiscalYearEnd,omitempty"`
	Contact        *CompanyContact                 `protobuf:"bytes,9,opt,name=contact,proto3" json:"contact,omitempty"`
	Address        *CompanyAddress                 `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Brokers        map[string]*CompanyReply_Broker `protobuf:"bytes,11,rep,name=brokers,proto3" json:"brokers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Active         bool                            `protobuf:"varint,12,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *CompanyReply) Reset() {
	*x = CompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyReply) ProtoMessage() {}

func (x *CompanyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyReply.ProtoReflect.Descriptor instead.
func (*CompanyReply) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{5}
}

func (x *CompanyReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CompanyReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanyReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CompanyReply) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *CompanyReply) GetIsin() string {
	if x != nil {
		return x.Isin
	}
	return ""
}

func (x *CompanyReply) GetActivity() int64 {
	if x != nil {
		return x.Activity
	}
	return 0
}

func (x *CompanyReply) GetCurrencyReport() string {
	if x != nil {
		return x.CurrencyReport
	}
	return ""
}

func (x *CompanyReply) GetFiscalYearEnd() int64 {
	if x != nil {
		return x.FiscalYearEnd
	}
	return 0
}

func (x *CompanyReply) GetContact() *CompanyContact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *CompanyReply) GetAddress() *CompanyAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CompanyReply) GetBrokers() map[string]*CompanyReply_Broker {
	if x != nil {
		return x.Brokers
	}
	return nil
}

func (x *CompanyReply) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type CompanyReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*CompanyReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CompanyReplies) Reset() {
	*x = CompanyReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyReplies) ProtoMessage() {}

func (x *CompanyReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyReplies.ProtoReflect.Descriptor instead.
func (*CompanyReplies) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{6}
}

func (x *CompanyReplies) GetResults() []*CompanyReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CompanyReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CompanyReply_Broker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ticker  string `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Routing string `protobuf:"bytes,3,opt,name=routing,proto3" json:"routing,omitempty"`
}

func (x *CompanyReply_Broker) Reset() {
	*x = CompanyReply_Broker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_v1_company_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyReply_Broker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyReply_Broker) ProtoMessage() {}

func (x *CompanyReply_Broker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_v1_company_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyReply_Broker.ProtoReflect.Descriptor instead.
func (*CompanyReply_Broker) Descriptor() ([]byte, []int) {
	return file_proto_company_v1_company_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CompanyReply_Broker) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CompanyReply_Broker) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CompanyReply_Broker) GetRouting() string {
	if x != nil {
		return x.Routing
	}
	return ""
}

var File_proto_company_v1_company_proto protoreflect.FileDescriptor

var file_proto_company_v1_company_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x80, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x08, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98,
	0x01, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x08,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x08, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4e, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x03, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x03, 0x42, 0x10, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa, 0x42,
	0x05, 0x92, 0x01, 0x02, 0x10, 0x14, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x12, 0x0a, 0x06, 0x42,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x42, 0x4b, 0x52, 0x10, 0x00, 0x22,
	0x84, 0x03, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x98, 0x01, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x56,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x4d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x1a, 0x05, 0x18, 0x96, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x3f, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6c, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0x56, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe4, 0x04, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x69, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x59, 0x65,
	0x61, 0x72, 0x45, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x3f, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x1a, 0x4a, 0x0a, 0x06, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x5b, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x5a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32,
	0xa6, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x6f, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x12, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x7b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x7d, 0x2f, 0x7b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x7d, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x7d, 0x12, 0x5e, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x4a, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x49, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_company_v1_company_proto_rawDescOnce sync.Once
	file_proto_company_v1_company_proto_rawDescData = file_proto_company_v1_company_proto_rawDesc
)

func file_proto_company_v1_company_proto_rawDescGZIP() []byte {
	file_proto_company_v1_company_proto_rawDescOnce.Do(func() {
		file_proto_company_v1_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_company_v1_company_proto_rawDescData)
	})
	return file_proto_company_v1_company_proto_rawDescData
}

var file_proto_company_v1_company_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_company_v1_company_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_company_v1_company_proto_goTypes = []interface{}{
	(CompanySearchByIDsRequest_Broker)(0), // 0: company.v1.CompanySearchByIDsRequest.Broker
	(*CompanyRequest)(nil),                // 1: company.v1.CompanyRequest
	(*CompanySearchByIDsRequest)(nil),     // 2: company.v1.CompanySearchByIDsRequest
	(*CompanySearchRequest)(nil),          // 3: company.v1.CompanySearchRequest
	(*CompanyAddress)(nil),                // 4: company.v1.CompanyAddress
	(*CompanyContact)(nil),                // 5: company.v1.CompanyContact
	(*CompanyReply)(nil),                  // 6: company.v1.CompanyReply
	(*CompanyReplies)(nil),                // 7: company.v1.CompanyReplies
	nil,                                   // 8: company.v1.CompanySearchRequest.FilterStringEntry
	nil,                                   // 9: company.v1.CompanySearchRequest.FilterIntEntry
	(*CompanyReply_Broker)(nil),           // 10: company.v1.CompanyReply.Broker
	nil,                                   // 11: company.v1.CompanyReply.BrokersEntry
	(*emptypb.Empty)(nil),                 // 12: google.protobuf.Empty
}
var file_proto_company_v1_company_proto_depIdxs = []int32{
	0,  // 0: company.v1.CompanySearchByIDsRequest.broker:type_name -> company.v1.CompanySearchByIDsRequest.Broker
	8,  // 1: company.v1.CompanySearchRequest.filterString:type_name -> company.v1.CompanySearchRequest.FilterStringEntry
	9,  // 2: company.v1.CompanySearchRequest.filterInt:type_name -> company.v1.CompanySearchRequest.FilterIntEntry
	5,  // 3: company.v1.CompanyReply.contact:type_name -> company.v1.CompanyContact
	4,  // 4: company.v1.CompanyReply.address:type_name -> company.v1.CompanyAddress
	11, // 5: company.v1.CompanyReply.brokers:type_name -> company.v1.CompanyReply.BrokersEntry
	6,  // 6: company.v1.CompanyReplies.results:type_name -> company.v1.CompanyReply
	10, // 7: company.v1.CompanyReply.BrokersEntry.value:type_name -> company.v1.CompanyReply.Broker
	1,  // 8: company.v1.Company.Get:input_type -> company.v1.CompanyRequest
	3,  // 9: company.v1.Company.Search:input_type -> company.v1.CompanySearchRequest
	12, // 10: company.v1.Company.Health:input_type -> google.protobuf.Empty
	6,  // 11: company.v1.Company.Get:output_type -> company.v1.CompanyReply
	7,  // 12: company.v1.Company.Search:output_type -> company.v1.CompanyReplies
	12, // 13: company.v1.Company.Health:output_type -> google.protobuf.Empty
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_company_v1_company_proto_init() }
func file_proto_company_v1_company_proto_init() {
	if File_proto_company_v1_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_company_v1_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyRequest); i {
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
		file_proto_company_v1_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanySearchByIDsRequest); i {
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
		file_proto_company_v1_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanySearchRequest); i {
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
		file_proto_company_v1_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyAddress); i {
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
		file_proto_company_v1_company_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyContact); i {
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
		file_proto_company_v1_company_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyReply); i {
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
		file_proto_company_v1_company_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyReplies); i {
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
		file_proto_company_v1_company_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyReply_Broker); i {
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
			RawDescriptor: file_proto_company_v1_company_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_company_v1_company_proto_goTypes,
		DependencyIndexes: file_proto_company_v1_company_proto_depIdxs,
		EnumInfos:         file_proto_company_v1_company_proto_enumTypes,
		MessageInfos:      file_proto_company_v1_company_proto_msgTypes,
	}.Build()
	File_proto_company_v1_company_proto = out.File
	file_proto_company_v1_company_proto_rawDesc = nil
	file_proto_company_v1_company_proto_goTypes = nil
	file_proto_company_v1_company_proto_depIdxs = nil
}
