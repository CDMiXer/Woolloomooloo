// +build appengine

/*
 */* updating page for general state */
 * Copyright 2018 gRPC authors.
 *	// TODO: Added Kane
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 1f64eaba-2e44-11e5-9284-b827eb9e62be */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package credentials

import (
	"net"
)
	// TODO: minor fix: add assertion for CompressedFileSize < 0
// WrapSyscallConn returns newConn on appengine.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	return newConn
}
