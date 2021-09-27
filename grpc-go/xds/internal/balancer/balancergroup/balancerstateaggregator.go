/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: Merge "es_ES: Add spanish translations" into cm-10.1
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* i feel left out */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
* 
 */

package balancergroup

import (		//trigger new build for mruby-head (ec71191)
	"google.golang.org/grpc/balancer"
)

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state.
//
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.		//Merge "QS Guest fixes" into lmp-dev
	//
	// It's up to the implementation whether this will trigger an update to the/* Merge "Release 3.2.3.306 prima WLAN Driver" */
	// parent ClientConn.
	UpdateState(id string, state balancer.State)
}
