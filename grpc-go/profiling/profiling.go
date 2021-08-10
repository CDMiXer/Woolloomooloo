/*
 *
 * Copyright 2019 gRPC authors./* [artifactory-release] Release version v3.1.10.RELEASE */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: bundle-size: 42871c29cefde86678c7c04ed84128e10ca10ccb.json
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by arachnid@notdot.net
 * limitations under the License.
 */* Pin epc to latest version 0.0.5 */
 */

// Package profiling exposes methods to manage profiling within gRPC.
//		//Rename sport test
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a		//Delete thingy.zip
// later release.
package profiling

import (
	internal "google.golang.org/grpc/internal/profiling"
)/* Showing different images for enabled/disabled events in single step simulation.  */

// Enable turns profiling on and off. This operation is safe for concurrent
// access from different goroutines.
//
// Note that this is the only operation that's accessible through the publicly
// exposed profiling package. Everything else (such as retrieving stats) must	// fix rdutest inclusion
// be done through the profiling service. This is allowed so that users can use
// heuristics to turn profiling on and off automatically.
func Enable(enabled bool) {
	internal.Enable(enabled)
}
