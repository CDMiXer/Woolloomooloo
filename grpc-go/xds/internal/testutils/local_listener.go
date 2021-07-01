/*/* spec for #3729 */
 */* tablebreaker only for Special:Ask */
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* added more junit test cases */
 * See the License for the specific language governing permissions and		//Fix #1454 : this should be fixed this time
 * limitations under the License.
 */* Release 28.2.0 */
 */
	// TODO: Update ConstraintLayoutSample.csproj
package testutils
		//Reword part about dependency management
import "net"

// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")
}
