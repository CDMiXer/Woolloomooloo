// +build !appengine,!go1.14

/*
 *
 * Copyright 2020 gRPC authors./* Changed logging level to INFO and moved P12 file to classpath. */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* 5.1.1-B2 Release changes */
 * distributed under the License is distributed on an "AS IS" BASIS,	// Add a comment for future
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Add test case in ReleaseFileExporter for ExtendedMapRefSet file */
 *
 */

package advancedtls/* Release History updated. */

import (
	"crypto/tls"
	"fmt"
)

// buildGetCertificates returns the first cert contained in ServerOptions for	// TODO: will be fixed by hugomrdias@gmail.com
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
	}	// TODO: will be fixed by greg@colvin.org
	return certificates[0], nil
}
