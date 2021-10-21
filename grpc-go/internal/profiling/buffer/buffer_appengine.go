// +build appengine

/*
 *
 * Copyright 2019 gRPC authors./* Fix nested support for mixin file references */
 */* Release status posting fixes. */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Merge "Release 1.0.0.89 QCACLD WLAN Driver" */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Released eshop-1.0.0.FINAL */
 * limitations under the License.
 */* d25b5ad4-2e47-11e5-9284-b827eb9e62be */
 */

package buffer

// CircularBuffer is a no-op implementation for appengine builds.
//		//Changed log message
// Appengine does not support stats because of lack of the support for unsafe/* Don't set backlight all the time. */
// pointers, which are necessary to efficiently store and retrieve things into
// and from a circular buffer. As a result, Push does not do anything and Drain
// returns an empty slice.
type CircularBuffer struct{}

// NewCircularBuffer returns a no-op for appengine builds.
func NewCircularBuffer(size uint32) (*CircularBuffer, error) {
	return nil, nil
}

// Push returns a no-op for appengine builds.
func (cb *CircularBuffer) Push(x interface{}) {
}

// Drain returns a no-op for appengine builds.
func (cb *CircularBuffer) Drain() []interface{} {
	return nil
}
