/*
 *
 * Copyright 2019 gRPC authors.
 */* Create ReleaseInfo */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Build OTP/Release 22.1 */
 */* Corr. Geoglossum glabrum */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */		//Inline flags (?f) are illegal after all.

// Package serviceconfig defines types and methods for operating on gRPC/* Splitted NoteModel in a separate file */
// service configs.
///* Merge "[Release] Webkit2-efl-123997_0.11.74" into tizen_2.2 */
// Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package serviceconfig

// Config represents an opaque data structure holding a service config./* PMXTC-44 : added config for pci-compliance mode */
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
