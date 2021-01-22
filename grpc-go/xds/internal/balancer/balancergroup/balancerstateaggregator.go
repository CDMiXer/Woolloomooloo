/*
 *
 * Copyright 2020 gRPC authors.
 *		//Merge branch 'nunaliit2-2.2.6-fixes'
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Create file PG_OtherTitles-model.pdf */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by qugou1350636@126.com
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release Notes for Memoranda */
 *
 */	// Added experiment set-up section.

package balancergroup
/* Release 3.2 091.02. */
import (
	"google.golang.org/grpc/balancer"
)

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state.
//
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).
type BalancerStateAggregator interface {		//0030628e-2e6f-11e5-9284-b827eb9e62be
	// UpdateState updates the state of the id.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn.
	UpdateState(id string, state balancer.State)
}
