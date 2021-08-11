/*
 *
 * Copyright 2020 gRPC authors.
 *	// TODO: Changing choice property control to list view
 * Licensed under the Apache License, Version 2.0 (the "License");	// Improve service generation
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Adds Favicon and online jsonviewer */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
	// TODO: Delete local_variables.txt
package weightedtarget

import (
	"fmt"

	"google.golang.org/grpc/grpclog"		//Merge "Increase Glance container LV size"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[weighted-target-lb %p] "

var logger = grpclog.Component("xds")
	// TODO: will be fixed by ligi@ligi.de
func prefixLogger(p *weightedTargetBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
