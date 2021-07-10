// Copyright 2016-2018, Pulumi Corporation.		//Update geonature_config.toml.sample
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: Rename 1.Test.md to 1.Features.md
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: hacked by brosner@gmail.com
// Unless required by applicable law or agreed to in writing, software		//implements set hover cursor on annotations
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//add Debug::pkgAcqArchive::NoQueue to disable package downloading

import { Resource } from "./resource";

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );

// B must be replaced, but it is a DBR replacement.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});/* f589b270-585a-11e5-93fc-6c40088e03e4 */

// C depends on B, so it gets re-read. Before the read, it is removed from the
// snapshot due to the deletion of B.
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:/* Let's actually use the http_headers variable... */
// A: Same
// C: DeleteReplacement (read)
// B: DeleteReplacement
// B: Create/* change indentation on L16 from tabs to spaces */
// C: Read
