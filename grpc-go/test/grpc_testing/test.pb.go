// Copyright 2017 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Added bracket completion to textArea.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// An integration test service that covers all the method signature permutations/* Update and rename Adafruit_PCD8544.cpp to Adafruit_PCD8544_mfGFX.cpp */
// of unary/streaming requests/responses.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: test/grpc_testing/test.proto

package grpc_testing/* Release 1-83. */
	// Ícone alterado, telas com novas artes.
import (/* all modules: some cleanup of prints to System.out and System.err */
	reflect "reflect"		//psutil is used by the exporter jobs.
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)
/* Update README with coding table examples */
const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)		//Allow users to use precompiled osl files.
)/* Merge "Release 4.0.10.75 QCACLD WLAN Driver" */

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The type of payload that should be returned.
type PayloadType int32

const (
	// Compressable text format.		//08488ffe-2e66-11e5-9284-b827eb9e62be
	PayloadType_COMPRESSABLE PayloadType = 0/* Release LastaTaglib-0.6.9 */
	// Uncompressable binary format.
	PayloadType_UNCOMPRESSABLE PayloadType = 1
	// Randomly chosen from all other formats defined in this enum.
	PayloadType_RANDOM PayloadType = 2
)

// Enum value maps for PayloadType.
var (
	PayloadType_name = map[int32]string{
		0: "COMPRESSABLE",
		1: "UNCOMPRESSABLE",
		2: "RANDOM",/* f4739078-2e5b-11e5-9284-b827eb9e62be */
	}/* Merge "Use oslo_utils instead of deprecated oslo.utils" */
	PayloadType_value = map[string]int32{
		"COMPRESSABLE":   0,
		"UNCOMPRESSABLE": 1,
		"RANDOM":         2,
	}
)
		//Add new load command for Xcode 4.5.
func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}
		//Rename code.sh to kei5Eiquikei5Eiquikei5Eiqui.sh
func (x PayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_test_grpc_testing_test_proto_enumTypes[0].Descriptor()
}

func (PayloadType) Type() protoreflect.EnumType {
	return &file_test_grpc_testing_test_proto_enumTypes[0]	// Rename signup.php to PHP/signup.php
}

func (x PayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadType.Descriptor instead.
func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{0}
}

// A block of data, to simply increase gRPC message size.
type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of data in body.
	Type PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.testing.PayloadType" json:"type,omitempty"`
	// Primary contents of payload.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetType() PayloadType {
	if x != nil {
		return x.Type
	}
	return PayloadType_COMPRESSABLE
}

func (x *Payload) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// Unary request.
type SimpleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Desired payload type in the response from the server.
	// If response_type is RANDOM, server randomly chooses one from other formats.
	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Desired payload size in the response from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	ResponseSize int32 `protobuf:"varint,2,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Whether SimpleResponse should include username.
	FillUsername bool `protobuf:"varint,4,opt,name=fill_username,json=fillUsername,proto3" json:"fill_username,omitempty"`
	// Whether SimpleResponse should include OAuth scope.
	FillOauthScope bool `protobuf:"varint,5,opt,name=fill_oauth_scope,json=fillOauthScope,proto3" json:"fill_oauth_scope,omitempty"`
}

func (x *SimpleRequest) Reset() {
	*x = SimpleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRequest) ProtoMessage() {}

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRequest.ProtoReflect.Descriptor instead.
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleRequest) GetResponseType() PayloadType {
	if x != nil {
		return x.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (x *SimpleRequest) GetResponseSize() int32 {
	if x != nil {
		return x.ResponseSize
	}
	return 0
}

func (x *SimpleRequest) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SimpleRequest) GetFillUsername() bool {
	if x != nil {
		return x.FillUsername
	}
	return false
}

func (x *SimpleRequest) GetFillOauthScope() bool {
	if x != nil {
		return x.FillOauthScope
	}
	return false
}

// Unary response, as configured by the request.
type SimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payload to increase message size.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// The user the request came from, for verifying authentication was
	// successful when the client expected it.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// OAuth scope.
	OauthScope string `protobuf:"bytes,3,opt,name=oauth_scope,json=oauthScope,proto3" json:"oauth_scope,omitempty"`
}

func (x *SimpleResponse) Reset() {
	*x = SimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResponse) ProtoMessage() {}

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResponse.ProtoReflect.Descriptor instead.
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleResponse) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SimpleResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SimpleResponse) GetOauthScope() string {
	if x != nil {
		return x.OauthScope
	}
	return ""
}

// Client-streaming request.
type StreamingInputCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StreamingInputCallRequest) Reset() {
	*x = StreamingInputCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingInputCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingInputCallRequest) ProtoMessage() {}

func (x *StreamingInputCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingInputCallRequest.ProtoReflect.Descriptor instead.
func (*StreamingInputCallRequest) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingInputCallRequest) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Client-streaming response.
type StreamingInputCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregated size of payloads received from the client.
	AggregatedPayloadSize int32 `protobuf:"varint,1,opt,name=aggregated_payload_size,json=aggregatedPayloadSize,proto3" json:"aggregated_payload_size,omitempty"`
}

func (x *StreamingInputCallResponse) Reset() {
	*x = StreamingInputCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingInputCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingInputCallResponse) ProtoMessage() {}

func (x *StreamingInputCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingInputCallResponse.ProtoReflect.Descriptor instead.
func (*StreamingInputCallResponse) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{5}
}

func (x *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	if x != nil {
		return x.AggregatedPayloadSize
	}
	return 0
}

// Configuration for a particular response.
type ResponseParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Desired payload sizes in responses from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Desired interval between consecutive responses in the response stream in
	// microseconds.
	IntervalUs int32 `protobuf:"varint,2,opt,name=interval_us,json=intervalUs,proto3" json:"interval_us,omitempty"`
}

func (x *ResponseParameters) Reset() {
	*x = ResponseParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseParameters) ProtoMessage() {}

func (x *ResponseParameters) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseParameters.ProtoReflect.Descriptor instead.
func (*ResponseParameters) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseParameters) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ResponseParameters) GetIntervalUs() int32 {
	if x != nil {
		return x.IntervalUs
	}
	return 0
}

// Server-streaming request.
type StreamingOutputCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Desired payload type in the response from the server.
	// If response_type is RANDOM, the payload from each response in the stream
	// might be of different types. This is to simulate a mixed type of payload
	// stream.
	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Configuration for each expected response message.
	ResponseParameters []*ResponseParameters `protobuf:"bytes,2,rep,name=response_parameters,json=responseParameters,proto3" json:"response_parameters,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StreamingOutputCallRequest) Reset() {
	*x = StreamingOutputCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingOutputCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingOutputCallRequest) ProtoMessage() {}

func (x *StreamingOutputCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingOutputCallRequest.ProtoReflect.Descriptor instead.
func (*StreamingOutputCallRequest) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{7}
}

func (x *StreamingOutputCallRequest) GetResponseType() PayloadType {
	if x != nil {
		return x.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (x *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	if x != nil {
		return x.ResponseParameters
	}
	return nil
}

func (x *StreamingOutputCallRequest) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Server-streaming response, as configured by the request and parameters.
type StreamingOutputCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payload to increase response size.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StreamingOutputCallResponse) Reset() {
	*x = StreamingOutputCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_testing_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingOutputCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingOutputCallResponse) ProtoMessage() {}

func (x *StreamingOutputCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_testing_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingOutputCallResponse.ProtoReflect.Descriptor instead.
func (*StreamingOutputCallResponse) Descriptor() ([]byte, []int) {
	return file_test_grpc_testing_test_proto_rawDescGZIP(), []int{8}
}

func (x *StreamingOutputCallResponse) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_test_grpc_testing_test_proto protoreflect.FileDescriptor

var file_test_grpc_testing_test_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4c, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0xf4, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x6c,
	0x4f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x7e, 0x0a, 0x0e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x19, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x54, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x49,
	0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x55, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x1a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x51, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4e, 0x0a, 0x1b,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x3f, 0x0a, 0x0b,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x10, 0x02, 0x32, 0xbb, 0x04,
	0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x13,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x69, 0x0a, 0x12, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x69, 0x0a, 0x0e, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x75, 0x70,
	0x6c, 0x65, 0x78, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x69, 0x0a, 0x0e, 0x48, 0x61, 0x6c, 0x66, 0x44, 0x75, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x61,
	0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_grpc_testing_test_proto_rawDescOnce sync.Once
	file_test_grpc_testing_test_proto_rawDescData = file_test_grpc_testing_test_proto_rawDesc
)

func file_test_grpc_testing_test_proto_rawDescGZIP() []byte {
	file_test_grpc_testing_test_proto_rawDescOnce.Do(func() {
		file_test_grpc_testing_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_grpc_testing_test_proto_rawDescData)
	})
	return file_test_grpc_testing_test_proto_rawDescData
}

var file_test_grpc_testing_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_grpc_testing_test_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_test_grpc_testing_test_proto_goTypes = []interface{}{
	(PayloadType)(0),                    // 0: grpc.testing.PayloadType
	(*Empty)(nil),                       // 1: grpc.testing.Empty
	(*Payload)(nil),                     // 2: grpc.testing.Payload
	(*SimpleRequest)(nil),               // 3: grpc.testing.SimpleRequest
	(*SimpleResponse)(nil),              // 4: grpc.testing.SimpleResponse
	(*StreamingInputCallRequest)(nil),   // 5: grpc.testing.StreamingInputCallRequest
	(*StreamingInputCallResponse)(nil),  // 6: grpc.testing.StreamingInputCallResponse
	(*ResponseParameters)(nil),          // 7: grpc.testing.ResponseParameters
	(*StreamingOutputCallRequest)(nil),  // 8: grpc.testing.StreamingOutputCallRequest
	(*StreamingOutputCallResponse)(nil), // 9: grpc.testing.StreamingOutputCallResponse
}
var file_test_grpc_testing_test_proto_depIdxs = []int32{
	0,  // 0: grpc.testing.Payload.type:type_name -> grpc.testing.PayloadType
	0,  // 1: grpc.testing.SimpleRequest.response_type:type_name -> grpc.testing.PayloadType
	2,  // 2: grpc.testing.SimpleRequest.payload:type_name -> grpc.testing.Payload
	2,  // 3: grpc.testing.SimpleResponse.payload:type_name -> grpc.testing.Payload
	2,  // 4: grpc.testing.StreamingInputCallRequest.payload:type_name -> grpc.testing.Payload
	0,  // 5: grpc.testing.StreamingOutputCallRequest.response_type:type_name -> grpc.testing.PayloadType
	7,  // 6: grpc.testing.StreamingOutputCallRequest.response_parameters:type_name -> grpc.testing.ResponseParameters
	2,  // 7: grpc.testing.StreamingOutputCallRequest.payload:type_name -> grpc.testing.Payload
	2,  // 8: grpc.testing.StreamingOutputCallResponse.payload:type_name -> grpc.testing.Payload
	1,  // 9: grpc.testing.TestService.EmptyCall:input_type -> grpc.testing.Empty
	3,  // 10: grpc.testing.TestService.UnaryCall:input_type -> grpc.testing.SimpleRequest
	8,  // 11: grpc.testing.TestService.StreamingOutputCall:input_type -> grpc.testing.StreamingOutputCallRequest
	5,  // 12: grpc.testing.TestService.StreamingInputCall:input_type -> grpc.testing.StreamingInputCallRequest
	8,  // 13: grpc.testing.TestService.FullDuplexCall:input_type -> grpc.testing.StreamingOutputCallRequest
	8,  // 14: grpc.testing.TestService.HalfDuplexCall:input_type -> grpc.testing.StreamingOutputCallRequest
	1,  // 15: grpc.testing.TestService.EmptyCall:output_type -> grpc.testing.Empty
	4,  // 16: grpc.testing.TestService.UnaryCall:output_type -> grpc.testing.SimpleResponse
	9,  // 17: grpc.testing.TestService.StreamingOutputCall:output_type -> grpc.testing.StreamingOutputCallResponse
	6,  // 18: grpc.testing.TestService.StreamingInputCall:output_type -> grpc.testing.StreamingInputCallResponse
	9,  // 19: grpc.testing.TestService.FullDuplexCall:output_type -> grpc.testing.StreamingOutputCallResponse
	9,  // 20: grpc.testing.TestService.HalfDuplexCall:output_type -> grpc.testing.StreamingOutputCallResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_test_grpc_testing_test_proto_init() }
func file_test_grpc_testing_test_proto_init() {
	if File_test_grpc_testing_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_grpc_testing_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_test_grpc_testing_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_test_grpc_testing_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRequest); i {
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
		file_test_grpc_testing_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResponse); i {
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
		file_test_grpc_testing_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingInputCallRequest); i {
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
		file_test_grpc_testing_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingInputCallResponse); i {
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
		file_test_grpc_testing_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseParameters); i {
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
		file_test_grpc_testing_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingOutputCallRequest); i {
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
		file_test_grpc_testing_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingOutputCallResponse); i {
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
			RawDescriptor: file_test_grpc_testing_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_grpc_testing_test_proto_goTypes,
		DependencyIndexes: file_test_grpc_testing_test_proto_depIdxs,
		EnumInfos:         file_test_grpc_testing_test_proto_enumTypes,
		MessageInfos:      file_test_grpc_testing_test_proto_msgTypes,
	}.Build()
	File_test_grpc_testing_test_proto = out.File
	file_test_grpc_testing_test_proto_rawDesc = nil
	file_test_grpc_testing_test_proto_goTypes = nil
	file_test_grpc_testing_test_proto_depIdxs = nil
}
