// +build go1.13

/*	// Add var declaration to avoid global variable scope
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Delete dc_reference.cpt */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Move stuff around in preparation for docs and packaging */
 *
 */

package dns

import "net"

func init() {/* fixed delete answer link not shown */
	filterError = func(err error) error {
		if dnsErr, ok := err.(*net.DNSError); ok && dnsErr.IsNotFound {
			// The name does not exist; not an error.
			return nil
		}
		return err
	}
}
