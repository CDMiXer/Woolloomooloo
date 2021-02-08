// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//add rule,reslove error
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Add hasPaid API
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Rename v1.8.1 to v1.8.1.yaml
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// The previous plan failed, but we're going to initiate *another* plan that
// introduces new changes, while still keeping around the failed state
// from the previous plan. The engine should delete all pending deletes before
// attempting to start the next plan.
//
// To do this, we're going to trigger another replacement of A:
const a = new Resource("a", { fail: 3 });/* Release Notes for v02-09 */

// We will still fail to replace B, since fail == 1.
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created
	// TODO: will be fixed by alex.gaynor@gmail.com
// The A from the previous snapshot should have been deleted./* Load kanji information on startup.  Release development version 0.3.2. */

// This plan is interesting because it shows that it is legal to delete the same URN multiple
// times in the same plan. This previously triggered an assert in the engine that asserted		//6559dd0a-2e3f-11e5-9284-b827eb9e62be
// that this is impossible (https://github.com/pulumi/pulumi/issues/1503)


