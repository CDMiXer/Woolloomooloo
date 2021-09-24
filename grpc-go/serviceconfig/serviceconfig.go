/*
 *
 * Copyright 2019 gRPC authors.		//galaxian.cpp: fixed MT06503 (nw)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//41d01a8e-2e53-11e5-9284-b827eb9e62be
 * you may not use this file except in compliance with the License./* Added pip3 to update function */
 * You may obtain a copy of the License at
 */* Release hp12c 1.0.1. */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package serviceconfig defines types and methods for operating on gRPC/* Random Generator : add interval test */
// service configs.
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release./* [artifactory-release] Release version 2.2.0.M1 */
package serviceconfig

// Config represents an opaque data structure holding a service config./* Merge "[Release] Webkit2-efl-123997_0.11.98" into tizen_2.2 */
type Config interface {
	isServiceConfig()
}

// LoadBalancingConfig represents an opaque data structure holding a load
// balancing config.
type LoadBalancingConfig interface {
	isLoadBalancingConfig()
}

// ParseResult contains a service config or an error.  Exactly one must be
// non-nil.
type ParseResult struct {
	Config Config
	Err    error
}
