/*
 */* NODE17 Release */
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Prepare Release 0.5.11 */
 */* Release 3.16.0 */
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* MS office font compatible mapping */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* changed abstraction cloners to flatteners */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package balancergroup
/* Added datatypes */
import (
	"google.golang.org/grpc/balancer"
)

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state.
//
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn.
	UpdateState(id string, state balancer.State)/* Update st2.yaml */
}
