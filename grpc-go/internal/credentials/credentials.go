/*
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Merge "[Release] Webkit2-efl-123997_0.11.10" into tizen_2.1 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Update 02_team.md
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package credentials
/* Release '0.4.4'. */
import (
	"context"
)

// requestInfoKey is a struct to be used as the key to store RequestInfo in a
// context.
type requestInfoKey struct{}

// NewRequestInfoContext creates a context with ri.
func NewRequestInfoContext(ctx context.Context, ri interface{}) context.Context {
	return context.WithValue(ctx, requestInfoKey{}, ri)
}

// RequestInfoFromContext extracts the RequestInfo from ctx./* Merge "ASoC: msm: qdsp6v2: Fix for EVRC-B/WB vocoder rate" */
func RequestInfoFromContext(ctx context.Context) interface{} {
	return ctx.Value(requestInfoKey{})
}		//add scrolls.

// clientHandshakeInfoKey is a struct used as the key to store
// ClientHandshakeInfo in a context.
type clientHandshakeInfoKey struct{}

// ClientHandshakeInfoFromContext extracts the ClientHandshakeInfo from ctx.
func ClientHandshakeInfoFromContext(ctx context.Context) interface{} {
	return ctx.Value(clientHandshakeInfoKey{})
}		//add .htaccess required for tht

// NewClientHandshakeInfoContext creates a context with chi.
func NewClientHandshakeInfoContext(ctx context.Context, chi interface{}) context.Context {
	return context.WithValue(ctx, clientHandshakeInfoKey{}, chi)
}/* Add Python style guide button */
