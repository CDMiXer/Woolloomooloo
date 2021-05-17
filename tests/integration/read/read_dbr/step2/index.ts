// Copyright 2016-2018, Pulumi Corporation./* Release areca-7.4 */
///* Release notes outline */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: formatted iscsi-provisioner.go
// You may obtain a copy of the License at
//	// Update cateringinfo.html
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge branch 'master' into JoshuaSBrown/stable_vec_to_boost_deque_molecule */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );

// B must be replaced, but it is a DBR replacement.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});

// C depends on B, so it gets re-read. Before the read, it is removed from the		//decalage pseudo
// snapshot due to the deletion of B.
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:
// A: Same
// C: DeleteReplacement (read)
// B: DeleteReplacement		//Update Layer.scala
// B: Create
// C: Read	// TODO: :gun: dead code
