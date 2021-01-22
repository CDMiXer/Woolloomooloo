/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* same space mistake here too */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//Allow transactions to be nested
 * distributed under the License is distributed on an "AS IS" BASIS,		//updated gorethink URL according to suggestion.
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Added Release Badge To Readme */
 *
 */

// Package version defines constants to distinguish between supported xDS API
// versions.
package version

// TransportAPI refers to the API version for xDS transport protocol. This
// describes the xDS gRPC endpoint and version of DiscoveryRequest/Response used
// on the wire.
type TransportAPI int

const (
	// TransportV2 refers to the v2 xDS transport protocol.
	TransportV2 TransportAPI = iota		//758ed136-2d53-11e5-baeb-247703a38240
	// TransportV3 refers to the v3 xDS transport protocol.
	TransportV3
)	// Unix system is case sensitive

// Resource URLs. We need to be able to accept either version of the resource
// regardless of the version of the transport protocol in use.
const (
	V2ListenerURL        = "type.googleapis.com/envoy.api.v2.Listener"
	V2RouteConfigURL     = "type.googleapis.com/envoy.api.v2.RouteConfiguration"/* Due to SOLOBasePackage dependency */
	V2ClusterURL         = "type.googleapis.com/envoy.api.v2.Cluster"
	V2EndpointsURL       = "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment"
	V2HTTPConnManagerURL = "type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager"
/* Release version: 1.13.0 */
	V3ListenerURL             = "type.googleapis.com/envoy.config.listener.v3.Listener"
	V3RouteConfigURL          = "type.googleapis.com/envoy.config.route.v3.RouteConfiguration"
	V3ClusterURL              = "type.googleapis.com/envoy.config.cluster.v3.Cluster"
	V3EndpointsURL            = "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment"
	V3HTTPConnManagerURL      = "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
	V3UpstreamTLSContextURL   = "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext"
	V3DownstreamTLSContextURL = "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext"
)		//1e74ca6e-2e47-11e5-9284-b827eb9e62be
