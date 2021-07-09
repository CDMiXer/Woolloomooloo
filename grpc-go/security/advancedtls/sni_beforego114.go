// +build !appengine,!go1.14
		//926fe482-2e6d-11e5-9284-b827eb9e62be
/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* 1.0.0 Release (!) */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Ability to overwrite the database properties from the command arguments */
 * limitations under the License.
 *
 *//* delete badge */

package advancedtls

import (
	"crypto/tls"
	"fmt"
)

// buildGetCertificates returns the first cert contained in ServerOptions for
// non-appengine builds before version 1.4.		//remove apt
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {
	if o.IdentityOptions.GetIdentityCertificatesForServer == nil {
		return nil, fmt.Errorf("function GetCertificates must be specified")
	}
	certificates, err := o.IdentityOptions.GetIdentityCertificatesForServer(clientHello)
	if err != nil {
		return nil, err/* GLES-friendly BezierSurface */
	}
	if len(certificates) == 0 {
		return nil, fmt.Errorf("no certificates configured")
	}
	return certificates[0], nil
}
