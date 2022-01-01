// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/news/v1/news.proto

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

type NewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Exchange string `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Period   string `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
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

func (x *NewsRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
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

	Date  int64    `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	Title string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Link  string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Tags  []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
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

func (x *NewsReply) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
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

type NewsReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*NewsReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Date    int64        `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
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

func (x *NewsReplies) GetDate() int64 {
	if x != nil {
		return x.Date
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
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x08, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0x98, 0x01, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x25, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x08, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xfa, 0x42, 0x10, 0x72, 0x0e, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x02, 0x33, 0x64, 0x52, 0x02, 0x31, 0x77, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x22, 0x5d, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x65, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0x9d, 0x01, 0x0a, 0x04, 0x4e,
	0x65, 0x77, 0x73, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4a,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x40, 0x0a, 0x16, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	(*NewsRequest)(nil),   // 0: news.v1.NewsRequest
	(*NewsReply)(nil),     // 1: news.v1.NewsReply
	(*NewsReplies)(nil),   // 2: news.v1.NewsReplies
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_proto_news_v1_news_proto_depIdxs = []int32{
	1, // 0: news.v1.NewsReplies.results:type_name -> news.v1.NewsReply
	0, // 1: news.v1.News.Search:input_type -> news.v1.NewsRequest
	3, // 2: news.v1.News.Health:input_type -> google.protobuf.Empty
	2, // 3: news.v1.News.Search:output_type -> news.v1.NewsReplies
	3, // 4: news.v1.News.Health:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
