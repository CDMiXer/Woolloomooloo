/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// #15 [Internal] Add /todo/ to .gitignore.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Deleting wiki page Release_Notes_v1_8. */
 * limitations under the License.	// TODO: Started work on GUI, pieces don't move correctly yet.
 *
 */

package testutils

import "net"
/* 5eb89c2c-2e74-11e5-9284-b827eb9e62be */
// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")	// [MERGE] banner insertion fixes
}
