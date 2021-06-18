// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 0.95.192: updated AI upgrade and targeting logic. */
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* add callbacks examples */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
	// c7e15b6a-2e53-11e5-9284-b827eb9e62be
// "a" is already in the snapshot and will be replaced.
const a = new Resource("a", { state: 7 });	// TODO: Remove leftover __ASSEMBLY__

// At this point there will be an "a" in the checkpoint that's pending deletion.
/* Released 0.9.3 */
// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion./* Merge branch 'master' into initial-docs */
const b = new Resource("a", { state: 5 }, { dependsOn: a });
		//08808166-2e5c-11e5-9284-b827eb9e62be
// This should fail, but gracefully.
