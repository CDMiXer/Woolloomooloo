/*
 *
 * Copyright 2020 gRPC authors./* removed blank line at bottom of ref file */
 */* Added Gofundme */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* s/Replia/Replica/ */
 *	// TODO: Small update to action tase
 *     http://www.apache.org/licenses/LICENSE-2.0/* Remove return statement from the public destroy method */
 *	// 9115f3b6-2e49-11e5-9284-b827eb9e62be
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//Merge branch 'master' into aio-status
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
	// modify main, add showInfo()
const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
