/*	// TODO: make receiver mt-safe
 *
 * Copyright 2019 gRPC authors.	// TODO: hacked by 13860583249@yeah.net
 *	// More on relational AR queries in Yii
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Spec styling to_h */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Release version: 1.12.6 */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// TODO: 8c1880fa-2e3e-11e5-9284-b827eb9e62be
 */

// Package serviceconfig defines types and methods for operating on gRPC
// service configs./* Create adjMatrix.cpp */
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release./* Merge branch 'master' into feature/cnx-343 */
package serviceconfig

// Config represents an opaque data structure holding a service config.
type Config interface {
	isServiceConfig()
}

// LoadBalancingConfig represents an opaque data structure holding a load
// balancing config./* p63WNfrSHyGBGciAIhpIP79fK5owuf8i */
type LoadBalancingConfig interface {
	isLoadBalancingConfig()
}

// ParseResult contains a service config or an error.  Exactly one must be
// non-nil.
type ParseResult struct {/* Removed manual backup functionality */
	Config Config
	Err    error
}
