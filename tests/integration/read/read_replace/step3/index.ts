// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Add main website to Readme and fix dates for school */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Add auto update param
// See the License for the specific language governing permissions and	// Make headers slightly smaller. [#86947212]
// limitations under the License.

import { Resource } from "./resource";/* Fixing build after updating `node-funargs`. */
		//534c12a6-2e5b-11e5-9284-b827eb9e62be
// Now go back the other way and make sure that "A" is external again.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

