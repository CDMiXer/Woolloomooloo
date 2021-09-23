/*
 *
 * Copyright 2020 gRPC authors.
 *	// TODO: hacked by alex.gaynor@gmail.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Moving method to compute scope to Variable Resolver */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* added Catalan language support */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 2569ce90-2e4d-11e5-9284-b827eb9e62be */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package balancer installs all the xds balancers.
package balancer

import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer/* Update htm.html */
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer	// Delete styleall.css
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer
)
