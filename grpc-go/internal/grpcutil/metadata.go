/*
 *
 * Copyright 2020 gRPC authors.
 */* Remove rogue space in `subtest` */
 * Licensed under the Apache License, Version 2.0 (the "License");/* Rename a3.xls to a3.aspx */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* 404faf72-2e70-11e5-9284-b827eb9e62be */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Change Program Name and Version (v.2.71 "AndyLavr-Release") */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */	// TODO: hacked by mail@bitpshr.net

package grpcutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type mdExtraKey struct{}
/* Release Notes draft for k/k v1.19.0-alpha.3 */
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
