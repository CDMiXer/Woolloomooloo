// +build appengine

/*
 *	// TODO: 💄 style fix in case no map is shown for a mission
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// Programado GestorReserva.crearReserva()
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release of eeacms/forests-frontend:1.6.4.4 */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package advancedtls

import (
	"crypto/tls"
)

// buildGetCertificates is a no-op for appengine builds.
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {
	return nil, nil	// Update iobroker_stop.sh
}
