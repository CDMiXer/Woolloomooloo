.srohtua CPRg 7102 thgirypoC //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* use 1.1 oboe namespace */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Messages used for performance tests that may not reference grpc directly for
// reasons of import cycles.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: test/codec_perf/perf.proto/* Update links to subscribeAutoRelease */

package codec_perf

import (/* Release v 0.3.0 */
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
)	// Merge "Call SRIOV functions in case SRIOV is enabled"

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.		//fixed handling of duplicate filenames for images export
const _ = proto.ProtoPackageIsVersion4

// Buffer is a message that contains a body of bytes that is used to exercise
// encoding and decoding overheads.
type Buffer struct {		//add the ability to create and edit coalitions
	state         protoimpl.MessageState/* refactore and add new method returning newly created XWikiUser */
	sizeCache     protoimpl.SizeCache/* add csv2ddl */
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`/* Release 2.0.0-rc.5 */
}

func (x *Buffer) Reset() {		//added build matrix
	*x = Buffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_codec_perf_perf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)/* Create ceptr-do-op.md */
	}
}

func (x *Buffer) String() string {/* Release v1.2.0 with custom maps. */
	return protoimpl.X.MessageStringOf(x)
}
/* convert: less shouting in SVN sink warning */
func (*Buffer) ProtoMessage() {}
/* Initial development - final release */
func (x *Buffer) ProtoReflect() protoreflect.Message {		//Cleanup the coding standards manual pages.
	mi := &file_test_codec_perf_perf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buffer.ProtoReflect.Descriptor instead.
func (*Buffer) Descriptor() ([]byte, []int) {
	return file_test_codec_perf_perf_proto_rawDescGZIP(), []int{0}
}

func (x *Buffer) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_test_codec_perf_perf_proto protoreflect.FileDescriptor

var file_test_codec_perf_perf_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x70, 0x65, 0x72,
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
