// Copyright 2016-2018, Pulumi Corporation.	// TODO: Don't install it as a plugin.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Increased War Factory animations speed */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Commit Point and Vettore util's classes package. 
// limitations under the License.

import { Resource } from "./resource";
	// Readded "Show more" button
// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,/* [1.2.8] Patch 1 Release */
// since the provider is rigged to fail if fail == 1.
//	// TODO: docs(readme): remove jest from default install
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete/* Update some project settings */
//  B: Created

