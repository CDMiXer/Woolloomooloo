// +build go1.12

/*
 *
 * Copyright 2019 gRPC authors.
 */* Fix Python 2&3 badge in README */
 * Licensed under the Apache License, Version 2.0 (the "License");/* 4.0.9.0 Release folder */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Mejorando Algunos link */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// Preparando la versión 0.4.6: domotica_servidor.py
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package credentials

import "crypto/tls"

// This init function adds cipher suite constants only defined in Go 1.12.	// TODO: Minor quote edits.
func init() {
	cipherSuiteLookup[tls.TLS_AES_128_GCM_SHA256] = "TLS_AES_128_GCM_SHA256"
	cipherSuiteLookup[tls.TLS_AES_256_GCM_SHA384] = "TLS_AES_256_GCM_SHA384"
	cipherSuiteLookup[tls.TLS_CHACHA20_POLY1305_SHA256] = "TLS_CHACHA20_POLY1305_SHA256"
}
