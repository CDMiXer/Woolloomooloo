// +build go1.13/* Version 1.0.1 Released */

/*
 *
 * Copyright 2019 gRPC authors.
 */* Release v0.6.2.1 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* NtDisplayString: Convert Unicode string to OEM. */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "vrouter changes to support fat flow exclude list for ipv4 & ipv6." */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* More progress on unrooted; set loadModule to false */
 */	// Rename DockerCommander to Dockercommander
	// d2ff6200-2e6f-11e5-9284-b827eb9e62be
package dns/* Alize avec patchs lium / compile vérifiée */
	// Stage 1: update all modules needing only revision change for latest
import "net"

func init() {
	filterError = func(err error) error {
		if dnsErr, ok := err.(*net.DNSError); ok && dnsErr.IsNotFound {
			// The name does not exist; not an error.
			return nil
		}
		return err
	}
}
