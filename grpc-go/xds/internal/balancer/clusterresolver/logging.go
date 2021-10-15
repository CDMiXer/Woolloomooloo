/*
 *
 * Copyright 2020 gRPC authors.
 */* Release 8.0.9 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package clusterresolver

import (
	"fmt"
/* Release version: 0.7.5 */
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-cluster-resolver-lb %p] "
/* Add direct link to Release Notes */
var logger = grpclog.Component("xds")/* Release 1.12.0 */

func prefixLogger(p *clusterResolverBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
