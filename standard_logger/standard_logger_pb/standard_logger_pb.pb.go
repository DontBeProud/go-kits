// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: standard_logger_pb/standard_logger_pb.proto

package standard_logger_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoggerLevel int32

const (
	LoggerLevel_DebugLevel  LoggerLevel = 0
	LoggerLevel_InfoLevel   LoggerLevel = 1
	LoggerLevel_WarnLevel   LoggerLevel = 2
	LoggerLevel_ErrorLevel  LoggerLevel = 3
	LoggerLevel_DPanicLevel LoggerLevel = 4
	LoggerLevel_PanicLevel  LoggerLevel = 5
	LoggerLevel_FatalLevel  LoggerLevel = 6
)

// Enum value maps for LoggerLevel.
var (
	LoggerLevel_name = map[int32]string{
		0: "DebugLevel",
		1: "InfoLevel",
		2: "WarnLevel",
		3: "ErrorLevel",
		4: "DPanicLevel",
		5: "PanicLevel",
		6: "FatalLevel",
	}
	LoggerLevel_value = map[string]int32{
		"DebugLevel":  0,
		"InfoLevel":   1,
		"WarnLevel":   2,
		"ErrorLevel":  3,
		"DPanicLevel": 4,
		"PanicLevel":  5,
		"FatalLevel":  6,
	}
)

func (x LoggerLevel) Enum() *LoggerLevel {
	p := new(LoggerLevel)
	*p = x
	return p
}

func (x LoggerLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggerLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_standard_logger_pb_standard_logger_pb_proto_enumTypes[0].Descriptor()
}

func (LoggerLevel) Type() protoreflect.EnumType {
	return &file_standard_logger_pb_standard_logger_pb_proto_enumTypes[0]
}

func (x LoggerLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoggerLevel.Descriptor instead.
func (LoggerLevel) EnumDescriptor() ([]byte, []int) {
	return file_standard_logger_pb_standard_logger_pb_proto_rawDescGZIP(), []int{0}
}

// 日志配置的proto
type LoggerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootDir         string               `protobuf:"bytes,1,opt,name=rootDir,proto3" json:"rootDir,omitempty"`
	LogLevel        LoggerLevel          `protobuf:"varint,2,opt,name=logLevel,proto3,enum=standard_logger_pb.LoggerLevel" json:"logLevel,omitempty"`
	StackTraceLevel *LoggerLevel         `protobuf:"varint,3,opt,name=stackTraceLevel,proto3,enum=standard_logger_pb.LoggerLevel,oneof" json:"stackTraceLevel,omitempty"`
	RotationTime    *durationpb.Duration `protobuf:"bytes,4,opt,name=rotationTime,proto3,oneof" json:"rotationTime,omitempty"`
	MaxAge          *durationpb.Duration `protobuf:"bytes,5,opt,name=maxAge,proto3,oneof" json:"maxAge,omitempty"`
}

func (x *LoggerConfig) Reset() {
	*x = LoggerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standard_logger_pb_standard_logger_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerConfig) ProtoMessage() {}

func (x *LoggerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_standard_logger_pb_standard_logger_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerConfig.ProtoReflect.Descriptor instead.
func (*LoggerConfig) Descriptor() ([]byte, []int) {
	return file_standard_logger_pb_standard_logger_pb_proto_rawDescGZIP(), []int{0}
}

func (x *LoggerConfig) GetRootDir() string {
	if x != nil {
		return x.RootDir
	}
	return ""
}

func (x *LoggerConfig) GetLogLevel() LoggerLevel {
	if x != nil {
		return x.LogLevel
	}
	return LoggerLevel_DebugLevel
}

func (x *LoggerConfig) GetStackTraceLevel() LoggerLevel {
	if x != nil && x.StackTraceLevel != nil {
		return *x.StackTraceLevel
	}
	return LoggerLevel_DebugLevel
}

func (x *LoggerConfig) GetRotationTime() *durationpb.Duration {
	if x != nil {
		return x.RotationTime
	}
	return nil
}

func (x *LoggerConfig) GetMaxAge() *durationpb.Duration {
	if x != nil {
		return x.MaxAge
	}
	return nil
}

var File_standard_logger_pb_standard_logger_pb_proto protoreflect.FileDescriptor

var file_standard_logger_pb_standard_logger_pb_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x5f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x12, 0x3b, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x4e, 0x0a, 0x0f, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0c, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a,
	0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d,
	0x61, 0x78, 0x41, 0x67, 0x65, 0x2a, 0x7c, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x61, 0x72, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x10, 0x06, 0x42, 0x29, 0x5a, 0x27, 0x2e, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_standard_logger_pb_standard_logger_pb_proto_rawDescOnce sync.Once
	file_standard_logger_pb_standard_logger_pb_proto_rawDescData = file_standard_logger_pb_standard_logger_pb_proto_rawDesc
)

func file_standard_logger_pb_standard_logger_pb_proto_rawDescGZIP() []byte {
	file_standard_logger_pb_standard_logger_pb_proto_rawDescOnce.Do(func() {
		file_standard_logger_pb_standard_logger_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_standard_logger_pb_standard_logger_pb_proto_rawDescData)
	})
	return file_standard_logger_pb_standard_logger_pb_proto_rawDescData
}

var file_standard_logger_pb_standard_logger_pb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_standard_logger_pb_standard_logger_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_standard_logger_pb_standard_logger_pb_proto_goTypes = []interface{}{
	(LoggerLevel)(0),            // 0: standard_logger_pb.LoggerLevel
	(*LoggerConfig)(nil),        // 1: standard_logger_pb.LoggerConfig
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_standard_logger_pb_standard_logger_pb_proto_depIdxs = []int32{
	0, // 0: standard_logger_pb.LoggerConfig.logLevel:type_name -> standard_logger_pb.LoggerLevel
	0, // 1: standard_logger_pb.LoggerConfig.stackTraceLevel:type_name -> standard_logger_pb.LoggerLevel
	2, // 2: standard_logger_pb.LoggerConfig.rotationTime:type_name -> google.protobuf.Duration
	2, // 3: standard_logger_pb.LoggerConfig.maxAge:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_standard_logger_pb_standard_logger_pb_proto_init() }
func file_standard_logger_pb_standard_logger_pb_proto_init() {
	if File_standard_logger_pb_standard_logger_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_standard_logger_pb_standard_logger_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerConfig); i {
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
	file_standard_logger_pb_standard_logger_pb_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_standard_logger_pb_standard_logger_pb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_standard_logger_pb_standard_logger_pb_proto_goTypes,
		DependencyIndexes: file_standard_logger_pb_standard_logger_pb_proto_depIdxs,
		EnumInfos:         file_standard_logger_pb_standard_logger_pb_proto_enumTypes,
		MessageInfos:      file_standard_logger_pb_standard_logger_pb_proto_msgTypes,
	}.Build()
	File_standard_logger_pb_standard_logger_pb_proto = out.File
	file_standard_logger_pb_standard_logger_pb_proto_rawDesc = nil
	file_standard_logger_pb_standard_logger_pb_proto_goTypes = nil
	file_standard_logger_pb_standard_logger_pb_proto_depIdxs = nil
}
