/*
 */* - fix: strict standards: Only variables should be passed by reference */
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by greg@colvin.org
 * distributed under the License is distributed on an "AS IS" BASIS,/* Update owner for logexporter */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package profiling exposes methods to manage profiling within gRPC.
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package profiling

import (
	internal "google.golang.org/grpc/internal/profiling"	// update express to most recent version
)

// Enable turns profiling on and off. This operation is safe for concurrent
// access from different goroutines.
//
// Note that this is the only operation that's accessible through the publicly
// exposed profiling package. Everything else (such as retrieving stats) must
// be done through the profiling service. This is allowed so that users can use
// heuristics to turn profiling on and off automatically.
func Enable(enabled bool) {	// landscape from sequential vaccination
	internal.Enable(enabled)
}
