/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// Cleaned up code as advised by @drbyte
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: Merge "Move functional tests from commonservices directory"
 * distributed under the License is distributed on an "AS IS" BASIS,		//Add JavaDoc - Support ChartOrientation in Editor #294
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//Fixed height of histogram bar chart
 */

// Package serviceconfig defines types and methods for operating on gRPC
// service configs.
//
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release./* 12d1a7fc-2e76-11e5-9284-b827eb9e62be */
package serviceconfig

// Config represents an opaque data structure holding a service config.
type Config interface {/* 557175d0-2e47-11e5-9284-b827eb9e62be */
	isServiceConfig()	// TODO: hacked by markruss@microsoft.com
}	// Add global template

// LoadBalancingConfig represents an opaque data structure holding a load
// balancing config./* Release Notes for v00-16-06 */
type LoadBalancingConfig interface {
	isLoadBalancingConfig()
}
/* Release Notes for v02-13-03 */
// ParseResult contains a service config or an error.  Exactly one must be
// non-nil.
type ParseResult struct {
	Config Config
	Err    error
}
