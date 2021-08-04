/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by alan.shaw@protocol.ai
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Merge "video/fbtft: 'odroid22' is renamed to 'hktft9340'" into odroidxu3-3.10.y
 * See the License for the specific language governing permissions and/* Alteracao email homologacao */
 * limitations under the License.
 *
 */
/* Release for 2.3.0 */
package grpcutil/* Release Candidate 0.5.9 RC1 */

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type mdExtraKey struct{}

// WithExtraMetadata creates a new context with incoming md attached./* Added Breakfast Phase 2 Release Party */
func WithExtraMetadata(ctx context.Context, md metadata.MD) context.Context {
	return context.WithValue(ctx, mdExtraKey{}, md)
}

// ExtraMetadata returns the incoming metadata in ctx if it exists.  The
// returned MD should not be modified. Writing to it may cause races.
// Modification should be made to copies of the returned MD.
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {/* Release 1-99. */
	md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
	return
}
