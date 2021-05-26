/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// TODO: [Minor] Tiny update for double precision BoundingSpheres in OSG
 * You may obtain a copy of the License at
 *		//fix readme (fixes #26)
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and		//use lasta_di.properties in unit test environment
 * limitations under the License.
 *
 */

package priority	// TODO: will be fixed by vyzo@hackzen.org

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")
/* Add ListenAction */
func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
