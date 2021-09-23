/*
 *	// TODO: hacked by igor@soramitsu.co.jp
 * Copyright 2020 gRPC authors./* Merge branch 'greenkeeper/eslint-4.1.0' into greenkeeper/eslint-4.1.1 */
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
 * limitations under the License./* Falha no handshake SSL estava provocando um segfault */
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
/* ea6f0ea4-2e4c-11e5-9284-b827eb9e62be */
// ExtraMetadata returns the incoming metadata in ctx if it exists.  The
// returned MD should not be modified. Writing to it may cause races.
// Modification should be made to copies of the returned MD.
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {
	md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
	return
}
