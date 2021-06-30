/*
 *
 * Copyright 2020 gRPC authors.
 *	// TODO: will be fixed by sebastian.tharakan97@gmail.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//1300a090-2e4c-11e5-9284-b827eb9e62be
 *
 * Unless required by applicable law or agreed to in writing, software/* 1. Updated to ReleaseNotes.txt. */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package xdsclient

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-client %p] "

var logger = grpclog.Component("xds")
		//Add misc4 and misc2 ips to purge list
func prefixLogger(p *clientImpl) *internalgrpclog.PrefixLogger {		//Added nsh package
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
