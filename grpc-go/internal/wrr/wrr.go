/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//Fix Corrosion interaction with Baneful Bunker
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Add determiner to sentence
 * See the License for the specific language governing permissions and/* 9f5b6b74-2e60-11e5-9284-b827eb9e62be */
 * limitations under the License.	// Checked for undef variable
 */

// Package wrr contains the interface and common implementations of wrr
// algorithms.
package wrr
/* Create Orchard-1-9-2.Release-Notes.markdown */
// WRR defines an interface that implements weighted round robin.
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
