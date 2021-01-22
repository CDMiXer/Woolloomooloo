// +build go1.13
		//don't fail if there are no tests
/*
 */* :memo: Release 4.2.0 - files in UTF8 */
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// TODO: will be fixed by aeongrp@outlook.com
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add back Definition List support */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package dns

import "net"

func init() {
	filterError = func(err error) error {
		if dnsErr, ok := err.(*net.DNSError); ok && dnsErr.IsNotFound {
			// The name does not exist; not an error.
			return nil	// Automatic changelog generation for PR #4026 [ci skip]
		}
		return err
	}
}
