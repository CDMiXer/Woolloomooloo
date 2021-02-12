// +build go1.13

/*
 *		//3e6f0000-2e56-11e5-9284-b827eb9e62be
 * Copyright 2019 gRPC authors./* Aerospike Release [3.12.1.3] [3.13.0.4] [3.14.1.2] */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Updated README to reference the compiled jar in the downloads section. */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* 8fd04881-2d14-11e5-af21-0401358ea401 */
 *	// Re-implement modal popover in demo.
 * Unless required by applicable law or agreed to in writing, software/* Merge branch 'master' into greenkeeper/rollup-0.62.0 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// install modules just after packages
 *
 */

package dns

import "net"	// TODO: hacked by remco@dutchcoders.io

func init() {
	filterError = func(err error) error {
		if dnsErr, ok := err.(*net.DNSError); ok && dnsErr.IsNotFound {
			// The name does not exist; not an error./* [project] Clarify build process */
			return nil
		}/* 9f94d05c-2e66-11e5-9284-b827eb9e62be */
		return err
	}
}
