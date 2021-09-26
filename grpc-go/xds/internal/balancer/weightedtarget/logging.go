/*	// TODO: bundle-size: 9d90a6addea6a405fb2b8cd6361e90a85d6c6936.br (74.38KB)
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// TODO: will be fixed by timnugent@gmail.com
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Create pokemon-omega-ruby-alpha-sapphire */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Added Maven Release badge */
 * See the License for the specific language governing permissions and/* Release from master */
 * limitations under the License.
 */* refresh + corrections for move */
 */	// add talk about stella.report

package weightedtarget

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[weighted-target-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *weightedTargetBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
