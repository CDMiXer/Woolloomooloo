// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Delete ZXCT1009F.lib */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Experimental removal of tmp.css
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release 4.2.4  */
// limitations under the License.

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,		//Update mock_server.js
// since the provider is rigged to fail if fail == 1./* Release version 2.0.0.M1 */
//
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });/* Create Exercise4_VariablesAndNames.py */
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created

