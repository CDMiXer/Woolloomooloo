/*
 *
 * Copyright 2020 gRPC authors./* Created requirements file. */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: New translations p01_ch04_pref.md (Italian)
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
/* Release of eeacms/plonesaas:5.2.1-57 */
package resolver	// TODO: Set required version of bash
/* Trunk is now 1.7 */
import (
	"fmt"
	// TODO: trigger new build for jruby-head (7252d05)
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-resolver %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *xdsResolver) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
