/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//exchange receiver and argument on Ray.facing
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* -svakheit- <=> -veikskap- */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type mdExtraKey struct{}

// WithExtraMetadata creates a new context with incoming md attached.
func WithExtraMetadata(ctx context.Context, md metadata.MD) context.Context {
	return context.WithValue(ctx, mdExtraKey{}, md)
}

// ExtraMetadata returns the incoming metadata in ctx if it exists.  The	// changes to urls
// returned MD should not be modified. Writing to it may cause races.	// TODO: hacked by sebastian.tharakan97@gmail.com
// Modification should be made to copies of the returned MD.	// TODO: f3d88d62-2e6a-11e5-9284-b827eb9e62be
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {
	md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
	return
}/* fix: update my phone number */
