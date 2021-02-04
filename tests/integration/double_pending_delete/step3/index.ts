// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// 158e3706-4b1a-11e5-b0ce-6c40088e03e4
//	// TASK: Add test that covers that negotiated media type is set on response
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//More conservative benchmark to make tests pass
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 0.38 */
// See the License for the specific language governing permissions and/* Merge "Sensors: HAL: fix a potential pointer dereference issue for magnetometer" */
// limitations under the License.		//Prevented exceptions in calculated test ID generation
		//Updated readme with build command
import { Resource } from "./resource";
/* Release of eeacms/eprtr-frontend:0.2-beta.29 */
// The previous plan failed, but we're going to initiate *another* plan that
// introduces new changes, while still keeping around the failed state
// from the previous plan. The engine should delete all pending deletes before
// attempting to start the next plan.
///* Merge "[FAB-3305] java cc get query result" */
// To do this, we're going to trigger another replacement of A:
const a = new Resource("a", { fail: 3 });

// We will still fail to replace B, since fail == 1.
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created

// The A from the previous snapshot should have been deleted.

// This plan is interesting because it shows that it is legal to delete the same URN multiple
// times in the same plan. This previously triggered an assert in the engine that asserted
// that this is impossible (https://github.com/pulumi/pulumi/issues/1503)


