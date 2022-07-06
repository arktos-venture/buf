// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/stats/v1/error_reason.proto

package v1Stats

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_CONTENT_MISSING ErrorReason = 0
	ErrorReason_UNAUTHORIZED    ErrorReason = 1
	ErrorReason_FORBIDDEN       ErrorReason = 2
	ErrorReason_NOT_FOUND       ErrorReason = 3
	ErrorReason_REQUEST_TIMEOUT ErrorReason = 4
	ErrorReason_CONFLICT        ErrorReason = 5
	ErrorReason_NOT_IMPLEMENTED ErrorReason = 6
	ErrorReason_UNAVAILABLE     ErrorReason = 7
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "CONTENT_MISSING",
		1: "UNAUTHORIZED",
		2: "FORBIDDEN",
		3: "NOT_FOUND",
		4: "REQUEST_TIMEOUT",
		5: "CONFLICT",
		6: "NOT_IMPLEMENTED",
		7: "UNAVAILABLE",
	}
	ErrorReason_value = map[string]int32{
		"CONTENT_MISSING": 0,
		"UNAUTHORIZED":    1,
		"FORBIDDEN":       2,
		"NOT_FOUND":       3,
		"REQUEST_TIMEOUT": 4,
		"CONFLICT":        5,
		"NOT_IMPLEMENTED": 6,
		"UNAVAILABLE":     7,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_stats_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_proto_stats_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_stats_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_proto_stats_v1_error_reason_proto protoreflect.FileDescriptor

var file_proto_stats_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xd1, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x16, 0x0a,
	0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x1a,
	0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44,
	0x45, 0x4e, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12,
	0x19, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x43, 0x4f,
	0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x19,
	0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45,
	0x44, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x15, 0x0a, 0x0b, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x5a, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b, 0x74, 0x6f, 0x73, 0x2d, 0x76,
	0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x53, 0x74, 0x61, 0x74,
	0x73, 0xa2, 0x02, 0x0e, 0x41, 0x50, 0x49, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stats_v1_error_reason_proto_rawDescOnce sync.Once
	file_proto_stats_v1_error_reason_proto_rawDescData = file_proto_stats_v1_error_reason_proto_rawDesc
)

func file_proto_stats_v1_error_reason_proto_rawDescGZIP() []byte {
	file_proto_stats_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_proto_stats_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stats_v1_error_reason_proto_rawDescData)
	})
	return file_proto_stats_v1_error_reason_proto_rawDescData
}

var file_proto_stats_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_stats_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: stats.v1.ErrorReason
}
var file_proto_stats_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_stats_v1_error_reason_proto_init() }
func file_proto_stats_v1_error_reason_proto_init() {
	if File_proto_stats_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_stats_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_stats_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_proto_stats_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_proto_stats_v1_error_reason_proto_enumTypes,
	}.Build()
	File_proto_stats_v1_error_reason_proto = out.File
	file_proto_stats_v1_error_reason_proto_rawDesc = nil
	file_proto_stats_v1_error_reason_proto_goTypes = nil
	file_proto_stats_v1_error_reason_proto_depIdxs = nil
}
