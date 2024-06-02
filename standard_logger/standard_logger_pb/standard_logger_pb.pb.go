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

type GormTracingLevel int32

const (
	GormTracingLevel_Silent GormTracingLevel = 0
	GormTracingLevel_Error  GormTracingLevel = 1
	GormTracingLevel_Warn   GormTracingLevel = 2
	GormTracingLevel_Info   GormTracingLevel = 3
)

// Enum value maps for GormTracingLevel.
var (
	GormTracingLevel_name = map[int32]string{
		0: "Silent",
		1: "Error",
		2: "Warn",
		3: "Info",
	}
	GormTracingLevel_value = map[string]int32{
		"Silent": 0,
		"Error":  1,
		"Warn":   2,
		"Info":   3,
	}
)

func (x GormTracingLevel) Enum() *GormTracingLevel {
	p := new(GormTracingLevel)
	*p = x
	return p
}

func (x GormTracingLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GormTracingLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_standard_logger_pb_standard_logger_pb_proto_enumTypes[1].Descriptor()
}

func (GormTracingLevel) Type() protoreflect.EnumType {
	return &file_standard_logger_pb_standard_logger_pb_proto_enumTypes[1]
}

func (x GormTracingLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GormTracingLevel.Descriptor instead.
func (GormTracingLevel) EnumDescriptor() ([]byte, []int) {
	return file_standard_logger_pb_standard_logger_pb_proto_rawDescGZIP(), []int{1}
}

// 日志配置的proto
type LoggerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootDir         string               `protobuf:"bytes,1,opt,name=rootDir,proto3" json:"rootDir,omitempty"`
	LogLevel        LoggerLevel          `protobuf:"varint,2,opt,name=logLevel,proto3,enum=standard_logger_pb.LoggerLevel" json:"logLevel,omitempty"`
	DirName         *string              `protobuf:"bytes,3,opt,name=DirName,proto3,oneof" json:"DirName,omitempty"`
	StackTraceLevel *LoggerLevel         `protobuf:"varint,4,opt,name=stackTraceLevel,proto3,enum=standard_logger_pb.LoggerLevel,oneof" json:"stackTraceLevel,omitempty"`
	RotationTime    *durationpb.Duration `protobuf:"bytes,5,opt,name=rotationTime,proto3,oneof" json:"rotationTime,omitempty"`
	MaxAge          *durationpb.Duration `protobuf:"bytes,6,opt,name=maxAge,proto3,oneof" json:"maxAge,omitempty"`
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

func (x *LoggerConfig) GetDirName() string {
	if x != nil && x.DirName != nil {
		return *x.DirName
	}
	return ""
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

// gorm tracing
type GormTracingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TracingLevel  GormTracingLevel     `protobuf:"varint,1,opt,name=tracingLevel,proto3,enum=standard_logger_pb.GormTracingLevel" json:"tracingLevel,omitempty"`
	SlowThreshold *durationpb.Duration `protobuf:"bytes,2,opt,name=SlowThreshold,proto3" json:"SlowThreshold,omitempty"`
	DontIgnoreRecordNotFoundError *bool                `protobuf:"varint,3,opt,name=DontIgnoreRecordNotFoundError,proto3,oneof" json:"DontIgnoreRecordNotFoundError,omitempty"`
	DontIgnoreKeyDuplicateError   *bool                `protobuf:"varint,4,opt,name=DontIgnoreKeyDuplicateError,proto3,oneof" json:"DontIgnoreKeyDuplicateError,omitempty"`
	FilterParams                  *bool                `protobuf:"varint,5,opt,name=FilterParams,proto3,oneof" json:"FilterParams,omitempty"`
}

func (x *GormTracingConfig) Reset() {
	*x = GormTracingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standard_logger_pb_standard_logger_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GormTracingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GormTracingConfig) ProtoMessage() {}

func (x *GormTracingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_standard_logger_pb_standard_logger_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GormTracingConfig.ProtoReflect.Descriptor instead.
func (*GormTracingConfig) Descriptor() ([]byte, []int) {
	return file_standard_logger_pb_standard_logger_pb_proto_rawDescGZIP(), []int{1}
}

func (x *GormTracingConfig) GetTracingLevel() GormTracingLevel {
	if x != nil {
		return x.TracingLevel
	}
	return GormTracingLevel_Silent
}

func (x *GormTracingConfig) GetSlowThreshold() *durationpb.Duration {
	if x != nil {
		return x.SlowThreshold
	}
	return nil
}

func (x *GormTracingConfig) GetDontIgnoreRecordNotFoundError() bool {
	if x != nil && x.DontIgnoreRecordNotFoundError != nil {
		return *x.DontIgnoreRecordNotFoundError
	}
	return false
}

func (x *GormTracingConfig) GetDontIgnoreKeyDuplicateError() bool {
	if x != nil && x.DontIgnoreKeyDuplicateError != nil {
		return *x.DontIgnoreKeyDuplicateError
	}
	return false
}

func (x *GormTracingConfig) GetFilterParams() bool {
	if x != nil && x.FilterParams != nil {
		return *x.FilterParams
	}
	return false
}

var File_standard_logger_pb_standard_logger_pb_proto protoreflect.FileDescriptor

var file_standard_logger_pb_standard_logger_pb_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x5f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x12, 0x3b, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x07, 0x44, 0x69, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x44, 0x69,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x48, 0x01, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x0c, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x06,
	0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x44, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x22, 0xac, 0x03, 0x0a, 0x11, 0x47, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x62, 0x2e, 0x47, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x3f, 0x0a, 0x0d, 0x53, 0x6c, 0x6f, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x53, 0x6c, 0x6f, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x49, 0x0a, 0x1d, 0x44, 0x6f, 0x6e, 0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x1d, 0x44, 0x6f, 0x6e, 0x74,
	0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x1b,
	0x44, 0x6f, 0x6e, 0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x44, 0x75, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x01, 0x52, 0x1b, 0x44, 0x6f, 0x6e, 0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x4b,
	0x65, 0x79, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0c, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x42, 0x20, 0x0a, 0x1e,
	0x5f, 0x44, 0x6f, 0x6e, 0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x1e,
	0x0a, 0x1c, 0x5f, 0x44, 0x6f, 0x6e, 0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79,
	0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a,
	0x7c, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e,
	0x0a, 0x0a, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x57, 0x61, 0x72, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x04, 0x12, 0x0e, 0x0a,
	0x0a, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x05, 0x12, 0x0e, 0x0a,
	0x0a, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x06, 0x2a, 0x3d, 0x0a,
	0x10, 0x47, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x61, 0x72, 0x6e,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x03, 0x42, 0x56, 0x5a, 0x54,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6e, 0x74, 0x42,
	0x65, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x62, 0x3b, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_standard_logger_pb_standard_logger_pb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_standard_logger_pb_standard_logger_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_standard_logger_pb_standard_logger_pb_proto_goTypes = []interface{}{
	(LoggerLevel)(0),            // 0: standard_logger_pb.LoggerLevel
	(GormTracingLevel)(0),       // 1: standard_logger_pb.GormTracingLevel
	(*LoggerConfig)(nil),        // 2: standard_logger_pb.LoggerConfig
	(*GormTracingConfig)(nil),   // 3: standard_logger_pb.GormTracingConfig
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_standard_logger_pb_standard_logger_pb_proto_depIdxs = []int32{
	0, // 0: standard_logger_pb.LoggerConfig.logLevel:type_name -> standard_logger_pb.LoggerLevel
	0, // 1: standard_logger_pb.LoggerConfig.stackTraceLevel:type_name -> standard_logger_pb.LoggerLevel
	4, // 2: standard_logger_pb.LoggerConfig.rotationTime:type_name -> google.protobuf.Duration
	4, // 3: standard_logger_pb.LoggerConfig.maxAge:type_name -> google.protobuf.Duration
	1, // 4: standard_logger_pb.GormTracingConfig.tracingLevel:type_name -> standard_logger_pb.GormTracingLevel
	4, // 5: standard_logger_pb.GormTracingConfig.SlowThreshold:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
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
		file_standard_logger_pb_standard_logger_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GormTracingConfig); i {
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
	file_standard_logger_pb_standard_logger_pb_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_standard_logger_pb_standard_logger_pb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
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
