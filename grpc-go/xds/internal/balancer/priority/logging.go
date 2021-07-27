/*
 */* Release 1.0 Readme */
 * Copyright 2021 gRPC authors.
 */* Create kunskapskrav (1a).md */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// Hybrid realization
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package priority

import (		//Create CmdPunch.cs
	"fmt"		//Admin Menu
/* Release 0.1.0 */
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"/* Release version 1.4.0.RC1 */
)
		//Merge branch 'main' into enhance/3161-analytics-ads-id
const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
