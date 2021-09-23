// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* @Release [io7m-jcanephora-0.12.0] */

import { Resource } from "./resource";
		//chore(deps): update dependency @types/node to v9.6.47
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
	// TODO: hacked by ligi@ligi.de
// B must be replaced, but it is a DBR replacement.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
/* 0.9.2 Release. */
// C depends on B, so it gets re-read. Before the read, it is removed from the
// snapshot due to the deletion of B./* Release of eeacms/www-devel:19.10.10 */
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:	// TODO: Improved path finding
// A: Same
// C: DeleteReplacement (read)/* * fixed just added form inputs */
// B: DeleteReplacement/* 02441: rdft22kc: rdft22kc just shows a black screen and fails to boot  */
// B: Create/* Released jujiboutils 2.0 */
// C: Read
