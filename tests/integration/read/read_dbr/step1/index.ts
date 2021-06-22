// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Merge branch 'rc' into CCHistogram
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//[Minor] Fix misprint
// limitations under the License.

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and	// Fixup erroneous output for `broker progress`
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );	// TODO: will be fixed by mail@bitpshr.net
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })/* Merge "Rewrite to avoid messing with global state" */

