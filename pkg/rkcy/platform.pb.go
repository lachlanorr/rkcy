// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: rkcy/platform.proto

package rkcy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Different types of partitioning mechanisms
type Platform_PartitionStrategy int32

const (
	// FNV-1 64 bit has followed by modulus of partition count
	Platform_FNV64_MOD Platform_PartitionStrategy = 0
)

// Enum value maps for Platform_PartitionStrategy.
var (
	Platform_PartitionStrategy_name = map[int32]string{
		0: "FNV64_MOD",
	}
	Platform_PartitionStrategy_value = map[string]int32{
		"FNV64_MOD": 0,
	}
)

func (x Platform_PartitionStrategy) Enum() *Platform_PartitionStrategy {
	p := new(Platform_PartitionStrategy)
	*p = x
	return p
}

func (x Platform_PartitionStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform_PartitionStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_rkcy_platform_proto_enumTypes[0].Descriptor()
}

func (Platform_PartitionStrategy) Type() protoreflect.EnumType {
	return &file_rkcy_platform_proto_enumTypes[0]
}

func (x Platform_PartitionStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform_PartitionStrategy.Descriptor instead.
func (Platform_PartitionStrategy) EnumDescriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 0}
}

// Type implies what topics will be present
type Platform_Concern_Type int32

const (
	// Required topics:
	//     * admin - handles control messages to coordinate distributed tasks
	//     * error - errors encountered during processing written here
	//     * ... - additional concern specific topics
	Platform_Concern_GENERAL Platform_Concern_Type = 0
	// Required topics:
	//     * admin - handles control messages to coordinate distributed tasks
	//     * error - errors encountered during processing written here
	//     * ... - additional concern specific topics for stages of batch processing
	// For telemetry purposes, stage topics are assumed to be in alphabetical order.
	// Batch jobs topics aren't created by default, but on demand with the timestamp
	// included in the names.
	Platform_Concern_BATCH Platform_Concern_Type = 1
	// Required topics:
	//     * admin - handles control messages to coordinate distributed tasks
	//     * process - handles messages that affect internal state of models
	//     * error - errors encountered during processing written here
	//     * complete - completed transactions written here for post processing
	//     * storage - handles messages sent to the persistence layer
	Platform_Concern_APECS Platform_Concern_Type = 2
)

// Enum value maps for Platform_Concern_Type.
var (
	Platform_Concern_Type_name = map[int32]string{
		0: "GENERAL",
		1: "BATCH",
		2: "APECS",
	}
	Platform_Concern_Type_value = map[string]int32{
		"GENERAL": 0,
		"BATCH":   1,
		"APECS":   2,
	}
)

func (x Platform_Concern_Type) Enum() *Platform_Concern_Type {
	p := new(Platform_Concern_Type)
	*p = x
	return p
}

func (x Platform_Concern_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform_Concern_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_rkcy_platform_proto_enumTypes[1].Descriptor()
}

func (Platform_Concern_Type) Type() protoreflect.EnumType {
	return &file_rkcy_platform_proto_enumTypes[1]
}

func (x Platform_Concern_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform_Concern_Type.Descriptor instead.
func (Platform_Concern_Type) EnumDescriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 0, 0}
}

// States surrounding current vs future topics and current to future transition
type Platform_Concern_Topics_State int32

const (
	// no future topic exists, everyting using current
	Platform_Concern_Topics_CURRENT Platform_Concern_Topics_State = 0
	// future topic added to concern and gets created
	Platform_Concern_Topics_FUTURE_INIT Platform_Concern_Topics_State = 1
	// producers all pause
	Platform_Concern_Topics_PRODUCER_PAUSE Platform_Concern_Topics_State = 2
	// consumers read until end and shutdown
	Platform_Concern_Topics_CONSUMER_SHUTDOWN Platform_Concern_Topics_State = 3
	// new consumers start on future, advanced passed newest
	Platform_Concern_Topics_CONSUMER_FUTURE_START Platform_Concern_Topics_State = 4
	// producers unpause and produce to future
	Platform_Concern_Topics_PRODUCER_FUTURE_START Platform_Concern_Topics_State = 5
	// future becomes current, future is nulled, state set to CURRENT
	Platform_Concern_Topics_FUTURE_TO_CURRENT Platform_Concern_Topics_State = 6
)

// Enum value maps for Platform_Concern_Topics_State.
var (
	Platform_Concern_Topics_State_name = map[int32]string{
		0: "CURRENT",
		1: "FUTURE_INIT",
		2: "PRODUCER_PAUSE",
		3: "CONSUMER_SHUTDOWN",
		4: "CONSUMER_FUTURE_START",
		5: "PRODUCER_FUTURE_START",
		6: "FUTURE_TO_CURRENT",
	}
	Platform_Concern_Topics_State_value = map[string]int32{
		"CURRENT":               0,
		"FUTURE_INIT":           1,
		"PRODUCER_PAUSE":        2,
		"CONSUMER_SHUTDOWN":     3,
		"CONSUMER_FUTURE_START": 4,
		"PRODUCER_FUTURE_START": 5,
		"FUTURE_TO_CURRENT":     6,
	}
)

func (x Platform_Concern_Topics_State) Enum() *Platform_Concern_Topics_State {
	p := new(Platform_Concern_Topics_State)
	*p = x
	return p
}

func (x Platform_Concern_Topics_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform_Concern_Topics_State) Descriptor() protoreflect.EnumDescriptor {
	return file_rkcy_platform_proto_enumTypes[2].Descriptor()
}

func (Platform_Concern_Topics_State) Type() protoreflect.EnumType {
	return &file_rkcy_platform_proto_enumTypes[2]
}

func (x Platform_Concern_Topics_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform_Concern_Topics_State.Descriptor instead.
func (Platform_Concern_Topics_State) EnumDescriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Concerns   []*Platform_Concern    `protobuf:"bytes,2,rep,name=concerns,proto3" json:"concerns,omitempty"`
	Clusters   []*Platform_Cluster    `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Platform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform) GetConcerns() []*Platform_Concern {
	if x != nil {
		return x.Concerns
	}
	return nil
}

func (x *Platform) GetClusters() []*Platform_Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *Platform) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type Program struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args   []string          `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Abbrev string            `protobuf:"bytes,3,opt,name=abbrev,proto3" json:"abbrev,omitempty"`
	Tags   map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Program) Reset() {
	*x = Program{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Program.ProtoReflect.Descriptor instead.
func (*Program) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{1}
}

func (x *Program) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Program) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Program) GetAbbrev() string {
	if x != nil {
		return x.Abbrev
	}
	return ""
}

func (x *Program) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type StorageSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsPrimary bool              `protobuf:"varint,2,opt,name=is_primary,json=isPrimary,proto3" json:"is_primary,omitempty"`
	Config    map[string]string `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StorageSystem) Reset() {
	*x = StorageSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageSystem) ProtoMessage() {}

func (x *StorageSystem) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageSystem.ProtoReflect.Descriptor instead.
func (*StorageSystem) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{2}
}

func (x *StorageSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StorageSystem) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

func (x *StorageSystem) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type Platform_Concern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     Platform_Concern_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=rkcy.Platform_Concern_Type" json:"type,omitempty"`
	Name     string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LogLevel Severity                   `protobuf:"varint,3,opt,name=log_level,json=logLevel,proto3,enum=rkcy.Severity" json:"log_level,omitempty"`
	Topics   []*Platform_Concern_Topics `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *Platform_Concern) Reset() {
	*x = Platform_Concern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform_Concern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform_Concern) ProtoMessage() {}

func (x *Platform_Concern) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform_Concern.ProtoReflect.Descriptor instead.
func (*Platform_Concern) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Platform_Concern) GetType() Platform_Concern_Type {
	if x != nil {
		return x.Type
	}
	return Platform_Concern_GENERAL
}

func (x *Platform_Concern) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform_Concern) GetLogLevel() Severity {
	if x != nil {
		return x.LogLevel
	}
	return Severity_DBG
}

func (x *Platform_Concern) GetTopics() []*Platform_Concern_Topics {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Platform_Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique name of cluster
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// kafka brokers for the cluster
	Brokers string `protobuf:"bytes,2,opt,name=brokers,proto3" json:"brokers,omitempty"`
}

func (x *Platform_Cluster) Reset() {
	*x = Platform_Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform_Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform_Cluster) ProtoMessage() {}

func (x *Platform_Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform_Cluster.ProtoReflect.Descriptor instead.
func (*Platform_Cluster) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Platform_Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform_Cluster) GetBrokers() string {
	if x != nil {
		return x.Brokers
	}
	return ""
}

type Platform_Concern_Topics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base name, it will get decorated with additional dot notated pieces
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// current vs future transition state
	State Platform_Concern_Topics_State `protobuf:"varint,2,opt,name=state,proto3,enum=rkcy.Platform_Concern_Topics_State" json:"state,omitempty"`
	// current physical topic
	Current *Platform_Concern_Topic `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	// topic we're in the process of migrating to, will be
	// null unless an active migration is taking place
	Future          *Platform_Concern_Topic `protobuf:"bytes,4,opt,name=future,proto3" json:"future,omitempty"`
	ConsumerProgram *Program                `protobuf:"bytes,5,opt,name=consumer_program,json=consumerProgram,proto3" json:"consumer_program,omitempty"`
}

func (x *Platform_Concern_Topics) Reset() {
	*x = Platform_Concern_Topics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform_Concern_Topics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform_Concern_Topics) ProtoMessage() {}

func (x *Platform_Concern_Topics) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform_Concern_Topics.ProtoReflect.Descriptor instead.
func (*Platform_Concern_Topics) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Platform_Concern_Topics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform_Concern_Topics) GetState() Platform_Concern_Topics_State {
	if x != nil {
		return x.State
	}
	return Platform_Concern_Topics_CURRENT
}

func (x *Platform_Concern_Topics) GetCurrent() *Platform_Concern_Topic {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *Platform_Concern_Topics) GetFuture() *Platform_Concern_Topic {
	if x != nil {
		return x.Future
	}
	return nil
}

func (x *Platform_Concern_Topics) GetConsumerProgram() *Program {
	if x != nil {
		return x.ConsumerProgram
	}
	return nil
}

type Platform_Concern_Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// incrementing generation id, +1 every time we migrate a logical topic
	Generation int32 `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	// kafka cluster topic exists within
	Cluster        string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	PartitionCount int32  `protobuf:"varint,3,opt,name=partition_count,json=partitionCount,proto3" json:"partition_count,omitempty"`
	// How to determine which partiton messages are produced to
	PartitionStrat Platform_PartitionStrategy `protobuf:"varint,4,opt,name=partition_strat,json=partitionStrat,proto3,enum=rkcy.Platform_PartitionStrategy" json:"partition_strat,omitempty"`
}

func (x *Platform_Concern_Topic) Reset() {
	*x = Platform_Concern_Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkcy_platform_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform_Concern_Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform_Concern_Topic) ProtoMessage() {}

func (x *Platform_Concern_Topic) ProtoReflect() protoreflect.Message {
	mi := &file_rkcy_platform_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform_Concern_Topic.ProtoReflect.Descriptor instead.
func (*Platform_Concern_Topic) Descriptor() ([]byte, []int) {
	return file_rkcy_platform_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Platform_Concern_Topic) GetGeneration() int32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *Platform_Concern_Topic) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Platform_Concern_Topic) GetPartitionCount() int32 {
	if x != nil {
		return x.PartitionCount
	}
	return 0
}

func (x *Platform_Concern_Topic) GetPartitionStrat() Platform_PartitionStrategy {
	if x != nil {
		return x.PartitionStrat
	}
	return Platform_FNV64_MOD
}

var File_rkcy_platform_proto protoreflect.FileDescriptor

var file_rkcy_platform_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x6b, 0x63, 0x79, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6b, 0x63, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x72, 0x6b,
	0x63, 0x79, 0x2f, 0x61, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda,
	0x08, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x63, 0x65,
	0x72, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x1a, 0xb7, 0x06, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e,
	0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x35, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x1a, 0x9f, 0x03, 0x0a, 0x06, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x66,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x6b,
	0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x63,
	0x65, 0x72, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x66, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x38, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6b,
	0x63, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22, 0x9d, 0x01, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x49, 0x4e, 0x49,
	0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x5f,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x53, 0x55,
	0x4d, 0x45, 0x52, 0x5f, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x46, 0x55, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x45, 0x52, 0x5f, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x54,
	0x4f, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x1a, 0xb5, 0x01, 0x0a, 0x05,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x22, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x47,
	0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x54, 0x43,
	0x48, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x50, 0x45, 0x43, 0x53, 0x10, 0x02, 0x1a, 0x37,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x22, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0d, 0x0a, 0x09,
	0x46, 0x4e, 0x56, 0x36, 0x34, 0x5f, 0x4d, 0x4f, 0x44, 0x10, 0x00, 0x22, 0xaf, 0x01, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb6, 0x01,
	0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x6b, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x63, 0x68, 0x6c, 0x61, 0x6e, 0x6f, 0x72, 0x72, 0x2f,
	0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x6b, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rkcy_platform_proto_rawDescOnce sync.Once
	file_rkcy_platform_proto_rawDescData = file_rkcy_platform_proto_rawDesc
)

func file_rkcy_platform_proto_rawDescGZIP() []byte {
	file_rkcy_platform_proto_rawDescOnce.Do(func() {
		file_rkcy_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_rkcy_platform_proto_rawDescData)
	})
	return file_rkcy_platform_proto_rawDescData
}

var file_rkcy_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_rkcy_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rkcy_platform_proto_goTypes = []interface{}{
	(Platform_PartitionStrategy)(0),    // 0: rkcy.Platform.PartitionStrategy
	(Platform_Concern_Type)(0),         // 1: rkcy.Platform.Concern.Type
	(Platform_Concern_Topics_State)(0), // 2: rkcy.Platform.Concern.Topics.State
	(*Platform)(nil),                   // 3: rkcy.Platform
	(*Program)(nil),                    // 4: rkcy.Program
	(*StorageSystem)(nil),              // 5: rkcy.StorageSystem
	(*Platform_Concern)(nil),           // 6: rkcy.Platform.Concern
	(*Platform_Cluster)(nil),           // 7: rkcy.Platform.Cluster
	(*Platform_Concern_Topics)(nil),    // 8: rkcy.Platform.Concern.Topics
	(*Platform_Concern_Topic)(nil),     // 9: rkcy.Platform.Concern.Topic
	nil,                                // 10: rkcy.Program.TagsEntry
	nil,                                // 11: rkcy.StorageSystem.ConfigEntry
	(*timestamppb.Timestamp)(nil),      // 12: google.protobuf.Timestamp
	(Severity)(0),                      // 13: rkcy.Severity
}
var file_rkcy_platform_proto_depIdxs = []int32{
	6,  // 0: rkcy.Platform.concerns:type_name -> rkcy.Platform.Concern
	7,  // 1: rkcy.Platform.clusters:type_name -> rkcy.Platform.Cluster
	12, // 2: rkcy.Platform.update_time:type_name -> google.protobuf.Timestamp
	10, // 3: rkcy.Program.tags:type_name -> rkcy.Program.TagsEntry
	11, // 4: rkcy.StorageSystem.config:type_name -> rkcy.StorageSystem.ConfigEntry
	1,  // 5: rkcy.Platform.Concern.type:type_name -> rkcy.Platform.Concern.Type
	13, // 6: rkcy.Platform.Concern.log_level:type_name -> rkcy.Severity
	8,  // 7: rkcy.Platform.Concern.topics:type_name -> rkcy.Platform.Concern.Topics
	2,  // 8: rkcy.Platform.Concern.Topics.state:type_name -> rkcy.Platform.Concern.Topics.State
	9,  // 9: rkcy.Platform.Concern.Topics.current:type_name -> rkcy.Platform.Concern.Topic
	9,  // 10: rkcy.Platform.Concern.Topics.future:type_name -> rkcy.Platform.Concern.Topic
	4,  // 11: rkcy.Platform.Concern.Topics.consumer_program:type_name -> rkcy.Program
	0,  // 12: rkcy.Platform.Concern.Topic.partition_strat:type_name -> rkcy.Platform.PartitionStrategy
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_rkcy_platform_proto_init() }
func file_rkcy_platform_proto_init() {
	if File_rkcy_platform_proto != nil {
		return
	}
	file_rkcy_apecs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rkcy_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
		file_rkcy_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Program); i {
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
		file_rkcy_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageSystem); i {
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
		file_rkcy_platform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform_Concern); i {
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
		file_rkcy_platform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform_Cluster); i {
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
		file_rkcy_platform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform_Concern_Topics); i {
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
		file_rkcy_platform_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform_Concern_Topic); i {
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
			RawDescriptor: file_rkcy_platform_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rkcy_platform_proto_goTypes,
		DependencyIndexes: file_rkcy_platform_proto_depIdxs,
		EnumInfos:         file_rkcy_platform_proto_enumTypes,
		MessageInfos:      file_rkcy_platform_proto_msgTypes,
	}.Build()
	File_rkcy_platform_proto = out.File
	file_rkcy_platform_proto_rawDesc = nil
	file_rkcy_platform_proto_goTypes = nil
	file_rkcy_platform_proto_depIdxs = nil
}
