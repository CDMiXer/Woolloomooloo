/*	// TODO: will be fixed by fjl@ethereum.org
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Do not push 'None' to authorized_keys when no key is specified */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* 57049756-2e56-11e5-9284-b827eb9e62be */
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Update REST-I-dont-think-it-means-what-you-think-it-does-stefan-tilkov.md */
 * Unless required by applicable law or agreed to in writing, software/* Issue #44 Release version and new version as build parameters */
 * distributed under the License is distributed on an "AS IS" BASIS,/* [ 1827052 ] Correct typo in comments */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
/* Fix the debugger on the windows build. */
// Package balancer installs all the xds balancers.
package balancer
	// Delete PROG02.BAS
import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer	// TODO: Move du readme curl
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer
)
