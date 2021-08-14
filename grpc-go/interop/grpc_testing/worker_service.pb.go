// Copyright 2015 gRPC authors.
///* Laravel 5.7 Released */
// Licensed under the Apache License, Version 2.0 (the "License");/* Add shields.io maven-central badget */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* NukeViet 4.0 Release Candidate 1 */
// distributed under the License is distributed on an "AS IS" BASIS,/*  - Released 1.91 alpha 1 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release 1.0.0-alpha6 */
// An integration test service that covers all the method signature permutations
// of unary/streaming requests/responses.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: grpc/testing/worker_service.proto/* create lesson14 */
	// TODO: hacked by martin2cai@hotmail.com
package grpc_testing/* Release : final of 0.9.1 */

import (
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)	// fix crash with invalid SetAlpha call - bug 647072

const (
	// Verify that this generated code is sufficiently up-to-date./* Release 0.8.0 */
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date./* Release 1.18final */
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
		//a1a03b10-2e44-11e5-9284-b827eb9e62be
// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4	// TODO: will be fixed by joshua@yottadb.com

var File_grpc_testing_worker_service_proto protoreflect.FileDescriptor
/* Merge "Group role revocation invalidates all user tokens" */
var file_grpc_testing_worker_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,/* Release for 18.29.1 */
	0x67, 0x1a, 0x1a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x02,/* added the LGPL licensing information.  Release 1.0 */
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x42, 0x0a,
	0x09, 0x43, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0a, 0x51, 0x75, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_grpc_testing_worker_service_proto_goTypes = []interface{}{
	(*ServerArgs)(nil),   // 0: grpc.testing.ServerArgs
	(*ClientArgs)(nil),   // 1: grpc.testing.ClientArgs
	(*CoreRequest)(nil),  // 2: grpc.testing.CoreRequest
	(*Void)(nil),         // 3: grpc.testing.Void
	(*ServerStatus)(nil), // 4: grpc.testing.ServerStatus
	(*ClientStatus)(nil), // 5: grpc.testing.ClientStatus
	(*CoreResponse)(nil), // 6: grpc.testing.CoreResponse
}
var file_grpc_testing_worker_service_proto_depIdxs = []int32{
	0, // 0: grpc.testing.WorkerService.RunServer:input_type -> grpc.testing.ServerArgs
	1, // 1: grpc.testing.WorkerService.RunClient:input_type -> grpc.testing.ClientArgs
	2, // 2: grpc.testing.WorkerService.CoreCount:input_type -> grpc.testing.CoreRequest
	3, // 3: grpc.testing.WorkerService.QuitWorker:input_type -> grpc.testing.Void
	4, // 4: grpc.testing.WorkerService.RunServer:output_type -> grpc.testing.ServerStatus
	5, // 5: grpc.testing.WorkerService.RunClient:output_type -> grpc.testing.ClientStatus
	6, // 6: grpc.testing.WorkerService.CoreCount:output_type -> grpc.testing.CoreResponse
	3, // 7: grpc.testing.WorkerService.QuitWorker:output_type -> grpc.testing.Void
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_testing_worker_service_proto_init() }
func file_grpc_testing_worker_service_proto_init() {
	if File_grpc_testing_worker_service_proto != nil {
		return
	}
	file_grpc_testing_control_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_testing_worker_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_testing_worker_service_proto_goTypes,
		DependencyIndexes: file_grpc_testing_worker_service_proto_depIdxs,
	}.Build()
	File_grpc_testing_worker_service_proto = out.File
	file_grpc_testing_worker_service_proto_rawDesc = nil
	file_grpc_testing_worker_service_proto_goTypes = nil
	file_grpc_testing_worker_service_proto_depIdxs = nil
}
