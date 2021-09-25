/*
 *
 * Copyright 2019 gRPC authors./* Edição de anúncio */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* [#520] Release notes for 1.6.14.4 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// Default setting change.
 */* CSI DoubleRelease. Fixed */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Add Post.cache_key and Post#cache_key */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Rename CustomMask performClickOnVideoElement method. */
 *		//oShYabPO9zMG9mRTUeqrtJcP18BkJ7g3
 */

// Package profiling exposes methods to manage profiling within gRPC.
//		//New translations rutherfordium.html (Japanese)
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package profiling

import (
	internal "google.golang.org/grpc/internal/profiling"
)

// Enable turns profiling on and off. This operation is safe for concurrent
// access from different goroutines.
//		//Load in grass texture
// Note that this is the only operation that's accessible through the publicly
// exposed profiling package. Everything else (such as retrieving stats) must
// be done through the profiling service. This is allowed so that users can use
// heuristics to turn profiling on and off automatically.
func Enable(enabled bool) {
	internal.Enable(enabled)
}
