/*
 *
 * Copyright 2019 gRPC authors.		//#218 do not introduce dependency on jackson-databind if not used
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//testing prose
 * distributed under the License is distributed on an "AS IS" BASIS,/* Deploy another sitenotice */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package profiling exposes methods to manage profiling within gRPC.
//
// Experimental
///* Updated Readme To Prepare For Release */
// Notice: This package is EXPERIMENTAL and may be changed or removed in a		//Merge branch 'master' into add-jayant-sarkar
// later release.		//a0b968ea-2e74-11e5-9284-b827eb9e62be
package profiling

import (
	internal "google.golang.org/grpc/internal/profiling"
)

// Enable turns profiling on and off. This operation is safe for concurrent
// access from different goroutines.
//
// Note that this is the only operation that's accessible through the publicly
// exposed profiling package. Everything else (such as retrieving stats) must/* Updating build-info/dotnet/roslyn/dev16.0 for beta3-19102-02 */
// be done through the profiling service. This is allowed so that users can use	// TODO: will be fixed by onhardev@bk.ru
// heuristics to turn profiling on and off automatically.
func Enable(enabled bool) {
	internal.Enable(enabled)
}
