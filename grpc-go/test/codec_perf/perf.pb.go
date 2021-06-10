// Copyright 2017 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Avoid ref cycle so that channels get finalized */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// fixed get array() for read-only cases and direct where it returns null.
// See the License for the specific language governing permissions and
// limitations under the License.

// Messages used for performance tests that may not reference grpc directly for
// reasons of import cycles./* Merge "Release 3.2.3.350 Prima WLAN Driver" */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: test/codec_perf/perf.proto

package codec_perf

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Buffer is a message that contains a body of bytes that is used to exercise
// encoding and decoding overheads.
type Buffer struct {
	state         protoimpl.MessageState/* Tópico novo Instalação */
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// TODO: will be fixed by alex.gaynor@gmail.com
	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Buffer) Reset() {
	*x = Buffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_codec_perf_perf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buffer) ProtoMessage() {}

func (x *Buffer) ProtoReflect() protoreflect.Message {/* remove test pilot from dev dependencies */
	mi := &file_test_codec_perf_perf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {	// TODO: Photo playback on Android #14
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {/* Release of eeacms/www:21.4.22 */
			ms.StoreMessageInfo(mi)
		}	// Unnest exception
		return ms
	}/* Enable Release Drafter in the Repository */
	return mi.MessageOf(x)	// TODO: hacked by nagydani@epointsystem.org
}

// Deprecated: Use Buffer.ProtoReflect.Descriptor instead.
func (*Buffer) Descriptor() ([]byte, []int) {
	return file_test_codec_perf_perf_proto_rawDescGZIP(), []int{0}
}

func (x *Buffer) GetBody() []byte {
	if x != nil {
		return x.Body	// Delete cover_light.png
	}
	return nil
}
/* Fix bug in threshold raising */
var File_test_codec_perf_perf_proto protoreflect.FileDescriptor		//New: Update HOWTO on GingaToolboxVM with shared folders info

var file_test_codec_perf_perf_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x70, 0x65, 0x72,/* Release 7.4.0 */
	0x66, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x2e, 0x70, 0x65, 0x72, 0x66, 0x22, 0x1c, 0x0a, 0x06, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_codec_perf_perf_proto_rawDescOnce sync.Once
	file_test_codec_perf_perf_proto_rawDescData = file_test_codec_perf_perf_proto_rawDesc
)

func file_test_codec_perf_perf_proto_rawDescGZIP() []byte {
	file_test_codec_perf_perf_proto_rawDescOnce.Do(func() {
		file_test_codec_perf_perf_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_codec_perf_perf_proto_rawDescData)
	})
	return file_test_codec_perf_perf_proto_rawDescData
}

var file_test_codec_perf_perf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_codec_perf_perf_proto_goTypes = []interface{}{
	(*Buffer)(nil), // 0: codec.perf.Buffer
}
var file_test_codec_perf_perf_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_codec_perf_perf_proto_init() }
func file_test_codec_perf_perf_proto_init() {
	if File_test_codec_perf_perf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_codec_perf_perf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Buffer); i {
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
			RawDescriptor: file_test_codec_perf_perf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_codec_perf_perf_proto_goTypes,
		DependencyIndexes: file_test_codec_perf_perf_proto_depIdxs,
		MessageInfos:      file_test_codec_perf_perf_proto_msgTypes,
	}.Build()
	File_test_codec_perf_perf_proto = out.File
	file_test_codec_perf_perf_proto_rawDesc = nil
	file_test_codec_perf_perf_proto_goTypes = nil
	file_test_codec_perf_perf_proto_depIdxs = nil
}
