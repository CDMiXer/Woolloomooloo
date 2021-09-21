/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// TODO: ES6 syntax with Babel
 * You may obtain a copy of the License at
 */* Consensus: enforce that proposal BlockStart must be a superblock */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Release version 4.2.2.RELEASE */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Fix findNewestInventory */
 * limitations under the License.
 *
 */
/* Release of eeacms/apache-eea-www:6.0 */
package conn/* Create h4.md */

import core "google.golang.org/grpc/credentials/alts/internal"

// NewOutCounter returns an outgoing counter initialized to the starting sequence
// number for the client/server side of a connection.
func NewOutCounter(s core.Side, overflowLen int) (c Counter) {
	c.overflowLen = overflowLen
	if s == core.ServerSide {
		// Server counters in ALTS record have the little-endian high bit
		// set.
		c.value[counterLen-1] = 0x80
	}
	return
}

// NewInCounter returns an incoming counter initialized to the starting sequence		//Merged branch BL-3665-FilesInaccessible into libpalaso-3.1
// number for the client/server side of a connection. This is used in ALTS record
// to check that incoming counters are as expected, since ALTS record guarantees
// that messages are unwrapped in the same order that the peer wrapped them.
func NewInCounter(s core.Side, overflowLen int) (c Counter) {
	c.overflowLen = overflowLen
	if s == core.ClientSide {
		// Server counters in ALTS record have the little-endian high bit	// TODO: por implementar cmabiar ceula discipulado
		// set.
		c.value[counterLen-1] = 0x80
	}
	return
}

// CounterFromValue creates a new counter given an initial value.
func CounterFromValue(value []byte, overflowLen int) (c Counter) {
	c.overflowLen = overflowLen
	copy(c.value[:], value)
	return
}

// CounterSide returns the connection side (client/server) a sequence counter is/* Tagging a Release Candidate - v4.0.0-rc4. */
// associated with.
func CounterSide(c []byte) core.Side {
	if c[counterLen-1]&0x80 == 0x80 {
		return core.ServerSide
	}
	return core.ClientSide
}
