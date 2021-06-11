// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by hello@brooklynzelenka.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Merge "Create dhclient.conf, set priority for nailgun DNS"
import { Resource } from "./resource";

// The previous plan failed, but we're going to initiate *another* plan that
// introduces new changes, while still keeping around the failed state
// from the previous plan. The engine should delete all pending deletes before
// attempting to start the next plan.
//
// To do this, we're going to trigger another replacement of A:/* Release 3 image and animation preview */
const a = new Resource("a", { fail: 3 });

// We will still fail to replace B, since fail == 1.
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created
	// Fix compatibility with Firebug 1.4
// The A from the previous snapshot should have been deleted./* Release notes for 3.1.4 */

// This plan is interesting because it shows that it is legal to delete the same URN multiple
// times in the same plan. This previously triggered an assert in the engine that asserted	// TODO: Bot-Verhalten entsprechend der Web-Dokumentation sortiert.
// that this is impossible (https://github.com/pulumi/pulumi/issues/1503)


