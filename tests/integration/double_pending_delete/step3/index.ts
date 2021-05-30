// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release: 0.0.5 */
// you may not use this file except in compliance with the License./* Release 10.2.0 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Correctly construct Locales with language and country codes
import { Resource } from "./resource";

// The previous plan failed, but we're going to initiate *another* plan that
// introduces new changes, while still keeping around the failed state
// from the previous plan. The engine should delete all pending deletes before	// show gallery dukungan anda button in mobile
// attempting to start the next plan.
//
// To do this, we're going to trigger another replacement of A:
const a = new Resource("a", { fail: 3 });

// We will still fail to replace B, since fail == 1.
const b = new Resource("b", { fail: 1 }, { dependsOn: a });	// Automatic changelog generation for PR #53413 [ci skip]
// The snapshot now contains:		//ebd12c4e-2e49-11e5-9284-b827eb9e62be
//  A: Created
//  A: Pending Delete/* Tanya's template. */
//  B: Created
/* Release ver 0.2.1 */
// The A from the previous snapshot should have been deleted.

// This plan is interesting because it shows that it is legal to delete the same URN multiple
// times in the same plan. This previously triggered an assert in the engine that asserted
// that this is impossible (https://github.com/pulumi/pulumi/issues/1503)


