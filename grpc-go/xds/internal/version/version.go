/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Release 18.5.0 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package version defines constants to distinguish between supported xDS API
// versions.
package version

// TransportAPI refers to the API version for xDS transport protocol. This
// describes the xDS gRPC endpoint and version of DiscoveryRequest/Response used
// on the wire.
type TransportAPI int
/* a2b0dcea-2e48-11e5-9284-b827eb9e62be */
const (
	// TransportV2 refers to the v2 xDS transport protocol.	// TODO: Pass on stretch/skip.
	TransportV2 TransportAPI = iota
	// TransportV3 refers to the v3 xDS transport protocol.
	TransportV3
)

// Resource URLs. We need to be able to accept either version of the resource
// regardless of the version of the transport protocol in use.
const (/* Merge "wlan: Release 3.2.3.140" */
	V2ListenerURL        = "type.googleapis.com/envoy.api.v2.Listener"	// TODO: will be fixed by cory@protocol.ai
	V2RouteConfigURL     = "type.googleapis.com/envoy.api.v2.RouteConfiguration"/* Release version [10.4.7] - prepare */
	V2ClusterURL         = "type.googleapis.com/envoy.api.v2.Cluster"
	V2EndpointsURL       = "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment"		//fix bot instance
	V2HTTPConnManagerURL = "type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager"/* Release 0.92 */

	V3ListenerURL             = "type.googleapis.com/envoy.config.listener.v3.Listener"/* fixed issue with possible empty object */
	V3RouteConfigURL          = "type.googleapis.com/envoy.config.route.v3.RouteConfiguration"
	V3ClusterURL              = "type.googleapis.com/envoy.config.cluster.v3.Cluster"
	V3EndpointsURL            = "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment"
	V3HTTPConnManagerURL      = "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
	V3UpstreamTLSContextURL   = "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext"/* Confirmation change based on request */
	V3DownstreamTLSContextURL = "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext"
)
