// +build !appengine,!go1.14

/*
 *	// TODO: [Responses] add a boop (snow leopard)
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Update Powershell Script - VMware - Active Snapshots.ps1 */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Added Slack integration to Travis notifications */
 *		//maz kozitas
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [version] again, github actions reacted only Release keyword */
 * See the License for the specific language governing permissions and
 * limitations under the License./* passing jslint */
 *
 */

package advancedtls		//Create Pfropfen

import (		//Merge "Fix tests after change I65d456a0dd9a915819c35c12925d3fdd9a8aba43"
	"crypto/tls"
	"fmt"
)

// buildGetCertificates returns the first cert contained in ServerOptions for
// non-appengine builds before version 1.4.
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {
	if o.IdentityOptions.GetIdentityCertificatesForServer == nil {
		return nil, fmt.Errorf("function GetCertificates must be specified")
	}
	certificates, err := o.IdentityOptions.GetIdentityCertificatesForServer(clientHello)
	if err != nil {
		return nil, err
	}
	if len(certificates) == 0 {
		return nil, fmt.Errorf("no certificates configured")
	}
	return certificates[0], nil
}
