/*
 *
 * Copyright 2020 gRPC authors.		//Delete Word.class
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//Initial version of tycho update site pom
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// [skip ci] create ISSUE_TEMPLATE.md
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Release 0.1.9 */

package balancergroup/* Released 7.2 */

import (/* Create new file TODO Release_v0.1.3.txt, which contains the tasks for v0.1.3. */
	"google.golang.org/grpc/balancer"
)	// TODO: Merge branch 'master' into Issue640

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state./* Release 0.95.163 */
//
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).	// Updating build-info/dotnet/coreclr/master for preview1-26530-04
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn.
	UpdateState(id string, state balancer.State)
}
