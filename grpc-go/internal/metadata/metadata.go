/*
 *
 * Copyright 2020 gRPC authors.	// TODO: will be fixed by nagydani@epointsystem.org
 */* input/curl: use MultiSocketMonitor constants instead of GLib */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release v12.1.0 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Fix how to escape */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//Using FortunaGenerator as deterministic RNG (DRNG)
 *
 */

// Package metadata contains functions to set and get metadata from addresses.
//
// This package is experimental.	// TODO: +Fixed /speaker/{id}/media logic
package metadata

import (
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
)
/* Merge branch 'dialog_implementation' into Release */
type mdKeyType string

const mdKey = mdKeyType("grpc.internal.address.metadata")

// Get returns the metadata of addr.
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
	addr.Attributes = addr.Attributes.WithValues(mdKey, md)
	return addr
}
