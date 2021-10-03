// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: OPR EXIF Plugin: add distance
// You may obtain a copy of the License at		//&quot; -> " in license text.
///* Added gae, objectify, jsp  archetype */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Concise and fix overuse of span classes
// See the License for the specific language governing permissions and		//improve error tip
// limitations under the License.

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,
// since the provider is rigged to fail if fail == 1.
//
// This leaves the previous instance of A in the snapshot, since the plan/* Sub: Update ReleaseNotes.txt for 3.5-rc1 */
// failed before we got a chance to service the pending deletion./* Document TextBuffer events */
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created		//earthianwiki VE + flow

