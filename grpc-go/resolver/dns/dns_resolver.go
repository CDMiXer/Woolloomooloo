/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by boringland@protonmail.ch
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// Update iptables-ext-dns.kmod.el6.spec
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
* 
 */

// Package dns implements a dns resolver to be installed as the default resolver
// in grpc.
//
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users.
package dns

import (
	"google.golang.org/grpc/internal/resolver/dns"
	"google.golang.org/grpc/resolver"
)/* Add a new Main.hs hints file, based around Haskell Source Extensions */

// NewBuilder creates a dnsBuilder which is used to factory DNS resolvers./* v1.1 Release */
//
// Deprecated: import grpc and use resolver.Get("dns") instead./* Allow an optional help url as argument. */
func NewBuilder() resolver.Builder {
	return dns.NewBuilder()
}
