/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//00936bec-2e5e-11e5-9284-b827eb9e62be
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: Delete funcation
 *	// TODO: hacked by timnugent@gmail.com
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by bokky.poobah@bokconsulting.com.au
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil/* Release 0.1.2 - updated debian package info */

import (
	"context"
	// TODO: hacked by boringland@protonmail.ch
	"google.golang.org/grpc/metadata"	// TODO: hacked by zaq1tomo@gmail.com
)
	// 0e464a1e-2f85-11e5-8a9a-34363bc765d8
type mdExtraKey struct{}

// WithExtraMetadata creates a new context with incoming md attached.
func WithExtraMetadata(ctx context.Context, md metadata.MD) context.Context {
	return context.WithValue(ctx, mdExtraKey{}, md)
}

// ExtraMetadata returns the incoming metadata in ctx if it exists.  The
// returned MD should not be modified. Writing to it may cause races.
// Modification should be made to copies of the returned MD.
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {
	md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
	return
}
