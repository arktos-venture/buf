// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/news/v1/news.proto

package news_v1

import (
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

type NewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Period   string `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *NewsRequest) Reset() {
	*x = NewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_v1_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsRequest) ProtoMessage() {}

func (x *NewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_v1_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsRequest.ProtoReflect.Descriptor instead.
func (*NewsRequest) Descriptor() ([]byte, []int) {
	return file_proto_news_v1_news_proto_rawDescGZIP(), []int{0}
}

func (x *NewsRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *NewsRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *NewsRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

type NewsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Link      string                 `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Tags      []string               `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *NewsReply) Reset() {
	*x = NewsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_v1_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsReply) ProtoMessage() {}

func (x *NewsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_v1_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsReply.ProtoReflect.Descriptor instead.
func (*NewsReply) Descriptor() ([]byte, []int) {
	return file_proto_news_v1_news_proto_rawDescGZIP(), []int{1}
}

func (x *NewsReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsReply) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *NewsReply) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *NewsReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type NewsReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*NewsReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *NewsReplies) Reset() {
	*x = NewsReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_v1_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsReplies) ProtoMessage() {}

func (x *NewsReplies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_v1_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsReplies.ProtoReflect.Descriptor instead.
func (*NewsReplies) Descriptor() ([]byte, []int) {
	return file_proto_news_v1_news_proto_rawDescGZIP(), []int{2}
}

func (x *NewsReplies) GetResults() []*NewsReply {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *NewsReplies) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_news_v1_news_proto protoreflect.FileDescriptor

var file_proto_news_v1_news_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x01, 0x18, 0x08, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x08, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x13, 0xfa, 0x42, 0x10, 0x72, 0x0e, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x52,
	0x02, 0x33, 0x64, 0x52, 0x02, 0x31, 0x77, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22,
	0x84, 0x01, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x51, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xa2, 0x01, 0x0a, 0x04, 0x4e, 0x65,
	0x77, 0x73, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x92, 0x41, 0x02, 0x62, 0x00, 0x42, 0xef,
	0x02, 0x0a, 0x16, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x50, 0x49, 0x4e, 0x65,
	0x77, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x76, 0x31,
	0x92, 0x41, 0x8c, 0x02, 0x12, 0x79, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x73, 0x12, 0x09, 0x4e, 0x65, 0x77, 0x73, 0x20, 0x41, 0x50, 0x49, 0x73, 0x22, 0x5c, 0x0a, 0x16,
	0x41, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x20, 0x56, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x20, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f,
	0x73, 0x2d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x6f,
	0x73, 0x73, 0x40, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_news_v1_news_proto_rawDescOnce sync.Once
	file_proto_news_v1_news_proto_rawDescData = file_proto_news_v1_news_proto_rawDesc
)

func file_proto_news_v1_news_proto_rawDescGZIP() []byte {
	file_proto_news_v1_news_proto_rawDescOnce.Do(func() {
		file_proto_news_v1_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_news_v1_news_proto_rawDescData)
	})
	return file_proto_news_v1_news_proto_rawDescData
}

var file_proto_news_v1_news_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_news_v1_news_proto_goTypes = []interface{}{
	(*NewsRequest)(nil),           // 0: news.v1.NewsRequest
	(*NewsReply)(nil),             // 1: news.v1.NewsReply
	(*NewsReplies)(nil),           // 2: news.v1.NewsReplies
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_proto_news_v1_news_proto_depIdxs = []int32{
	3, // 0: news.v1.NewsReply.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: news.v1.NewsReplies.results:type_name -> news.v1.NewsReply
	0, // 2: news.v1.News.Search:input_type -> news.v1.NewsRequest
	4, // 3: news.v1.News.Health:input_type -> google.protobuf.Empty
	2, // 4: news.v1.News.Search:output_type -> news.v1.NewsReplies
	4, // 5: news.v1.News.Health:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_news_v1_news_proto_init() }
func file_proto_news_v1_news_proto_init() {
	if File_proto_news_v1_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_news_v1_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsRequest); i {
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
		file_proto_news_v1_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsReply); i {
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
		file_proto_news_v1_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsReplies); i {
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
			RawDescriptor: file_proto_news_v1_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_news_v1_news_proto_goTypes,
		DependencyIndexes: file_proto_news_v1_news_proto_depIdxs,
		MessageInfos:      file_proto_news_v1_news_proto_msgTypes,
	}.Build()
	File_proto_news_v1_news_proto = out.File
	file_proto_news_v1_news_proto_rawDesc = nil
	file_proto_news_v1_news_proto_goTypes = nil
	file_proto_news_v1_news_proto_depIdxs = nil
}
