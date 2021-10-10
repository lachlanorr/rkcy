// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: rkcy/admin.proto

package rkcy

import (
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

type DecodeInstanceArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concern   string `protobuf:"bytes,1,opt,name=concern,proto3" json:"concern,omitempty"`
	Payload64 string `protobuf:"bytes,2,opt,name=payload64,proto3" json:"payload64,omitempty"`
}

func (x *DecodeInstanceArgs) Reset() {
	*x = DecodeInstanceArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeInstanceArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeInstanceArgs) ProtoMessage() {}

func (x *DecodeInstanceArgs) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeInstanceArgs.ProtoReflect.Descriptor instead.
func (*DecodeInstanceArgs) Descriptor() ([]byte, []int) {
	return file_rkcy_admin_proto_rawDescGZIP(), []int{0}
}

func (x *DecodeInstanceArgs) GetConcern() string {
	if x != nil {
		return x.Concern
	}
	return ""
}

func (x *DecodeInstanceArgs) GetPayload64() string {
	if x != nil {
		return x.Payload64
	}
	return ""
}

type DecodePayloadArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concern   string `protobuf:"bytes,1,opt,name=concern,proto3" json:"concern,omitempty"`
	System    System `protobuf:"varint,2,opt,name=system,proto3,enum=rkcy.System" json:"system,omitempty"`
	Command   string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Payload64 string `protobuf:"bytes,4,opt,name=payload64,proto3" json:"payload64,omitempty"`
}

func (x *DecodePayloadArgs) Reset() {
	*x = DecodePayloadArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodePayloadArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodePayloadArgs) ProtoMessage() {}

func (x *DecodePayloadArgs) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodePayloadArgs.ProtoReflect.Descriptor instead.
func (*DecodePayloadArgs) Descriptor() ([]byte, []int) {
	return file_rkcy_admin_proto_rawDescGZIP(), []int{1}
}

func (x *DecodePayloadArgs) GetConcern() string {
	if x != nil {
		return x.Concern
	}
	return ""
}

func (x *DecodePayloadArgs) GetSystem() System {
	if x != nil {
		return x.System
	}
	return System_NO_SYSTEM
}

func (x *DecodePayloadArgs) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *DecodePayloadArgs) GetPayload64() string {
	if x != nil {
		return x.Payload64
	}
	return ""
}

type DecodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *DecodeResponse) Reset() {
	*x = DecodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResponse) ProtoMessage() {}

func (x *DecodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResponse.ProtoReflect.Descriptor instead.
func (*DecodeResponse) Descriptor() ([]byte, []int) {
	return file_rkcy_admin_proto_rawDescGZIP(), []int{2}
}

func (x *DecodeResponse) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type TrackedProducers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicProducers []*TrackedProducers_ProducerInfo `protobuf:"bytes,1,rep,name=topic_producers,json=topicProducers,proto3" json:"topic_producers,omitempty"`
}

func (x *TrackedProducers) Reset() {
	*x = TrackedProducers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackedProducers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedProducers) ProtoMessage() {}

func (x *TrackedProducers) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedProducers.ProtoReflect.Descriptor instead.
func (*TrackedProducers) Descriptor() ([]byte, []int) {
	return file_rkcy_admin_proto_rawDescGZIP(), []int{3}
}

func (x *TrackedProducers) GetTopicProducers() []*TrackedProducers_ProducerInfo {
	if x != nil {
		return x.TopicProducers
	}
	return nil
}

type TrackedProducers_ProducerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic           string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Id              string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TimeSinceUpdate string `protobuf:"bytes,3,opt,name=time_since_update,json=timeSinceUpdate,proto3" json:"time_since_update,omitempty"`
}

func (x *TrackedProducers_ProducerInfo) Reset() {
	*x = TrackedProducers_ProducerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackedProducers_ProducerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedProducers_ProducerInfo) ProtoMessage() {}

func (x *TrackedProducers_ProducerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedProducers_ProducerInfo.ProtoReflect.Descriptor instead.
func (*TrackedProducers_ProducerInfo) Descriptor() ([]byte, []int) {
	return file_rkcy_admin_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TrackedProducers_ProducerInfo) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TrackedProducers_ProducerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackedProducers_ProducerInfo) GetTimeSinceUpdate() string {
	if x != nil {
		return x.TimeSinceUpdate
	}
	return ""
}

var File_rkcy_admin_proto protoreflect.FileDescriptor

var file_rkcy_admin_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x6b, 0x63, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x6b, 0x63, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x72, 0x6b, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x65,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x6b, 0x63, 0x79, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a,
	0x12, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x36, 0x34, 0x22, 0x8b, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x67,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x72, 0x6b,
	0x63, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x36, 0x34, 0x22, 0x24, 0x0a, 0x0e, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22,
	0xc2, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x72, 0x6b, 0x63, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x73, 0x1a, 0x60, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x32, 0xc2, 0x03, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x0a, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0e, 0x2e,
	0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x4b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x16, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x73, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x73,
	0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x60, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x14, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x41, 0x72, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x2e, 0x72, 0x6b,
	0x63, 0x79, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x14, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x61,
	0x72, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x13, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x2e, 0x72,
	0x6b, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x14, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x63, 0x68, 0x6c, 0x61, 0x6e, 0x6f,
	0x72, 0x72, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x6b, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rkcy_admin_proto_rawDescOnce sync.Once
	file_rkcy_admin_proto_rawDescData = file_rkcy_admin_proto_rawDesc
)

func file_rkcy_admin_proto_rawDescGZIP() []byte {
	file_rkcy_admin_proto_rawDescOnce.Do(func() {
		file_rkcy_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_rkcy_admin_proto_rawDescData)
	})
	return file_rkcy_admin_proto_rawDescData
}

var file_rkcy_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rkcy_admin_proto_goTypes = []interface{}{
	(*DecodeInstanceArgs)(nil),            // 0: rkcy.DecodeInstanceArgs
	(*DecodePayloadArgs)(nil),             // 1: rkcy.DecodePayloadArgs
	(*DecodeResponse)(nil),                // 2: rkcy.DecodeResponse
	(*TrackedProducers)(nil),              // 3: rkcy.TrackedProducers
	(*TrackedProducers_ProducerInfo)(nil), // 4: rkcy.TrackedProducers.ProducerInfo
	(System)(0),                           // 5: rkcy.System
	(*Void)(nil),                          // 6: rkcy.Void
	(*Platform)(nil),                      // 7: rkcy.Platform
}
var file_rkcy_admin_proto_depIdxs = []int32{
	5, // 0: rkcy.DecodePayloadArgs.system:type_name -> rkcy.System
	4, // 1: rkcy.TrackedProducers.topic_producers:type_name -> rkcy.TrackedProducers.ProducerInfo
	6, // 2: rkcy.AdminService.Platform:input_type -> rkcy.Void
	6, // 3: rkcy.AdminService.Producers:input_type -> rkcy.Void
	0, // 4: rkcy.AdminService.DecodeInstance:input_type -> rkcy.DecodeInstanceArgs
	1, // 5: rkcy.AdminService.DecodeArgPayload:input_type -> rkcy.DecodePayloadArgs
	1, // 6: rkcy.AdminService.DecodeResultPayload:input_type -> rkcy.DecodePayloadArgs
	7, // 7: rkcy.AdminService.Platform:output_type -> rkcy.Platform
	3, // 8: rkcy.AdminService.Producers:output_type -> rkcy.TrackedProducers
	2, // 9: rkcy.AdminService.DecodeInstance:output_type -> rkcy.DecodeResponse
	2, // 10: rkcy.AdminService.DecodeArgPayload:output_type -> rkcy.DecodeResponse
	2, // 11: rkcy.AdminService.DecodeResultPayload:output_type -> rkcy.DecodeResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rkcy_admin_proto_init() }
func file_rkcy_admin_proto_init() {
	if File_rkcy_admin_proto != nil {
		return
	}
	file_rkcy_apecs_proto_init()
	file_rkcy_platform_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rkcy_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeInstanceArgs); i {
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
		file_rkcy_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodePayloadArgs); i {
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
		file_rkcy_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResponse); i {
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
		file_rkcy_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackedProducers); i {
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
		file_rkcy_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackedProducers_ProducerInfo); i {
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
			RawDescriptor: file_rkcy_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rkcy_admin_proto_goTypes,
		DependencyIndexes: file_rkcy_admin_proto_depIdxs,
		MessageInfos:      file_rkcy_admin_proto_msgTypes,
	}.Build()
	File_rkcy_admin_proto = out.File
	file_rkcy_admin_proto_rawDesc = nil
	file_rkcy_admin_proto_goTypes = nil
	file_rkcy_admin_proto_depIdxs = nil
}
