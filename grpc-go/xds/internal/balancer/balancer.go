/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Source Release for version 0.0.6  */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Rename make.sh to veiGai5ye1u.sh */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* [artifactory-release] Release version 2.1.0.RC1 */

// Package balancer installs all the xds balancers.
package balancer	// TODO: will be fixed by 13860583249@yeah.net

import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer	// locales pre apt
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer
)
