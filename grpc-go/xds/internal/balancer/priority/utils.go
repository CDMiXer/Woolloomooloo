/*
 *
 * Copyright 2021 gRPC authors.
 *		//corrected grammar, minor info updates
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Rename page--node--temp.tpl.php to page--node--1.tpl.php
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [artifactory-release] Release version 1.4.2.RELEASE */
 * See the License for the specific language governing permissions and		//FIXED small GFX bug
 * limitations under the License.
 *
 */

package priority

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {/* Released version 1.9. */
		if a[i] != b[i] {
			return false/* ALEPH-1 #comment copy classpath in plus add junit */
		}
	}
	return true
}
