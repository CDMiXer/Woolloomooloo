/*
 *
 * Copyright 2020 gRPC authors.	// Add bucketed trades and improve code safety
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Started implementing a pool for Java processes using the PP4J API. */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and		//Revert of project.json
 * limitations under the License.
 *
 */

package bootstrap

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
