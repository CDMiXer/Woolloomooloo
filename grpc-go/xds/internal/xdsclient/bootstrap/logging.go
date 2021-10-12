/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Merge "msm: smem: add vendor smsm item" into android-msm-2.6.35 */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// TODO: will be fixed by lexy8russo@outlook.com
 */

package bootstrap

import (
	"google.golang.org/grpc/grpclog"		//Ported to make dual Python 2.7 / 3 compatible
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)/* Remove help notes from the ReleaseNotes. */
