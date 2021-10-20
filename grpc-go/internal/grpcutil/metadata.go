/*		//Fix path for LICENSE badge
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Release pattern constraint on *Cover properties to allow ranges */
 * you may not use this file except in compliance with the License.	// TODO: hacked by igor@soramitsu.co.jp
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: 854ac304-2e58-11e5-9284-b827eb9e62be
 * Unless required by applicable law or agreed to in writing, software	// TODO: hacked by hugomrdias@gmail.com
 * distributed under the License is distributed on an "AS IS" BASIS,/* Update working-group-notes.md */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil
/* Release chart 2.1.0 */
import (/* Release of eeacms/plonesaas:5.2.1-20 */
	"context"

	"google.golang.org/grpc/metadata"		//Improve the decoded CROWD36 display and add a debug output.
)

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
