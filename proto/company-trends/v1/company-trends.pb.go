// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/company-trends/v1/company-trends.proto

package company_trends_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type TrendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Exchange string `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *TrendRequest) Reset() {
	*x = TrendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrendRequest) ProtoMessage() {}

func (x *TrendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrendRequest.ProtoReflect.Descriptor instead.
func (*TrendRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_trends_v1_company_trends_proto_rawDescGZIP(), []int{0}
}

func (x *TrendRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *TrendRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TrendRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type TrendIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string                `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Currency string                `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Exchange string                `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Result   *TrendIdReply_Trend   `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	Related  []*TrendIdReply_Trend `protobuf:"bytes,5,rep,name=related,proto3" json:"related,omitempty"`
}

func (x *TrendIdReply) Reset() {
	*x = TrendIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrendIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrendIdReply) ProtoMessage() {}

func (x *TrendIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrendIdReply.ProtoReflect.Descriptor instead.
func (*TrendIdReply) Descriptor() ([]byte, []int) {
	return file_proto_company_trends_v1_company_trends_proto_rawDescGZIP(), []int{1}
}

func (x *TrendIdReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *TrendIdReply) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TrendIdReply) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *TrendIdReply) GetResult() *TrendIdReply_Trend {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TrendIdReply) GetRelated() []*TrendIdReply_Trend {
	if x != nil {
		return x.Related
	}
	return nil
}

type TrendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref     string                `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Related []*TrendReply_Related `protobuf:"bytes,2,rep,name=related,proto3" json:"related,omitempty"`
	Geomap  []*TrendReply_GeoMap  `protobuf:"bytes,3,rep,name=geomap,proto3" json:"geomap,omitempty"`
}

func (x *TrendReply) Reset() {
	*x = TrendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrendReply) ProtoMessage() {}

func (x *TrendReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrendReply.ProtoReflect.Descriptor instead.
func (*TrendReply) Descriptor() ([]byte, []int) {
	return file_proto_company_trends_v1_company_trends_proto_rawDescGZIP(), []int{2}
}

func (x *TrendReply) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *TrendReply) GetRelated() []*TrendReply_Related {
	if x != nil {
		return x.Related
	}
	return nil
}

func (x *TrendReply) GetGeomap() []*TrendReply_GeoMap {
	if x != nil {
		return x.Geomap
	}
	return nil
}

type TrendIdReply_Trend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ref  string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *TrendIdReply_Trend) Reset() {
	*x = TrendIdReply_Trend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrendIdReply_Trend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrendIdReply_Trend) ProtoMessage() {}

func (x *TrendIdReply_Trend) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrendIdReply_Trend.ProtoReflect.Descriptor instead.
func (*TrendIdReply_Trend) Descriptor() ([]byte, []int) {
	return file_proto_company_trends_v1_company_trends_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TrendIdReply_Trend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrendIdReply_Trend) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

type TrendReply_Related struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Mid   string `protobuf:"bytes,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Link  string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Value int32  `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TrendReply_Related) Reset() {
	*x = TrendReply_Related{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrendReply_Related) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrendReply_Related) ProtoMessage() {}

func (x *TrendReply_Related) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrendReply_Related.ProtoReflect.Descriptor instead.
func (*TrendReply_Related) Descriptor() ([]byte, []int) {
	return file_proto_company_trends_v1_company_trends_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TrendReply_Related) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TrendReply_Related) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *TrendReply_Related) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *TrendReply_Related) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TrendReply_GeoMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TrendReply_GeoMap) Reset() {
	*x = TrendReply_GeoMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrendReply_GeoMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrendReply_GeoMap) ProtoMessage() {}

func (x *TrendReply_GeoMap) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_trends_v1_company_trends_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrendReply_GeoMap.ProtoReflect.Descriptor instead.
func (*TrendReply_GeoMap) Descriptor() ([]byte, []int) {
	return file_proto_company_trends_v1_company_trends_proto_rawDescGZIP(), []int{2, 1}
}

func (x *TrendReply_GeoMap) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TrendReply_GeoMap) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_company_trends_v1_company_trends_proto protoreflect.FileDescriptor

var file_proto_company_trends_v1_company_trends_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d,
	0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2d, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0c, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x08,
	0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x98, 0x01, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25,
	0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x08, 0x52, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x8d, 0x02, 0x0a, 0x0c, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64,
	0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64,
	0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x2d, 0x0a, 0x05, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0xae, 0x02, 0x0a, 0x0a, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x52, 0x07,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x67, 0x65, 0x6f, 0x6d, 0x61,
	0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x65, 0x6f, 0x4d, 0x61, 0x70, 0x52, 0x06, 0x67,
	0x65, 0x6f, 0x6d, 0x61, 0x70, 0x1a, 0x5b, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x32, 0x0a, 0x06, 0x47, 0x65, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb2, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x67, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x3a, 0x01,
	0x2a, 0x12, 0x67, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2f, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x92, 0x41, 0x02, 0x62, 0x00, 0x42, 0xb9, 0x03, 0x0a, 0x20,
	0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x17, 0x41, 0x50, 0x49, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x72, 0x65, 0x6e,
	0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x74, 0x72, 0x65, 0x6e, 0x64,
	0x73, 0x5f, 0x76, 0x31, 0x92, 0x41, 0xaf, 0x02, 0x12, 0x9b, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x20, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73,
	0x12, 0x21, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x20, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x72, 0x65,
	0x6e, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65,
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
	file_proto_company_trends_v1_company_trends_proto_rawDescOnce sync.Once
	file_proto_company_trends_v1_company_trends_proto_rawDescData = file_proto_company_trends_v1_company_trends_proto_rawDesc
)

func file_proto_company_trends_v1_company_trends_proto_rawDescGZIP() []byte {
	file_proto_company_trends_v1_company_trends_proto_rawDescOnce.Do(func() {
		file_proto_company_trends_v1_company_trends_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_company_trends_v1_company_trends_proto_rawDescData)
	})
	return file_proto_company_trends_v1_company_trends_proto_rawDescData
}

var file_proto_company_trends_v1_company_trends_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_company_trends_v1_company_trends_proto_goTypes = []interface{}{
	(*TrendRequest)(nil),       // 0: company.trends.v1.TrendRequest
	(*TrendIdReply)(nil),       // 1: company.trends.v1.TrendIdReply
	(*TrendReply)(nil),         // 2: company.trends.v1.TrendReply
	(*TrendIdReply_Trend)(nil), // 3: company.trends.v1.TrendIdReply.Trend
	(*TrendReply_Related)(nil), // 4: company.trends.v1.TrendReply.Related
	(*TrendReply_GeoMap)(nil),  // 5: company.trends.v1.TrendReply.GeoMap
	(*emptypb.Empty)(nil),      // 6: google.protobuf.Empty
}
var file_proto_company_trends_v1_company_trends_proto_depIdxs = []int32{
	3, // 0: company.trends.v1.TrendIdReply.result:type_name -> company.trends.v1.TrendIdReply.Trend
	3, // 1: company.trends.v1.TrendIdReply.related:type_name -> company.trends.v1.TrendIdReply.Trend
	4, // 2: company.trends.v1.TrendReply.related:type_name -> company.trends.v1.TrendReply.Related
	5, // 3: company.trends.v1.TrendReply.geomap:type_name -> company.trends.v1.TrendReply.GeoMap
	0, // 4: company.trends.v1.CompanyTrends.Get:input_type -> company.trends.v1.TrendRequest
	0, // 5: company.trends.v1.CompanyTrends.Search:input_type -> company.trends.v1.TrendRequest
	6, // 6: company.trends.v1.CompanyTrends.Health:input_type -> google.protobuf.Empty
	1, // 7: company.trends.v1.CompanyTrends.Get:output_type -> company.trends.v1.TrendIdReply
	2, // 8: company.trends.v1.CompanyTrends.Search:output_type -> company.trends.v1.TrendReply
	6, // 9: company.trends.v1.CompanyTrends.Health:output_type -> google.protobuf.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_company_trends_v1_company_trends_proto_init() }
func file_proto_company_trends_v1_company_trends_proto_init() {
	if File_proto_company_trends_v1_company_trends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_company_trends_v1_company_trends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrendRequest); i {
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
		file_proto_company_trends_v1_company_trends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrendIdReply); i {
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
		file_proto_company_trends_v1_company_trends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrendReply); i {
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
		file_proto_company_trends_v1_company_trends_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrendIdReply_Trend); i {
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
		file_proto_company_trends_v1_company_trends_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrendReply_Related); i {
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
		file_proto_company_trends_v1_company_trends_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrendReply_GeoMap); i {
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
			RawDescriptor: file_proto_company_trends_v1_company_trends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_company_trends_v1_company_trends_proto_goTypes,
		DependencyIndexes: file_proto_company_trends_v1_company_trends_proto_depIdxs,
		MessageInfos:      file_proto_company_trends_v1_company_trends_proto_msgTypes,
	}.Build()
	File_proto_company_trends_v1_company_trends_proto = out.File
	file_proto_company_trends_v1_company_trends_proto_rawDesc = nil
	file_proto_company_trends_v1_company_trends_proto_goTypes = nil
	file_proto_company_trends_v1_company_trends_proto_depIdxs = nil
}
