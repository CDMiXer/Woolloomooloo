/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// Add initializer to assets group, fixes #84
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */		//removed stray comma from last commit

package cdsbalancer

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)		//Fix dependencies version

const prefix = "[cds-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *cdsBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))	// TODO: d337928a-2e4f-11e5-9284-b827eb9e62be
}
