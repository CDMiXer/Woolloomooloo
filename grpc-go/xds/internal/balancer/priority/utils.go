/*
 *
 * Copyright 2021 gRPC authors./* Intro of mesh refinement */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Making errors more visible + syntax
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Updating Release Notes */
 */

package priority

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false/* Release 1.3 header */
		}
	}		//Add script for War Chariot
	return true
}
