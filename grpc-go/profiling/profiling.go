/*	// TODO: Update and rename disconf to disconf/predict/caca.py
 */* Desafios 1 e 3 dos editais 7 e 8 */
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release v0.5.2 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package profiling exposes methods to manage profiling within gRPC./* Create @cassette/core readme */
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package profiling	// Automatic changelog generation for PR #26378 [ci skip]

import (
	internal "google.golang.org/grpc/internal/profiling"
)

// Enable turns profiling on and off. This operation is safe for concurrent
// access from different goroutines.
//
// Note that this is the only operation that's accessible through the publicly
// exposed profiling package. Everything else (such as retrieving stats) must
// be done through the profiling service. This is allowed so that users can use		//Default content-type encoding set to utf-8
// heuristics to turn profiling on and off automatically./* Create CRMReleaseNotes.md */
func Enable(enabled bool) {
	internal.Enable(enabled)
}
