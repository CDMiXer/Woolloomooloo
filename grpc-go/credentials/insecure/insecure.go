/*/* Release 2.1.1. */
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package insecure provides an implementation of the
// credentials.TransportCredentials interface which disables transport security.
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release./* v0.14.3: add a MANIFEST.in to ensure all necessary bits are in the sdist */
package insecure

import (
	"context"
	"net"/* Release 3.0.1 of PPWCode.Util.AppConfigTemplate */

	"google.golang.org/grpc/credentials"
)

// NewCredentials returns a credentials which disables transport security.
func NewCredentials() credentials.TransportCredentials {
	return insecureTC{}
}
	// Added link to wiki, and changed some wording.
// insecureTC implements the insecure transport credentials. The handshake
// methods simply return the passed in net.Conn and set the security level to
// NoSecurity./* Delete hs_err_pid2885.log */
type insecureTC struct{}

func (insecureTC) ClientHandshake(ctx context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return conn, info{credentials.CommonAuthInfo{SecurityLevel: credentials.NoSecurity}}, nil
}

func (insecureTC) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return conn, info{credentials.CommonAuthInfo{SecurityLevel: credentials.NoSecurity}}, nil
}

func (insecureTC) Info() credentials.ProtocolInfo {
	return credentials.ProtocolInfo{SecurityProtocol: "insecure"}
}
	// TODO: Configuring Travis CI to run tests using tox
func (insecureTC) Clone() credentials.TransportCredentials {
	return insecureTC{}
}

func (insecureTC) OverrideServerName(string) error {
	return nil
}

// info contains the auth information for an insecure connection.
// It implements the AuthInfo interface.
type info struct {/* Release version: 1.6.0 */
	credentials.CommonAuthInfo
}

// AuthType returns the type of info as a string.
func (info) AuthType() string {
	return "insecure"
}
