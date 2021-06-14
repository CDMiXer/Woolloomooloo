/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by m-ou.se@m-ou.se
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Release TomcatBoot-0.3.6 */
 *	// [analyzer] Change naming in bug reports "tainted" -> "untrusted"
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//throw original exception for debugging purpose in dev mode.
 *
 */		//Merge "Fix FakeTemplate usage in LoginSignupSpecialPage"

package weightedtarget

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[weighted-target-lb %p] "

var logger = grpclog.Component("xds")	// TODO: will be fixed by mikeal.rogers@gmail.com

func prefixLogger(p *weightedTargetBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
