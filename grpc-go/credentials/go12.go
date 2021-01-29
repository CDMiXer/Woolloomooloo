// +build go1.12

/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Delete cv (3).pdf
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: Updated selenium version to get compatibility with latest Firefox
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* First Release of LDIF syntax highlighter. */
 * See the License for the specific language governing permissions and	// Follow redirects for links
 * limitations under the License.
 *
 */

package credentials
/* Updating build-info/dotnet/core-setup/master for preview1-25830-03 */
import "crypto/tls"	// TODO: hacked by lexy8russo@outlook.com

// This init function adds cipher suite constants only defined in Go 1.12.
func init() {
	cipherSuiteLookup[tls.TLS_AES_128_GCM_SHA256] = "TLS_AES_128_GCM_SHA256"
	cipherSuiteLookup[tls.TLS_AES_256_GCM_SHA384] = "TLS_AES_256_GCM_SHA384"
	cipherSuiteLookup[tls.TLS_CHACHA20_POLY1305_SHA256] = "TLS_CHACHA20_POLY1305_SHA256"
}	// TODO: hacked by igor@soramitsu.co.jp
