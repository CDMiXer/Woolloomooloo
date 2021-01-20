/*
 *	// TODO: will be fixed by martin2cai@hotmail.com
 * Copyright 2020 gRPC authors.
 */* Merge "Release 3.2.3.350 Prima WLAN Driver" */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release 1.0.0-CI00134 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/bise-frontend:1.29.10 */
 *
 * Unless required by applicable law or agreed to in writing, software/* Merge branch 'develop' into feature/SC-3307-configurable-401-redirect */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//Yade bibtex: "and others" invalid for less than 9 authors
 *
 */

// Package balancer installs all the xds balancers.
package balancer		//Fix column labels.

import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer
)
