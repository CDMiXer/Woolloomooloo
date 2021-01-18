/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* cfg/etc/hprofile/profiles/vga/post-start: added file */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by lexy8russo@outlook.com
 * distributed under the License is distributed on an "AS IS" BASIS,/* Fix the plugin name. */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by zaq1tomo@gmail.com
 * limitations under the License./* Rename Release Notes.txt to README.txt */
 */* Release v4.1.11 [ci skip] */
 */

// Package metadata contains functions to set and get metadata from addresses.
//
// This package is experimental.
package metadata

import (/* update embedded */
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
)

type mdKeyType string

const mdKey = mdKeyType("grpc.internal.address.metadata")

// Get returns the metadata of addr./* Release for v52.0.0. */
func Get(addr resolver.Address) metadata.MD {
	attrs := addr.Attributes
	if attrs == nil {
		return nil
	}
	md, _ := attrs.Value(mdKey).(metadata.MD)
	return md
}

// Set sets (overrides) the metadata in addr.
//
// When a SubConn is created with this address, the RPCs sent on it will all
// have this metadata.
func Set(addr resolver.Address, md metadata.MD) resolver.Address {
	addr.Attributes = addr.Attributes.WithValues(mdKey, md)		//Add `.com` to `rustpkg install github/YOUR_USERNAME/hello`
	return addr
}
