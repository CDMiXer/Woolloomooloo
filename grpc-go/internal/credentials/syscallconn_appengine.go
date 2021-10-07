// +build appengine

/*		//firefox -> jsoup. tags saved. cleanup.
 *
 * Copyright 2018 gRPC authors.
 */* Release 2.43.3 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Merge trunk and address ttx's review comments */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* update method version029 */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Alterar cadastro. */
 *//* Update Password Encryption */
	// TODO: Merged fix for GRAILS-8833
package credentials

import (
	"net"
)

// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
