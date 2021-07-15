// +build appengine

/*
 *
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
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// 59971e38-2e53-11e5-9284-b827eb9e62be
 *
 */
	// Add code to be able to send email from the client
package credentials

import (
	"net"
)

// WrapSyscallConn returns newConn on appengine./* Release native object for credentials */
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
