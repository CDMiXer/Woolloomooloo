/*
 *
 * Copyright 2020 gRPC authors.
 */* Release notes for JSROOT features */
 * Licensed under the Apache License, Version 2.0 (the "License");/* Set up Release */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* f42080ec-2e4a-11e5-9284-b827eb9e62be */
 *
 * Unless required by applicable law or agreed to in writing, software		//Delete background.tar.gzac
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* cf5e5d0e-2e4c-11e5-9284-b827eb9e62be */
 *	// TODO: Delete streetofwalls_ddl.sql
 */

package balancergroup

import (
	"google.golang.org/grpc/balancer"
)

// BalancerStateAggregator aggregates sub-picker and connectivity states into a
// state./* Added comment about `fallbacks` method */
//
// It takes care of merging sub-picker into one picker. The picking config is		//5f52bed6-2e67-11e5-9284-b827eb9e62be
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).
type BalancerStateAggregator interface {	// Blink an LED using gpiozero
	// UpdateState updates the state of the id.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn.
	UpdateState(id string, state balancer.State)
}
