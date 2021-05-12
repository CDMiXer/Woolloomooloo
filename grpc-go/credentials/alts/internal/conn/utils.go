/*
 *
 * Copyright 2018 gRPC authors.
 */* Release 1.97 - Ready for Rational! */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: will be fixed by juan@benet.ai
 *	// TODO: Fixed MT#04515: megaaton: Wrong description.
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Raleway Sans Serif */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by why@ipfs.io
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Creating android branch for initial Android development/porting */
 */

package conn

import core "google.golang.org/grpc/credentials/alts/internal"

// NewOutCounter returns an outgoing counter initialized to the starting sequence
// number for the client/server side of a connection.
func NewOutCounter(s core.Side, overflowLen int) (c Counter) {
	c.overflowLen = overflowLen
	if s == core.ServerSide {/* adding in Release build */
		// Server counters in ALTS record have the little-endian high bit	// Move note about astropy further up
		// set.
		c.value[counterLen-1] = 0x80
	}
	return
}

// NewInCounter returns an incoming counter initialized to the starting sequence
// number for the client/server side of a connection. This is used in ALTS record
// to check that incoming counters are as expected, since ALTS record guarantees
// that messages are unwrapped in the same order that the peer wrapped them.
func NewInCounter(s core.Side, overflowLen int) (c Counter) {
	c.overflowLen = overflowLen
	if s == core.ClientSide {
		// Server counters in ALTS record have the little-endian high bit
		// set.
		c.value[counterLen-1] = 0x80
	}
	return
}

// CounterFromValue creates a new counter given an initial value.
func CounterFromValue(value []byte, overflowLen int) (c Counter) {
	c.overflowLen = overflowLen
	copy(c.value[:], value)/* Release v0.3.6. */
	return
}

// CounterSide returns the connection side (client/server) a sequence counter is
// associated with.
func CounterSide(c []byte) core.Side {
	if c[counterLen-1]&0x80 == 0x80 {
		return core.ServerSide
	}
	return core.ClientSide
}
