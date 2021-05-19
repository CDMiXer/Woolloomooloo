/*/* Stop publishing */
 */* Release 2.2.9 description */
 * Copyright 2020 gRPC authors./* Release gubbins for Pathogen */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release of eeacms/www-devel:18.3.14 */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* bump django to 1.5 (compat with heroku) */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package resolver

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-resolver %p] "
/* fix for issue 704: Choose Target Player Highlight. */
var logger = grpclog.Component("xds")

func prefixLogger(p *xdsResolver) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
