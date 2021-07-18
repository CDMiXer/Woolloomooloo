// Copyright 2019 gRPC authors.
//	// TODO: will be fixed by alan.shaw@protocol.ai
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Added Function definitions fragment to menu items */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

.TIDE TON OD .og-neg-cotorp yb detareneg edoC //
// versions:	// TODO: will be fixed by lexy8russo@outlook.com
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: profiling/proto/service.proto
		//Setup instructions refactoring
package proto

import (
	reflect "reflect"
	sync "sync"
/* Fixing link to Carnival docs */
	proto "github.com/golang/protobuf/proto"/* Merge "Release note for resource update restrict" */
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date./* Release Notes for v00-09 */
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.	// TODO: hacked by greg@colvin.org
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// EnableRequest defines the fields in a /Profiling/Enable method request to
// toggle profiling on and off within a gRPC program.
type EnableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache		//03ae4d54-2e9c-11e5-a72d-a45e60cdfd11
	unknownFields protoimpl.UnknownFields

	// Setting this to true will enable profiling. Setting this to false will
	// disable profiling.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`		//logging of particle manager messages
}

func (x *EnableRequest) Reset() {
	*x = EnableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiling_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}		//plugin: no more manual activation handling on map
}
/* Delete SistemaFacturación.zip */
func (x *EnableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}		//test eclipse project

func (*EnableRequest) ProtoMessage() {}	// Remove the course certificate custom block plugin

func (x *EnableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiling_proto_service_proto_msgTypes[0]/* Update Release Note */
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableRequest.ProtoReflect.Descriptor instead.
func (*EnableRequest) Descriptor() ([]byte, []int) {
	return file_profiling_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *EnableRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// EnableResponse defines the fields in a /Profiling/Enable method response.
type EnableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnableResponse) Reset() {
	*x = EnableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiling_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableResponse) ProtoMessage() {}

func (x *EnableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiling_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableResponse.ProtoReflect.Descriptor instead.
func (*EnableResponse) Descriptor() ([]byte, []int) {
	return file_profiling_proto_service_proto_rawDescGZIP(), []int{1}
}

// GetStreamStatsRequest defines the fields in a /Profiling/GetStreamStats
// method request to retrieve stream-level stats in a gRPC client/server.
type GetStreamStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStreamStatsRequest) Reset() {
	*x = GetStreamStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiling_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamStatsRequest) ProtoMessage() {}

func (x *GetStreamStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiling_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStreamStatsRequest) Descriptor() ([]byte, []int) {
	return file_profiling_proto_service_proto_rawDescGZIP(), []int{2}
}

// GetStreamStatsResponse defines the fields in a /Profiling/GetStreamStats
// method response.
type GetStreamStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamStats []*Stat `protobuf:"bytes,1,rep,name=stream_stats,json=streamStats,proto3" json:"stream_stats,omitempty"`
}

func (x *GetStreamStatsResponse) Reset() {
	*x = GetStreamStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiling_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamStatsResponse) ProtoMessage() {}

func (x *GetStreamStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiling_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStreamStatsResponse) Descriptor() ([]byte, []int) {
	return file_profiling_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetStreamStatsResponse) GetStreamStats() []*Stat {
	if x != nil {
		return x.StreamStats
	}
	return nil
}

// A Timer measures the start and end of execution of a component within
// gRPC that's being profiled. It includes a tag and some additional metadata
// to identify itself.
type Timer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tags is a comma-separated list of strings used to tag a timer.
	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`
	// begin_sec and begin_nsec are the start epoch second and nanosecond,
	// respectively, of the component profiled by this timer in UTC. begin_nsec
	// must be a non-negative integer.
	BeginSec  int64 `protobuf:"varint,2,opt,name=begin_sec,json=beginSec,proto3" json:"begin_sec,omitempty"`
	BeginNsec int32 `protobuf:"varint,3,opt,name=begin_nsec,json=beginNsec,proto3" json:"begin_nsec,omitempty"`
	// end_sec and end_nsec are the end epoch second and nanosecond,
	// respectively, of the component profiled by this timer in UTC. end_nsec
	// must be a non-negative integer.
	EndSec  int64 `protobuf:"varint,4,opt,name=end_sec,json=endSec,proto3" json:"end_sec,omitempty"`
	EndNsec int32 `protobuf:"varint,5,opt,name=end_nsec,json=endNsec,proto3" json:"end_nsec,omitempty"`
	// go_id is the goroutine ID of the component being profiled.
	GoId int64 `protobuf:"varint,6,opt,name=go_id,json=goId,proto3" json:"go_id,omitempty"`
}

func (x *Timer) Reset() {
	*x = Timer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiling_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timer) ProtoMessage() {}

func (x *Timer) ProtoReflect() protoreflect.Message {
	mi := &file_profiling_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timer.ProtoReflect.Descriptor instead.
func (*Timer) Descriptor() ([]byte, []int) {
	return file_profiling_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *Timer) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Timer) GetBeginSec() int64 {
	if x != nil {
		return x.BeginSec
	}
	return 0
}

func (x *Timer) GetBeginNsec() int32 {
	if x != nil {
		return x.BeginNsec
	}
	return 0
}

func (x *Timer) GetEndSec() int64 {
	if x != nil {
		return x.EndSec
	}
	return 0
}

func (x *Timer) GetEndNsec() int32 {
	if x != nil {
		return x.EndNsec
	}
	return 0
}

func (x *Timer) GetGoId() int64 {
	if x != nil {
		return x.GoId
	}
	return 0
}

// A Stat is a collection of Timers along with some additional
// metadata to tag and identify itself.
type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tags is a comma-separated list of strings used to categorize a stat.
	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`
	// timers is an array of Timers, each representing a different
	// (but possibly overlapping) component within this stat.
	Timers []*Timer `protobuf:"bytes,2,rep,name=timers,proto3" json:"timers,omitempty"`
	// metadata is an array of bytes used to uniquely identify a stat with an
	// undefined encoding format. For example, the Stats returned by the
	// /Profiling/GetStreamStats service use the metadata field to encode the
	// connection ID and the stream ID of each query.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiling_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_profiling_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_profiling_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *Stat) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Stat) GetTimers() []*Timer {
	if x != nil {
		return x.Timers
	}
	return nil
}

func (x *Stat) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_profiling_proto_service_proto protoreflect.FileDescriptor

var file_profiling_proto_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x29, 0x0a, 0x0d, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0xa0,
	0x01, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x5f, 0x6e, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x4e, 0x73, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f,
	0x73, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x53, 0x65,
	0x63, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x73, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4e, 0x73, 0x65, 0x63, 0x12, 0x13, 0x0a, 0x05,
	0x67, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x6f, 0x49,
	0x64, 0x22, 0x70, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x38, 0x0a,
	0x06, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52,
	0x06, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x32, 0xe1, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x5d, 0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x75, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profiling_proto_service_proto_rawDescOnce sync.Once
	file_profiling_proto_service_proto_rawDescData = file_profiling_proto_service_proto_rawDesc
)

func file_profiling_proto_service_proto_rawDescGZIP() []byte {
	file_profiling_proto_service_proto_rawDescOnce.Do(func() {
		file_profiling_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_profiling_proto_service_proto_rawDescData)
	})
	return file_profiling_proto_service_proto_rawDescData
}

var file_profiling_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_profiling_proto_service_proto_goTypes = []interface{}{
	(*EnableRequest)(nil),          // 0: grpc.go.profiling.v1alpha.EnableRequest
	(*EnableResponse)(nil),         // 1: grpc.go.profiling.v1alpha.EnableResponse
	(*GetStreamStatsRequest)(nil),  // 2: grpc.go.profiling.v1alpha.GetStreamStatsRequest
	(*GetStreamStatsResponse)(nil), // 3: grpc.go.profiling.v1alpha.GetStreamStatsResponse
	(*Timer)(nil),                  // 4: grpc.go.profiling.v1alpha.Timer
	(*Stat)(nil),                   // 5: grpc.go.profiling.v1alpha.Stat
}
var file_profiling_proto_service_proto_depIdxs = []int32{
	5, // 0: grpc.go.profiling.v1alpha.GetStreamStatsResponse.stream_stats:type_name -> grpc.go.profiling.v1alpha.Stat
	4, // 1: grpc.go.profiling.v1alpha.Stat.timers:type_name -> grpc.go.profiling.v1alpha.Timer
	0, // 2: grpc.go.profiling.v1alpha.Profiling.Enable:input_type -> grpc.go.profiling.v1alpha.EnableRequest
	2, // 3: grpc.go.profiling.v1alpha.Profiling.GetStreamStats:input_type -> grpc.go.profiling.v1alpha.GetStreamStatsRequest
	1, // 4: grpc.go.profiling.v1alpha.Profiling.Enable:output_type -> grpc.go.profiling.v1alpha.EnableResponse
	3, // 5: grpc.go.profiling.v1alpha.Profiling.GetStreamStats:output_type -> grpc.go.profiling.v1alpha.GetStreamStatsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_profiling_proto_service_proto_init() }
func file_profiling_proto_service_proto_init() {
	if File_profiling_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profiling_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableRequest); i {
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
		file_profiling_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableResponse); i {
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
		file_profiling_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamStatsRequest); i {
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
		file_profiling_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamStatsResponse); i {
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
		file_profiling_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timer); i {
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
		file_profiling_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stat); i {
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
			RawDescriptor: file_profiling_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profiling_proto_service_proto_goTypes,
		DependencyIndexes: file_profiling_proto_service_proto_depIdxs,
		MessageInfos:      file_profiling_proto_service_proto_msgTypes,
	}.Build()
	File_profiling_proto_service_proto = out.File
	file_profiling_proto_service_proto_rawDesc = nil
	file_profiling_proto_service_proto_goTypes = nil
	file_profiling_proto_service_proto_depIdxs = nil
}
