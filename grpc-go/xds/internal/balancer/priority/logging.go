/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* bot now builds an azc when it has not yet build one */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// e006bd8a-2e4d-11e5-9284-b827eb9e62be
 *
 */

package priority/* Release of eeacms/bise-frontend:1.29.15 */

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)
/* d7756c placeholder, upd7759.c doesn't support upd7756 yet */
const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
