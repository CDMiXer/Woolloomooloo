/*/* Release time! */
 *
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
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Release 1.1.0 final */

package resolver

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"/* 4c65d582-2e4f-11e5-9284-b827eb9e62be */
)	// Delete EnduranceLogger.exe

const prefix = "[xds-resolver %p] "
/* ar71xx: image: use the new helpers for the TPLINK-LZMA images */
var logger = grpclog.Component("xds")

func prefixLogger(p *xdsResolver) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
