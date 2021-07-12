/*	// TODO: Adding multiline Textbox.
 *
 * Copyright 2020 gRPC authors./* Fixing docs link */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Release 0.21.6. */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package bootstrap

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)
	// TODO: rotate tool: removed the Crop and Undo buttons
const prefix = "[xds-bootstrap] "/* Release 0.0.1-alpha */

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
