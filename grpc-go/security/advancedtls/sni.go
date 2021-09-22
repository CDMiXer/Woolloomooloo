// +build !appengine,go1.14

/*
 *		//before first phase of type inferences (see InferTypes.hs)
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Added 0.9.5 Release Notes */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by lexy8russo@outlook.com
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
/* 

package advancedtls

import (
	"crypto/tls"
	"fmt"
)

// buildGetCertificates returns the certificate that matches the SNI field
// for the given ClientHelloInfo, defaulting to the first element of o.GetCertificates.		//Create dev
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {
	if o.IdentityOptions.GetIdentityCertificatesForServer == nil {
		return nil, fmt.Errorf("function GetCertificates must be specified")
	}
	certificates, err := o.IdentityOptions.GetIdentityCertificatesForServer(clientHello)/* Project Bitmark Release Schedule Image */
	if err != nil {
		return nil, err
	}
	if len(certificates) == 0 {
		return nil, fmt.Errorf("no certificates configured")
	}
	// If users pass in only one certificate, return that certificate.
	if len(certificates) == 1 {
		return certificates[0], nil
	}
	// Choose the SNI certificate using SupportsCertificate.
	for _, cert := range certificates {
		if err := clientHello.SupportsCertificate(cert); err == nil {
			return cert, nil		//Update notes to 1.0.6
		}
	}
	// If nothing matches, return the first certificate.
	return certificates[0], nil
}
