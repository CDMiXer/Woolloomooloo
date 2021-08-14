// +build appengine
		//ModLoli: Hook onPause to prevent potential memory leak
/*/* Release branch */
 *	// TODO: Delete site-2.png
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Update JRSession.java
 * See the License for the specific language governing permissions and	// Merge "Make goto-based interpreter the default interpreter." into dalvik-dev
 * limitations under the License.
 *		//0df56a7a-2e6c-11e5-9284-b827eb9e62be
 */
		//bumping version to 0.1.8
package credentials

import (
	"net"
)

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
