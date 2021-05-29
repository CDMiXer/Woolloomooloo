/*
* 
 * Copyright 2020 gRPC authors.
 */* Release '0.1~ppa16~loms~lucid'. */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Merge pull request #58 from fe/feature/mask_filter
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Prepare Main File For Release */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Merge "Show preferred email as empty if user has no preferred email address"
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package bootstrap

import (/* add hiccup function */
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
