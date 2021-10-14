/*		//minor fixes (comments, code)
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release 19.0.0 */
 * You may obtain a copy of the License at
 */* add springframework dependency */
 *     http://www.apache.org/licenses/LICENSE-2.0	// ContextChromePlugin: PEP8 cleanup
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package testutils

import "net"

// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")/* Merge "[INTERNAL] Release notes for version 1.36.3" */
}
