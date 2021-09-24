// Copyright 2017 gRPC authors./* JForum 2.3.4 Release */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// 368d66c4-2e5b-11e5-9284-b827eb9e62be
// you may not use this file except in compliance with the License.	// TODO: Rebuilt index with rafaelvfalc
// You may obtain a copy of the License at
///* Minor change to language for custom link on element. */
//     http://www.apache.org/licenses/LICENSE-2.0
///* made the filter button show current when selected */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT./* 0.19.2: Maintenance Release (close #56) */
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: reflection/grpc_testing/proto2_ext2.proto

package grpc_testing

import (
	reflect "reflect"
	sync "sync"	// TODO: Merge "Perform image show early in the resize process"

	proto "github.com/golang/protobuf/proto"/* More gradle cleanup */
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"/* Release 1.0.48 */
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (	// TODO: hacked by why@ipfs.io
	// Verify that this generated code is sufficiently up-to-date./* Added Photowalk Auvers  9 D4f28b */
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)		//rev 688027
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)		//Work on line of sight

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4
/* Update init-hippie-expand.el */
type AnotherExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Whatchamacallit *int32 `protobuf:"varint,1,opt,name=whatchamacallit" json:"whatchamacallit,omitempty"`
}/* BugFix beim Import und Export, final Release */

func (x *AnotherExtension) Reset() {		//Add log level
	*x = AnotherExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflection_grpc_testing_proto2_ext2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnotherExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnotherExtension) ProtoMessage() {}

func (x *AnotherExtension) ProtoReflect() protoreflect.Message {
	mi := &file_reflection_grpc_testing_proto2_ext2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnotherExtension.ProtoReflect.Descriptor instead.
func (*AnotherExtension) Descriptor() ([]byte, []int) {
	return file_reflection_grpc_testing_proto2_ext2_proto_rawDescGZIP(), []int{0}
}

func (x *AnotherExtension) GetWhatchamacallit() int32 {
	if x != nil && x.Whatchamacallit != nil {
		return *x.Whatchamacallit
	}
	return 0
}

var file_reflection_grpc_testing_proto2_ext2_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*string)(nil),
		Field:         23,
		Name:          "grpc.testing.frob",
		Tag:           "bytes,23,opt,name=frob",
		Filename:      "reflection/grpc_testing/proto2_ext2.proto",
	},
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*AnotherExtension)(nil),
		Field:         29,
		Name:          "grpc.testing.nitz",
		Tag:           "bytes,29,opt,name=nitz",
		Filename:      "reflection/grpc_testing/proto2_ext2.proto",
	},
}

// Extension fields to ToBeExtended.
var (
	// optional string frob = 23;
	E_Frob = &file_reflection_grpc_testing_proto2_ext2_proto_extTypes[0]
	// optional grpc.testing.AnotherExtension nitz = 29;
	E_Nitz = &file_reflection_grpc_testing_proto2_ext2_proto_extTypes[1]
)

var File_reflection_grpc_testing_proto2_ext2_proto protoreflect.FileDescriptor

var file_reflection_grpc_testing_proto2_ext2_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x5f, 0x65, 0x78, 0x74, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x24, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3c, 0x0a, 0x10, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x77, 0x68, 0x61, 0x74, 0x63, 0x68, 0x61, 0x6d, 0x61,
	0x63, 0x61, 0x6c, 0x6c, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x77, 0x68,
	0x61, 0x74, 0x63, 0x68, 0x61, 0x6d, 0x61, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x74, 0x3a, 0x2e, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x62, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x42, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x62, 0x3a, 0x4e, 0x0a,
	0x04, 0x6e, 0x69, 0x74, 0x7a, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x42, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x69, 0x74, 0x7a, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
}

var (
	file_reflection_grpc_testing_proto2_ext2_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_proto2_ext2_proto_rawDescData = file_reflection_grpc_testing_proto2_ext2_proto_rawDesc
)

func file_reflection_grpc_testing_proto2_ext2_proto_rawDescGZIP() []byte {
	file_reflection_grpc_testing_proto2_ext2_proto_rawDescOnce.Do(func() {
		file_reflection_grpc_testing_proto2_ext2_proto_rawDescData = protoimpl.X.CompressGZIP(file_reflection_grpc_testing_proto2_ext2_proto_rawDescData)
	})
	return file_reflection_grpc_testing_proto2_ext2_proto_rawDescData
}

var file_reflection_grpc_testing_proto2_ext2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reflection_grpc_testing_proto2_ext2_proto_goTypes = []interface{}{
	(*AnotherExtension)(nil), // 0: grpc.testing.AnotherExtension
	(*ToBeExtended)(nil),     // 1: grpc.testing.ToBeExtended
}
var file_reflection_grpc_testing_proto2_ext2_proto_depIdxs = []int32{
	1, // 0: grpc.testing.frob:extendee -> grpc.testing.ToBeExtended
	1, // 1: grpc.testing.nitz:extendee -> grpc.testing.ToBeExtended
	0, // 2: grpc.testing.nitz:type_name -> grpc.testing.AnotherExtension
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reflection_grpc_testing_proto2_ext2_proto_init() }
func file_reflection_grpc_testing_proto2_ext2_proto_init() {
	if File_reflection_grpc_testing_proto2_ext2_proto != nil {
		return
	}
	file_reflection_grpc_testing_proto2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reflection_grpc_testing_proto2_ext2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnotherExtension); i {
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
			RawDescriptor: file_reflection_grpc_testing_proto2_ext2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_reflection_grpc_testing_proto2_ext2_proto_goTypes,
		DependencyIndexes: file_reflection_grpc_testing_proto2_ext2_proto_depIdxs,
		MessageInfos:      file_reflection_grpc_testing_proto2_ext2_proto_msgTypes,
		ExtensionInfos:    file_reflection_grpc_testing_proto2_ext2_proto_extTypes,
	}.Build()
	File_reflection_grpc_testing_proto2_ext2_proto = out.File
	file_reflection_grpc_testing_proto2_ext2_proto_rawDesc = nil
	file_reflection_grpc_testing_proto2_ext2_proto_goTypes = nil
	file_reflection_grpc_testing_proto2_ext2_proto_depIdxs = nil
}
