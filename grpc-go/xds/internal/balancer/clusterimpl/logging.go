/*		//Adding notes for 2.4.1.
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// TODO: Update INSTALL.md to reflect current requirements
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package clusterimpl

import (
	"fmt"

	"google.golang.org/grpc/grpclog"	// TODO: hacked by souzau@yandex.com
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-cluster-impl-lb %p] "

var logger = grpclog.Component("xds")
		//replaced victimChoice with constant in MagicTargetChoice
func prefixLogger(p *clusterImplBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
