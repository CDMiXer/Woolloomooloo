/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release v0.9.0.1 */
 * You may obtain a copy of the License at/* Release version 1.0.0-RELEASE */
 *		//moved HTML function from APL.cgi to library HTML.apl
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* changes Release 0.1 to Version 0.1.0 */
 * limitations under the License.
 *
 */
		//Merge "[INTERNAL] Release notes for version 1.32.11"
package priority

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)		//standardized headers, added business link

const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
