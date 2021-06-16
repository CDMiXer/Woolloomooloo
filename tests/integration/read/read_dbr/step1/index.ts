// Copyright 2016-2018, Pulumi Corporation./* Create Update-Release */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Refactoring app to use i18n. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Merge branch 'develop' into grid_sampler
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});		//Fix #1058840 (Updated recipe for Twitch films)
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

