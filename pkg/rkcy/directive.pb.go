// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: directive.proto

package rkcy

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Type implies what topics will be present
type Directive int32

const (
	Directive_UNSPECIFIED            Directive = 0
	Directive_PLATFORM               Directive = 65536
	Directive_ADMIN_PRODUCER         Directive = 131072
	Directive_ADMIN_PRODUCER_STOP    Directive = 131073
	Directive_ADMIN_PRODUCER_STOPPED Directive = 131074
	Directive_ADMIN_PRODUCER_START   Directive = 131076
	Directive_ADMIN_PRODUCER_STARTED Directive = 131080
	Directive_ADMIN_PRODUCER_STATUS  Directive = 131088
	Directive_ADMIN_CONSUMER         Directive = 262144
	Directive_ADMIN_CONSUMER_STOP    Directive = 262145
	Directive_ADMIN_CONSUMER_STOPPED Directive = 262146
	Directive_ADMIN_CONSUMER_START   Directive = 262148
	Directive_ADMIN_CONSUMER_STARTED Directive = 262152
	Directive_APECS_TXN              Directive = 524288
	Directive_ALL                    Directive = -1
)

// Enum value maps for Directive.
var (
	Directive_name = map[int32]string{
		0:      "UNSPECIFIED",
		65536:  "PLATFORM",
		131072: "ADMIN_PRODUCER",
		131073: "ADMIN_PRODUCER_STOP",
		131074: "ADMIN_PRODUCER_STOPPED",
		131076: "ADMIN_PRODUCER_START",
		131080: "ADMIN_PRODUCER_STARTED",
		131088: "ADMIN_PRODUCER_STATUS",
		262144: "ADMIN_CONSUMER",
		262145: "ADMIN_CONSUMER_STOP",
		262146: "ADMIN_CONSUMER_STOPPED",
		262148: "ADMIN_CONSUMER_START",
		262152: "ADMIN_CONSUMER_STARTED",
		524288: "APECS_TXN",
		-1:     "ALL",
	}
	Directive_value = map[string]int32{
		"UNSPECIFIED":            0,
		"PLATFORM":               65536,
		"ADMIN_PRODUCER":         131072,
		"ADMIN_PRODUCER_STOP":    131073,
		"ADMIN_PRODUCER_STOPPED": 131074,
		"ADMIN_PRODUCER_START":   131076,
		"ADMIN_PRODUCER_STARTED": 131080,
		"ADMIN_PRODUCER_STATUS":  131088,
		"ADMIN_CONSUMER":         262144,
		"ADMIN_CONSUMER_STOP":    262145,
		"ADMIN_CONSUMER_STOPPED": 262146,
		"ADMIN_CONSUMER_START":   262148,
		"ADMIN_CONSUMER_STARTED": 262152,
		"APECS_TXN":              524288,
		"ALL":                    -1,
	}
)

func (x Directive) Enum() *Directive {
	p := new(Directive)
	*p = x
	return p
}

func (x Directive) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Directive) Descriptor() protoreflect.EnumDescriptor {
	return file_directive_proto_enumTypes[0].Descriptor()
}

func (Directive) Type() protoreflect.EnumType {
	return &file_directive_proto_enumTypes[0]
}

func (x Directive) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Directive.Descriptor instead.
func (Directive) EnumDescriptor() ([]byte, []int) {
	return file_directive_proto_rawDescGZIP(), []int{0}
}

type AdminProducerDirective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concern    string `protobuf:"bytes,1,opt,name=concern,proto3" json:"concern,omitempty"`
	Topic      string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Generation int32  `protobuf:"varint,3,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *AdminProducerDirective) Reset() {
	*x = AdminProducerDirective{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminProducerDirective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminProducerDirective) ProtoMessage() {}

func (x *AdminProducerDirective) ProtoReflect() protoreflect.Message {
	mi := &file_directive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminProducerDirective.ProtoReflect.Descriptor instead.
func (*AdminProducerDirective) Descriptor() ([]byte, []int) {
	return file_directive_proto_rawDescGZIP(), []int{0}
}

func (x *AdminProducerDirective) GetConcern() string {
	if x != nil {
		return x.Concern
	}
	return ""
}

func (x *AdminProducerDirective) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *AdminProducerDirective) GetGeneration() int32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

type AdminConsumerDirective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program *Program `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
}

func (x *AdminConsumerDirective) Reset() {
	*x = AdminConsumerDirective{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminConsumerDirective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminConsumerDirective) ProtoMessage() {}

func (x *AdminConsumerDirective) ProtoReflect() protoreflect.Message {
	mi := &file_directive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminConsumerDirective.ProtoReflect.Descriptor instead.
func (*AdminConsumerDirective) Descriptor() ([]byte, []int) {
	return file_directive_proto_rawDescGZIP(), []int{1}
}

func (x *AdminConsumerDirective) GetProgram() *Program {
	if x != nil {
		return x.Program
	}
	return nil
}

var File_directive_proto protoreflect.FileDescriptor

var file_directive_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x72, 0x6b, 0x63, 0x79, 0x1a, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x16, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x41, 0x0a, 0x16, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72,
	0x6b, 0x63, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x2a, 0xfe, 0x02, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x08, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x10,
	0x80, 0x80, 0x04, 0x12, 0x14, 0x0a, 0x0e, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x45, 0x52, 0x10, 0x80, 0x80, 0x08, 0x12, 0x19, 0x0a, 0x13, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x4f, 0x50,
	0x10, 0x81, 0x80, 0x08, 0x12, 0x1c, 0x0a, 0x16, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x82,
	0x80, 0x08, 0x12, 0x1a, 0x0a, 0x14, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x84, 0x80, 0x08, 0x12, 0x1c,
	0x0a, 0x16, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x88, 0x80, 0x08, 0x12, 0x1b, 0x0a, 0x15,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x90, 0x80, 0x08, 0x12, 0x14, 0x0a, 0x0e, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x80, 0x80, 0x10, 0x12,
	0x19, 0x0a, 0x13, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x81, 0x80, 0x10, 0x12, 0x1c, 0x0a, 0x16, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x4f,
	0x50, 0x50, 0x45, 0x44, 0x10, 0x82, 0x80, 0x10, 0x12, 0x1a, 0x0a, 0x14, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x84, 0x80, 0x10, 0x12, 0x1c, 0x0a, 0x16, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x43, 0x4f,
	0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x88,
	0x80, 0x10, 0x12, 0x0f, 0x0a, 0x09, 0x41, 0x50, 0x45, 0x43, 0x53, 0x5f, 0x54, 0x58, 0x4e, 0x10,
	0x80, 0x80, 0x20, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x63, 0x68, 0x6c, 0x61, 0x6e, 0x6f, 0x72, 0x72, 0x2f, 0x72,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x6b, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_directive_proto_rawDescOnce sync.Once
	file_directive_proto_rawDescData = file_directive_proto_rawDesc
)

func file_directive_proto_rawDescGZIP() []byte {
	file_directive_proto_rawDescOnce.Do(func() {
		file_directive_proto_rawDescData = protoimpl.X.CompressGZIP(file_directive_proto_rawDescData)
	})
	return file_directive_proto_rawDescData
}

var file_directive_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_directive_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_directive_proto_goTypes = []interface{}{
	(Directive)(0),                 // 0: rkcy.Directive
	(*AdminProducerDirective)(nil), // 1: rkcy.AdminProducerDirective
	(*AdminConsumerDirective)(nil), // 2: rkcy.AdminConsumerDirective
	(*Program)(nil),                // 3: rkcy.Program
}
var file_directive_proto_depIdxs = []int32{
	3, // 0: rkcy.AdminConsumerDirective.program:type_name -> rkcy.Program
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_directive_proto_init() }
func file_directive_proto_init() {
	if File_directive_proto != nil {
		return
	}
	file_platform_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_directive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminProducerDirective); i {
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
		file_directive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminConsumerDirective); i {
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
			RawDescriptor: file_directive_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_directive_proto_goTypes,
		DependencyIndexes: file_directive_proto_depIdxs,
		EnumInfos:         file_directive_proto_enumTypes,
		MessageInfos:      file_directive_proto_msgTypes,
	}.Build()
	File_directive_proto = out.File
	file_directive_proto_rawDesc = nil
	file_directive_proto_goTypes = nil
	file_directive_proto_depIdxs = nil
}
