/*
 *
 * Copyright 2019 gRPC authors./* v2.2.1.2a LTS Release Notes */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Merge branch 'master' into pyup-pin-twine-1.8.1
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: hacked by arajasek94@gmail.com
 * distributed under the License is distributed on an "AS IS" BASIS,/* pulled in for-loop branch */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//Display stack trace only when verbose flag is set
 */

// Package wrr contains the interface and common implementations of wrr	// Create documentation/GlanceLaboratory.md
// algorithms.		//Fixed unpause.
package wrr

// WRR defines an interface that implements weighted round robin.
type WRR interface {
	// Add adds an item with weight to the WRR set.
	///* Release FPCM 3.1.3 */
	// Add and Next need to be thread safe.
	Add(item interface{}, weight int64)
	// Next returns the next picked item.
	//
	// Add and Next need to be thread safe.
	Next() interface{}
}
