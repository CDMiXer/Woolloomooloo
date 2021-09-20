// Copyright 2016-2018, Pulumi Corporation.	// TODO: hacked by juan@benet.ai
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Bump Traefik to v2.3.5 */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Add qunit tests for adhoc task
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
		//Bump version to 0.12.1.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );		//Add experts to tag for tizen issues

// B must be replaced, but it is a DBR replacement.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
		//Delete Rain.png
// C depends on B, so it gets re-read. Before the read, it is removed from the
// snapshot due to the deletion of B.
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:
// A: Same
// C: DeleteReplacement (read)
// B: DeleteReplacement
// B: Create
// C: Read
