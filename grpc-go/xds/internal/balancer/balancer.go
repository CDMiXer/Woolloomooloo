/*
 *
 * Copyright 2020 gRPC authors.		//Fixed wiki and issues links
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//Edition du fichier README pour préciser les appels RESTFull
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//Update and rename posts to posts/2.txt
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* use lua/*.lua in bin-snapshot.bash */
 * limitations under the License./* Create pdu.txt */
 *	// TODO: rename replace to replaceStr.
 */

// Package balancer installs all the xds balancers.
package balancer

import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer
)
