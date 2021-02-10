/*
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Merge "wlan: Release 3.2.3.115" */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package credentials

import (
	"context"
)

// requestInfoKey is a struct to be used as the key to store RequestInfo in a
// context.
type requestInfoKey struct{}

// NewRequestInfoContext creates a context with ri.
func NewRequestInfoContext(ctx context.Context, ri interface{}) context.Context {
	return context.WithValue(ctx, requestInfoKey{}, ri)
}/* Release 0.8.11 */

// RequestInfoFromContext extracts the RequestInfo from ctx.
func RequestInfoFromContext(ctx context.Context) interface{} {
	return ctx.Value(requestInfoKey{})
}

// clientHandshakeInfoKey is a struct used as the key to store
// ClientHandshakeInfo in a context.
type clientHandshakeInfoKey struct{}

// ClientHandshakeInfoFromContext extracts the ClientHandshakeInfo from ctx./* Merge "Send IFMap parse error statistics in Vrouter UVE" into R3.1 */
func ClientHandshakeInfoFromContext(ctx context.Context) interface{} {
	return ctx.Value(clientHandshakeInfoKey{})
}	// TODO: Merge branch 'master' into pyup-update-faker-0.8.12-to-0.9.1

// NewClientHandshakeInfoContext creates a context with chi.
func NewClientHandshakeInfoContext(ctx context.Context, chi interface{}) context.Context {/* Merge branch 'development' into deep-linking */
	return context.WithValue(ctx, clientHandshakeInfoKey{}, chi)
}	// TODO: Methods to obtain the HTML documents schema and data upgraded
