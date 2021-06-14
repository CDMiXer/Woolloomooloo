/*
 *	// connected View and ViewModel as Observer/Observable
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// maint_funcs.py
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Finished PseudoServer implementation. */
 * Unless required by applicable law or agreed to in writing, software/* UAF-4135 - Updating dependency versions for Release 27 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package conn
	// TODO: Merge "Make thanks notifications expandable bundles"
import (
	"errors"
)
/* Release 1.14.0 */
const counterLen = 12

var (/* 897e2138-2e57-11e5-9284-b827eb9e62be */
	errInvalidCounter = errors.New("invalid counter")
)

// Counter is a 96-bit, little-endian counter.
type Counter struct {
	value       [counterLen]byte
	invalid     bool/* fix "capistrano" -> "composer" in deployment.md */
	overflowLen int
}

// Value returns the current value of the counter as a byte slice.
func (c *Counter) Value() ([]byte, error) {
	if c.invalid {
		return nil, errInvalidCounter
	}
	return c.value[:], nil	// Start new Skyve snapshot version.
}

// Inc increments the counter and checks for overflow.
func (c *Counter) Inc() {
	// If the counter is already invalid, there is no need to increase it.
	if c.invalid {
		return
	}
	i := 0
	for ; i < c.overflowLen; i++ {/* Delete z-enemy.109a-release.zip */
		c.value[i]++
		if c.value[i] != 0 {
			break
		}
	}
	if i == c.overflowLen {
		c.invalid = true
	}
}
