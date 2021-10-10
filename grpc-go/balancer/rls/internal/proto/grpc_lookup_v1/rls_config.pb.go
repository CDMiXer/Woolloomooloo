// Copyright 2020 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");/* TreeChopper 1.0 Release, REQUEST-DarkriftX */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* spring provided */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0/* Release areca-5.4 */
// source: grpc/lookup/v1/rls_config.proto

package grpc_lookup_v1/* [IMP] Github style Release */

import (
	reflect "reflect"
	sync "sync"/* upgrade support annotations to 23.3.0 */
	// TODO: hacked by lexy8russo@outlook.com
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"/* ivi - fix map action target */
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
		//Added .zip capability
// Extract a key based on a given name (e.g. header name or query parameter
// name).  The name must match one of the names listed in the "name" field.  If
// the "required_match" field is true, one of the specified names must be
// present for the keybuilder to match.
type NameMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields/* fixed PhReleaseQueuedLockExclusiveFast */

	// The name that will be used in the RLS key_map to refer to this value.
	// If required_match is true, you may omit this field or set it to an empty
	// string, in which case the matcher will require a match, but won't update
	// the key_map.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Ordered list of names (headers or query parameter names) that can supply
	// this value; the first one with a non-empty value is used.		//use weissi.github.io
	Names []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	// If true, make this extraction required; the key builder will not match
	// if no value is found.
	RequiredMatch bool `protobuf:"varint,3,opt,name=required_match,json=requiredMatch,proto3" json:"required_match,omitempty"`	// fix xwi parsing
}

func (x *NameMatcher) Reset() {
	*x = NameMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))/* Release 2.0.0-RC4 */
		ms.StoreMessageInfo(mi)
	}
}/* #379 find after submit */

func (x *NameMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)	// Added loading of the history using readline/history library.
}
/* Released gem 2.1.3 */
func (*NameMatcher) ProtoMessage() {}

func (x *NameMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {		//rm Core load object trace log messages
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameMatcher.ProtoReflect.Descriptor instead.
func (*NameMatcher) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_config_proto_rawDescGZIP(), []int{0}
}

func (x *NameMatcher) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NameMatcher) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *NameMatcher) GetRequiredMatch() bool {
	if x != nil {
		return x.RequiredMatch
	}
	return false
}

// A GrpcKeyBuilder applies to a given gRPC service, name, and headers.
type GrpcKeyBuilder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names     []*GrpcKeyBuilder_Name    `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	ExtraKeys *GrpcKeyBuilder_ExtraKeys `protobuf:"bytes,3,opt,name=extra_keys,json=extraKeys,proto3" json:"extra_keys,omitempty"`
	// Extract keys from all listed headers.
	// For gRPC, it is an error to specify "required_match" on the NameMatcher
	// protos.
	Headers []*NameMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// You can optionally set one or more specific key/value pairs to be added to
	// the key_map.  This can be useful to identify which builder built the key,
	// for example if you are suppressing the actual method, but need to
	// separately cache and request all the matched methods.
	ConstantKeys map[string]string `protobuf:"bytes,4,rep,name=constant_keys,json=constantKeys,proto3" json:"constant_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GrpcKeyBuilder) Reset() {
	*x = GrpcKeyBuilder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcKeyBuilder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcKeyBuilder) ProtoMessage() {}

func (x *GrpcKeyBuilder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcKeyBuilder.ProtoReflect.Descriptor instead.
func (*GrpcKeyBuilder) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_config_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcKeyBuilder) GetNames() []*GrpcKeyBuilder_Name {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *GrpcKeyBuilder) GetExtraKeys() *GrpcKeyBuilder_ExtraKeys {
	if x != nil {
		return x.ExtraKeys
	}
	return nil
}

func (x *GrpcKeyBuilder) GetHeaders() []*NameMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *GrpcKeyBuilder) GetConstantKeys() map[string]string {
	if x != nil {
		return x.ConstantKeys
	}
	return nil
}

// An HttpKeyBuilder applies to a given HTTP URL and headers.
//
// Path and host patterns use the matching syntax from gRPC transcoding to
// extract named key/value pairs from the path and host components of the URL:
// https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
//
// It is invalid to specify the same key name in multiple places in a pattern.
//
// For a service where the project id can be expressed either as a subdomain or
// in the path, separate HttpKeyBuilders must be used:
//     host_pattern: 'example.com' path_pattern: '/{id}/{object}/**'
//     host_pattern: '{id}.example.com' path_pattern: '/{object}/**'
// If the host is exactly 'example.com', the first path segment will be used as
// the id and the second segment as the object. If the host has a subdomain, the
// subdomain will be used as the id and the first segment as the object. If
// neither pattern matches, no keys will be extracted.
type HttpKeyBuilder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host_pattern is an ordered list of host template patterns for the desired
	// value.  If any host_pattern values are specified, then at least one must
	// match, and the last one wins and sets any specified variables.  A host
	// consists of labels separated by dots. Each label is matched against the
	// label in the pattern as follows:
	//   - "*": Matches any single label.
	//   - "**": Matches zero or more labels (first or last part of host only).
	//   - "{<name>=...}": One or more label capture, where "..." can be any
	//      template that does not include a capture.
	//   - "{<name>}": A single label capture. Identical to {<name>=*}.
	//
	// Examples:
	//   - "example.com": Only applies to the exact host example.com.
	//   - "*.example.com": Matches subdomains of example.com.
	//   - "**.example.com": matches example.com, and all levels of subdomains.
	//   - "{project}.example.com": Extracts the third level subdomain.
	//   - "{project=**}.example.com": Extracts the third level+ subdomains.
	//   - "{project=**}": Extracts the entire host.
	HostPatterns []string `protobuf:"bytes,1,rep,name=host_patterns,json=hostPatterns,proto3" json:"host_patterns,omitempty"`
	// path_pattern is an ordered list of path template patterns for the desired
	// value.  If any path_pattern values are specified, then at least one must
	// match, and the last one wins and sets any specified variables.  A path
	// consists of segments separated by slashes. Each segment is matched against
	// the segment in the pattern as follows:
	//   - "*": Matches any single segment.
	//   - "**": Matches zero or more segments (first or last part of path only).
	//   - "{<name>=...}": One or more segment capture, where "..." can be any
	//      template that does not include a capture.
	//   - "{<name>}": A single segment capture. Identical to {<name>=*}.
	// A custom method may also be specified by appending ":" and the custom
	// method name or "*" to indicate any custom method (including no custom
	// method).  For example, "/*/projects/{project_id}/**:*" extracts
	// `{project_id}` for any version, resource and custom method that includes
	// it.  By default, any custom method will be matched.
	//
	// Examples:
	//   - "/v1/{name=messages/*}": extracts a name like "messages/12345".
	//   - "/v1/messages/{message_id}": extracts a message_id like "12345".
	//   - "/v1/users/{user_id}/messages/{message_id}": extracts two key values.
	PathPatterns []string `protobuf:"bytes,2,rep,name=path_patterns,json=pathPatterns,proto3" json:"path_patterns,omitempty"`
	// List of query parameter names to try to match.
	// For example: ["parent", "name", "resource.name"]
	// We extract all the specified query_parameters (case-sensitively).  If any
	// are marked as "required_match" and are not present, this keybuilder fails
	// to match.  If a given parameter appears multiple times (?foo=a&foo=b) we
	// will report it as a comma-separated string (foo=a,b).
	QueryParameters []*NameMatcher `protobuf:"bytes,3,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	// List of headers to try to match.
	// We extract all the specified header values (case-insensitively).  If any
	// are marked as "required_match" and are not present, this keybuilder fails
	// to match.  If a given header appears multiple times in the request we will
	// report it as a comma-separated string, in standard HTTP fashion.
	Headers []*NameMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// You can optionally set one or more specific key/value pairs to be added to
	// the key_map.  This can be useful to identify which builder built the key,
	// for example if you are suppressing a lot of information from the URL, but
	// need to separately cache and request URLs with that content.
	ConstantKeys map[string]string `protobuf:"bytes,5,rep,name=constant_keys,json=constantKeys,proto3" json:"constant_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HttpKeyBuilder) Reset() {
	*x = HttpKeyBuilder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpKeyBuilder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpKeyBuilder) ProtoMessage() {}

func (x *HttpKeyBuilder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpKeyBuilder.ProtoReflect.Descriptor instead.
func (*HttpKeyBuilder) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_config_proto_rawDescGZIP(), []int{2}
}

func (x *HttpKeyBuilder) GetHostPatterns() []string {
	if x != nil {
		return x.HostPatterns
	}
	return nil
}

func (x *HttpKeyBuilder) GetPathPatterns() []string {
	if x != nil {
		return x.PathPatterns
	}
	return nil
}

func (x *HttpKeyBuilder) GetQueryParameters() []*NameMatcher {
	if x != nil {
		return x.QueryParameters
	}
	return nil
}

func (x *HttpKeyBuilder) GetHeaders() []*NameMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HttpKeyBuilder) GetConstantKeys() map[string]string {
	if x != nil {
		return x.ConstantKeys
	}
	return nil
}

type RouteLookupConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ordered specifications for constructing keys for HTTP requests.  Last
	// match wins.  If no HttpKeyBuilder matches, an empty key_map will be sent to
	// the lookup service; it should likely reply with a global default route
	// and raise an alert.
	HttpKeybuilders []*HttpKeyBuilder `protobuf:"bytes,1,rep,name=http_keybuilders,json=httpKeybuilders,proto3" json:"http_keybuilders,omitempty"`
	// Unordered specifications for constructing keys for gRPC requests.  All
	// GrpcKeyBuilders on this list must have unique "name" fields so that the
	// client is free to prebuild a hash map keyed by name.  If no GrpcKeyBuilder
	// matches, an empty key_map will be sent to the lookup service; it should
	// likely reply with a global default route and raise an alert.
	GrpcKeybuilders []*GrpcKeyBuilder `protobuf:"bytes,2,rep,name=grpc_keybuilders,json=grpcKeybuilders,proto3" json:"grpc_keybuilders,omitempty"`
	// The name of the lookup service as a gRPC URI.  Typically, this will be
	// a subdomain of the target, such as "lookup.datastore.googleapis.com".
	LookupService string `protobuf:"bytes,3,opt,name=lookup_service,json=lookupService,proto3" json:"lookup_service,omitempty"`
	// Configure a timeout value for lookup service requests.
	// Defaults to 10 seconds if not specified.
	LookupServiceTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=lookup_service_timeout,json=lookupServiceTimeout,proto3" json:"lookup_service_timeout,omitempty"`
	// How long are responses valid for (like HTTP Cache-Control).
	// If omitted or zero, the longest valid cache time is used.
	// This value is clamped to 5 minutes to avoid unflushable bad responses.
	MaxAge *durationpb.Duration `protobuf:"bytes,5,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// After a response has been in the client cache for this amount of time
	// and is re-requested, start an asynchronous RPC to re-validate it.
	// This value should be less than max_age by at least the length of a
	// typical RTT to the Route Lookup Service to fully mask the RTT latency.
	// If omitted, keys are only re-requested after they have expired.
	StaleAge *durationpb.Duration `protobuf:"bytes,6,opt,name=stale_age,json=staleAge,proto3" json:"stale_age,omitempty"`
	// Rough indicator of amount of memory to use for the client cache.  Some of
	// the data structure overhead is not accounted for, so actual memory consumed
	// will be somewhat greater than this value.  If this field is omitted or set
	// to zero, a client default will be used.  The value may be capped to a lower
	// amount based on client configuration.
	CacheSizeBytes int64 `protobuf:"varint,7,opt,name=cache_size_bytes,json=cacheSizeBytes,proto3" json:"cache_size_bytes,omitempty"`
	// This is a list of all the possible targets that can be returned by the
	// lookup service.  If a target not on this list is returned, it will be
	// treated the same as an unhealthy target.
	ValidTargets []string `protobuf:"bytes,8,rep,name=valid_targets,json=validTargets,proto3" json:"valid_targets,omitempty"`
	// This value provides a default target to use if needed.  If set, it will be
	// used if RLS returns an error, times out, or returns an invalid response.
	// Note that requests can be routed only to a subdomain of the original
	// target, e.g. "us_east_1.cloudbigtable.googleapis.com".
	DefaultTarget string `protobuf:"bytes,9,opt,name=default_target,json=defaultTarget,proto3" json:"default_target,omitempty"`
}

func (x *RouteLookupConfig) Reset() {
	*x = RouteLookupConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteLookupConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteLookupConfig) ProtoMessage() {}

func (x *RouteLookupConfig) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteLookupConfig.ProtoReflect.Descriptor instead.
func (*RouteLookupConfig) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_config_proto_rawDescGZIP(), []int{3}
}

func (x *RouteLookupConfig) GetHttpKeybuilders() []*HttpKeyBuilder {
	if x != nil {
		return x.HttpKeybuilders
	}
	return nil
}

func (x *RouteLookupConfig) GetGrpcKeybuilders() []*GrpcKeyBuilder {
	if x != nil {
		return x.GrpcKeybuilders
	}
	return nil
}

func (x *RouteLookupConfig) GetLookupService() string {
	if x != nil {
		return x.LookupService
	}
	return ""
}

func (x *RouteLookupConfig) GetLookupServiceTimeout() *durationpb.Duration {
	if x != nil {
		return x.LookupServiceTimeout
	}
	return nil
}

func (x *RouteLookupConfig) GetMaxAge() *durationpb.Duration {
	if x != nil {
		return x.MaxAge
	}
	return nil
}

func (x *RouteLookupConfig) GetStaleAge() *durationpb.Duration {
	if x != nil {
		return x.StaleAge
	}
	return nil
}

func (x *RouteLookupConfig) GetCacheSizeBytes() int64 {
	if x != nil {
		return x.CacheSizeBytes
	}
	return 0
}

func (x *RouteLookupConfig) GetValidTargets() []string {
	if x != nil {
		return x.ValidTargets
	}
	return nil
}

func (x *RouteLookupConfig) GetDefaultTarget() string {
	if x != nil {
		return x.DefaultTarget
	}
	return ""
}

// To match, one of the given Name fields must match; the service and method
// fields are specified as fixed strings.  The service name is required and
// includes the proto package name.  The method name may be omitted, in
// which case any method on the given service is matched.
type GrpcKeyBuilder_Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *GrpcKeyBuilder_Name) Reset() {
	*x = GrpcKeyBuilder_Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcKeyBuilder_Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcKeyBuilder_Name) ProtoMessage() {}

func (x *GrpcKeyBuilder_Name) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcKeyBuilder_Name.ProtoReflect.Descriptor instead.
func (*GrpcKeyBuilder_Name) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GrpcKeyBuilder_Name) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GrpcKeyBuilder_Name) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

// If you wish to include the host, service, or method names as keys in the
// generated RouteLookupRequest, specify key names to use in the extra_keys
// submessage. If a key name is empty, no key will be set for that value.
// If this submessage is specified, the normal host/path fields will be left
// unset in the RouteLookupRequest. We are deprecating host/path in the
// RouteLookupRequest, so services should migrate to the ExtraKeys approach.
type GrpcKeyBuilder_ExtraKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host    string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *GrpcKeyBuilder_ExtraKeys) Reset() {
	*x = GrpcKeyBuilder_ExtraKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcKeyBuilder_ExtraKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcKeyBuilder_ExtraKeys) ProtoMessage() {}

func (x *GrpcKeyBuilder_ExtraKeys) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcKeyBuilder_ExtraKeys.ProtoReflect.Descriptor instead.
func (*GrpcKeyBuilder_ExtraKeys) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GrpcKeyBuilder_ExtraKeys) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GrpcKeyBuilder_ExtraKeys) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GrpcKeyBuilder_ExtraKeys) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

var File_grpc_lookup_v1_rls_config_proto protoreflect.FileDescriptor

var file_grpc_lookup_v1_rls_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5c, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x22,
	0xf0, 0x03, 0x0a, 0x0e, 0x47, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x09, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x55, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x1a, 0x38, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x51,
	0x0a, 0x09, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xf1, 0x02, 0x0a, 0x0e, 0x48, 0x74, 0x74, 0x70, 0x4b, 0x65, 0x79, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f,
	0x73, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12,
	0x46, 0x0a, 0x10, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x55,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4b, 0x65, 0x79, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa6, 0x04, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x10,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4b, 0x65, 0x79, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x4b, 0x65, 0x79, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x16, 0x6c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x61,
	0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x74,
	0x61, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4a, 0x04, 0x08, 0x0a,
	0x10, 0x0b, 0x52, 0x1b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x42,
	0x53, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_lookup_v1_rls_config_proto_rawDescOnce sync.Once
	file_grpc_lookup_v1_rls_config_proto_rawDescData = file_grpc_lookup_v1_rls_config_proto_rawDesc
)

func file_grpc_lookup_v1_rls_config_proto_rawDescGZIP() []byte {
	file_grpc_lookup_v1_rls_config_proto_rawDescOnce.Do(func() {
		file_grpc_lookup_v1_rls_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_lookup_v1_rls_config_proto_rawDescData)
	})
	return file_grpc_lookup_v1_rls_config_proto_rawDescData
}

var file_grpc_lookup_v1_rls_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_lookup_v1_rls_config_proto_goTypes = []interface{}{
	(*NameMatcher)(nil),              // 0: grpc.lookup.v1.NameMatcher
	(*GrpcKeyBuilder)(nil),           // 1: grpc.lookup.v1.GrpcKeyBuilder
	(*HttpKeyBuilder)(nil),           // 2: grpc.lookup.v1.HttpKeyBuilder
	(*RouteLookupConfig)(nil),        // 3: grpc.lookup.v1.RouteLookupConfig
	(*GrpcKeyBuilder_Name)(nil),      // 4: grpc.lookup.v1.GrpcKeyBuilder.Name
	(*GrpcKeyBuilder_ExtraKeys)(nil), // 5: grpc.lookup.v1.GrpcKeyBuilder.ExtraKeys
	nil,                              // 6: grpc.lookup.v1.GrpcKeyBuilder.ConstantKeysEntry
	nil,                              // 7: grpc.lookup.v1.HttpKeyBuilder.ConstantKeysEntry
	(*durationpb.Duration)(nil),      // 8: google.protobuf.Duration
}
var file_grpc_lookup_v1_rls_config_proto_depIdxs = []int32{
	4,  // 0: grpc.lookup.v1.GrpcKeyBuilder.names:type_name -> grpc.lookup.v1.GrpcKeyBuilder.Name
	5,  // 1: grpc.lookup.v1.GrpcKeyBuilder.extra_keys:type_name -> grpc.lookup.v1.GrpcKeyBuilder.ExtraKeys
	0,  // 2: grpc.lookup.v1.GrpcKeyBuilder.headers:type_name -> grpc.lookup.v1.NameMatcher
	6,  // 3: grpc.lookup.v1.GrpcKeyBuilder.constant_keys:type_name -> grpc.lookup.v1.GrpcKeyBuilder.ConstantKeysEntry
	0,  // 4: grpc.lookup.v1.HttpKeyBuilder.query_parameters:type_name -> grpc.lookup.v1.NameMatcher
	0,  // 5: grpc.lookup.v1.HttpKeyBuilder.headers:type_name -> grpc.lookup.v1.NameMatcher
	7,  // 6: grpc.lookup.v1.HttpKeyBuilder.constant_keys:type_name -> grpc.lookup.v1.HttpKeyBuilder.ConstantKeysEntry
	2,  // 7: grpc.lookup.v1.RouteLookupConfig.http_keybuilders:type_name -> grpc.lookup.v1.HttpKeyBuilder
	1,  // 8: grpc.lookup.v1.RouteLookupConfig.grpc_keybuilders:type_name -> grpc.lookup.v1.GrpcKeyBuilder
	8,  // 9: grpc.lookup.v1.RouteLookupConfig.lookup_service_timeout:type_name -> google.protobuf.Duration
	8,  // 10: grpc.lookup.v1.RouteLookupConfig.max_age:type_name -> google.protobuf.Duration
	8,  // 11: grpc.lookup.v1.RouteLookupConfig.stale_age:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_grpc_lookup_v1_rls_config_proto_init() }
func file_grpc_lookup_v1_rls_config_proto_init() {
	if File_grpc_lookup_v1_rls_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_lookup_v1_rls_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameMatcher); i {
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
		file_grpc_lookup_v1_rls_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcKeyBuilder); i {
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
		file_grpc_lookup_v1_rls_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpKeyBuilder); i {
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
		file_grpc_lookup_v1_rls_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteLookupConfig); i {
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
		file_grpc_lookup_v1_rls_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcKeyBuilder_Name); i {
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
		file_grpc_lookup_v1_rls_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcKeyBuilder_ExtraKeys); i {
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
			RawDescriptor: file_grpc_lookup_v1_rls_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_lookup_v1_rls_config_proto_goTypes,
		DependencyIndexes: file_grpc_lookup_v1_rls_config_proto_depIdxs,
		MessageInfos:      file_grpc_lookup_v1_rls_config_proto_msgTypes,
	}.Build()
	File_grpc_lookup_v1_rls_config_proto = out.File
	file_grpc_lookup_v1_rls_config_proto_rawDesc = nil
	file_grpc_lookup_v1_rls_config_proto_goTypes = nil
	file_grpc_lookup_v1_rls_config_proto_depIdxs = nil
}
