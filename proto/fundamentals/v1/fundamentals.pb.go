// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.4
// source: proto/fundamentals/v1/fundamentals.proto

package v1Fundamentals

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FundamentalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *FundamentalRequest) Reset() {
	*x = FundamentalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fundamentals_v1_fundamentals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundamentalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundamentalRequest) ProtoMessage() {}

func (x *FundamentalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fundamentals_v1_fundamentals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundamentalRequest.ProtoReflect.Descriptor instead.
func (*FundamentalRequest) Descriptor() ([]byte, []int) {
	return file_proto_fundamentals_v1_fundamentals_proto_rawDescGZIP(), []int{0}
}

func (x *FundamentalRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *FundamentalRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type FundamentalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string                          `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Exchange string                          `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Results  []*FundamentalReply_Fundamental `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *FundamentalReply) Reset() {
	*x = FundamentalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fundamentals_v1_fundamentals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundamentalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundamentalReply) ProtoMessage() {}

func (x *FundamentalReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fundamentals_v1_fundamentals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundamentalReply.ProtoReflect.Descriptor instead.
func (*FundamentalReply) Descriptor() ([]byte, []int) {
	return file_proto_fundamentals_v1_fundamentals_proto_rawDescGZIP(), []int{1}
}

func (x *FundamentalReply) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *FundamentalReply) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *FundamentalReply) GetResults() []*FundamentalReply_Fundamental {
	if x != nil {
		return x.Results
	}
	return nil
}

type FundamentalReply_Fundamental struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *FundamentalReply_Fundamental) Reset() {
	*x = FundamentalReply_Fundamental{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fundamentals_v1_fundamentals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundamentalReply_Fundamental) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundamentalReply_Fundamental) ProtoMessage() {}

func (x *FundamentalReply_Fundamental) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fundamentals_v1_fundamentals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundamentalReply_Fundamental.ProtoReflect.Descriptor instead.
func (*FundamentalReply_Fundamental) Descriptor() ([]byte, []int) {
	return file_proto_fundamentals_v1_fundamentals_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FundamentalReply_Fundamental) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

var File_proto_fundamentals_v1_fundamentals_proto protoreflect.FileDescriptor

var file_proto_fundamentals_v1_fundamentals_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x75, 0x6e, 0x64,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x12, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x01, 0x18, 0x08, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x08, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x66,
	0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x1a, 0x1d, 0x0a, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6f, 0x6b, 0x32, 0x7f, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x73, 0x12, 0x6f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x6e,
	0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x73, 0x42, 0x9f, 0x03, 0x0a, 0x1e, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x41, 0x50, 0x49, 0x46, 0x75, 0x6e, 0x64,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72,
	0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x73, 0x92, 0x41, 0x9d, 0x02, 0x12, 0x89, 0x01, 0x0a, 0x11, 0x46, 0x75,
	0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x12,
	0x11, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65, 0x6e,
	0x74, 0x75, 0x72, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f, 0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a,
	0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79,
	0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20,
	0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fundamentals_v1_fundamentals_proto_rawDescOnce sync.Once
	file_proto_fundamentals_v1_fundamentals_proto_rawDescData = file_proto_fundamentals_v1_fundamentals_proto_rawDesc
)

func file_proto_fundamentals_v1_fundamentals_proto_rawDescGZIP() []byte {
	file_proto_fundamentals_v1_fundamentals_proto_rawDescOnce.Do(func() {
		file_proto_fundamentals_v1_fundamentals_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fundamentals_v1_fundamentals_proto_rawDescData)
	})
	return file_proto_fundamentals_v1_fundamentals_proto_rawDescData
}

var file_proto_fundamentals_v1_fundamentals_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_fundamentals_v1_fundamentals_proto_goTypes = []interface{}{
	(*FundamentalRequest)(nil),           // 0: fundamentals.v1.FundamentalRequest
	(*FundamentalReply)(nil),             // 1: fundamentals.v1.FundamentalReply
	(*FundamentalReply_Fundamental)(nil), // 2: fundamentals.v1.FundamentalReply.Fundamental
}
var file_proto_fundamentals_v1_fundamentals_proto_depIdxs = []int32{
	2, // 0: fundamentals.v1.FundamentalReply.results:type_name -> fundamentals.v1.FundamentalReply.Fundamental
	0, // 1: fundamentals.v1.Fundamentals.Get:input_type -> fundamentals.v1.FundamentalRequest
	1, // 2: fundamentals.v1.Fundamentals.Get:output_type -> fundamentals.v1.FundamentalReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_fundamentals_v1_fundamentals_proto_init() }
func file_proto_fundamentals_v1_fundamentals_proto_init() {
	if File_proto_fundamentals_v1_fundamentals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fundamentals_v1_fundamentals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundamentalRequest); i {
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
		file_proto_fundamentals_v1_fundamentals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundamentalReply); i {
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
		file_proto_fundamentals_v1_fundamentals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundamentalReply_Fundamental); i {
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
			RawDescriptor: file_proto_fundamentals_v1_fundamentals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fundamentals_v1_fundamentals_proto_goTypes,
		DependencyIndexes: file_proto_fundamentals_v1_fundamentals_proto_depIdxs,
		MessageInfos:      file_proto_fundamentals_v1_fundamentals_proto_msgTypes,
	}.Build()
	File_proto_fundamentals_v1_fundamentals_proto = out.File
	file_proto_fundamentals_v1_fundamentals_proto_rawDesc = nil
	file_proto_fundamentals_v1_fundamentals_proto_goTypes = nil
	file_proto_fundamentals_v1_fundamentals_proto_depIdxs = nil
}
