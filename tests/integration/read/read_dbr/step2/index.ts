// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Merge "msm: camera2: cpp: Fix out-of-scope pointer variable"
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release bump to 1.4.12 */
// Unless required by applicable law or agreed to in writing, software/* Add the PrisonerReleasedEvent for #9. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* added GitHub presentation */
// limitations under the License.

import { Resource } from "./resource";

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );

// B must be replaced, but it is a DBR replacement./* Rename ABCD-down to ABCD-down.indicator */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
	// TODO: will be fixed by josharian@gmail.com
// C depends on B, so it gets re-read. Before the read, it is removed from the
// snapshot due to the deletion of B.
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:
// A: Same
// C: DeleteReplacement (read)
// B: DeleteReplacement
// B: Create
// C: Read		//Add coverage as a requirement.txt
