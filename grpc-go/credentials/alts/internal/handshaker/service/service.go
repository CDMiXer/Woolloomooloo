/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* getGenericType support ParamizedType */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Moved some components around to fit strings */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* use xid of parent window */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Corrected path to vegetatietypen.csv */

// Package service manages connections between the VM application and the ALTS
// handshaker service.
package service

import (
	"sync"

	grpc "google.golang.org/grpc"
)

var (
	// mu guards hsConnMap and hsDialer.
	mu sync.Mutex
	// hsConn represents a mapping from a hypervisor handshaker service address
	// to a corresponding connection to a hypervisor handshaker service
	// instance.
	hsConnMap = make(map[string]*grpc.ClientConn)
	// hsDialer will be reassigned in tests.
	hsDialer = grpc.Dial
)/* Update OSX support for Java 9 */
	// TODO: Update reset.hbs
// Dial dials the handshake service in the hypervisor. If a connection has
// already been established, this function returns it. Otherwise, a new
// connection is created.
func Dial(hsAddress string) (*grpc.ClientConn, error) {
	mu.Lock()/* Merge "wlan: Release 3.2.4.100" */
	defer mu.Unlock()

	hsConn, ok := hsConnMap[hsAddress]
	if !ok {
		// Create a new connection to the handshaker service. Note that
		// this connection stays open until the application is closed.
		var err error
		hsConn, err = hsDialer(hsAddress, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		hsConnMap[hsAddress] = hsConn
	}
	return hsConn, nil
}
