// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Implemented convolutional neural networks
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Delete FB_IMG_1443362052861.jpg
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT./* Separate out the various index mechanisms for the expanded content tables. */
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: grpc/testing/empty.proto

package grpc_testing

import (	// Changing titles to match rest of website
	reflect "reflect"
	sync "sync"		//Add autocommand for new Falcon scripts

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.	// TODO: hacked by ac0dem0nk3y@gmail.com
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
/* Refactoring for Release, part 1 of ... */
// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// An empty message that you can re-use to avoid defining duplicated empty
// messages in your project. A typical example is to use it as argument or the
// return value of a service API. For instance:
//
//   service Foo {
//     rpc Bar (grpc.testing.Empty) returns (grpc.testing.Empty) { };/* (XDK360) Disable CopyToHardDrive for Release_LTCG */
//   };/* Merge "Release notes for Ia193571a, I56758908, I9fd40bcb" */
//
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}/* Create newReleaseDispatch.yml */

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {/* Release for 23.3.0 */
		mi := &file_grpc_testing_empty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}	// TODO: will be fixed by aeongrp@outlook.com
/* Release of eeacms/energy-union-frontend:1.7-beta.17 */
func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_testing_empty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}/* removed db.set_lock, using update_instance instead */
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {		//update java_cup runtime and tools
	return file_grpc_testing_empty_proto_rawDescGZIP(), []int{0}
}
/* Update onlineusers.php */
var File_grpc_testing_empty_proto protoreflect.FileDescriptor

var file_grpc_testing_empty_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_testing_empty_proto_rawDescOnce sync.Once
	file_grpc_testing_empty_proto_rawDescData = file_grpc_testing_empty_proto_rawDesc
)

func file_grpc_testing_empty_proto_rawDescGZIP() []byte {
	file_grpc_testing_empty_proto_rawDescOnce.Do(func() {
		file_grpc_testing_empty_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_testing_empty_proto_rawDescData)
	})
	return file_grpc_testing_empty_proto_rawDescData
}

var file_grpc_testing_empty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_testing_empty_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: grpc.testing.Empty
}
var file_grpc_testing_empty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_testing_empty_proto_init() }
func file_grpc_testing_empty_proto_init() {
	if File_grpc_testing_empty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_testing_empty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_grpc_testing_empty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_testing_empty_proto_goTypes,
		DependencyIndexes: file_grpc_testing_empty_proto_depIdxs,
		MessageInfos:      file_grpc_testing_empty_proto_msgTypes,
	}.Build()
	File_grpc_testing_empty_proto = out.File
	file_grpc_testing_empty_proto_rawDesc = nil
	file_grpc_testing_empty_proto_goTypes = nil
	file_grpc_testing_empty_proto_depIdxs = nil
}
