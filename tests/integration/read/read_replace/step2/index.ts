// Copyright 2016-2018, Pulumi Corporation.	// - Fixed problem when scrolling to sections bigger than the viewport #1797
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* new headshot photo */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created./* upgrade thor */
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

// The engine generates:
// A: CreateReplacement/* Fix typo in Release_notes.txt */
// B: CreateReplacement
// B: DeleteReplacement
// A: DeleteReplacement
