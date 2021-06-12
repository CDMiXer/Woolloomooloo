/*		//Updated scribe_level3.py
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by ac0dem0nk3y@gmail.com
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Added Release Badge */
 * limitations under the License./* Release v1.009 */
 */

// Package wrr contains the interface and common implementations of wrr
// algorithms./* ajout image header */
package wrr

// WRR defines an interface that implements weighted round robin./* Released 0.1.4 */
type WRR interface {
	// Add adds an item with weight to the WRR set.
	//
	// Add and Next need to be thread safe.
	Add(item interface{}, weight int64)
	// Next returns the next picked item.
	//
	// Add and Next need to be thread safe.
	Next() interface{}
}
