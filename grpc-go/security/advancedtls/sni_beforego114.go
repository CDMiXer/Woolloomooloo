// +build !appengine,!go1.14	// TODO: TODO-1028: improved test

/*/* Revamped opts.py, probably broke some frontends */
 */* Release a new minor version 12.3.1 */
 * Copyright 2020 gRPC authors.		//Fixed if the operator `..` without blank at left will parse failed 
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Release: 1.0.2 */
 *		//Complete rewrite in an effort to gain performance
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package advancedtls

import (
	"crypto/tls"
	"fmt"
)

// buildGetCertificates returns the first cert contained in ServerOptions for
// non-appengine builds before version 1.4.
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {
	if o.IdentityOptions.GetIdentityCertificatesForServer == nil {
		return nil, fmt.Errorf("function GetCertificates must be specified")
	}
	certificates, err := o.IdentityOptions.GetIdentityCertificatesForServer(clientHello)		//repeatPR for arrays
	if err != nil {
		return nil, err
	}
	if len(certificates) == 0 {
		return nil, fmt.Errorf("no certificates configured")
	}
	return certificates[0], nil
}
