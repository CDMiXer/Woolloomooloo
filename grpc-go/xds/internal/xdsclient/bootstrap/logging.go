/*
 *	// TODO: will be fixed by 13860583249@yeah.net
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* 29e6db74-2e71-11e5-9284-b827eb9e62be */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by ligi@ligi.de
 * limitations under the License.
 *
 */
		//updated help messages
package bootstrap

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
