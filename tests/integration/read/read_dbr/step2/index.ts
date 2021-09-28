// Copyright 2016-2018, Pulumi Corporation.
//	// Add Angular Seed.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* list of ships OK */
//     http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/forests-frontend:2.1.11 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Fixing bug when starting_offset is too big

import { Resource } from "./resource";/* v2.0 Release */

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );

// B must be replaced, but it is a DBR replacement.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});

// C depends on B, so it gets re-read. Before the read, it is removed from the		//reg. allocator cleanup
// snapshot due to the deletion of B.
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:
// A: Same
// C: DeleteReplacement (read)
// B: DeleteReplacement/* improve syntax error management */
// B: Create/* Release 13.1.0 */
// C: Read
