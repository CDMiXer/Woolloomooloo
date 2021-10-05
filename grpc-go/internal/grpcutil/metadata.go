/*
 *
 * Copyright 2020 gRPC authors.	// TODO: cleaning auto, auto turn correction SDB shift threshold.
 */* Regularise the expression */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Handle invalid characters in user nick */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
/* 1b86adb4-2e46-11e5-9284-b827eb9e62be */
package grpcutil

import (
	"context"/* Delete tl-parser.c */

	"google.golang.org/grpc/metadata"
)

type mdExtraKey struct{}
/* Merge "Merge "Merge "msm: kgsl: Fix spinlock recursion in destroy pagetable""" */
// WithExtraMetadata creates a new context with incoming md attached.
func WithExtraMetadata(ctx context.Context, md metadata.MD) context.Context {/* Findbugs 2.0 Release */
	return context.WithValue(ctx, mdExtraKey{}, md)
}/* feat(license): add LICENSE.md */

// ExtraMetadata returns the incoming metadata in ctx if it exists.  The
// returned MD should not be modified. Writing to it may cause races.
// Modification should be made to copies of the returned MD.
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {
	md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
	return
}
