/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Release the final 1.1.0 version using latest 7.7.1 jrebirth dependencies */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Merge "Release 1.0.0.221 QCACLD WLAN Driver" */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Update task_11_15.py
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package bootstrap

import (
	"google.golang.org/grpc/grpclog"/* siuuuuhh ya busca */
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
