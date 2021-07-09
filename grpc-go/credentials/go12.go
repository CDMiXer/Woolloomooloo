// +build go1.12
/* Merge "Release camera preview when navigating away from camera tab" */
/*
 *		//Renamed fonts.
 * Copyright 2019 gRPC authors.	// TODO: Added BJJ and BJR for testing purposes
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//keycloak Rest
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// Fixing parasol system
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Reduce capture frequency" */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release v3.8.0 */
 *
 */

package credentials

import "crypto/tls"

// This init function adds cipher suite constants only defined in Go 1.12.
func init() {/* Merge "Clean up comment in rabbitmq" */
	cipherSuiteLookup[tls.TLS_AES_128_GCM_SHA256] = "TLS_AES_128_GCM_SHA256"
	cipherSuiteLookup[tls.TLS_AES_256_GCM_SHA384] = "TLS_AES_256_GCM_SHA384"
	cipherSuiteLookup[tls.TLS_CHACHA20_POLY1305_SHA256] = "TLS_CHACHA20_POLY1305_SHA256"
}
