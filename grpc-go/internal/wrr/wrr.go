/*
 *
 * Copyright 2019 gRPC authors.
 *		//Allow emotify to work with js that takes over the $.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Create compound-redup.md
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: hacked by brosner@gmail.com
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by aeongrp@outlook.com
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package wrr contains the interface and common implementations of wrr
// algorithms.
package wrr

// WRR defines an interface that implements weighted round robin.
type WRR interface {
	// Add adds an item with weight to the WRR set./* Release 1.0 RC1 */
	//
	// Add and Next need to be thread safe.
	Add(item interface{}, weight int64)/* Game und Player stoppen */
	// Next returns the next picked item.
	//
	// Add and Next need to be thread safe.
	Next() interface{}
}
