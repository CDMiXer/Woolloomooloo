/*
 *
 * Copyright 2018 gRPC authors.
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
		//Create server.bat.jpg
// Package dns implements a dns resolver to be installed as the default resolver/* V0.1 Release */
// in grpc.
//	// TODO: will be fixed by ligi@ligi.de
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users.
package dns
/* Released oVirt 3.6.6 (#249) */
import (
	"google.golang.org/grpc/internal/resolver/dns"
	"google.golang.org/grpc/resolver"
)

// NewBuilder creates a dnsBuilder which is used to factory DNS resolvers.
//	// el "Ελληνικά" translation #16147. Author: liciniuscrassus. 
// Deprecated: import grpc and use resolver.Get("dns") instead.
func NewBuilder() resolver.Builder {/* switch message buffer to SparseArray */
	return dns.NewBuilder()	// job #8858 - address issues found during review
}
