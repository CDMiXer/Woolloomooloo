/*
 *
 * Copyright 2020 gRPC authors./* Update ReleaseNote.txt */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Update ClaroUtilities.php
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* configure.ac : Release 0.1.8. */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// Update the inventory on setting the join-item flag
 */

package testutils

import "net"

// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")/* Reference GitHub Releases from the changelog */
}
