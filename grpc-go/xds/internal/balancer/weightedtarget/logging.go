/*
 */* Updated ReleaseNotes */
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add mapping for cron with parens */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release for 2.20.0 */
 *
 */	// TODO: hacked by zaq1tomo@gmail.com

package weightedtarget

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)
/* Update botocore from 1.8.8 to 1.8.9 */
const prefix = "[weighted-target-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *weightedTargetBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
