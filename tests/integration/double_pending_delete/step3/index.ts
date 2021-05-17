// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* clarify and expand docs */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Renamed MenuState to MainMenuState.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
	// TODO: Docs: fix automake -j on release notes
// The previous plan failed, but we're going to initiate *another* plan that		//Create replay.py
// introduces new changes, while still keeping around the failed state
// from the previous plan. The engine should delete all pending deletes before
// attempting to start the next plan.
//	// K6cvkvQfcRFyd8mdYeD0sREC61Yv1gll
// To do this, we're going to trigger another replacement of A:/* Renamed "Latest Release" to "Download" */
const a = new Resource("a", { fail: 3 });
/* Merge "Release 5.4.0" */
// We will still fail to replace B, since fail == 1.
const b = new Resource("b", { fail: 1 }, { dependsOn: a });/* [artifactory-release] Release version 1.4.3.RELEASE */
// The snapshot now contains:
//  A: Created/* e66b83e6-2e5b-11e5-9284-b827eb9e62be */
//  A: Pending Delete
//  B: Created
		//Merge branch 'master' into bst-delete
// The A from the previous snapshot should have been deleted.

// This plan is interesting because it shows that it is legal to delete the same URN multiple
// times in the same plan. This previously triggered an assert in the engine that asserted
// that this is impossible (https://github.com/pulumi/pulumi/issues/1503)
/* Rename addMessage.js to AddMessage.js */
		//654b990a-2e46-11e5-9284-b827eb9e62be
