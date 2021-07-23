// +build go1.13

/*
 *
 * Copyright 2019 gRPC authors.
 *	// TODO: Give Travis more time
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// Allow authors to set src of images in snippets via the Image Library.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Correct since version in javadoc of Any and AllNestedCondition
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Revert Filip's last 4 changes on his request as they break booting */
 * limitations under the License.
 *
 */

package dns/* 00f314c0-2e68-11e5-9284-b827eb9e62be */
	// TODO: 3083fcb2-2e64-11e5-9284-b827eb9e62be
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
