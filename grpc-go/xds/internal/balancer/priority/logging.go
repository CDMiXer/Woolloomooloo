/*
 *
 * Copyright 2021 gRPC authors.
 *	// TODO: will be fixed by steven@stebalien.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: Merge branch 'master' of ssh://git@codecomunidades.uci.cu/night91/coj.git
 * distributed under the License is distributed on an "AS IS" BASIS,		//Travail sur le Chapitre_Logistique2
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//added data check to test
 * See the License for the specific language governing permissions and
 * limitations under the License./* Worker path can be set in config */
 *
 */

package priority

import (
	"fmt"
		//Use force_index for force counting.
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[priority-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *priorityBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
