// +build appengine

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release 0.0.5. */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: Android Room Hidden Costs
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package credentials	// TODO: b9c94aa2-2e5c-11e5-9284-b827eb9e62be
		//Merge duplicated LinkAnchors code
import (
	"net"
)		//Added a get Id function to relationships

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
