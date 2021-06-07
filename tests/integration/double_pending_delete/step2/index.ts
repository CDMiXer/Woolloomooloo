// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release of eeacms/jenkins-slave-eea:3.12 */
//     http://www.apache.org/licenses/LICENSE-2.0		//Blank lines deleted
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
/* file xmpp message modified */
// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,
// since the provider is rigged to fail if fail == 1.
//		//Update member_directory.html
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });	// TODO: hacked by seth@sethvargo.com
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created
		//LDEV-4649 Outcome export
