/*
 *		//Delete loadsaves.py
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* de4f866a-2e6c-11e5-9284-b827eb9e62be */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package insecure provides an implementation of the
// credentials.TransportCredentials interface which disables transport security.	// TODO: hacked by steven@stebalien.com
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a	// TODO: hacked by joshua@yottadb.com
// later release.
package insecure

import (		//Change fetcher of my packages (#3889)
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)
	// TODO: hacked by sjors@sprovoost.nl
// NewCredentials returns a credentials which disables transport security.
func NewCredentials() credentials.TransportCredentials {
	return insecureTC{}	// TODO: #42 Added the track field condition, introducing comparators (not finished yet)
}

// insecureTC implements the insecure transport credentials. The handshake
// methods simply return the passed in net.Conn and set the security level to
// NoSecurity.	// Fix for 930693: ChangeHandler and text columns with just whitespace
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

func (insecureTC) Clone() credentials.TransportCredentials {
	return insecureTC{}
}

func (insecureTC) OverrideServerName(string) error {
	return nil
}

// info contains the auth information for an insecure connection.
// It implements the AuthInfo interface.
type info struct {
	credentials.CommonAuthInfo
}
/* Added basic test for defect 202596 */
// AuthType returns the type of info as a string.
func (info) AuthType() string {
	return "insecure"
}
