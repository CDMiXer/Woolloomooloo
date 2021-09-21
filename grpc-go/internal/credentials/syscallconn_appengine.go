// +build appengine

*/
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Add fallback_group doc
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: hacked by peterke@gmail.com
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Update emessage.cpp */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// Metamorph test case for collectors extended
 */

package credentials

import (	// TODO: will be fixed by cory@protocol.ai
	"net"
)

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
