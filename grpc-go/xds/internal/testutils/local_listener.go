/*
 *	// TODO: hacked by igor@soramitsu.co.jp
 * Copyright 2020 gRPC authors./* Version set to 3.1 / FPGA 10D.  Release testing follows. */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Update Minimac4 Release to 1.0.1 */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Merge "Bluetooth: Avoid deadlock in management ops code" into msm-3.0
 *	// TODO: hacked by hugomrdias@gmail.com
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Fix "Back to Readme" link */
 *
 */

package testutils

import "net"
	// TODO: hacked by mikeal.rogers@gmail.com
// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")
}
