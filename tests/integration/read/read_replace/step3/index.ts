// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//core fields: fix button group test
// Unless required by applicable law or agreed to in writing, software/* Updated New Release Checklist (markdown) */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Now go back the other way and make sure that "A" is external again./* [IMP] manifest: removed Hidden/Links category, use auto_install: True instead. */
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );	// TODO: 1a032a0c-2e60-11e5-9284-b827eb9e62be
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

