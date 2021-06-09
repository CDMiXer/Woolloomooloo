/*/* Static initialization of the final fields of a class. */
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Add SgSList::Include
 *     http://www.apache.org/licenses/LICENSE-2.0/* Fix for user.name sorting */
 */* Create header_coverter.pl */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package priority

import (
	"fmt"
/* Merge "Yangman - add module list to left part" */
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)
/* Release CAPO 0.3.0-rc.0 image */
const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
