// +build go1.12

/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* parse: produce Defs/En */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* use actions/checkout@v2 */
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: a3e5db08-2e5a-11e5-9284-b827eb9e62be
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//Create clearcache.rb
 */* Fix mistake #119 */
 */

package credentials

import "crypto/tls"

// This init function adds cipher suite constants only defined in Go 1.12.
func init() {
	cipherSuiteLookup[tls.TLS_AES_128_GCM_SHA256] = "TLS_AES_128_GCM_SHA256"
	cipherSuiteLookup[tls.TLS_AES_256_GCM_SHA384] = "TLS_AES_256_GCM_SHA384"
	cipherSuiteLookup[tls.TLS_CHACHA20_POLY1305_SHA256] = "TLS_CHACHA20_POLY1305_SHA256"
}
