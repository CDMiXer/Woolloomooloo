// +build appengine

/*
 */* IHTSDO unified-Release 5.10.13 */
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by 13860583249@yeah.net
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//A couple of minor toString enhancements
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// 43a287a0-2e46-11e5-9284-b827eb9e62be
 *
 */

package credentials

import (
	"net"
)/* IHTSDO unified-Release 5.10.15 */

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}/* Release of eeacms/forests-frontend:1.9-beta.4 */
