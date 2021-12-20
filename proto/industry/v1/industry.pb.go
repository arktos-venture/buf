// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/industry/v1/industry.proto

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

type IndustryReply_Ref_Classication int32

const (
	IndustryReply_Ref_EconomicSector IndustryReply_Ref_Classication = 0
	IndustryReply_Ref_BusinessSector IndustryReply_Ref_Classication = 1
	IndustryReply_Ref_IndustryGroup  IndustryReply_Ref_Classication = 2
	IndustryReply_Ref_Industry       IndustryReply_Ref_Classication = 3
	IndustryReply_Ref_Activity       IndustryReply_Ref_Classication = 4
)

// Enum value maps for IndustryReply_Ref_Classication.
var (
	IndustryReply_Ref_Classication_name = map[int32]string{
		0: "EconomicSector",
		1: "BusinessSector",
		2: "IndustryGroup",
		3: "Industry",
		4: "Activity",
	}
	IndustryReply_Ref_Classication_value = map[string]int32{
		"EconomicSector": 0,
		"BusinessSector": 1,
		"IndustryGroup":  2,
		"Industry":       3,
		"Activity":       4,
	}
)

func (x IndustryReply_Ref_Classication) Enum() *IndustryReply_Ref_Classication {
	p := new(IndustryReply_Ref_Classication)
	*p = x
	return p
}

func (x IndustryReply_Ref_Classication) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndustryReply_Ref_Classication) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_industry_v1_industry_proto_enumTypes[0].Descriptor()
}

func (IndustryReply_Ref_Classication) Type() protoreflect.EnumType {
	return &file_proto_industry_v1_industry_proto_enumTypes[0]
}

func (x IndustryReply_Ref_Classication) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndustryReply_Ref_Classication.Descriptor instead.
func (IndustryReply_Ref_Classication) EnumDescriptor() ([]byte, []int) {
	return file_proto_industry_v1_industry_proto_rawDescGZIP(), []int{1, 0, 0}
}

type IndustryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref int64 `protobuf:"varint,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *IndustryRequest) Reset() {
	*x = IndustryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_industry_v1_industry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndustryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndustryRequest) ProtoMessage() {}

func (x *IndustryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_industry_v1_industry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndustryRequest.ProtoReflect.Descriptor instead.
func (*IndustryRequest) Descriptor() ([]byte, []int) {
	return file_proto_industry_v1_industry_proto_rawDescGZIP(), []int{0}
}

func (x *IndustryRequest) GetRef() int64 {
	if x != nil {
		return x.Ref
	}
	return 0
}

type IndustryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[int64]*IndustryReply_Ref `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IndustryReply) Reset() {
	*x = IndustryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_industry_v1_industry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndustryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndustryReply) ProtoMessage() {}

func (x *IndustryReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_industry_v1_industry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndustryReply.ProtoReflect.Descriptor instead.
func (*IndustryReply) Descriptor() ([]byte, []int) {
	return file_proto_industry_v1_industry_proto_rawDescGZIP(), []int{1}
}

func (x *IndustryReply) GetResults() map[int64]*IndustryReply_Ref {
	if x != nil {
		return x.Results
	}
	return nil
}

type IndustryReply_Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classication IndustryReply_Ref_Classication `protobuf:"varint,1,opt,name=classication,proto3,enum=industry.v1.IndustryReply_Ref_Classication" json:"classication,omitempty"`
	Name         string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Next         map[int64]*IndustryReply_Ref   `protobuf:"bytes,3,rep,name=next,proto3" json:"next,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IndustryReply_Ref) Reset() {
	*x = IndustryReply_Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_industry_v1_industry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndustryReply_Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndustryReply_Ref) ProtoMessage() {}

func (x *IndustryReply_Ref) ProtoReflect() protoreflect.Message {
	mi := &file_proto_industry_v1_industry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndustryReply_Ref.ProtoReflect.Descriptor instead.
func (*IndustryReply_Ref) Descriptor() ([]byte, []int) {
	return file_proto_industry_v1_industry_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IndustryReply_Ref) GetClassication() IndustryReply_Ref_Classication {
	if x != nil {
		return x.Classication
	}
	return IndustryReply_Ref_EconomicSector
}

func (x *IndustryReply_Ref) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndustryReply_Ref) GetNext() map[int64]*IndustryReply_Ref {
	if x != nil {
		return x.Next
	}
	return nil
}

var File_proto_industry_v1_industry_proto protoreflect.FileDescriptor

var file_proto_industry_v1_industry_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0f, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x22, 0x0c, 0x18, 0xd2, 0x9a, 0xfe, 0xc0, 0x17,
	0x28, 0x92, 0xa6, 0x80, 0xd5, 0x12, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x99, 0x04, 0x0a, 0x0d,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x1a, 0xe8, 0x02, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x12, 0x4f, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x66, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x66, 0x2e, 0x4e, 0x65, 0x78, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x1a, 0x57, 0x0a, 0x09, 0x4e,
	0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x65, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x10, 0x04, 0x1a, 0x5a, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x87, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x79, 0x12, 0x5b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x69, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x7b, 0x72, 0x65, 0x66,
	0x7d, 0x12, 0x52, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x7a, 0x42, 0x4c, 0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42,
	0x0f, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31,
	0x50, 0x01, 0x5a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_industry_v1_industry_proto_rawDescOnce sync.Once
	file_proto_industry_v1_industry_proto_rawDescData = file_proto_industry_v1_industry_proto_rawDesc
)

func file_proto_industry_v1_industry_proto_rawDescGZIP() []byte {
	file_proto_industry_v1_industry_proto_rawDescOnce.Do(func() {
		file_proto_industry_v1_industry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_industry_v1_industry_proto_rawDescData)
	})
	return file_proto_industry_v1_industry_proto_rawDescData
}

var file_proto_industry_v1_industry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_industry_v1_industry_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_industry_v1_industry_proto_goTypes = []interface{}{
	(IndustryReply_Ref_Classication)(0), // 0: industry.v1.IndustryReply.Ref.Classication
	(*IndustryRequest)(nil),             // 1: industry.v1.IndustryRequest
	(*IndustryReply)(nil),               // 2: industry.v1.IndustryReply
	(*IndustryReply_Ref)(nil),           // 3: industry.v1.IndustryReply.Ref
	nil,                                 // 4: industry.v1.IndustryReply.ResultsEntry
	nil,                                 // 5: industry.v1.IndustryReply.Ref.NextEntry
	(*emptypb.Empty)(nil),               // 6: google.protobuf.Empty
}
var file_proto_industry_v1_industry_proto_depIdxs = []int32{
	4, // 0: industry.v1.IndustryReply.results:type_name -> industry.v1.IndustryReply.ResultsEntry
	0, // 1: industry.v1.IndustryReply.Ref.classication:type_name -> industry.v1.IndustryReply.Ref.Classication
	5, // 2: industry.v1.IndustryReply.Ref.next:type_name -> industry.v1.IndustryReply.Ref.NextEntry
	3, // 3: industry.v1.IndustryReply.ResultsEntry.value:type_name -> industry.v1.IndustryReply.Ref
	3, // 4: industry.v1.IndustryReply.Ref.NextEntry.value:type_name -> industry.v1.IndustryReply.Ref
	1, // 5: industry.v1.Industry.Get:input_type -> industry.v1.IndustryRequest
	6, // 6: industry.v1.Industry.List:input_type -> google.protobuf.Empty
	6, // 7: industry.v1.Industry.Health:input_type -> google.protobuf.Empty
	2, // 8: industry.v1.Industry.Get:output_type -> industry.v1.IndustryReply
	2, // 9: industry.v1.Industry.List:output_type -> industry.v1.IndustryReply
	6, // 10: industry.v1.Industry.Health:output_type -> google.protobuf.Empty
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_industry_v1_industry_proto_init() }
func file_proto_industry_v1_industry_proto_init() {
	if File_proto_industry_v1_industry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_industry_v1_industry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndustryRequest); i {
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
		file_proto_industry_v1_industry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndustryReply); i {
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
		file_proto_industry_v1_industry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndustryReply_Ref); i {
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
			RawDescriptor: file_proto_industry_v1_industry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_industry_v1_industry_proto_goTypes,
		DependencyIndexes: file_proto_industry_v1_industry_proto_depIdxs,
		EnumInfos:         file_proto_industry_v1_industry_proto_enumTypes,
		MessageInfos:      file_proto_industry_v1_industry_proto_msgTypes,
	}.Build()
	File_proto_industry_v1_industry_proto = out.File
	file_proto_industry_v1_industry_proto_rawDesc = nil
	file_proto_industry_v1_industry_proto_goTypes = nil
	file_proto_industry_v1_industry_proto_depIdxs = nil
}
