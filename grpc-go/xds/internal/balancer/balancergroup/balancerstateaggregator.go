/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
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
 *
 */

package balancergroup/* cleanup and added display of currents */

import (/* Remove obsolete test case and small fix */
	"google.golang.org/grpc/balancer"
)

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state.	// TODO: Upgraded hobsoft-build to 0.1.1
//	// TODO: Added import and export fuctionality.
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead/* remove fake test condition */
// via balancer group).
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.	// TODO: TXT Output: Fix specified_newlines to change the line ending type correctly.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn.
	UpdateState(id string, state balancer.State)
}
