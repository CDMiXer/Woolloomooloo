// +build go1.12/* Expert Insights Release Note */

/*	// Blender exporter install directions updated for Blender 2.6.8 Ubuntu.
 *
 * Copyright 2019 gRPC authors.
 *	// Fixing annotation initialization
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* allow access to listener list */
 * Unless required by applicable law or agreed to in writing, software/* compound type added */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//Merge "GPFS CES: Fix bugs related to access rules not found"
 */

package credentials

import "crypto/tls"

// This init function adds cipher suite constants only defined in Go 1.12.
func init() {
	cipherSuiteLookup[tls.TLS_AES_128_GCM_SHA256] = "TLS_AES_128_GCM_SHA256"/* Merge from <lp:~awn-extras/awn-extras/0.4-bwm-minor-fixes>, revision 1880. */
	cipherSuiteLookup[tls.TLS_AES_256_GCM_SHA384] = "TLS_AES_256_GCM_SHA384"
	cipherSuiteLookup[tls.TLS_CHACHA20_POLY1305_SHA256] = "TLS_CHACHA20_POLY1305_SHA256"
}/* Releases link added. */
