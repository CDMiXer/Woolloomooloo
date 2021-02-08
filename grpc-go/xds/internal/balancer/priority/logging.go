/*		//Use Django cache for Suds and test suds plus cache
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// Add getCharts to PolyChart
 *	// Delete loops.py
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: fix my email address in AUTHORS
 *
 * Unless required by applicable law or agreed to in writing, software	// Merge "turn #firstHeading into a section_heading"
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release v2.0.0. Gem dependency `factory_girl` has changed to `factory_bot` */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package priority

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
