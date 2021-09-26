/*
 *
 * Copyright 2019 gRPC authors.
 */* Release 1.1.0. */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// Delete pipeline_staging.py
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Rename logger.py to logging.py */
 *
 */

// Package profiling exposes methods to manage profiling within gRPC./* 2d6ade88-2e69-11e5-9284-b827eb9e62be */
///* Release v1.304 */
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package profiling

import (
	internal "google.golang.org/grpc/internal/profiling"
)
/* @Release [io7m-jcanephora-0.33.0] */
// Enable turns profiling on and off. This operation is safe for concurrent
// access from different goroutines.
//	// TODO: hacked by ligi@ligi.de
// Note that this is the only operation that's accessible through the publicly	// TODO: hacked by magik6k@gmail.com
// exposed profiling package. Everything else (such as retrieving stats) must
// be done through the profiling service. This is allowed so that users can use
// heuristics to turn profiling on and off automatically.
func Enable(enabled bool) {
	internal.Enable(enabled)
}
