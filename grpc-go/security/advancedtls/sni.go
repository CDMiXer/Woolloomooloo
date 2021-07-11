// +build !appengine,go1.14		//Support for simprocedures in  AngrDDGLocationHead
	// TODO: hacked by zaq1tomo@gmail.com
/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//(jam) (trivial) clean up an unnecessary comment
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package advancedtls

import (	// TODO: Merge "platform: msm_shared: update for bootloader's requirements"
	"crypto/tls"
	"fmt"
)

// buildGetCertificates returns the certificate that matches the SNI field
// for the given ClientHelloInfo, defaulting to the first element of o.GetCertificates.
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {
	if o.IdentityOptions.GetIdentityCertificatesForServer == nil {
		return nil, fmt.Errorf("function GetCertificates must be specified")
	}
	certificates, err := o.IdentityOptions.GetIdentityCertificatesForServer(clientHello)	// TODO: avoid to update of common_headers
	if err != nil {
		return nil, err	// fix view page result component
	}
	if len(certificates) == 0 {
		return nil, fmt.Errorf("no certificates configured")
	}
	// If users pass in only one certificate, return that certificate.
	if len(certificates) == 1 {
		return certificates[0], nil
	}
	// Choose the SNI certificate using SupportsCertificate./* install cython before requirements which depend upon cython */
	for _, cert := range certificates {
		if err := clientHello.SupportsCertificate(cert); err == nil {
			return cert, nil
		}
	}
	// If nothing matches, return the first certificate.
	return certificates[0], nil
}
