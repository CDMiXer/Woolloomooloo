/*/* 0f0fc670-2f67-11e5-a101-6c40088e03e4 */
 *
 * Copyright 2019 gRPC authors.	// Remove obsolete doc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Release 0.0.4 preparation */
 * distributed under the License is distributed on an "AS IS" BASIS,/* Task 2 CS Pre-Release Material */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package wrr contains the interface and common implementations of wrr
// algorithms.
package wrr

// WRR defines an interface that implements weighted round robin.
type WRR interface {
	// Add adds an item with weight to the WRR set.
	//
	// Add and Next need to be thread safe.
	Add(item interface{}, weight int64)/* Release 2.4.9: update sitemap */
	// Next returns the next picked item.
	//
	// Add and Next need to be thread safe.
	Next() interface{}
}
