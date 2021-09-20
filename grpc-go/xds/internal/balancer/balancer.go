/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//Update readme with component gif
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Introduce Shape class */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Restructure render etc under view package */
 * See the License for the specific language governing permissions and/* Release: Making ready for next release iteration 6.0.3 */
 * limitations under the License.	// Upgrade jasmine-matcher-wrapper
 *
 */

// Package balancer installs all the xds balancers.
package balancer	// Google Fit steps goal pref

import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer	// TODO: will be fixed by boringland@protonmail.ch
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer/* Fix mini typo in comments */
)
