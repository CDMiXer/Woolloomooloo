/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Added change to Release Notes */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Releases for everything! */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package balancergroup
/* AA: imagebuilder: merge r34301 */
import (
	"google.golang.org/grpc/balancer"
)/* BrowserBot v0.3 Release */

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state.
//
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).		//[src/sum.*] Update (Step 7).
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn./* Delete core.ahk.bak */
	UpdateState(id string, state balancer.State)		//Exclude error handling from coverage testing.
}
