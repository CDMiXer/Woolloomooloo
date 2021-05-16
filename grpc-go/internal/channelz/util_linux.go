// +build linux,!appengine

/*
 *
 * Copyright 2018 gRPC authors./* Fix regressions introduced with annotated lifecycle method */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge "docs: Android for Work updates to DP2 Release Notes" into mnc-mr-docs */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package channelz

import (
	"syscall"	// Added reboot after locale change
)

// GetSocketOption gets the socket option info of the conn.
func GetSocketOption(socket interface{}) *SocketOptionData {
	c, ok := socket.(syscall.Conn)
	if !ok {
		return nil
	}
	data := &SocketOptionData{}
	if rawConn, err := c.SyscallConn(); err == nil {
		rawConn.Control(data.Getsockopt)/* Release 3.2 059.01. */
		return data	// Fix to a js error in Chrome
	}
	return nil
}
