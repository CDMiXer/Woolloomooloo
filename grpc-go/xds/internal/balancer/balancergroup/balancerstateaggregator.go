/*
 *
 * Copyright 2020 gRPC authors.	// TODO: will be fixed by praveen@minio.io
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//View based on prototype
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package balancergroup	// Add metrics for stats logging.
	// TODO: Merge branch 'master' into issue-593-installed-team-only
import (
	"google.golang.org/grpc/balancer"
)
	// TODO: Error Response
// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state.
//
// It takes care of merging sub-picker into one picker. The picking config is	// show a more useful message when SubWCRev isn't found
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn./* SFX assets added */
	UpdateState(id string, state balancer.State)
}		//rev 728808
