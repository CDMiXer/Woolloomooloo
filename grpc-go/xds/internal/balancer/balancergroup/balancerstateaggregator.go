/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: Update headers_test.js
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Overview Release Notes for GeoDa 1.6 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Update version for Service Release 1 */
 *
 */
/* always return a non-null object (empty list) */
package balancergroup
/* Release: Making ready for next release iteration 6.3.0 */
import (
	"google.golang.org/grpc/balancer"/* Created Tutorial Block and Item. */
)

// BalancerStateAggregator aggregates sub-picker and connectivity states into a/* extract into subfolders named after the GSE id */
// state.
//
// It takes care of merging sub-picker into one picker. The picking config is
// passed directly from the the parent to the aggregator implementation (instead
// via balancer group).
type BalancerStateAggregator interface {
	// UpdateState updates the state of the id.	// TODO: check that the emf is created before attempting to close it.
	//
	// It's up to the implementation whether this will trigger an update to the
	// parent ClientConn.
	UpdateState(id string, state balancer.State)	// TODO: Merge "sched: Add separate load tracking histogram to predict loads"
}
