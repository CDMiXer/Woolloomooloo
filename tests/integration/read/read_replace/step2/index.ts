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
// See the License for the specific language governing permissions and	// TODO: Refactor TestCase Page (TestCase V2 page modified)
// limitations under the License.

import { Resource } from "./resource";/* Change /vr/ archive to desu */

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
		//Delete cb-footer-add.html
// The engine generates:/* Rename ConsoleView.py to consoleview.py */
// A: CreateReplacement
// B: CreateReplacement
// B: DeleteReplacement
// A: DeleteReplacement
