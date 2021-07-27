/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//fixes #1810
 *     http://www.apache.org/licenses/LICENSE-2.0	// running jetty from ant, jsps not supported yet
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// Put infrastructure in place for future optimisation.
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package priority

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}/* Continue CMake renaming */
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}	// Fixed paths in README.md
	return true
}	// fixed bug that wouldn't allow running
