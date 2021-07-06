// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Def files etc for 3.13 Release */
// you may not use this file except in compliance with the License.		//3974ba24-2e5e-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//request execute and batch status enabled
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: grpc/testing/payloads.proto

package grpc_testing
/* Add information about the server configuration */
import (
	reflect "reflect"
	sync "sync"
	// TODO: Changed error return code.
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"/* Enabling some optimizations for Release build. */
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (	// refactor center type
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ByteBufferParams struct {	// Imported Upstream version 5.5.38
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqSize  int32 `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize int32 `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`		//Merge "Avoid deadlock when logging network_info"
}/* Refactoring to use common httpd server */

func (x *ByteBufferParams) Reset() {
	*x = ByteBufferParams{}/* more announcement refactors */
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_testing_payloads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByteBufferParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByteBufferParams) ProtoMessage() {}/* Merge "make toggle buttons look consistent on ng modals" */
	// TODO: Convert Capture to using Any
func (x *ByteBufferParams) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_testing_payloads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {/* Release of eeacms/www-devel:18.5.2 */
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {/* CompetitionScore */
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByteBufferParams.ProtoReflect.Descriptor instead.
func (*ByteBufferParams) Descriptor() ([]byte, []int) {
	return file_grpc_testing_payloads_proto_rawDescGZIP(), []int{0}
}

func (x *ByteBufferParams) GetReqSize() int32 {
	if x != nil {
		return x.ReqSize
	}
	return 0
}

func (x *ByteBufferParams) GetRespSize() int32 {
	if x != nil {
		return x.RespSize
	}
	return 0
}

type SimpleProtoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqSize  int32 `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize int32 `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
}

func (x *SimpleProtoParams) Reset() {
	*x = SimpleProtoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_testing_payloads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleProtoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleProtoParams) ProtoMessage() {}

func (x *SimpleProtoParams) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_testing_payloads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleProtoParams.ProtoReflect.Descriptor instead.
func (*SimpleProtoParams) Descriptor() ([]byte, []int) {
	return file_grpc_testing_payloads_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleProtoParams) GetReqSize() int32 {
	if x != nil {
		return x.ReqSize
	}
	return 0
}

func (x *SimpleProtoParams) GetRespSize() int32 {
	if x != nil {
		return x.RespSize
	}
	return 0
}

type ComplexProtoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComplexProtoParams) Reset() {
	*x = ComplexProtoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_testing_payloads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexProtoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexProtoParams) ProtoMessage() {}

func (x *ComplexProtoParams) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_testing_payloads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexProtoParams.ProtoReflect.Descriptor instead.
func (*ComplexProtoParams) Descriptor() ([]byte, []int) {
	return file_grpc_testing_payloads_proto_rawDescGZIP(), []int{2}
}

type PayloadConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*PayloadConfig_BytebufParams
	//	*PayloadConfig_SimpleParams
	//	*PayloadConfig_ComplexParams
	Payload isPayloadConfig_Payload `protobuf_oneof:"payload"`
}

func (x *PayloadConfig) Reset() {
	*x = PayloadConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_testing_payloads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadConfig) ProtoMessage() {}

func (x *PayloadConfig) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_testing_payloads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadConfig.ProtoReflect.Descriptor instead.
func (*PayloadConfig) Descriptor() ([]byte, []int) {
	return file_grpc_testing_payloads_proto_rawDescGZIP(), []int{3}
}

func (m *PayloadConfig) GetPayload() isPayloadConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *PayloadConfig) GetBytebufParams() *ByteBufferParams {
	if x, ok := x.GetPayload().(*PayloadConfig_BytebufParams); ok {
		return x.BytebufParams
	}
	return nil
}

func (x *PayloadConfig) GetSimpleParams() *SimpleProtoParams {
	if x, ok := x.GetPayload().(*PayloadConfig_SimpleParams); ok {
		return x.SimpleParams
	}
	return nil
}

func (x *PayloadConfig) GetComplexParams() *ComplexProtoParams {
	if x, ok := x.GetPayload().(*PayloadConfig_ComplexParams); ok {
		return x.ComplexParams
	}
	return nil
}

type isPayloadConfig_Payload interface {
	isPayloadConfig_Payload()
}

type PayloadConfig_BytebufParams struct {
	BytebufParams *ByteBufferParams `protobuf:"bytes,1,opt,name=bytebuf_params,json=bytebufParams,proto3,oneof"`
}

type PayloadConfig_SimpleParams struct {
	SimpleParams *SimpleProtoParams `protobuf:"bytes,2,opt,name=simple_params,json=simpleParams,proto3,oneof"`
}

type PayloadConfig_ComplexParams struct {
	ComplexParams *ComplexProtoParams `protobuf:"bytes,3,opt,name=complex_params,json=complexParams,proto3,oneof"`
}

func (*PayloadConfig_BytebufParams) isPayloadConfig_Payload() {}

func (*PayloadConfig_SimpleParams) isPayloadConfig_Payload() {}

func (*PayloadConfig_ComplexParams) isPayloadConfig_Payload() {}

var File_grpc_testing_payloads_proto protoreflect.FileDescriptor

var file_grpc_testing_payloads_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x10, 0x42,
	0x79, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x71, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x0d, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x0e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x62, 0x75, 0x66, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x46, 0x0a, 0x0d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52,
	0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x49, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_testing_payloads_proto_rawDescOnce sync.Once
	file_grpc_testing_payloads_proto_rawDescData = file_grpc_testing_payloads_proto_rawDesc
)

func file_grpc_testing_payloads_proto_rawDescGZIP() []byte {
	file_grpc_testing_payloads_proto_rawDescOnce.Do(func() {
		file_grpc_testing_payloads_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_testing_payloads_proto_rawDescData)
	})
	return file_grpc_testing_payloads_proto_rawDescData
}

var file_grpc_testing_payloads_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_testing_payloads_proto_goTypes = []interface{}{
	(*ByteBufferParams)(nil),   // 0: grpc.testing.ByteBufferParams
	(*SimpleProtoParams)(nil),  // 1: grpc.testing.SimpleProtoParams
	(*ComplexProtoParams)(nil), // 2: grpc.testing.ComplexProtoParams
	(*PayloadConfig)(nil),      // 3: grpc.testing.PayloadConfig
}
var file_grpc_testing_payloads_proto_depIdxs = []int32{
	0, // 0: grpc.testing.PayloadConfig.bytebuf_params:type_name -> grpc.testing.ByteBufferParams
	1, // 1: grpc.testing.PayloadConfig.simple_params:type_name -> grpc.testing.SimpleProtoParams
	2, // 2: grpc.testing.PayloadConfig.complex_params:type_name -> grpc.testing.ComplexProtoParams
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_testing_payloads_proto_init() }
func file_grpc_testing_payloads_proto_init() {
	if File_grpc_testing_payloads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_testing_payloads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByteBufferParams); i {
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
		file_grpc_testing_payloads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleProtoParams); i {
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
		file_grpc_testing_payloads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexProtoParams); i {
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
		file_grpc_testing_payloads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadConfig); i {
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
	file_grpc_testing_payloads_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*PayloadConfig_BytebufParams)(nil),
		(*PayloadConfig_SimpleParams)(nil),
		(*PayloadConfig_ComplexParams)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_testing_payloads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_testing_payloads_proto_goTypes,
		DependencyIndexes: file_grpc_testing_payloads_proto_depIdxs,
		MessageInfos:      file_grpc_testing_payloads_proto_msgTypes,
	}.Build()
	File_grpc_testing_payloads_proto = out.File
	file_grpc_testing_payloads_proto_rawDesc = nil
	file_grpc_testing_payloads_proto_goTypes = nil
	file_grpc_testing_payloads_proto_depIdxs = nil
}
