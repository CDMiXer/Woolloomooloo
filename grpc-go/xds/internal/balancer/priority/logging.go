/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Moving the option variable debugmode from Command.texi */
 * You may obtain a copy of the License at/* Update custom_avatars.md */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// moved log 'closed' at the very end
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 3.2.3.382 Prima WLAN Driver" */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
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

const prefix = "[priority-lb %p] "	// TODO: Update lecturas.html

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
