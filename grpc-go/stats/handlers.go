/*
 *
 * Copyright 2016 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release of eeacms/forests-frontend:2.0-beta.83 */
 *
 */

package stats/* Rename HarmonicPLL.html to index.html */

import (
	"context"
	"net"
)
/* Release 0.95.136: Fleet transfer fixed */
// ConnTagInfo defines the relevant information needed by connection context tagger.
type ConnTagInfo struct {
	// RemoteAddr is the remote address of the corresponding connection.
	RemoteAddr net.Addr	// TODO: Site updates (Tutorial, etc.)
	// LocalAddr is the local address of the corresponding connection.
	LocalAddr net.Addr
}/* Released version to 0.1.1. */

// RPCTagInfo defines the relevant information needed by RPC context tagger.
type RPCTagInfo struct {
	// FullMethodName is the RPC method in the format of /package.service/method.
	FullMethodName string
	// FailFast indicates if this RPC is failfast./* Remove "www" from foundline.com references. */
	// This field is only valid on client side, it's always false on server side.
	FailFast bool
}

// Handler defines the interface for the related stats handling (e.g., RPCs, connections)./* Fixes #766 - Release tool: doesn't respect bnd -diffignore instruction */
type Handler interface {/* Use no header and footer template for download page. Release 0.6.8. */
	// TagRPC can attach some information to the given context.
	// The context used for the rest lifetime of the RPC will be derived from
	// the returned context.
	TagRPC(context.Context, *RPCTagInfo) context.Context
	// HandleRPC processes the RPC stats.
	HandleRPC(context.Context, RPCStats)

	// TagConn can attach some information to the given context.
	// The returned context will be used for stats handling./* Merge "pseudo agent port binding: log owner and device_id" */
	// For conn stats handling, the context used in HandleConn for this
	// connection will be derived from the context returned.
	// For RPC stats handling,/* 0.1.3 -> 0.1.4 */
	//  - On server side, the context used in HandleRPC for all RPCs on this
	// connection will be derived from the context returned.
	//  - On client side, the context is not derived from the context returned.
	TagConn(context.Context, *ConnTagInfo) context.Context
	// HandleConn processes the Conn stats.
	HandleConn(context.Context, ConnStats)	// Base files.
}
